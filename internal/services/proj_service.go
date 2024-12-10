package services

import (
	"errors"
	"fmt"
	"svi_danie/internal/repositories"
	"svi_danie/internal/repositories/models"

	"github.com/google/uuid"
)

type ProjService struct {
	ProjRepo *repositories.ProjectRepository
	PageRepo *repositories.PageRepository
}

func (p *ProjService) CreateProj(proj *models.Project) error {
	err := p.ProjRepo.Create(proj)
	if err != nil {
		return errors.New(fmt.Sprintf("proj service: create proj: %s", err))
	}
	return nil
}

func (p *ProjService) DeleteProj(projId uuid.UUID) error {
	err := p.ProjRepo.Delete(projId)
	if err != nil {
		return errors.New(fmt.Sprintf("proj service: delete proj: %s", err))
	}
	return nil
}

func (p *ProjService) GetProj(projId uuid.UUID) (*models.Project, error) {
	proj, err := p.ProjRepo.Read(projId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("proj service: get proj: %s", err))
	}

	pages, err := p.PageRepo.ReadAllProjectPages(projId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("proj service: get proj pages: %s", err))
	}

	proj.Pages = pages
	return proj, nil
}

func (p *ProjService) GetAllUserProj(userId uuid.UUID) ([]*models.Project, error) {
	proj, err := p.ProjRepo.ReadAllUserProjects(userId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("proj service: get all user %s projects: %s", userId, err))
	}

	//for _, proj := range proj {
	//	pages, err := p.PageRepo.ReadAllProjectPages(proj.Id)
	//	if err != nil {
	//		return nil, errors.New(fmt.Sprintf("proj service: get proj pages: %s", err))
	//	}
	//
	//	proj.Pages = pages
	//}
	return proj, nil
}

func (p *ProjService) CheckOwnership(userId, projectId uuid.UUID) error {
	project, err := p.GetProj(projectId)
	if err != nil {
		return errors.New("project not found")
	}
	if project.OwnerId != userId {
		return errors.New("project ownership does not match")
	}
	return nil
}
