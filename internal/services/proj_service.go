package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"svi_danie/internal/repositories"
	"svi_danie/internal/repositories/models"
)

type ProjService struct {
	ProRepo *repositories.ProjectRepository
}

func (p *ProjService) CreateProj(proj *models.Project) error {
	err := p.ProRepo.Create(proj)
	if err != nil {
		return errors.New(fmt.Sprintf("proj service: create proj: %s", err))
	}
	return nil
}

func (p *ProjService) DeleteProj(projId uuid.UUID) error {
	err := p.ProRepo.Delete(projId)
	if err != nil {
		return errors.New(fmt.Sprintf("proj service: delete proj: %s", err))
	}
	return nil
}

func (p *ProjService) GetProj(projId uuid.UUID) (*models.Project, error) {
	proj, err := p.ProRepo.Read(projId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("proj service: delete proj: %s", err))
	}
	return proj, nil
}

func (p *ProjService) GetAllUserProj(userId uuid.UUID) ([]*models.Project, error) {
	proj, err := p.ProRepo.ReadAllUserProjects(userId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("proj service: get all user %s projects: %s", userId, err))
	}
	return proj, nil
}
