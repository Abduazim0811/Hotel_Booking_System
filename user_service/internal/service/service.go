package service

import (
	"user_service/internal/entity/user"
	"user_service/internal/infrastructura/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Createuser(req user.UserRequest) (*user.UserResponse, error) {
	return u.repo.AddUser(req)
}

func (u *UserService) GetByEmailUser(email string)(*user.User, error){
	return u.repo.GetbyEmail(email)
}

func (u *UserService) GetByIduser(req user.GetUserRequest) (*user.User, error) {
	return u.repo.GetbyIdUser(req)
}

func (u *UserService) GetAlluser()(*user.ListUser, error){
	return u.repo.GetAll()
}
func (u *UserService) Update(req user.UpdateUserReq) error {
	return u.repo.UpdateUser(req)
}

func (u *UserService)UpdatePassworduser(req user.UpdatePasswordReq)error{
	return u.repo.UpdatePassword(req)
}

func (u *UserService)Delete(req user.GetUserRequest) error{
	return u.repo.DeleteUser(req)
}
