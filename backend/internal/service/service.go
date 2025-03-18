package service

import (
	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/AntonZatsepilin/music-library.git/internal/repository"
)

type SongService interface {
	CreateSong(song models.Song) error
}

type Service struct {
	SongService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		SongService: NewSongService(repos.SongRepository),
	}
}