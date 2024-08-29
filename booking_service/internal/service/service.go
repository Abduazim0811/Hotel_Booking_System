package service

import (
	"booking_service/internal/entity/booking"
	"booking_service/internal/infrastructura/repository"
)

type BookingService struct{
	repo repository.BookingRepository
}

func NewBookingService(r repository.BookingRepository) *BookingService{
	return &BookingService{repo: r}
}

func (b *BookingService) Createbooking(req booking.BookingRequest)(*booking.BookingResponse, error){
	return b.repo.Create(req)
}

func (b *BookingService) Getbyidbooking(req booking.GetRequest)(*booking.BookingResponse, error){
	return b.repo.GetById(req)
}

func (b *BookingService) Updatebooking(req booking.UpdateRequest)(*booking.BookingResponse, error){
	return b.repo.Update(req)
}

func (b *BookingService) Deletebooking(req booking.GetRequest)(*booking.DeleteResponse, error){
	return b.repo.Delete(req)
}

func (b *BookingService) Getbyidusers(req booking.GetUsersRequest)([]*booking.BookingResponse, error){
	return b.repo.GetByUserId(req)
}
func (b *BookingService) CreateWaiting(req booking.CreateWaitingList) error {
	return b.repo.AddWaiting(req)
}

func (b *BookingService) GetbyIdwaitingList(id string) (*booking.GetWaitingResponse, error) {
	return b.repo.GetbyIdwaitingList(id)
}

func (b *BookingService) GetAllWaiting() (*booking.WaitingList, error) {
	return b.repo.GetAllWaiting()
}

func (b *BookingService) UpdateWaiting(req booking.UpdateWaitingListRequest) error {
	return b.repo.UpdateWaiting(req)
}

func (b *BookingService) DeleteWaiting(req booking.GetWaitingRequest) error {
	return b.repo.DeleteWaiting(req)
}
