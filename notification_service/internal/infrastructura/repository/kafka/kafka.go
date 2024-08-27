package kafka

import (
	"context"
	"encoding/json"
	"log"
	"notification_service/notificationproto"
	notificationservice "notification_service/service/notification_service"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

type NotificationKafka struct {
	kafkaConsumer *kgo.Client
	service       notificationservice.NotificationService
}

func NewNotificationKafka(con *kgo.Client, svc notificationservice.NotificationService) *NotificationKafka {
	return &NotificationKafka{
		kafkaConsumer: con,
		service:       svc,
	}
}

func (ns *NotificationKafka) ConsumeMessages(topic string) {
	ctx := context.Background()

	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.FetchMaxBytes(1<<20),
		kgo.FetchMaxWait(1),
	)
	if err != nil {
		log.Fatalf("Failed to create Kafka client: %v", err)
	}
	defer client.Close()

	client.AssignPartitions(kgo.SeedTopics(topic))

	for {
		fetches := client.FetchPartitions(ctx, kgo.FetchPartitionOffsets())
		if fetches.Err() != nil {
			log.Printf("Error fetching Kafka messages: %v", fetches.Err())
			time.Sleep(1 * time.Second)
			continue
		}

		for _, partition := range fetches.Partitions() {
			for _, message := range partition.Messages {
				var msg notificationproto.NotificationRequest
				if err := json.Unmarshal(message.Value, &msg); err != nil {
					log.Printf("Failed to unmarshal Kafka message: %v", err)
					continue
				}
				if _, err := ns.service.SendNotification(ctx, &msg); err != nil {
					log.Printf("Failed to send notification: %v", err)
				}
			}
		}

		time.Sleep(1 * time.Second)
	}
}
