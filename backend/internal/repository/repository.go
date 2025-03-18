package repository

import (
	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/jmoiron/sqlx"
)

type  SongRepository interface {
	CreateSong(song models.Song) error
}

type Repository struct {
	SongRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		SongRepository: NewSongPostgres(db),
		
	}
}
