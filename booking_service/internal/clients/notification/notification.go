package notification

import (
	"booking_service/protos/notificationproto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialNotificationGrpc() notificationproto.NotificationClient{
	conn, err := grpc.NewClient("localhost:8887", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client notification:",err)
	}
	return notificationproto.NewNotificationClient(conn)
}
