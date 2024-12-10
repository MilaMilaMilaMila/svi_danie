package repositories

import (
	"database/sql"
	"svi_danie/internal/repositories/models"

	"github.com/google/uuid"
)

type ImgRepository struct {
	Db *sql.DB
}

func (p *ImgRepository) Create(img *models.Img) error {
	_, err := p.Db.Exec(`
        INSERT INTO img (id, data)
        VALUES ($1, $2)
    `, img.Id, img.Data)
	if err != nil {
		return err
	}

	return nil
}

func (p *ImgRepository) Read(imgID uuid.UUID) (*models.Img, error) {
	var img models.Img
	err := p.Db.QueryRow(`
        SELECT id, data
        FROM img
        WHERE id = $1
    `, imgID).Scan(&img.Id, &img.Data)
	if err != nil {
		return nil, err
	}

	return &img, nil
}
