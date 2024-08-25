package storage

import (
	"context"
	"hotel_service/hotelproto"
	"hotel_service/internal/infrastructura/repository/mongodb"
	"hotel_service/internal/service"
	hotelservice "hotel_service/service/hotel_service"
	"log"
	"net"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func NewMongodb() (*mongo.Client, *mongo.Collection, *mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("mongo_url"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	hotelCollection := client.Database("Hotels").Collection("hotel")
	roomCollection := client.Database("Hotels").Collection("room")

	return client, hotelCollection, roomCollection, nil
}

func Connection() {
	client, hotelCollection, roomCollection, err := NewMongodb()
	if err != nil {
		log.Println("connection mongodb error:", err)
		return
	}

	repo := mongodb.NewHotelMongodb(client, hotelCollection, roomCollection)
	service := service.NewHotelService(repo)
	handler := hotelservice.HotelGrpcService(service)

	server := grpc.NewServer()
	hotelproto.RegisterHotelServiceServer(server, handler)

	lis, err := net.Listen("tcp", os.Getenv("server_url"))
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	log.Println("Server is listening on port ", os.Getenv("server_url"))
	if err = server.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}

}
