package userhandler

import (
	"api-gateway/internal/protos/userproto"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Userhandler struct {
	Clientuser userproto.UserServiceClient
	ctx context.Context
}

// @title Artisan Connect
// @version 1.0
// @description This is a sample server for a restaurant reservation system.
// @securityDefinitions.apikey Bearer
// @in 				header
// @name Authorization
// @description Enter the token in the format `Bearer {token}`
// @host localhost:7777
// @BasePath /


// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body userproto.UserRequest true "User request body"
// @Success 200 {object} userproto.UpdateUserRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /register [post]
func (u *Userhandler) CreateUser(c *gin.Context) {
	var req userproto.UserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := u.Clientuser.Register(u.ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// VerfiyCode godoc
// @Summary VerifyCode a user
// @Description Login a user and get a JWT token
// @Tags user
// @Accept json
// @Produce json
// @Param login body userproto.LoginRequest true "Login request body"
// @Success 200 {object} userproto.UserResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /verifycode [post]
func (u *Userhandler) VerifyCode(c *gin.Context) {
	var req userproto.Req
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := u.Clientuser.VerifyCode(u.ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (u *Userhandler) Login(c *gin.Context){
	var req userproto.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := u.Clientuser.Login(u.ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (u *Userhandler) GetbyIdUser(c *gin.Context){
	id := c.Param("id")
	var req userproto.GetUserRequest
	userid, _ := strconv.Atoi(id)
	req.Id =int32(userid)

	res, err :=u.Clientuser.GetByIdUser(u.ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (u *Userhandler) GetAllUsers(c *gin.Context){
	res, err := u.Clientuser.GetUsers(u.ctx, &userproto.UserEmpty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (u *Userhandler) UpdateUsers(c *gin.Context){
	id := c.Param("id")
	userid, _ := strconv.Atoi(id)
	var req userproto.UpdateUserReq
	req.Id = int32(userid)
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := u.Clientuser.UpdateUser(u.ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (u *Userhandler) UpdatePasswordUsers(c *gin.Context){
	id := c.Param("id")
	userid, _ := strconv.Atoi(id)

	var req userproto.UpdatePasswordReq

	req.Id = int32(userid)

	res, err := u.Clientuser.UpdatePassword(u.ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (u *Userhandler) DeleteUsers(c *gin.Context){
	id := c.Param("id")
	userid, _ := strconv.Atoi(id)

	var req userproto.GetUserRequest
	req.Id = int32(userid)

	res, err := u.Clientuser.DeleteUser(u.ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)	
}
