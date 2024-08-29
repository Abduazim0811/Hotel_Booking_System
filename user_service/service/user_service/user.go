package userservice

import (
	"context"
	"fmt"
	"log"
	"user_service/internal/entity/user"
	"user_service/internal/infrastructura/redis"
	pkg "user_service/internal/pkg/email"
	"user_service/internal/pkg/jwt"
	"user_service/internal/service"
	"user_service/userproto"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

type Service struct {
	*userproto.UnimplementedUserServiceServer
	service *service.UserService
	redis   *redis.RedisClient
}

func NewGrpcService(service *service.UserService, redis *redis.RedisClient) *Service {
	return &Service{service: service, redis: redis}
}

func (s *Service) Register(ctx context.Context, req *userproto.UserRequest) (*userproto.UpdateUserRes, error) {
	var userreq user.UserRequest
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	if req.Password != req.ConfirmPassword {
		return nil, fmt.Errorf("password and confirm password do not match")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %v", err)
	}

	req.Password = string(bytes)

	code := 10000 + rand.Intn(90000)
	err = pkg.SendEmail(req.Email, pkg.SendClientCode(code, req.Username))
	if err != nil {
		return nil, fmt.Errorf("error sending email to user: %v", err)
	}

	userreq.Username = req.Username
	userreq.Email = req.Email
	userreq.Password = req.Password

	userData := map[string]interface{}{
		"userName": userreq.Username,
		"email":    userreq.Email,
		"password": userreq.Password,
		"age":      userreq.Age,
		"code":     code,
	}

	if s.redis == nil {
		log.Println("redis error", err)
		return nil, fmt.Errorf("redis client is not initialized")
	}

	err = s.redis.SetHash(req.Email, userData)
	if err != nil {
		return nil, fmt.Errorf("failed to save user data in Redis: %v", err)
	}

	return &userproto.UpdateUserRes{Message: "Verify code"}, nil
}

func (s *Service) VerifyCode(ctx context.Context, req *userproto.Req) (*userproto.UserResponse, error) {
	res, err := s.redis.VerifyEmail(ctx, req.Email, int64(req.Code))
	if err != nil {
		log.Println("verify code error: ")
		return nil, fmt.Errorf("verify code error: %v", err)
	}

	var userreq user.UserRequest

	userreq.Username = res.Username
	userreq.Email = res.Email
	userreq.Password = res.Password
	userreq.Age = res.Age

	userres, err := s.service.Createuser(userreq)
	if err != nil {
		log.Println("error")
		return nil, fmt.Errorf("error: %v", err)
	}

	return &userproto.UserResponse{
		Id: userres.Id,
		Username: userres.Username,
		Email: userres.Email,
		Age: userres.Age,
	}, nil
}

func (s *Service) Login(ctx context.Context, req *userproto.LoginRequest) (*userproto.LoginResponse, error) {
	res, err := s.service.GetByEmailUser(req.Email)
	if err != nil {
		log.Println("login error")
		return nil, fmt.Errorf("login erro")
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("incorrect password")
	}

	token, err := jwt.GenerateJWTToken(req.Email)
	if err != nil {
		return nil, err
	}

	return &userproto.LoginResponse{Token: token, ExpiresIn: "s"}, nil
}

func (s *Service) GetByIdUser(ctx context.Context, req *userproto.GetUserRequest) (*userproto.User, error) {
	var userreq user.GetUserRequest
	userreq.ID = req.Id
	res, err := s.service.GetByIduser(userreq)
	if err != nil {
		return nil, fmt.Errorf("error fetching user by ID: %v", err)
	}

	return &userproto.User{
		Id:       res.ID,
		Username: res.Username,
		Email:    res.Email,
		Password: res.Password,
	}, nil
}

func (s *Service) GetUsers(ctx context.Context, _ *userproto.UserEmpty) (*userproto.ListUser, error) {
	res, err := s.service.GetAlluser()
	if err != nil {
		return nil, fmt.Errorf("error fetching all users: %v", err)
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

func (s *Service) UpdateUser(ctx context.Context, req *userproto.UpdateUserReq) (*userproto.UpdateUserRes, error) {
	var users user.UpdateUserReq

	users.Id = req.Id
	users.Username = req.Username
	users.Age = req.Age
	users.Email = req.Email

	err := s.service.Update(users)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}

	return &userproto.UpdateUserRes{Message: "user updated successfully"}, nil
}

func (s *Service) UpdatePassword(ctx context.Context, req *userproto.UpdatePasswordReq) (*userproto.UpdateUserRes, error) {
	var users user.UpdatePasswordReq

	users.NewPassword = req.NewPassword
	users.OldPassword = req.OldPassword
	users.Id = req.Id
	err := s.service.UpdatePassworduser(users)
	if err != nil {
		return nil, fmt.Errorf("error updating password: %v", err)
	}

	return &userproto.UpdateUserRes{Message: "password updated successfully"}, nil
}

func (s *Service) DeleteUser(ctx context.Context, req *userproto.GetUserRequest) (*userproto.UpdateUserRes, error) {
	var users user.GetUserRequest
	users.ID = req.Id
	err := s.service.Delete(users)
	if err != nil {
		return nil, fmt.Errorf("error deleting user: %v", err)
	}

	return &userproto.UpdateUserRes{Message: "user deleted successfully"}, nil
}
