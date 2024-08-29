package userhandler

import (
	_ "api-gateway/docs"
	producer "api-gateway/internal/kafka"
	"api-gateway/internal/protos/userproto"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Userhandler struct {
	Clientuser userproto.UserServiceClient
}

// @title Hotel Booking System
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.Clientuser.Register(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	byted, err := json.Marshal(&req)
	if err != nil {
		log.Println(err)
	}
	if err := producer.Producer("create", byted); err != nil {
		log.Println(err)
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := u.Clientuser.VerifyCode(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Login godoc
// @Summary Login a user
// @Description Login a user and get a JWT token
// @Tags user
// @Accept json
// @Produce json
// @Param login body userproto.LoginRequest true "Login request body"
// @Success 200 {object} userproto.LoginResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /login [post]
func (u *Userhandler) Login(c *gin.Context) {
	var req userproto.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := u.Clientuser.Login(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetByIdUser godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "User ID"
// @Success 200 {object} userproto.User
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /users/{id} [get]
func (u *Userhandler) GetbyIdUser(c *gin.Context) {
	id := c.Param("id")
	var req userproto.GetUserRequest
	userid, _ := strconv.Atoi(id)
	req.Id = int32(userid)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := u.Clientuser.GetByIdUser(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} userproto.ListUser
// @Failure 500 {object} string
// @Security Bearer
// @Router /users [get]
func (u *Userhandler) GetAllUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := u.Clientuser.GetUsers(ctx, &userproto.UserEmpty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateUsers godoc
// @Summary Update a user
// @Description Update a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body userproto.UpdateUserReq true "User request body"
// @Success 200 {object} userproto.UpdateUserRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /users [put]
func (u *Userhandler) UpdateUsers(c *gin.Context) {
	id := c.Param("id")
	userid, _ := strconv.Atoi(id)
	var req userproto.UpdateUserReq
	req.Id = int32(userid)
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.Clientuser.UpdateUser(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	byted, err := json.Marshal(&req)
	if err != nil {
		log.Println(err)
	}
	if err := producer.Producer("update", byted); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, res)
}

// UpdatePassword godoc
// @Summary Update a user's password
// @Description Update a user's password
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param user body userproto.UpdatePasswordReq true "Password update request body"
// @Success 200 {object} userproto.UpdateUserRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /users/password/{id} [put]
func (u *Userhandler) UpdatePasswordUsers(c *gin.Context) {
	id := c.Param("id")
	userid, _ := strconv.Atoi(id)

	var req userproto.UpdatePasswordReq

	req.Id = int32(userid)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := u.Clientuser.UpdatePassword(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body userproto.GetUserRequest true "Delete user request body"
// @Success 200 {object} userproto.UpdateUserRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /users/{id} [delete]
func (u *Userhandler) DeleteUsers(c *gin.Context) {
	id := c.Param("id")
	userid, _ := strconv.Atoi(id)

	var req userproto.GetUserRequest
	req.Id = int32(userid)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := u.Clientuser.DeleteUser(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	byted, err := json.Marshal(&req)
	if err != nil {
		log.Println(err)
	}
	if err := producer.Producer("delete", byted); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, res)
}
