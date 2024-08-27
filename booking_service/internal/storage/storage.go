package storage

import (
	"booking_service/internal/infrastructura/repository/mongodb"
	"booking_service/internal/service"
	"booking_service/protos/bookingproto"
	bookingservice "booking_service/service/booking_service"
	"context"
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

func NewMongodb() (*mongo.Client, *mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("mongo_url"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("Booking").Collection("booking")
	return client, collection, nil
}

func Connection() {
	client, collection, err := NewMongodb()
	if err != nil {
		log.Println("connection mongodb error:", err)
		return
	}

	repo := mongodb.NewBookingMongodb(client, collection)

	s := service.NewBookingService(repo)

	handler := bookingservice.NewGrpcService(s)
	server := grpc.NewServer()

	bookingproto.RegisterBookingServiceServer(server, handler)

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
