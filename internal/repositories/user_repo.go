package repositories

import (
	"database/sql"
	"svi_danie/internal/repositories/models"

	"github.com/google/uuid"
)

type UserRepository struct {
	Db *sql.DB
}

func (p *UserRepository) Create(user *models.User) error {
	_, err := p.Db.Exec(`
        INSERT INTO users (id, login, password)
        VALUES ($1, $2, $3)
    `, user.Id, user.Login, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *UserRepository) FindByLogin(login string) *models.User {
	var user models.User
	err := p.Db.QueryRow(`
        SELECT id, login, password
        FROM users
        WHERE login = $1
    `, login).Scan(&user.Id, &user.Login, &user.Password)
	if err != nil {
		return nil
	}

	return &user
}

func (p *UserRepository) Read(userID uuid.UUID) (*models.User, error) {
	var user models.User
	err := p.Db.QueryRow(`
        SELECT id, login, password
        FROM users
        WHERE id = $1
    `, userID).Scan(&user.Id, &user.Login, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *UserRepository) Update(user models.User) error {
	_, err := p.Db.Exec(`
        UPDATE users
        SET login = $1, password = $2
        WHERE id = $3
    `, user.Login, user.Password, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (p *UserRepository) Delete(userId uuid.UUID) error {
	_, err := p.Db.Exec(`
        DELETE FROM users
        WHERE id = $1
    `, userId)
	if err != nil {
		return err
	}

	return nil
}
