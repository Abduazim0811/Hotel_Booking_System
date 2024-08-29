package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"notification_service/internal/clients/bookingclient"
	"notification_service/internal/clients/hotelclient"
	"notification_service/internal/pkg/email"
	"notification_service/internal/protos/bookingproto"
	"notification_service/internal/protos/hotelproto"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/twmb/franz-go/pkg/kgo"
)

type HandlerWebSocket struct {
	Map     map[string]*websocket.Conn
	Mutex   *sync.Mutex
	Ctx     context.Context
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (u *HandlerWebSocket) HandleWebSocket(c *gin.Context) {
	fmt.Println("WebSocket is working")
	userID := c.Request.Header.Get("id")
	fmt.Println(userID)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Current connections map:", u.Map)

	kafkaReader, err := kgo.NewClient(
		kgo.SeedBrokers("broker:29092"),
		kgo.ConsumeTopics("notification"),
	)
	if err != nil {
		log.Println("Kafka client creation error:", err)
		return
	}
	defer kafkaReader.Close()

	for {
		fetches := kafkaReader.PollFetches(c.Request.Context())
		if fetches.IsClientClosed() {
			break
		}
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			message := record.Value
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("Error writing message to WebSocket:", err)
				return
			}
			u.RoomChecker(conn)
			fmt.Println("Room check is over")
		}
	}
}

func (u *HandlerWebSocket) AddUser(userID string, conn *websocket.Conn) error {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()
	if _, exists := u.Map[userID]; exists {
		log.Printf("User %s is already connected", userID)
		return errors.New("user already exists")
	}
	u.Map[userID] = conn
	log.Printf("User %s added to the map", userID)
	return nil
}

func (u *HandlerWebSocket) RoomChecker(conn *websocket.Conn) {
	res, err := hotelclient.DialHotelGrpc().GetAllHotels(u.Ctx, &hotelproto.HotelEmpty{})
	if err != nil {
		log.Println(err)
	}
	for _, v := range res.Hotel {
		u.RoomDetective(v.HotelId, conn)
	}
}

func (u *HandlerWebSocket) RoomDetective(id string, conn *websocket.Conn) {
	rooms, err := hotelclient.DialHotelGrpc().GetAllRooms(u.Ctx, &hotelproto.HotelResponse{HotelId: id})
	if err != nil {
		log.Println(err)
	} else {
		for _, v := range rooms.Room {
			if v.Availability {
				message := fmt.Sprintf("Hotel ID %v\nRoom ID %v\nRoom Type %v\nRoom Price Per Night %v\nRoom Available %v\n",
					v.HotelId, v.RoomId, v.RoomType, v.PricePerNight, v.Availability)
				if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					log.Println("Error writing message to WebSocket:", err)
				}
				u.WaitingUsers(message)
			}
		}
	}
}

func (u *HandlerWebSocket) WaitingUsers(body string) {
	users, err := bookingclient.DialGrpcBooking().GetAllWaiting(u.Ctx, &bookingproto.Empty{})
	if err != nil {
		log.Println(err)
	} else {
		for _, v := range users.Users {
			fmt.Println(v.UserEmail)
			if err := email.SendEmail(v.UserEmail, body); err != nil {
				log.Println(err)
			}
		}
	}
}
