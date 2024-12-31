package service

import (
	"ecommerce-inventory/model"
	"ecommerce-inventory/repository"
	"fmt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{UserRepository: repo}
}

func (service *UserService) RegisterUser(user *model.User) error {
	return service.UserRepository.AddUser(user)
}

func (service *UserService) AuthenticateUser(username, password string) (*model.User, error) {
	user, err := service.UserRepository.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	if user.Password != password {
		return nil, fmt.Errorf("invalid credentials")
	}

	return user, nil
}
