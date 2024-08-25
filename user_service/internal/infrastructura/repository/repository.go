package repository

import "user_service/internal/entity/user"

type UserRepository interface {
	AddUser(req user.UserRequest) (*user.UserResponse, error)
	GetbyIdUser(req user.GetUserRequest) (*user.User, error)
	GetAll() (*user.ListUser, error)
	UpdateUser(req user.UpdateUserReq)error
	UpdatePassword(req user.UpdatePasswordReq)error
	DeleteUser(req user.GetUserRequest)error
}
