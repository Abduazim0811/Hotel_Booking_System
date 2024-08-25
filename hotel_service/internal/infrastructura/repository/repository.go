package repository

import "hotel_service/internal/entity/hotel"

type HotelRepository interface{
	AddHotel(req hotel.HotelRequest)(*hotel.HotelResponse, error)
	GetbyId(req hotel.HotelResponse) (*hotel.Hotel, error)
	GetAll()(*[]hotel.Hotel, error)
	UpdateHotel(req hotel.Hotel) error
	DeleteHotel(hotelID string) error
	CreateRoom(req hotel.RoomRequest) (*hotel.RoomResponse, error)
	GetRoomById(req hotel.RoomResponse) (*hotel.Room, error)
	GetAllRooms() (*[]hotel.Room, error)
	UpdateRoom(req hotel.Room) error
	DeleteRoom(roomID string) error
}