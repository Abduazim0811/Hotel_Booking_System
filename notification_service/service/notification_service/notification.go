package notificationservice

import (
	"context"
	"log"
	"notification_service/notificationproto"

	"github.com/twmb/franz-go/pkg/kgo"
)

type NotificationService struct {
	notificationproto.UnimplementedNotificationServiceServer
	kafkaConsumer *kgo.Client
}

func NewNotificationService(kafkaConsumer *kgo.Client) *NotificationService {
	return &NotificationService{kafkaConsumer: kafkaConsumer}
}

func (ns *NotificationService) SendNotification(ctx context.Context, req *notificationproto.NotificationRequest) (*notificationproto.NotificationResponse, error) {
	switch req.Action {
	case "register":
		ns.handleRegister(req.Details)
	case "update":
		ns.handleUpdate(req.Details)
	case "delete":
		ns.handleDelete(req.Details)
	default:
		log.Printf("Unknown action: %s", req.Action)
		return &notificationproto.NotificationResponse{Message: "Unknown action"}, nil
	}

	return &notificationproto.NotificationResponse{Message: "Notification sent successfully"}, nil
}

func (ns *NotificationService) handleRegister(details string) {
	subject := "User Registration"
	body := email.GenerateEmailBody("Registration", details)
	err := email.SendEmail("recipient@example.com", subject, body)
	if err != nil {
		log.Printf("Error sending registration email: %v", err)
	}
}

func (ns *NotificationService) handleUpdate(details string) {
	subject := "User Update"
	body := email.GenerateEmailBody("Update", details)
	err := email.SendEmail("recipient@example.com", subject, body)
	if err != nil {
		log.Printf("Error sending update email: %v", err)
	}
}

func (ns *NotificationService) handleDelete(details string) {
	subject := "User Deletion"
	body := email.GenerateEmailBody("Deletion", details)
	err := email.SendEmail("recipient@example.com", subject, body)
	if err != nil {
		log.Printf("Error sending deletion email: %v", err)
	}
}
