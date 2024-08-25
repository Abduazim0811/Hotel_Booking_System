package mongodb

import (
	"context"
	"hotel_service/internal/entity/hotel"
	"hotel_service/internal/infrastructura/repository"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelMongodb struct {
	client     *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
}

func NewHotelMongodb(client *mongo.Client, collection *mongo.Collection) repository.HotelRepository {
	return &HotelMongodb{client: client, collection: collection}
}

func (h *HotelMongodb) AddHotel(req hotel.HotelRequest) (*hotel.HotelResponse, error) {
	res, err := h.collection.InsertOne(h.ctx, req)
	if err != nil {
		log.Println("Error inserting hotel:", err)
		return nil, err
	}

	insertedId := res.InsertedID.(primitive.ObjectID).Hex()

	return &hotel.HotelResponse{HotelID: insertedId}, nil
}

func (h *HotelMongodb) GetbyId(req hotel.HotelResponse) (*hotel.Hotel, error) {
	var res hotel.Hotel
	id, err := primitive.ObjectIDFromHex(req.HotelID)
	if err != nil {
		log.Println("objectid error")
		return nil, err
	}

	err = h.collection.FindOne(h.ctx, bson.M{"_id": id}).Decode(&res)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			log.Println("hotel not found")
			return nil, err
		}
	}
	return &res, nil
}

func (h *HotelMongodb) GetAll()(*[]hotel.Hotel, error){
	var resHotels []hotel.Hotel

	cursor, err := h.collection.Find(h.ctx, bson.M{})
	if err != nil {
		log.Println("failed to get all hotel")
		return nil, err
	}

	for cursor.Next(h.ctx){
		var res hotel.Hotel
		if err := cursor.Decode(&res); err != nil {
			log.Println("Failed to decode item:", err)
			return nil, err
		}

		resHotels =  append(resHotels, res)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}
	return &resHotels, nil
}

func (h *HotelMongodb) UpdateHotel(req hotel.Hotel) error {
	id, err := primitive.ObjectIDFromHex(req.HotelID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":     req.Name,
			"location": req.Location,
			"rating":   req.Rating,
			"address":  req.Address,
		},
	}

	res, err := h.collection.UpdateOne(h.ctx, filter, update)
	if err != nil {
		log.Println("Error updating hotel:", err)
		return err
	}

	if res.MatchedCount == 0 {
		log.Println("No hotel found with the given ID")
		return mongo.ErrNoDocuments
	}

	return nil
}

func (h *HotelMongodb) DeleteHotel(hotelID string) error {
	id, err := primitive.ObjectIDFromHex(hotelID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return err
	}

	filter := bson.M{"_id": id}

	res, err := h.collection.DeleteOne(h.ctx, filter)
	if err != nil {
		log.Println("Error deleting hotel:", err)
		return err
	}

	if res.DeletedCount == 0 {
		log.Println("No hotel found with the given ID")
		return mongo.ErrNoDocuments
	}

	return nil
}

func (h *HotelMongodb) CreateRoom(req hotel.RoomRequest) (*hotel.RoomResponse, error) {
	res, err := h.collection.InsertOne(h.ctx, req)
	if err != nil {
		log.Println("Error inserting room:", err)
		return nil, err
	}

	insertedId := res.InsertedID.(primitive.ObjectID).Hex()

	return &hotel.RoomResponse{RoomID: insertedId}, nil
}

func (h *HotelMongodb) GetRoomById(req hotel.RoomResponse) (*hotel.Room, error) {
	var res hotel.Room
	id, err := primitive.ObjectIDFromHex(req.RoomID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return nil, err
	}

	err = h.collection.FindOne(h.ctx, bson.M{"_id": id}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Room not found")
			return nil, err
		}
		log.Println("Error finding room:", err)
		return nil, err
	}

	return &res, nil
}

func (h *HotelMongodb) GetAllRooms() (*[]hotel.Room, error) {
	var resRooms []hotel.Room

	cursor, err := h.collection.Find(h.ctx, bson.M{})
	if err != nil {
		log.Println("Failed to get all rooms:", err)
		return nil, err
	}

	for cursor.Next(h.ctx) {
		var res hotel.Room
		if err := cursor.Decode(&res); err != nil {
			log.Println("Failed to decode room:", err)
			return nil, err
		}
		resRooms = append(resRooms, res)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return &resRooms, nil
}

func (h *HotelMongodb) UpdateRoom(req hotel.Room) error {
	id, err := primitive.ObjectIDFromHex(req.RoomID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"roomType":       req.RoomType,
			"pricePerNight":  req.PricePerNight,
			"availability":   req.Availability,
		},
	}

	res, err := h.collection.UpdateOne(h.ctx, filter, update)
	if err != nil {
		log.Println("Error updating room:", err)
		return err
	}

	if res.MatchedCount == 0 {
		log.Println("No room found with the given ID")
		return mongo.ErrNoDocuments
	}

	return nil
}

func (h *HotelMongodb) DeleteRoom(roomID string) error {
	id, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return err
	}

	filter := bson.M{"_id": id}

	res, err := h.collection.DeleteOne(h.ctx, filter)
	if err != nil {
		log.Println("Error deleting room:", err)
		return err
	}

	if res.DeletedCount == 0 {
		log.Println("No room found with the given ID")
		return mongo.ErrNoDocuments
	}

	return nil
}


