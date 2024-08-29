package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	bookingclients "api-gateway/internal/clients/booking_clients"
	hotelclients "api-gateway/internal/clients/hotel_clients"
	userclients "api-gateway/internal/clients/user_clients"
	bookinghandler "api-gateway/internal/https/api/handlers/booking-handler"
	hotelhandler "api-gateway/internal/https/api/handlers/hotel-handler"
	userhandler "api-gateway/internal/https/api/handlers/user-handler"
	"api-gateway/internal/pkg/jwt"
	middleware "api-gateway/internal/rate-limiting"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *http.Server {
	userclient := userclients.DialUserGrpc()
	userhandler := &userhandler.Userhandler{Clientuser: userclient}

	hotelclient := hotelclients.DialHotelGrpc()
	hotelhandler := &hotelhandler.HotelHandler{ClientHotel: hotelclient}

	bookingclient := bookingclients.DialBookingGrpc()
	bookinghandler := &bookinghandler.BookingHandler{ClientBooking: bookingclient}

	// Rate limiter o‘rnatish
	userRateLimiter := middleware.NewRateLimiter(2, time.Minute)
	hotelRateLimiter := middleware.NewRateLimiter(5, time.Minute)
	bookingRateLimiter := middleware.NewRateLimiter(3, time.Minute)

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// user-handler
	userRoutes := router.Group("/")
	userRoutes.Use(userRateLimiter.Limit())
	{
		userRoutes.POST("/register", userhandler.CreateUser)
		userRoutes.POST("/verifycode", userhandler.VerifyCode)
		userRoutes.POST("/login", userhandler.Login)
		userRoutes.GET("/users/:id", jwt.Protected(), userhandler.GetbyIdUser)
		userRoutes.GET("/users", userhandler.GetAllUsers)
		userRoutes.PUT("/users/:id", jwt.Protected(), userhandler.UpdateUsers)
		userRoutes.PUT("/users/password/:id", jwt.Protected(), userhandler.UpdatePasswordUsers)
		userRoutes.DELETE("/users/:id", jwt.Protected(), userhandler.DeleteUsers)
	}

	// hotel-handler
	hotelRoutes := router.Group("/")
	hotelRoutes.Use(hotelRateLimiter.Limit())
	{
		hotelRoutes.POST("/hotels", jwt.Protected(), hotelhandler.Createhotel)
		hotelRoutes.GET("/hotels/:id", jwt.Protected(), hotelhandler.GetByIdHotel)
		hotelRoutes.GET("/hotels", jwt.Protected(), hotelhandler.GetAllHotel)
		hotelRoutes.PUT("/hotels/:id", jwt.Protected(), hotelhandler.UpdateHotels)
		hotelRoutes.DELETE("/hotels/:id", jwt.Protected(), hotelhandler.DeleteHotels)

		hotelRoutes.POST("/rooms", jwt.Protected(), hotelhandler.CreateRooms)
		hotelRoutes.GET("/rooms/:id", jwt.Protected(), hotelhandler.GetByIDRoom)
		hotelRoutes.GET("/rooms", jwt.Protected(), hotelhandler.GetAllRoom)
		hotelRoutes.PUT("/rooms/:id", jwt.Protected(), hotelhandler.UpdateRooms)
		hotelRoutes.DELETE("/rooms/:id", jwt.Protected(), hotelhandler.DeleteRooms)
	}

	// booking-handler
	bookingRoutes := router.Group("/")
	bookingRoutes.Use(bookingRateLimiter.Limit())
	{
		bookingRoutes.POST("/bookings", jwt.Protected(), bookinghandler.Createbooking)
		bookingRoutes.GET("/bookings/:id", jwt.Protected(), bookinghandler.GetbyidBooking)
		bookingRoutes.GET("/users/:id/bookings", jwt.Protected(), bookinghandler.GetUserbyIdBooking)
		bookingRoutes.PUT("/bookings/:id", jwt.Protected(), bookinghandler.Updatebooking)
		bookingRoutes.DELETE("/bookings/:id", jwt.Protected(), bookinghandler.Deletebooking)
		bookingRoutes.POST("/waitinglist", jwt.Protected(), bookinghandler.CreateWaiting)
		bookingRoutes.GET("/waitinglist/:id", jwt.Protected(), bookinghandler.GetWaitinglist)
		bookingRoutes.GET("/waitinglist", jwt.Protected(), bookinghandler.GetAllWaiting)
		bookingRoutes.PUT("/waitinglist/:id", jwt.Protected(), bookinghandler.Updatewaiting)
		bookingRoutes.DELETE("/waitinglist/:id", jwt.Protected(), bookinghandler.Deletewaiting)
	}

	server := &http.Server{
		Addr:    "api_gateway:7777",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServeTLS("./internal/tls/items.pem", "./internal/tls/items-key.pem"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to run HTTPS server: %v", err)
		}
	}()

	GracefulShutdown(server, log.Default())

	return server
}

func GracefulShutdown(srv *http.Server, logger *log.Logger) {
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	<-shutdownCh
	logger.Println("Shutdown signal received, initiating graceful shutdown...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Printf("Server shutdown encountered an error: %v", err)
	} else {
		logger.Println("Server gracefully stopped")
	}

	select {
	case <-shutdownCtx.Done():
		if shutdownCtx.Err() == context.DeadlineExceeded {
			logger.Println("Shutdown deadline exceeded, forcing server to stop")
		}
	default:
		logger.Println("Shutdown completed within the timeout period")
	}
}
