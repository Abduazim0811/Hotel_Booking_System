package hotelclient

import (
	"log"
	"notification_service/internal/protos/hotelproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialHotelGrpc() hotelproto.HotelServiceClient {
	conn, err := grpc.NewClient("localhost:9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client hotel:", err)
	}

	return hotelproto.NewHotelServiceClient(conn)
}
