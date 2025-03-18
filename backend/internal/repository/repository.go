package repository

import (
	"github.com/jmoiron/sqlx"
)

type  SongRepository interface {
	Create(song *models.Song) error
	GetByID(id int) (*models.Song, error)
	GetAll(filter models.SongFilter, page, limit int) ([]models.Song, int, error)
	Update(song *models.Song) error
	Delete(id int) error
	GetSongText(id int, page, limit int) ([]string, int, error)
}

type Repository struct {
	SongRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		SongRepository: NewSongPostgres(db),
		
	}
}
