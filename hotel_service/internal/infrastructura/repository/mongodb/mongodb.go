package mongodb

import (
	"fmt" // qo'shing
	"log"
	"context"
	"hotel_service/internal/entity/hotel"
	"hotel_service/internal/infrastructura/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelMongodb struct {
	client     *mongo.Client
	hotelcollection *mongo.Collection
	roomcollection *mongo.Collection
	ctx        context.Context
}

func NewHotelMongodb(client *mongo.Client, hotelcollection *mongo.Collection, roomcollection *mongo.Collection) repository.HotelRepository {
	return &HotelMongodb{client: client, hotelcollection: hotelcollection, roomcollection: roomcollection}
}

func (h *HotelMongodb) AddHotel(req hotel.HotelRequest) (*hotel.HotelResponse, error) {
	res, err := h.hotelcollection.InsertOne(h.ctx, req)
	if err != nil {
		log.Println("Error inserting hotel:", err)
		return nil, fmt.Errorf("failed to insert hotel: %w", err)
	}

	insertedId := res.InsertedID.(primitive.ObjectID).Hex()

	return &hotel.HotelResponse{HotelID: insertedId}, nil
}

func (h *HotelMongodb) GetbyId(req string) (*hotel.Hotel, error) {
	var res hotel.Hotel
	id, err := primitive.ObjectIDFromHex(req)
	if err != nil {
		log.Println("objectid error")
		return nil, fmt.Errorf("invalid objectid format: %w", err)
	}

	err = h.hotelcollection.FindOne(h.ctx, bson.M{"_id": id}).Decode(&res)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			log.Println("hotel not found")
			return nil, fmt.Errorf("hotel not found: %w", err)
		}
		return nil, fmt.Errorf("error finding hotel: %w", err)
	}
	return &res, nil
}

func (h *HotelMongodb) GetAll()(*[]hotel.Hotel, error){
	var resHotels []hotel.Hotel

	cursor, err := h.hotelcollection.Find(h.ctx, bson.M{})
	if err != nil {
		log.Println("failed to get all hotel")
		return nil, fmt.Errorf("failed to get all hotels: %w", err)
	}

	for cursor.Next(h.ctx){
		var res hotel.Hotel
		if err := cursor.Decode(&res); err != nil {
			log.Println("Failed to decode item:", err)
			return nil, fmt.Errorf("failed to decode hotel: %w", err)
		}

		resHotels =  append(resHotels, res)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, fmt.Errorf("cursor error: %w", err)
	}
	return &resHotels, nil
}

func (h *HotelMongodb) UpdateHotel(req hotel.Hotel) error {
	id, err := primitive.ObjectIDFromHex(req.HotelID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return fmt.Errorf("invalid objectid format: %w", err)
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

	res, err := h.hotelcollection.UpdateOne(h.ctx, filter, update)
	if err != nil {
		log.Println("Error updating hotel:", err)
		return fmt.Errorf("failed to update hotel: %w", err)
	}

	if res.MatchedCount == 0 {
		log.Println("No hotel found with the given ID")
		return fmt.Errorf("no hotel found with the given ID: %w", mongo.ErrNoDocuments)
	}

	return nil
}

func (h *HotelMongodb) DeleteHotel(hotelID string) error {
	id, err := primitive.ObjectIDFromHex(hotelID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return fmt.Errorf("invalid objectid format: %w", err)
	}

	filter := bson.M{"_id": id}

	res, err := h.hotelcollection.DeleteOne(h.ctx, filter)
	if err != nil {
		log.Println("Error deleting hotel:", err)
		return fmt.Errorf("failed to delete hotel: %w", err)
	}

	if res.DeletedCount == 0 {
		log.Println("No hotel found with the given ID")
		return fmt.Errorf("no hotel found with the given ID: %w", mongo.ErrNoDocuments)
	}

	return nil
}

func (h *HotelMongodb) CreateRoom(req hotel.RoomRequest) (*hotel.RoomResponse, error) {
	res, err := h.roomcollection.InsertOne(h.ctx, req)
	if err != nil {
		log.Println("Error inserting room:", err)
		return nil, fmt.Errorf("failed to insert room: %w", err)
	}

	insertedId := res.InsertedID.(primitive.ObjectID).Hex()

	return &hotel.RoomResponse{RoomID: insertedId}, nil
}

func (h *HotelMongodb) GetRoomById(req hotel.RoomResponse) (*hotel.Room, error) {
	var res hotel.Room
	id, err := primitive.ObjectIDFromHex(req.RoomID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return nil, fmt.Errorf("invalid objectid format: %w", err)
	}

	err = h.roomcollection.FindOne(h.ctx, bson.M{"_id": id}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Room not found")
			return nil, fmt.Errorf("room not found: %w", err)
		}
		log.Println("Error finding room:", err)
		return nil, fmt.Errorf("failed to find room: %w", err)
	}

	return &res, nil
}



func (h *HotelMongodb) GetAllRooms(id string) (*[]hotel.Room, error) {
	var resRooms []hotel.Room

	cursor, err := h.roomcollection.Find(h.ctx, bson.M{"hotel_id": id})
	if err != nil {
		log.Println("Failed to get all rooms:", err)
		return nil, fmt.Errorf("failed to get all rooms: %w", err)
	}

	for cursor.Next(h.ctx) {
		var res hotel.Room
		if err := cursor.Decode(&res); err != nil {
			log.Println("Failed to decode room:", err)
			return nil, fmt.Errorf("failed to decode room: %w", err)
		}
		resRooms = append(resRooms, res)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return &resRooms, nil
}

func (h *HotelMongodb) UpdateRoom(req hotel.Room) error {
	id, err := primitive.ObjectIDFromHex(req.RoomID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return fmt.Errorf("invalid objectid format: %w", err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"roomType":       req.RoomType,
			"pricePerNight":  req.PricePerNight,
			"availability":   req.Availability,
		},
	}

	res, err := h.roomcollection.UpdateOne(h.ctx, filter, update)
	if err != nil {
		log.Println("Error updating room:", err)
		return fmt.Errorf("failed to update room: %w", err)
	}

	if res.MatchedCount == 0 {
		log.Println("No room found with the given ID")
		return fmt.Errorf("no room found with the given ID: %w", mongo.ErrNoDocuments)
	}

	return nil
}

func (h *HotelMongodb) DeleteRoom(roomID string) error {
	id, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return fmt.Errorf("invalid objectid format: %w", err)
	}

	filter := bson.M{"_id": id}

	res, err := h.roomcollection.DeleteOne(h.ctx, filter)
	if err != nil {
		log.Println("Error deleting room:", err)
		return fmt.Errorf("failed to delete room: %w", err)
	}

	if res.DeletedCount == 0 {
		log.Println("No room found with the given ID")
		return fmt.Errorf("no room found with the given ID: %w", mongo.ErrNoDocuments)
	}

	return nil
}
