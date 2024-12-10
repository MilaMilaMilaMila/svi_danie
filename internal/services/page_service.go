package services

import (
	"errors"
	"fmt"
	"svi_danie/internal/repositories"
	"svi_danie/internal/repositories/models"

	"github.com/google/uuid"
)

type PageService struct {
	PageRepo *repositories.PageRepository
}

func (p *PageService) CreatePage(page *models.Page) error {
	err := p.PageRepo.Create(page)
	if err != nil {
		return errors.New(fmt.Sprintf("page service: create page: %s", err))
	}
	return nil
}

func (p *PageService) UpdatePage(page *models.Page) error {
	err := p.PageRepo.Update(page)
	if err != nil {
		return errors.New(fmt.Sprintf("page service: update page: %s", err))
	}
	return nil
}

func (p *PageService) DeletePage(projId uuid.UUID) error {
	err := p.PageRepo.Delete(projId)
	if err != nil {
		return errors.New(fmt.Sprintf("page service: delete page: %s", err))
	}
	return nil
}

func (p *PageService) GetPage(pageId uuid.UUID) (*models.Page, error) {
	proj, err := p.PageRepo.Read(pageId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("page service: get page: %s", err))
	}
	return proj, nil
}

func (p *PageService) GetAllProjectPages(projId uuid.UUID) ([]*models.Page, error) {
	pages, err := p.PageRepo.ReadAllProjectPages(projId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("page service: get all proj %s pages: %s", projId, err))
	}
	return pages, nil
}
