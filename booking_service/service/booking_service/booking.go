package bookingservice

import (
	"booking_service/internal/clients/hotels"
	"booking_service/internal/clients/users"
	"booking_service/internal/entity/booking"
	"booking_service/internal/service"
	"booking_service/protos/bookingproto"
	"booking_service/protos/hotelproto"
	"booking_service/protos/notificationproto"
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	bookingproto.UnimplementedBookingServiceServer
	service *service.BookingService
	h       hotelproto.HotelServiceClient
	n       notificationproto.NotificationClient
}

func NewGrpcService(s *service.BookingService) *Service {
	return &Service{service: s}
}

var userid int

func (s *Service) CreateBooking(ctx context.Context, req *bookingproto.BookingRequest) (*bookingproto.BookingResponse, error) {
	var bookingreq booking.BookingRequest
	bookingreq.HotelID = req.Hotelid
	bookingreq.RoomID = req.RoomId
	bookingreq.UserID = req.Userid
	bookingreq.RoomType = req.Roomtype
	bookingreq.CheckInDate = req.CheckInDate.AsTime()
	bookingreq.CheckOutDate = req.CheckOutDate.AsTime()
	bookingreq.TotalAmount = req.TotalAmount

	err := hotels.Hotels(ctx, bookingreq.HotelID)
	if err != nil {
		_, err := s.n.Notification(ctx, &notificationproto.ProduceMessageRequest{Id: req.Hotelid, Message: err.Error()})
		if err != nil {
			log.Println(err)
		}
		log.Println("hotel not found")
		return nil, fmt.Errorf("hotel not found: %v", err)
	}

	err = hotels.Rooms(ctx, bookingreq.RoomID)
	if err != nil {
		_, err := s.n.Notification(ctx, &notificationproto.ProduceMessageRequest{Id: req.RoomId, Message: err.Error()})
		if err != nil {
			log.Println(err)
		}
		log.Println("room not found")
		return nil, fmt.Errorf("room not found: %v", err)
	}

	err = users.Users(ctx, bookingreq.UserID)
	if err != nil {
		_, err := s.n.Notification(ctx, &notificationproto.ProduceMessageRequest{Id: string(req.Userid), Message: err.Error()})
		if err != nil {
			log.Println(err)
		}
		log.Println("user not found")
		return nil, fmt.Errorf("user not found: %v", err)
	}

	res, err := s.service.Createbooking(bookingreq)
	if err != nil {
		log.Println("error creat booking")
		return nil, fmt.Errorf("error creat booking: %v", err)
	}

	// _, err = s.n.Notification(ctx, &notificationproto.ProduceMessageRequest{Id: string(req.Userid), Message: res. })
	// if err != nil {
	// 	log.Println(err)
	// }
	return &bookingproto.BookingResponse{
		BookingId:    res.BookingID,
		UserId:       res.UserID,
		RoomId:       res.RoomID,
		HotelId:      res.HotelID,
		Roomtype:     res.RoomType,
		CheckInDate:  timestamppb.New(res.CheckInDate),
		CheckOutDate: timestamppb.New(res.CheckOutDate),
		TotalAmount:  res.TotalAmount,
		Status:       res.Status,
	}, nil
}

func (s *Service) GetbyIdBooking(ctx context.Context, req *bookingproto.GetRequest) (*bookingproto.BookingResponse, error) {
	var bookingreq booking.GetRequest
	bookingreq.BookingID = req.BookingId
	res, err := s.service.Getbyidbooking(bookingreq)
	if err != nil {
		log.Println("error get by id booking")
		return nil, fmt.Errorf("error get by id booking: %v", err)
	}

	return &bookingproto.BookingResponse{
		BookingId:    res.BookingID,
		UserId:       res.UserID,
		HotelId:      res.HotelID,
		RoomId:       res.RoomID,
		Roomtype:     res.RoomType,
		CheckInDate:  timestamppb.New(res.CheckInDate),
		CheckOutDate: timestamppb.New(res.CheckOutDate),
		TotalAmount:  res.TotalAmount,
		Status:       res.Status,
	}, nil
}

func (s *Service) UpdateBooking(ctx context.Context, req *bookingproto.UpdateRequest) (*bookingproto.BookingResponse, error) {
	var bookingreq booking.UpdateRequest

	bookingreq.BookingID = req.BookingId
	bookingreq.CheckInDate = req.CheckInDate.AsTime()
	bookingreq.CheckOutDate = req.CheckOutDate.AsTime()
	bookingreq.RoomID = req.RoomId
	bookingreq.RoomType = req.Roomtype
	bookingreq.Status = req.Status

	err := hotels.Rooms(ctx, bookingreq.RoomID)
	if err != nil {
		log.Println("room not found")
		return nil, fmt.Errorf("room not found: %v", err)
	}

	res, err := s.service.Updatebooking(bookingreq)
	if err != nil {
		log.Println("error update booking")
		return nil, fmt.Errorf("error update booking: %v", err)
	}

	_, err = s.n.Notification(ctx, &notificationproto.ProduceMessageRequest{Id: string(userid), Message: "Room updated"})
	if err != nil {
		log.Println(err)
	}

	return &bookingproto.BookingResponse{
		BookingId:    res.BookingID,
		UserId:       res.UserID,
		HotelId:      res.HotelID,
		RoomId:       res.RoomID,
		Roomtype:     res.RoomType,
		CheckInDate:  timestamppb.New(res.CheckInDate),
		CheckOutDate: timestamppb.New(res.CheckOutDate),
		TotalAmount:  res.TotalAmount,
		Status:       res.Status,
	}, nil
}

func (s *Service) DeleteBooking(ctx context.Context, req *bookingproto.GetRequest) (*bookingproto.DeleteResponse, error) {
	var bookingreq booking.GetRequest
	bookingreq.BookingID = req.BookingId

	res, err := s.service.Deletebooking(bookingreq)
	if err != nil {
		log.Println("error delete booking")
		return nil, fmt.Errorf("error delete booking: %v", err)
	}
	booking, err := s.GetbyIdBooking(ctx, req)
	if err != nil {
		log.Println(err)
	}
	room, err := s.h.GetbyIdRoom(ctx, &hotelproto.RoomResponse{RoomId: booking.RoomId})
	if err != nil {
		log.Println(err)
	}
	_, err = s.h.UpdateRoom(ctx, &hotelproto.Room{Availability: true, RoomId: room.RoomId,
		HotelId:       room.HotelId,
		RoomType:      room.RoomType,
		PricePerNight: room.PricePerNight})
	if err != nil {
		log.Println(err)
	}
	_, err = s.n.Notification(ctx, &notificationproto.ProduceMessageRequest{Id: string(userid), Message: "Room Deleted"})
	if err != nil {
		log.Println(err)
	}

	return &bookingproto.DeleteResponse{Message: "Booking deleted", BookingId: res.BookingID}, nil
}

func (s *Service) CreateWaiting(ctx context.Context, req *bookingproto.CreateWaitingList) (*bookingproto.WaitingResponse, error) {
	var bookingreq booking.CreateWaitingList
	bookingreq.UserID = req.UserId
	bookingreq.UserEmail = req.UserEmail
	bookingreq.RoomType = req.RoomType
	bookingreq.HotelID = req.HotelId
	bookingreq.CheckInDate = req.CheckInDate.AsTime()
	bookingreq.CheckOutDate = req.CheckOutDate.AsTime()
	err := s.service.CreateWaiting(bookingreq)
	if err != nil {
		log.Println("error waiting created")
		return nil, fmt.Errorf("error waiting created: %v", err)
	}

	_, err = s.n.Notification(ctx, &notificationproto.ProduceMessageRequest{Id: string(userid), Message: "waiting user list"})
	if err != nil {
		log.Println(err)
	}
	return &bookingproto.WaitingResponse{Message: "waiting created"}, nil
}

func (s *Service) GetWaitingList(ctx context.Context, req *bookingproto.GetWaitingRequest) (*bookingproto.GetWaitingResponse, error) {
	res, err := s.service.GetbyIdwaitingList(req.Id)
	if err != nil {
		log.Println("erro get waiting")
		return nil, fmt.Errorf("error get waiting: %v", err)
	}

	return &bookingproto.GetWaitingResponse{
		Id:           res.ID,
		UserId:       res.UserID,
		UserEmail:    res.UserEmail,
		HotelId:      res.HotelID,
		RoomType:     res.RoomType,
		CheckInDate:  timestamppb.New(res.CheckInDate),
		CheckOutDate: timestamppb.New(res.CheckOutDate),
		Status:       res.Status,
	}, nil
}

func (s *Service) GetAllWaiting(ctx context.Context, _ *bookingproto.Empty) (*bookingproto.WaitingList, error) {
	res, err := s.service.GetAllWaiting()
	if err != nil {
		log.Println("error get all waiting")
		return nil, fmt.Errorf("error get all waiting: %v", err)
	}

	var users []*bookingproto.GetWaitingResponse
	for _, user := range res.Users {
		users = append(users, &bookingproto.GetWaitingResponse{
			UserId:       user.UserID,
			Id:           user.ID,
			UserEmail:    user.UserEmail,
			RoomType:     user.RoomType,
			HotelId:      user.HotelID,
			CheckInDate:  timestamppb.New(user.CheckInDate),
			CheckOutDate: timestamppb.New(user.CheckOutDate),
			Status:       user.Status,
		})
	}

	return &bookingproto.WaitingList{Users: users}, nil
}

func (s *Service) UpdateWaiting(ctx context.Context, req *bookingproto.UpdateWaitingListRequest) (*bookingproto.WaitingResponse, error) {
	var bookingreq booking.UpdateWaitingListRequest
	bookingreq.ID = req.Id
	bookingreq.UserID = req.UserId
	bookingreq.RoomType = req.RoomType
	bookingreq.HotelID = req.HotelId
	bookingreq.CheckInDate = req.CheckInDate.AsTime()
	bookingreq.CheckOutDate = req.CheckOutDate.AsTime()
	err := s.service.UpdateWaiting(bookingreq)
	if err != nil {
		log.Println("error update waiting")
		return nil, fmt.Errorf("error update waiting: %v", err)
	}

	return &bookingproto.WaitingResponse{Message: "Updated"}, nil
}

func (s *Service) DeleteWaiting(ctx context.Context, req *bookingproto.GetWaitingRequest) (*bookingproto.WaitingResponse, error) {
	var bookingreq booking.GetWaitingRequest
	bookingreq.ID = req.Id
	err := s.service.DeleteWaiting(bookingreq)
	if err != nil {
		log.Println("error deleted waiting")
		return nil, fmt.Errorf("error deleted waiting: %v", err)
	}

	return &bookingproto.WaitingResponse{Message: "Deleted"}, nil
}
