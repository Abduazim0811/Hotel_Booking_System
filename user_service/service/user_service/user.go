package userservice

import (
	"context"
	"fmt"
	"log"
	"user_service/internal/entity/user"
	pkg "user_service/internal/pkg/email"
	"user_service/internal/service"
	"user_service/userproto"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

type Service struct {
	*userproto.UnimplementedUserServiceServer
	service *service.UserService
}

func NewGrpcService(service *service.UserService) *Service {
	return &Service{service: service}
}

func (s *Service) CreateUser(ctx context.Context, req *userproto.UserRequest) (*userproto.UserResponse, error) {
	var userreq user.UserRequest
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.Password != req.ConfirmPassword {
		return nil, fmt.Errorf("confirm password error")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	req.Password = string(bytes)

	code := 10000 + rand.Intn(90000)
	err = pkg.SendEmail(req.Email, pkg.SendClientCode(code, req.Username))
	if err != nil {
		log.Println("ERROR: sending email to user !!", err)
	}

	userreq.Username = req.Username
	userreq.Email = req.Email
	userreq.Password = req.Password

	res, err := s.service.Createuser(userreq)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return &userproto.UserResponse{
		UserId: res.UserID,
		Username: res.Username,
		Email: res.Email,}, nil
}

func (s *Service) GetbyIdUser(ctx context.Context, req *userproto.GetUserRequest)(*userproto.User, error){
	var userreq user.GetUserRequest
	userreq.ID = req.Id
	res, err := s.service.GetByIdUser(userreq)
	if err != nil{
		log.Println("error:", err)
		return nil, err
	}

	return &userproto.User{
		Id: res.ID,
		Username: res.Username,
		Email: res.Email,
		Password: res.Password}, nil
}

func (s *Service) GetUsers(ctx context.Context, _ *userproto.Empty) (*userproto.ListUser, error) {
	res, err := s.service.GetAlluser()
	if err != nil {
		log.Println("Get all user error:", err)
		return nil, err
	}
	var protoUsers []*userproto.User
	for _, u := range res.User {
		protoUser := &userproto.User{
			Id:       u.ID,
			Username: u.Username,
			Age:      u.Age,
			Email:    u.Email,
		}
		protoUsers = append(protoUsers, protoUser)
	}

	return &userproto.ListUser{User: protoUsers}, nil
}

func (s *Service) UpdateUser(ctx context.Context, req *userproto.UpdateUserReq) (*userproto.UpdateUserRes, error){
	var users user.UpdateUserReq

	users.UserID = req.UserId
	users.Username = req.Username
	users.Age = req.Age
	users.Email = req.Email

	err := s.service.Update(users)
	if err != nil {
		log.Println("Update user error:", err)
		return nil, err
	}

	return &userproto.UpdateUserRes{Message: "users updated"}, nil
}

func (s *Service) UpdatePassword(ctx context.Context, req *userproto.UpdatePasswordReq)(*userproto.UpdateUserRes, error){
	var users user.UpdatePasswordReq

	users.NewPassword  = req.NewPassword
	users.OldPassword = req.OldPassword
	users.UserID = req.UserId
	err := s.service.UpdatePassworduser(users)
	if err != nil {
		log.Println("update password error")
		return nil, err
	}

	return &userproto.UpdateUserRes{Message: "Password updated"}, nil
}

func (s *Service) DeleteUser(ctx context.Context, req *userproto.GetUserRequest) (*userproto.UpdateUserRes, error){
	var users user.GetUserRequest
	users.ID = req.Id
	err := s.service.Delete(users)
	if err != nil {
		log.Println("delete user error")
		return nil, err
	}

	return &userproto.UpdateUserRes{Message: "users deleted"}, nil
}

