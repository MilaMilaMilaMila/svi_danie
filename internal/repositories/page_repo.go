package repositories

import (
	"database/sql"
	"github.com/google/uuid"
	"svi_danie/internal/repositories/models"
)

type PageRepository struct {
	Db *sql.DB
}

func (p *PageRepository) Create(page models.Page) error {
	_, err := p.Db.Exec(`
        INSERT INTO pages (id, owner_id, title, data)
        VALUES ($1, $2, $3, $4)
    `, page.Id, page.OwnerId, page.Title, page.Data)
	if err != nil {
		return err
	}

	return nil
}

func (p *PageRepository) Read(pageID uuid.UUID) (*models.Page, error) {
	var page models.Page
	err := p.Db.QueryRow(`
        SELECT id, owner_id, title, data
        FROM pages
        WHERE id = $1
    `, pageID).Scan(&page.Id, &page.OwnerId, &page.Title, &page.Data)
	if err != nil {
		return nil, err
	}

	return &page, nil
}

func (p *PageRepository) Update(page models.Page) error {
	_, err := p.Db.Exec(`
        UPDATE pages
        SET owner_id = $1, title = $2, data = $3
        WHERE id = $4
    `, page.OwnerId, page.Title, page.Data, page.Id)
	if err != nil {
		return err
	}

	return nil
}

func (p *PageRepository) Delete(pageID uuid.UUID) error {
	_, err := p.Db.Exec(`
        DELETE FROM pages
        WHERE id = $1
    `, pageID)
	if err != nil {
		return err
	}

	return nil
}
