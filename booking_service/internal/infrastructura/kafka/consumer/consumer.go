package consumer

import (
	"booking_service/internal/entity/booking"
	"booking_service/protos/bookingproto"
	bookingservice "booking_service/service/booking_service"
	"context"
	"encoding/json"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookingConsumer struct {
	Ctx context.Context
	service *bookingservice.Service
}

func NewBookingConsumer(service *bookingservice.Service) *BookingConsumer{
	return &BookingConsumer{service: service}
}

func (u *BookingConsumer) Consumer() {
	client, err := kgo.NewClient(
		kgo.SeedBrokers("broker:29092"),
		kgo.ConsumeTopics("booking"),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	for {
		fetches := client.PollFetches(ctx)
		if err := fetches.Errors(); len(err) > 0 {
			log.Fatal(err)
		}
		fetches.EachPartition(func(ftp kgo.FetchTopicPartition) {
			for _, record := range ftp.Records {
				if err := u.Adjust(record); err != nil {
					log.Println(err)
				}
			}
		})
	}
}

func (u *BookingConsumer) Adjust(record *kgo.Record) error {
	switch string(record.Key) {
	case "create":
		if err := u.Create(record.Value); err != nil {
			log.Println(err)
			return nil
		}
	case "update":
		if err := u.Update(record.Value); err != nil {
			log.Println(err)
			return err
		}
	case "delete":
		if err := u.Delete(record.Value); err != nil {
			log.Println(err)
			return err
		}
	case "createW":
		if err := u.CreateW(record.Value); err != nil {
			log.Println(err)
			return nil
		}
	case "updateW":
		if err := u.UpdateW(record.Value); err != nil {
			log.Println(err)
			return nil
		}
	case "deleteW":
		if err := u.DeleteW(record.Value); err != nil {
			log.Println(err)
			return nil
		}
	}
	return nil
}

func (u *BookingConsumer) Create(req []byte) error {
	var req1 booking.BookingRequest

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = bookingproto.BookingRequest{
		Userid:       req1.UserID,
		Hotelid:      req1.HotelID,
		RoomId:       req1.RoomID,
		Roomtype:     req1.RoomType,
		CheckInDate:  timestamppb.New(req1.CheckInDate),
		CheckOutDate: timestamppb.New(req1.CheckOutDate),
	}
	_, err := u.service.CreateBooking(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *BookingConsumer) Update(req []byte) error {
	var req1 booking.UpdateRequest

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = bookingproto.UpdateRequest{
		BookingId:    req1.BookingID,
		RoomId:       req1.RoomID,
		Roomtype:     req1.RoomType,
		CheckInDate:  timestamppb.New(req1.CheckInDate),
		CheckOutDate: timestamppb.New(req1.CheckOutDate),
	}
	_, err := u.service.UpdateBooking(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *BookingConsumer) Delete(req []byte) error {
	var req1 booking.GetRequest
	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}

	var newreq = bookingproto.GetRequest{
		BookingId: req1.BookingID,
	}

	_, err := u.service.DeleteBooking(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *BookingConsumer) CreateW(req []byte) error {
	var req1 booking.CreateWaitingList

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = bookingproto.CreateWaitingList{
		UserId:       req1.UserID,
		UserEmail:    req1.UserEmail,
		RoomType:     req1.RoomType,
		HotelId:      req1.HotelID,
		CheckInDate:  timestamppb.New(req1.CheckInDate),
		CheckOutDate: timestamppb.New(req1.CheckOutDate),
	}
	_, err := u.service.CreateWaiting(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *BookingConsumer) UpdateW(req []byte) error {
	var req1 booking.UpdateWaitingListRequest
	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}

	var newreq = bookingproto.UpdateWaitingListRequest{
		Id:           req1.ID,
		UserId:       req1.UserID,
		RoomType:     req1.RoomType,
		HotelId:      req1.HotelID,
		CheckInDate:  timestamppb.New(req1.CheckInDate),
		CheckOutDate: timestamppb.New(req1.CheckOutDate),
	}
	_, err := u.service.UpdateWaiting(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *BookingConsumer) DeleteW(req []byte) error {
	var req1 booking.GetWaitingRequest
	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}

	var newreq = bookingproto.GetWaitingRequest{
		Id: req1.ID,
	}

	_, err := u.service.DeleteWaiting(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}