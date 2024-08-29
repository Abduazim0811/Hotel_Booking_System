package repository

import "booking_service/internal/entity/booking"

type BookingRepository interface {
	Create(req booking.BookingRequest) (*booking.BookingResponse, error)
	GetById(req booking.GetRequest) (*booking.BookingResponse, error)
	Update(req booking.UpdateRequest) (*booking.BookingResponse, error)
	Delete(req booking.GetRequest) (*booking.DeleteResponse, error)
	GetByUserId(req booking.GetUsersRequest) ([]*booking.BookingResponse, error)

	AddWaiting(req booking.CreateWaitingList) error
	GetbyIdwaitingList(id string) (*booking.GetWaitingResponse, error)
	GetAllWaiting() (*booking.WaitingList, error)
	UpdateWaiting(req booking.UpdateWaitingListRequest) error
	DeleteWaiting(req booking.GetWaitingRequest)error
}
