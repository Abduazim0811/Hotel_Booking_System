package users

import (
	"booking_service/protos/userproto"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialUserGrpc() userproto.UserServiceClient {
	conn, err := grpc.NewClient("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client hotel:", err)
	}

	return userproto.NewUserServiceClient(conn)
}

func Users(ctx context.Context, id int32) error {
	var req userproto.GetUserRequest
	req.Id = id
	_, err := DialUserGrpc().GetByIdUser(ctx, &req)
	return err
}
