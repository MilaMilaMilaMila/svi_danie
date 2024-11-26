package repositories

import (
	"database/sql"
	"github.com/google/uuid"
	"svi_danie/internal/repositories/models"
)

type ProjectRepository struct {
	Db *sql.DB
}

func (p *ProjectRepository) Create(proj models.Project) error {
	_, err := p.Db.Exec(`
        INSERT INTO projects (id, owner_id, title)
        VALUES ($1, $2, $3)
    `, proj.Id, proj.OwnerId, proj.Title)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectRepository) Read(projId uuid.UUID) (*models.Project, error) {
	var proj models.Project
	err := p.Db.QueryRow(`
        SELECT id, owner_id, title
        FROM projects
        WHERE id = $1
    `, projId).Scan(&proj.Id, &proj.OwnerId, &proj.Title)
	if err != nil {
		return nil, err
	}

	return &proj, nil
}

func (p *ProjectRepository) Update(proj models.Project) error {
	_, err := p.Db.Exec(`
        UPDATE projects
        SET owner_id = $1, title = $2
        WHERE id = $3
    `, proj.OwnerId, proj.Title, proj.Id)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectRepository) Delete(projId uuid.UUID) error {
	_, err := p.Db.Exec(`
        DELETE FROM projects
        WHERE id = $1
    `, projId)
	if err != nil {
		return err
	}

	return nil
}
