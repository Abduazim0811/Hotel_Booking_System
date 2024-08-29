package service

import (
	"hotel_service/internal/entity/hotel"
	"hotel_service/internal/infrastructura/repository"
)

type HotelService struct{
	repo repository.HotelRepository
}

func NewHotelService(repo repository.HotelRepository) *HotelService{
	return &HotelService{repo: repo}
}

func (h *HotelService) Createhotel(req hotel.HotelRequest)(*hotel.HotelResponse, error){
	return h.repo.AddHotel(req)
}

func (h *HotelService) Getbyidhotel(req string)(*hotel.Hotel, error){
	return h.repo.GetbyId(req)
}

func (h *HotelService) Getallhotel()(*[]hotel.Hotel, error){
	return h.repo.GetAll()
}

func (h *HotelService) Updatehotel(req hotel.Hotel)error{
	return h.repo.UpdateHotel(req)
}

func (h *HotelService) Deletehotel(id string)error{
	return h.repo.DeleteHotel(id)
}

func (h *HotelService) Createroom(req hotel.RoomRequest)(*hotel.RoomResponse, error){
	return h.repo.CreateRoom(req)
}

func(h *HotelService) Getbyidroom(req hotel.RoomResponse)(*hotel.Room, error){
	return h.repo.GetRoomById(req)
}


func (h *HotelService) Getallroom(id string)(*[]hotel.Room, error){
	return h.repo.GetAllRooms(id)
}

func (h *HotelService) Updateroom(req hotel.Room)error{
	return h.repo.UpdateRoom(req)
} 

func (h *HotelService) Deleteroom(id string) error{
	return h.repo.DeleteRoom(id)
}