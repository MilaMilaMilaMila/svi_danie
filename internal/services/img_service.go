package services

import (
	"errors"
	"fmt"
	"svi_danie/internal/repositories"
	"svi_danie/internal/repositories/models"

	"github.com/google/uuid"
)

type ImgService struct {
	ImgRepo *repositories.ImgRepository
}

func (s *ImgService) GetImageById(imgId uuid.UUID) (*models.Img, error) {
	image, err := s.ImgRepo.Read(imgId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("image service: get image by id: %s", err))
	}
	return image, nil
}

func (s *ImgService) CreateImage(img models.Img) error {
	err := s.ImgRepo.Create(img)
	if err != nil {
		return errors.New(fmt.Sprintf("image service: create image: %s", err))
	}
	return nil
}
