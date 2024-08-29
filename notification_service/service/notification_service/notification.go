package notificationservice

import (
	"context"
	"log"
	"notification_service/internal/http/handler"
	"notification_service/notificationproto"
)

type Service struct {
	notificationproto.UnimplementedNotificationServer
	W handler.HandlerWebSocket
}

func (s *Service) AddUser(ctx context.Context, req *notificationproto.AddUserRequest) (*notificationproto.EmailResponse, error) {
	if err := s.W.AddUser(req.Id, nil); err != nil {
		log.Println(err)
		return nil, err
	}
	return &notificationproto.EmailResponse{
		Message: "User added successfully",
	}, nil
}
