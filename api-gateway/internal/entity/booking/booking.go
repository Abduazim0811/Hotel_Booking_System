package booking

import "time"

type BookingRequest struct {
	UserID       int32     `json:"userid" bson:"userid"`
	HotelID      string    `json:"hotelid" bson:"hotelid"`
	RoomID       string    `json:"roomId" bson:"roomId"`
	RoomType     string    `json:"roomtype" bson:"roomtype"`
	CheckInDate  time.Time `json:"checkInDate" bson:"checkInDate"`
	CheckOutDate time.Time `json:"checkOutDate" bson:"checkOutDate"`
	TotalAmount  int32     `json:"totalAmount" bson:"totalAmount"`
}

type BookingResponse struct {
	BookingID    string    `json:"bookingId" bson:"_id,omitempty"`
	UserID       int32     `json:"userId" bson:"userid"`
	HotelID      string    `json:"hotelId" bson:"hotelid"`
	RoomID       string    `json:"roomId" bson:"roomId"`
	RoomType     string    `json:"roomtype" bson:"roomtype"`
	CheckInDate  time.Time `json:"checkInDate" bson:"checkInDate"`
	CheckOutDate time.Time `json:"checkOutDate" bson:"checkOutDate"`
	TotalAmount  int32     `json:"totalAmount" bson:"totalAmount"`
	Status       string    `json:"status" bson:"status"`
}

type GetRequest struct {
	BookingID string `json:"bookingId" bson:"_id,omitempty"`
}

type UpdateRequest struct {
	BookingID    string    `json:"bookingId" bson:"_id,omitempty"`
	RoomID       string    `json:"roomId" bson:"roomId"`
	RoomType     string    `json:"roomtype" bson:"roomtype"`
	CheckInDate  time.Time `json:"checkInDate" bson:"checkInDate"`
	CheckOutDate time.Time `json:"checkOutDate" bson:"checkOutDate"`
	Status       string    `json:"status" bson:"status"`
}

type DeleteResponse struct {
	Message   string `json:"message" bson:"message"`
	BookingID string `json:"bookingId" bson:"bookingId"`
}

type GetUsersResponse struct {
	BookingID    string    `json:"bookingId" bson:"_id,omitempty"`
	HotelID      string    `json:"hotelId" bson:"hotelId"`
	RoomID       string    `json:"roomId" bson:"roomId"`
	RoomType     string    `json:"roomType" bson:"roomType"`
	CheckInDate  time.Time `json:"checkInDate" bson:"checkInDate"`
	CheckOutDate time.Time `json:"checkOutDate" bson:"checkOutDate"`
	TotalAmount  int32     `json:"totalAmount" bson:"totalAmount"`
	Status       string    `json:"status" bson:"status"`
}

type GetUsersRequest struct {
	UserID int32 `json:"userId" bson:"userid"`
}

type CreateWaitingList struct {
	UserID       int32     `json:"user_id" bson:"user_id"`
	UserEmail    string    `json:"user_email" bson:"user_email"`
	RoomType     string    `json:"room_type" bson:"room_type"`
	HotelID      string    `json:"hotel_id" bson:"hotel_id"`
	CheckInDate  time.Time `json:"check_in_date" bson:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date" bson:"check_out_date"`
}

type UpdateWaitingListRequest struct {
	UserID       int32     `json:"user_id" bson:"user_id"`
	ID           string    `json:"id" bson:"id"`
	RoomType     string    `json:"room_type" bson:"room_type"`
	HotelID      string    `json:"hotel_id" bson:"hotel_id"`
	CheckInDate  time.Time `json:"check_in_date" bson:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date" bson:"check_out_date"`
}

type GetWaitingResponse struct {
	UserID       int32     `json:"user_id" bson:"user_id"`
	UserEmail    string    `json:"user_email" bson:"user_email"`
	RoomType     string    `json:"room_type" bson:"room_type"`
	HotelID      string    `json:"hotel_id" bson:"hotel_id"`
	CheckInDate  time.Time `json:"check_in_date" bson:"check_in_date"`
	Status       string    `json:"status" bson:"status"`
	ID           string    `json:"id" bson:"id"`
	CheckOutDate time.Time `json:"check_out_date" bson:"check_out_date"`
}

type WaitingList struct {
	Users []GetWaitingResponse `json:"users" bson:"users"`
}

type WaitingResponse struct {
	Message string `json:"message" bson:"message"`
}

type GetWaitingRequest struct {
	ID string `json:"id" bson:"id"`
}

type Empty struct{}
