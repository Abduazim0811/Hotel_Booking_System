package hotels

import (
	"booking_service/protos/hotelproto"
	"context"
	"log"

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

func Hotels(ctx context.Context, id string) error {
	var req hotelproto.HotelResponse
	req.HotelId = id
	_, err := DialHotelGrpc().GetbyIdHotel(ctx, &req)
	return err
}

func Rooms(ctx context.Context, id string) error {
	var req hotelproto.RoomResponse
	req.RoomId = id
	_, err := DialHotelGrpc().GetbyIdRoom(ctx, &req)
	return err
}
