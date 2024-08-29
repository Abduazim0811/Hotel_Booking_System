package router

import (
	"fmt"
	"log"
	"net"
	"notification_service/internal/storage"
	"notification_service/notificationproto"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewRouter() {
	r := gin.Default()
	handler := storage.NewService().W
	r.GET("/ws", handler.HandleWebSocket)

	go Grpc()

	fmt.Println("Server started on port 8083")
	if err := r.Run("localhost:8083"); err != nil {
		log.Fatal(err)
	}
}

func Grpc() {
	listener, err := net.Listen("tcp", os.Getenv("server_url"))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	s := grpc.NewServer()
	server := storage.NewService()
	notificationproto.RegisterNotificationServer(s, server)
	reflection.Register(s)

	fmt.Printf("gRPC server started on port %s\n", os.Getenv("server_url"))

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
