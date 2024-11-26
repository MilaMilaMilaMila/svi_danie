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
