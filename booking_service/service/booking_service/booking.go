package bookingservice

import (
	"booking_service/internal/clients/hotels"
	"booking_service/internal/clients/users"
	"booking_service/internal/entity/booking"
	"booking_service/internal/service"
	"booking_service/protos/bookingproto"
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	bookingproto.UnimplementedBookingServiceServer
	service *service.BookingService
}

func NewGrpcService(s *service.BookingService) *Service {
	return &Service{service: s}
}

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
		log.Println("hotel not found")
		return nil, fmt.Errorf("hotel not found: %v", err)
	}

	err = hotels.Rooms(ctx, bookingreq.RoomID)
	if err != nil {
		log.Println("room not found")
		return nil, fmt.Errorf("room not found: %v", err)
	}

	err = users.Users(ctx, bookingreq.UserID)
	if err != nil {
		log.Println("user not found")
		return nil, fmt.Errorf("user not found: %v", err)
	}

	res, err := s.service.Createbooking(bookingreq)
	if err != nil {
		log.Println("error creat booking")
		return nil, fmt.Errorf("error creat booking: %v", err)
	}
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

func (s *Service) GetbyIdBooking(ctx context.Context, req *bookingproto.GetRequest) (*bookingproto.BookingResponse, error){
	var bookingreq booking.GetRequest
	bookingreq.BookingID = req.BookingId
	res, err := s.service.Getbyidbooking(bookingreq)
	if err != nil {
		log.Println("error get by id booking")
		return nil, fmt.Errorf("error get by id booking: %v", err)
	}

	return &bookingproto.BookingResponse{
		BookingId:  res.BookingID,
		UserId:  res.UserID,
		HotelId: res.HotelID,
		RoomId:  res.RoomID,
		Roomtype: res.RoomType,
		CheckInDate:  timestamppb.New(res.CheckInDate),
		CheckOutDate: timestamppb.New(res.CheckOutDate),
		TotalAmount:  res.TotalAmount,
		Status:       res.Status,
	}, nil
}

func (s *Service) UpdateBooking(ctx context.Context,req *bookingproto.UpdateRequest) (*bookingproto.BookingResponse, error){
	var bookingreq booking.UpdateRequest

	bookingreq.BookingID = req.BookingId
	bookingreq.CheckInDate = req.CheckInDate.AsTime()
	bookingreq.CheckOutDate = req.CheckOutDate.AsTime()
	bookingreq.RoomID = req.RoomId
	bookingreq.RoomType = req.Roomtype
	bookingreq.Status = req.Status
	res, err := s.service.Updatebooking(bookingreq)
	if err != nil {
		log.Println("error update booking")
		return nil, fmt.Errorf("error update booking: %v", err)
	}

	return &bookingproto.BookingResponse{
		BookingId: res.BookingID,
		UserId:  res.UserID,
		HotelId: res.HotelID,
		RoomId:  res.RoomID,
		Roomtype: res.RoomType,
		CheckInDate:  timestamppb.New(res.CheckInDate),
		CheckOutDate: timestamppb.New(res.CheckOutDate),
		TotalAmount:  res.TotalAmount,
		Status:       res.Status,
	}, nil
}

func (s *Service) DeleteBooking(ctx context.Context, req *bookingproto.GetRequest) (*bookingproto.DeleteResponse, error){
	var bookingreq booking.GetRequest
	bookingreq.BookingID = req.BookingId

	res, err := s.service.Deletebooking(bookingreq)
	if err != nil {
		log.Println("error delete booking")
		return nil, fmt.Errorf("error delete booking: %v", err)
	}

	return &bookingproto.DeleteResponse{Message: "Booking deleted", BookingId: res.BookingID}, nil
}
