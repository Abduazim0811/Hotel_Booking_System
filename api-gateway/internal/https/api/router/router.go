package router

import (
	userclients "api-gateway/internal/clients/user_clients"
	userhandler "api-gateway/internal/https/api/handlers/user-handler"
	"api-gateway/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() {
	userclient := userclients.DialUserGrpc()
	userhandler := &userhandler.Userhandler{Clientuser: userclient}

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/register", userhandler.CreateUser)
	router.POST("/verifycode", userhandler.VerifyCode)
	router.POST("/login", userhandler.Login)
	router.GET("/users/:id", jwt.Protected(), userhandler.GetbyIdUser)
	router.GET("/users", userhandler.GetAllUsers)
	router.PUT("/users/:id", jwt.Protected(), userhandler.UpdateUsers)
	router.PUT("/users/password/:id", jwt.Protected(), userhandler.UpdatePasswordUsers)
	router.DELETE("/users/:id", jwt.Protected(), userhandler.DeleteUsers)

	router.Run(":7777")
}
