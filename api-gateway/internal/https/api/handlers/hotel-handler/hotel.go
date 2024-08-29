package hotelhandler

import (
	"api-gateway/internal/protos/hotelproto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HotelHandler struct {
	ClientHotel hotelproto.HotelServiceClient
}

// CreateHotel godoc
// @Summary Create a new hotel
// @Description Create a new hotel
// @Tags hotel
// @Accept json
// @Produce json
// @Param hotel body hotelproto.HotelRequest true "Hotel request body"
// @Success 200 {object} hotelproto.HotelResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /hotels [post]
func (h *HotelHandler) Createhotel(c *gin.Context) {
	var req hotelproto.HotelRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := h.ClientHotel.CreateHotel(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetByIdHotel godoc
// @Summary Get hotel by ID
// @Description Get hotel by ID
// @Tags hotel
// @Accept json
// @Produce json
// @Param id path string true "Hotel ID"
// @Success 200 {object} hotelproto.Hotel
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /hotels/{id} [get]
func (h *HotelHandler) GetByIdHotel(c *gin.Context) {
	id := c.Param("id")
	var req hotelproto.HotelResponse
	req.HotelId = id

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := h.ClientHotel.GetbyIdHotel(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllHotels godoc
// @Summary Get all hotels
// @Description Get all hotels
// @Tags hotel
// @Produce json
// @Success 200 {object} hotelproto.ListHotels
// @Failure 500 {object} string
// @Security Bearer
// @Router /hotels [get]
func (h *HotelHandler) GetAllHotel(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := h.ClientHotel.GetAllHotels(ctx, &hotelproto.HotelEmpty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateHotel godoc
// @Summary Update hotel by ID
// @Description Update hotel by ID
// @Tags hotel
// @Accept json
// @Produce json
// @Param id path string true "Hotel ID"
// @Param hotel body hotelproto.Hotel true "Hotel update request body"
// @Success 200 {object} hotelproto.HotelRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /hotels/{id} [put]
func (h *HotelHandler) UpdateHotels(c *gin.Context) {
	id := c.Param("id")
	var req hotelproto.Hotel
	req.HotelId = id

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := h.ClientHotel.UpdateHotel(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteHotel godoc
// @Summary Delete hotel by ID
// @Description Delete hotel by ID
// @Tags hotel
// @Param id path string true "Hotel ID"
// @Success 200 {object} hotelproto.HotelRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /hotels/{id} [delete]
func (h *HotelHandler) DeleteHotels(c *gin.Context) {
	id := c.Param("id")
	var req hotelproto.HotelResponse
	req.HotelId = id

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := h.ClientHotel.DeleteHotel(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateRoom godoc
// @Summary Create a new room
// @Description Create a new room
// @Tags room
// @Accept json
// @Produce json
// @Param room body hotelproto.RoomRequest true "Room request body"
// @Success 200 {object} hotelproto.RoomResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /rooms [post]
func (h *HotelHandler) CreateRooms(c *gin.Context) {
	var req hotelproto.RoomRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := h.ClientHotel.CreateRoom(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetByIdRoom godoc
// @Summary Get room by ID
// @Description Get room by ID
// @Tags room
// @Produce json
// @Param id path string true "Room ID"
// @Success 200 {object} hotelproto.Room
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /rooms/{id} [get]
func (h *HotelHandler) GetByIDRoom(c *gin.Context) {
	id := c.Param("id")
	var req hotelproto.RoomResponse
	req.RoomId = id

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := h.ClientHotel.GetbyIdRoom(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllRooms godoc
// @Summary Get all rooms
// @Description Get all rooms
// @Tags room
// @Produce json
// @Success 200 {object} hotelproto.ListRooms
// @Failure 500 {object} string
// @Security Bearer
// @Router /rooms [get]
func (h *HotelHandler) GetAllRoom(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := h.ClientHotel.GetAllRooms(ctx, &hotelproto.HotelEmpty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateRoom godoc
// @Summary Update room by ID
// @Description Update room by ID
// @Tags room
// @Accept json
// @Produce json
// @Param id path string true "Room ID"
// @Param room body hotelproto.Room true "Room update request body"
// @Success 200 {object} hotelproto.RoomRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /rooms/{id} [put]
func (h *HotelHandler) UpdateRooms(c *gin.Context) {
	id := c.Param("id")
	var req hotelproto.Room
	req.RoomId = id

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := h.ClientHotel.UpdateRoom(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteRoom godoc
// @Summary Delete room by ID
// @Description Delete room by ID
// @Tags room
// @Param id path string true "Room ID"
// @Success 200 {object} hotelproto.RoomRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /rooms/{id} [delete]
func (h *HotelHandler) DeleteRooms(c *gin.Context) {
	id := c.Param("id")
	var req hotelproto.RoomResponse
	req.RoomId = id

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := h.ClientHotel.DeleteRoom(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
