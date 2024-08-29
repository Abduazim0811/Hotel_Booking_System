package mongodb

import (
	"booking_service/internal/entity/booking"
	"booking_service/internal/infrastructura/repository"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingMongodb struct {
	client       *mongo.Client
	b_collection *mongo.Collection
	w_collection *mongo.Collection
	ctx          context.Context
}

func NewBookingMongodb(client *mongo.Client, collection *mongo.Collection, waitingcollection *mongo.Collection) repository.BookingRepository {
	return &BookingMongodb{client: client, b_collection: collection, w_collection: waitingcollection}
}

func (b *BookingMongodb) Create(req booking.BookingRequest) (*booking.BookingResponse, error) {
	inserted, err := b.b_collection.InsertOne(b.ctx, req)
	if err != nil {
		log.Println("Error creating booking")
		return nil, fmt.Errorf("error creating booking: %v", err)
	}
	bookingId := inserted.InsertedID.(primitive.ObjectID)
	res := &booking.BookingResponse{
		BookingID:    bookingId.Hex(),
		UserID:       req.UserID,
		HotelID:      req.HotelID,
		RoomID:       req.RoomID,
		RoomType:     req.RoomType,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		TotalAmount:  req.TotalAmount,
		Status:       "Created",
	}

	return res, nil
}

func (b *BookingMongodb) GetById(req booking.GetRequest) (*booking.BookingResponse, error) {
	var res booking.BookingResponse
	bookingID, err := primitive.ObjectIDFromHex(req.BookingID)
	if err != nil {
		log.Println("Invalid booking ID:", err)
		return nil, err
	}
	err = b.b_collection.FindOne(b.ctx, bson.M{"_id": bookingID}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Booking not found")
			return nil, err
		}
		log.Println("Error finding booking:", err)
		return nil, err
	}

	return &res, nil
}

func (b *BookingMongodb) Update(req booking.UpdateRequest) (*booking.BookingResponse, error) {
	bookingID, err := primitive.ObjectIDFromHex(req.BookingID)
	if err != nil {
		log.Println("Invalid booking ID:", err)
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"roomId":       req.RoomID,
			"roomType":     req.RoomType,
			"checkInDate":  req.CheckInDate,
			"checkOutDate": req.CheckOutDate,
			"status":       req.Status,
		},
	}

	_, err = b.b_collection.UpdateOne(b.ctx, bson.M{"_id": bookingID}, update)
	if err != nil {
		log.Println("Error updating booking:", err)
		return nil, err
	}

	updatedBooking := &booking.BookingResponse{
		BookingID:    req.BookingID,
		RoomID:       req.RoomID,
		RoomType:     req.RoomType,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		Status:       req.Status,
	}

	return updatedBooking, nil
}

func (b *BookingMongodb) Delete(req booking.GetRequest) (*booking.DeleteResponse, error) {
	bookingID, err := primitive.ObjectIDFromHex(req.BookingID)
	if err != nil {
		log.Println("Invalid booking ID:", err)
		return nil, err
	}

	res, err := b.b_collection.DeleteOne(b.ctx, bson.M{"_id": bookingID})
	if err != nil {
		log.Println("Error deleting booking:", err)
		return nil, err
	}

	if res.DeletedCount == 0 {
		log.Println("Booking not found")
		return nil, mongo.ErrNoDocuments
	}

	return &booking.DeleteResponse{
		Message:   "Booking deleted successfully",
		BookingID: req.BookingID,
	}, nil
}

func (b *BookingMongodb) GetByUserId(req booking.GetUsersRequest) ([]*booking.BookingResponse, error) {
	var res []*booking.BookingResponse

	cursor, err := b.b_collection.Find(b.ctx, bson.M{"userid": req.UserID})
	if err != nil {
		log.Println("Error finding bookings for user:", err)
		return nil, err
	}

	defer cursor.Close(b.ctx)

	for cursor.Next(b.ctx) {
		var booking *booking.BookingResponse
		if err := cursor.Decode(&booking); err != nil {
			log.Println("Error decoding booking:", err)
			return nil, err
		}
		res = append(res, booking)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return res, nil
}

func (b *BookingMongodb) AddWaiting(req booking.CreateWaitingList) error {
	_, err := b.w_collection.InsertOne(b.ctx, req)
	if err != nil {
		log.Println("error create waiting")
		return fmt.Errorf("error create waiting")
	}

	return nil
}

func (b *BookingMongodb) GetbyIdwaitingList(id string) (*booking.GetWaitingResponse, error) {
	waitingid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid waiting list ID:", err)
		return nil, fmt.Errorf("invalid waiting list  id error: %v", err)
	}
	var res booking.GetWaitingResponse

	err = b.w_collection.FindOne(b.ctx, bson.M{"_id": waitingid}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Booking not found")
			return nil, fmt.Errorf("waiting list not found")
		}
		log.Println("Error finding waiting list:", err)
		return nil, fmt.Errorf("error waiting list")
	}

	return &res, nil
}

func (b *BookingMongodb) GetAllWaiting() (*booking.WaitingList, error) {
	var res []booking.GetWaitingResponse

	cursor, err := b.w_collection.Find(b.ctx, bson.M{})
	if err != nil {
		log.Println("Error finding all waiting lists:", err)
		return nil, fmt.Errorf("error finding all waiting lists")
	}
	defer cursor.Close(b.ctx)

	for cursor.Next(b.ctx) {
		var waiting booking.GetWaitingResponse
		if err := cursor.Decode(&waiting); err != nil {
			log.Println("Error decoding waiting list:", err)
			return nil, fmt.Errorf("error decoding waiting list")
		}
		res = append(res, waiting)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, fmt.Errorf("cursor error")
	}

	return &booking.WaitingList{Users: res}, nil
}

func (b *BookingMongodb) UpdateWaiting(req booking.UpdateWaitingListRequest) error {
	waitingID, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		log.Println("Invalid waiting list ID:", err)
		return fmt.Errorf("invalid waiting list ID")
	}

	update := bson.M{
		"$set": bson.M{
			"user_id":       req.UserID,
			"room_type":     req.RoomType,
			"hotel_id":      req.HotelID,
			"check_in_date": req.CheckInDate,
			"check_out_date": req.CheckOutDate,
		},
	}

	_, err = b.w_collection.UpdateOne(b.ctx, bson.M{"_id": waitingID}, update)
	if err != nil {
		log.Println("Error updating waiting list:", err)
		return fmt.Errorf("error updating waiting list: %v", err)
	}

	return nil
}

func (b *BookingMongodb) DeleteWaiting(req booking.GetWaitingRequest)error{
	waitingID, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		log.Println("Invalid waiting list ID:", err)
		return fmt.Errorf("invalid waiting list ID: %v", err)
	}

	res, err := b.w_collection.DeleteOne(b.ctx, bson.M{"_id": waitingID})
	if err != nil {
		log.Println("Error deleting booking:", err)
		return fmt.Errorf("error deleting booking:%v", err)
	}

	if res.DeletedCount == 0 {
		log.Println("Booking not found")
		return  mongo.ErrNoDocuments
	}
	return nil

}
