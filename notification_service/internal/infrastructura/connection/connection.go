package connection

import (
	"log"
	"net"
	"notification_service/internal/infrastructura/repository/kafka"
	"notification_service/notificationproto"
	notificationservice "notification_service/service/notification_service"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/grpc"
)

func StartServer() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	kafkaConsumer := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.FetchMaxBytes(1<<20),
	)

	notificationService := notificationservice.NewNotificationService(kafkaConsumer)
	notificationproto.RegisterNotificationServiceServer(grpcServer, notificationService)

	go kafka.NewNOtificationKafka().ConsumeMessages("notification-topic")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
