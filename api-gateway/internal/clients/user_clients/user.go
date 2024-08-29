package userclients

import (
	"api-gateway/internal/protos/userproto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialUserGrpc() userproto.UserServiceClient {
	conn, err := grpc.NewClient("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("gRPC serverga ulanishda xato:", err)
	}

	return userproto.NewUserServiceClient(conn)
}
