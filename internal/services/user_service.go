package services

import (
	"errors"
	"fmt"
	"svi_danie/internal/repositories"
	"svi_danie/internal/repositories/models"

	"golang.org/x/crypto/bcrypt"
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
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("user service: wrong password: %s", err))
	}
	return user, nil
}
