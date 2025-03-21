package repository

import (
	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/jmoiron/sqlx"
)

type  SongRepository interface {
	CreateSong(song models.Song) error
	DeleteSongById(id int) error
	UpdateSongById(id int, input models.UpdateSongRequest) error
	GetSongById(id int) (models.Song, error)
	GetSongs(filter models.SongFilter, page, limit int) ([]models.Song, int, error)
}

type Repository struct {
	SongRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		SongRepository: NewSongPostgres(db),
		
	}
}
