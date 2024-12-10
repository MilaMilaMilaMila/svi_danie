package services

import (
	"errors"
	"fmt"
	"svi_danie/internal/repositories"
	"svi_danie/internal/repositories/models"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func (u *UserService) AddUser(user *models.User) error {
	err := u.UserRepo.Create(user)
	if err != nil {
		return errors.New(fmt.Sprintf("user service: create user: %s", err))
	}

	return nil
}

func (u *UserService) AuthUser(login, password string) (*models.User, error) {
	user := u.UserRepo.FindByLogin(login)
	if user == nil {
		return nil, errors.New(fmt.Sprintf("user service: no such user: %s", login))
	}
	if user.Password != password {
		return nil, errors.New("user service: wrong password")
	}
	return user, nil
}
