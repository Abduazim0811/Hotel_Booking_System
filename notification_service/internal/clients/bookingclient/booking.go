package bookingclient

import (
	"log"
	"notification_service/internal/protos/bookingproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialGrpcBooking() bookingproto.BookingServiceClient {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client booking:", err)
	}

	return bookingproto.NewBookingServiceClient(conn)
}
