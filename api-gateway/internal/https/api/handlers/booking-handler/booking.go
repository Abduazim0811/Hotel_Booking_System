package bookinghandler

import (
	"api-gateway/internal/entity/booking"
	"api-gateway/internal/protos/bookingproto"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookingHandler struct {
	ClientBooking bookingproto.BookingServiceClient
}

// CreateBooking godoc
// @Summary Create a new booking
// @Description Create a new Booking
// @Tags Booking
// @Accept json
// @Produce json
// @Param hotel body booking.BookingRequest true "Booking request body"
// @Success 200 {object} booking.BookingResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /bookings [post]
func (b *BookingHandler) Createbooking(c *gin.Context) {
	var req booking.BookingRequest
	var bookingreq bookingproto.BookingRequest

	var bookingres booking.BookingResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookingreq.Hotelid = req.HotelID
	bookingreq.RoomId = req.RoomID
	bookingreq.CheckInDate = timestamppb.New(req.CheckInDate)
	bookingreq.CheckOutDate = timestamppb.New(req.CheckOutDate)
	bookingreq.Roomtype = req.RoomType
	bookingreq.TotalAmount = req.TotalAmount
	bookingreq.Userid = req.UserID

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := b.ClientBooking.CreateBooking(ctx, &bookingreq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bookingres.BookingID = res.BookingId
	bookingres.HotelID = res.HotelId
	bookingres.RoomID = res.RoomId
	bookingres.RoomType = res.Roomtype
	bookingres.Status = res.Status
	bookingres.CheckInDate = res.CheckInDate.AsTime()
	bookingres.CheckOutDate = res.CheckOutDate.AsTime()
	bookingres.UserID = res.UserId
	bookingres.TotalAmount = res.TotalAmount

	c.JSON(http.StatusOK, bookingres)
}

// GetBookingById godoc
// @Summary Get booking by ID
// @Description Retrieve a specific booking by its ID
// @Tags Booking
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} booking.BookingResponse
// @Failure 500 {object} string
// @Security Bearer
// @Router /bookings/{id} [get]
func (b *BookingHandler) GetbyidBooking(c *gin.Context) {
	id := c.Param("id")
	var req bookingproto.GetRequest

	req.BookingId = id
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := b.ClientBooking.GetbyIdBooking(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var bookingres booking.BookingResponse
	bookingres.BookingID = res.BookingId
	bookingres.HotelID = res.HotelId
	bookingres.RoomID = res.RoomId
	bookingres.RoomType = res.Roomtype
	bookingres.Status = res.Status
	bookingres.CheckInDate = res.CheckInDate.AsTime()
	bookingres.CheckOutDate = res.CheckOutDate.AsTime()
	bookingres.UserID = res.UserId
	bookingres.TotalAmount = res.TotalAmount

	c.JSON(http.StatusOK, bookingres)
}

// GetUserBookings godoc
// @Summary Get bookings for a user
// @Description Retrieve all bookings made by a specific user
// @Tags Booking
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} booking.GetUsersResponse
// @Failure 500 {object} string
// @Security Bearer
// @Router /users/{id}/bookings [get]
func (b *BookingHandler) GetUserbyIdBooking(c *gin.Context) {
	id := c.Param("id")
	userid, _ := strconv.Atoi(id)
	var req bookingproto.GetUsersRequst
	var bookingres booking.GetUsersResponse

	req.UserId = int32(userid)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := b.ClientBooking.GetUsersBooking(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	bookingres.BookingID = res.BookingId
	bookingres.HotelID = res.HotelId
	bookingres.RoomID = res.RoomId
	bookingres.RoomType = res.RoomType
	bookingres.Status = res.Status
	bookingres.CheckInDate = res.CheckInDate.AsTime()
	bookingres.CheckOutDate = res.CheckOutDate.AsTime()
	bookingres.TotalAmount = res.TotalAmount
	c.JSON(http.StatusOK, bookingres)
}

// UpdateBooking godoc
// @Summary Update a booking
// @Description Update the details of an existing booking
// @Tags Booking
// @Accept json
// @Produce json
// @Param id path string true "Booking ID"
// @Param booking body booking.UpdateRequest true "Updated booking details"
// @Success 200 {object} booking.BookingResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /bookings/{id} [put]
func (b *BookingHandler) Updatebooking(c *gin.Context) {
	id := c.Param("id")
	var req bookingproto.UpdateRequest
	var bookingreq booking.UpdateRequest
	bookingreq.BookingID = id
	if err := c.ShouldBindJSON(&bookingreq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.BookingId = bookingreq.BookingID
	req.CheckInDate = timestamppb.New(bookingreq.CheckInDate)
	req.CheckOutDate = timestamppb.New(bookingreq.CheckOutDate)
	req.RoomId = bookingreq.RoomID
	req.Roomtype = bookingreq.RoomType
	req.Status = bookingreq.Status
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var bookingres booking.BookingResponse
	res, err := b.ClientBooking.UpdateBooking(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	bookingres.BookingID = res.BookingId
	bookingres.HotelID = res.HotelId
	bookingres.RoomID = res.RoomId
	bookingres.RoomType = res.Roomtype
	bookingres.Status = res.Status
	bookingres.CheckInDate = res.CheckInDate.AsTime()
	bookingres.CheckOutDate = res.CheckOutDate.AsTime()
	bookingres.UserID = res.UserId
	bookingres.TotalAmount = res.TotalAmount

	c.JSON(http.StatusOK, bookingres)
}

// DeleteBooking godoc
// @Summary Delete a booking
// @Description Remove a booking from the system
// @Tags Booking
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} booking.DeleteResponse
// @Failure 500 {object} string
// @Security Bearer
// @Router /bookings/{id} [delete]
func (b *BookingHandler) Deletebooking(c *gin.Context) {
	id := c.Param("id")
	var req bookingproto.GetRequest
	req.BookingId = id

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var bookingres booking.DeleteResponse
	res, err := b.ClientBooking.DeleteBooking(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	bookingres.BookingID = res.BookingId
	bookingres.Message = res.Message

	c.JSON(http.StatusOK, bookingres)
}

// CreateWaitingList godoc
// @Summary Add a booking to the waiting list
// @Description Add a new booking to the waiting list
// @Tags WaitingList
// @Accept json
// @Produce json
// @Param waitinglist body booking.CreateWaitingList true "Waiting list request body"
// @Success 200 {object} booking.WaitingResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /waitinglist [post]
func (b *BookingHandler) CreateWaiting(c *gin.Context) {
	var req bookingproto.CreateWaitingList
	var bookingreq booking.CreateWaitingList
	if err := c.ShouldBindJSON(&bookingreq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.CheckInDate = timestamppb.New(bookingreq.CheckInDate)
	req.CheckOutDate = timestamppb.New(bookingreq.CheckOutDate)
	req.HotelId = bookingreq.HotelID
	req.RoomType = bookingreq.RoomType
	req.UserEmail = bookingreq.UserEmail
	req.UserId = bookingreq.UserID
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var bookingres booking.WaitingResponse
	res, err := b.ClientBooking.CreateWaiting(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bookingres.Message = res.Message
	c.JSON(http.StatusOK, bookingres)
}

// GetWaitingListById godoc
// @Summary Get waiting list by ID
// @Description Retrieve a specific entry from the waiting list by its ID
// @Tags WaitingList
// @Produce json
// @Param id path string true "Waiting list ID"
// @Success 200 {object} booking.GetWaitingResponse
// @Failure 500 {object} string
// @Security Bearer
// @Router /waitinglist/{id} [get]
func (b *BookingHandler) GetWaitinglist(c *gin.Context) {
	id := c.Param("id")
	var req bookingproto.GetWaitingRequest
	var bookingres booking.GetWaitingResponse
	req.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := b.ClientBooking.GetWaitingList(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	bookingres.UserID = res.UserId
	bookingres.UserEmail = res.UserEmail
	bookingres.Status = res.Status
	bookingres.RoomType = res.RoomType
	bookingres.ID = res.Id
	bookingres.HotelID = res.HotelId
	bookingres.CheckInDate = res.CheckInDate.AsTime()
	bookingres.CheckOutDate = res.CheckOutDate.AsTime()

	c.JSON(http.StatusOK, bookingres)
}


func (b *BookingHandler) GetAllWaiting(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := b.ClientBooking.GetAllWaiting(ctx, &bookingproto.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateWaitingList godoc
// @Summary Update a waiting list entry
// @Description Update the details of an entry in the waiting list
// @Tags WaitingList
// @Accept json
// @Produce json
// @Param id path string true "Waiting list ID"
// @Param waitinglist body booking.UpdateWaitingListRequest true "Updated waiting list details"
// @Success 200 {object} booking.WaitingResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /waitinglist/{id} [put]
func (b *BookingHandler) Updatewaiting(c *gin.Context) {
	id := c.Param("id")
	var req bookingproto.UpdateWaitingListRequest
	var bookingreq booking.UpdateWaitingListRequest
	bookingreq.ID = id
	if err := c.ShouldBindJSON(&bookingreq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.CheckInDate = timestamppb.New(bookingreq.CheckInDate)
	req.CheckOutDate = timestamppb.New(bookingreq.CheckOutDate)
	req.HotelId = bookingreq.HotelID
	req.Id = bookingreq.ID
	req.RoomType = bookingreq.RoomType
	req.UserId = bookingreq.UserID
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var bookingres booking.WaitingResponse
	res, err := b.ClientBooking.UpdateWaiting(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	bookingres.Message = res.Message

	c.JSON(http.StatusOK, bookingres)
}

// DeleteWaitingList godoc
// @Summary Delete a waiting list entry
// @Description Remove an entry from the waiting list
// @Tags WaitingList
// @Produce json
// @Param id path string true "Waiting list ID"
// @Success 200 {object} booking.WaitingResponse
// @Failure 500 {object} string
// @Security Bearer
// @Router /waitinglist/{id} [delete]
func (b *BookingHandler) Deletewaiting(c *gin.Context) {
	id := c.Param("id")
	var req bookingproto.GetWaitingRequest
	req.Id = id
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var bookingres booking.WaitingResponse
	res, err := b.ClientBooking.DeleteWaiting(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	bookingres.Message =res.Message

	c.JSON(http.StatusOK, bookingres)
}
