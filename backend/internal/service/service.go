package service

import (
	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/AntonZatsepilin/music-library.git/internal/repository"
)

type SongService interface {
	CreateSong(song models.CreateSongRequest) error
	GenerateFakeSongs(count int) error
	DeleteSongById(id int) error
	UpdateSongById(id int, input models.UpdateSongRequest) error
}

type Service struct {
	SongService
}

func NewService(repos *repository.Repository, infoClient *MusicInfoClient) *Service {
	return &Service{
		SongService: NewSongService(repos.SongRepository, infoClient),
	}
}