package hotelservice

import (
	"context"
	"fmt"
	"hotel_service/hotelproto"
	"hotel_service/internal/entity/hotel"
	"hotel_service/internal/service"
	"log"
)

type Service struct {
	hotelproto.UnimplementedHotelServiceServer
	service *service.HotelService
}

func HotelGrpcService(service *service.HotelService) *Service {
	return &Service{service: service}
}

func (s *Service) CreateHotel(ctx context.Context, req *hotelproto.HotelRequest) (*hotelproto.HotelResponse, error) {
	var hotelreq hotel.HotelRequest
	hotelreq.Address = req.Address
	hotelreq.Location = req.Location
	hotelreq.Name = req.Name
	hotelreq.Rating = req.Rating

	res, err := s.service.Createhotel(hotelreq)
	if err != nil {
		return nil, fmt.Errorf("error creating hotel: %v", err)
	}

	return &hotelproto.HotelResponse{HotelId: res.HotelID}, nil
}

func (s *Service) GetbyIdHotel(ctx context.Context, req *hotelproto.HotelResponse) (*hotelproto.Hotel, error){
	var hotelreq hotel.HotelResponse
	hotelreq.HotelID = req.HotelId
	res, err := s.service.Getbyidhotel(hotelreq)
	if err != nil {
		log.Println("get by id hotel error")
		return nil, fmt.Errorf("error get by id hotel: %v", err)
	}

	return &hotelproto.Hotel{
		HotelId: res.HotelID,
		Location: res.Location,
		Address: res.Address,
		Name: res.Name,
		Rating: res.Rating,
	}, nil
}

func (s *Service) GetAllHotels(ctx context.Context, _ *hotelproto.Empty) (*hotelproto.ListHotels, error) {
	res, err := s.service.Getallhotel()
	if err != nil {
		log.Println("get all hotels error")
		return nil, fmt.Errorf("error get all hotels: %v", err)
	}

	var protoHotels []*hotelproto.Hotel
	for _, hotel := range *res {
		protoHotels = append(protoHotels, &hotelproto.Hotel{
			HotelId:  hotel.HotelID,
			Name:     hotel.Name,
			Location: hotel.Location,
			Rating:   hotel.Rating,
			Address:  hotel.Address,
		})
	}

	return &hotelproto.ListHotels{Hotel: protoHotels}, nil
}

func (s *Service) UpdateHotel(ctx context.Context,req *hotelproto.Hotel) (*hotelproto.HotelRes, error){
	var hotelreq hotel.Hotel

	hotelreq.HotelID = req.HotelId
	hotelreq.Address = req.Address
	hotelreq.Location = req.Location
	hotelreq.Name = req.Name
	hotelreq.Rating = req.Rating

	err := s.service.Updatehotel(hotelreq)
	if err != nil {
		log.Println("update hotel error")
		return nil, fmt.Errorf("update hotel error: %v", err)
	}

	return &hotelproto.HotelRes{Message: "Hotel updated"}, nil
}

func (s *Service) DeleteHotel(ctx context.Context,req *hotelproto.HotelResponse) (*hotelproto.HotelRes, error){
	err := s.service.Deletehotel(req.HotelId)
	if err != nil {
		log.Println("deleted hotel error")
		return nil, fmt.Errorf("deleted hotel error: %v", err)
	}

	return &hotelproto.HotelRes{Message: "Hotel Deleted"}, nil
}

func (s *Service) CreateRoom(ctx context.Context, req *hotelproto.RoomRequest)(*hotelproto.RoomResponse, error){
	var roomreq hotel.RoomRequest
	roomreq.Availability = req.Availability
	roomreq.PricePerNight = req.PricePerNight
	roomreq.RoomType = req.RoomType

	res, err := s.service.Createroom(roomreq)
	if err != nil {
		log.Println("room created error")
		return nil, fmt.Errorf("room created error: %v", err)
	}

	return &hotelproto.RoomResponse{RoomId: res.RoomID}, nil
}

func (s *Service) GetbyIdRoom(ctx context.Context,req *hotelproto.RoomResponse) (*hotelproto.Room, error){
	var roomreq hotel.RoomResponse
	roomreq.RoomID = req.RoomId
	res, err := s.service.Getbyidroom(roomreq)
	if err != nil {
		log.Println("room get by id error")
		return nil, fmt.Errorf("room get by id error: %v", err)
	}

	return &hotelproto.Room{
		RoomId: res.RoomID,
		RoomType: res.RoomType,
		Availability: res.Availability,
		PricePerNight: res.PricePerNight,
	}, nil
}

func (s *Service) GetAllRooms(ctx context.Context, _ *hotelproto.Empty) (*hotelproto.ListRooms, error){
	res, err := s.service.Getallroom()
	if err != nil {
		log.Println("get all room error")
		return nil, fmt.Errorf("get all room error: %v", err)
	}
	var protoroom []*hotelproto.Room
	for _, hotel := range *res{
		protoroom = append(protoroom, &hotelproto.Room{
			RoomId: hotel.RoomID,
			Availability: hotel.Availability,
			RoomType: hotel.RoomType,
			PricePerNight: hotel.PricePerNight,
		})
	}

	return &hotelproto.ListRooms{Room: protoroom}, nil
}

func (s *Service) UpdateRoom(ctx context.Context,req *hotelproto.Room) (*hotelproto.RoomRes, error){
	var roomreq hotel.Room
	roomreq.RoomID = req.RoomId
	roomreq.RoomType = req.RoomType
	roomreq.Availability = req.Availability
	roomreq.PricePerNight = req.PricePerNight

	err := s.service.Updateroom(roomreq)
	if err != nil {
		log.Println("update room error")
		return nil, fmt.Errorf("update room error: %v", err)
	}

	return &hotelproto.RoomRes{Message: "room updated"}, nil
}

func (s *Service) DeleteRoom(ctx context.Context, req *hotelproto.RoomResponse) (*hotelproto.RoomRes, error){
	err := s.service.Deleteroom(req.RoomId)
	if err != nil {
		log.Println("delete room error")
		return nil, fmt.Errorf("delete room error: %v", err)
	}

	return &hotelproto.RoomRes{Message: "deleted room"}, nil
}