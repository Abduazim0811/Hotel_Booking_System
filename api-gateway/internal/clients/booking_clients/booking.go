package bookingclients

import (
	"api-gateway/internal/protos/bookingproto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialBookingGrpc() bookingproto.BookingServiceClient {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client Booking:", err)
	}
	return bookingproto.NewBookingServiceClient(conn)
}
