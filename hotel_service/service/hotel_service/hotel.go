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