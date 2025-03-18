package service

import (
	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/AntonZatsepilin/music-library.git/internal/repository"
)

type SongServiceImpl struct {
	repo repository.SongRepository
}

func NewSongService(repo repository.SongRepository) *SongServiceImpl {
	return &SongServiceImpl{repo: repo}
}

func (s *SongServiceImpl) CreateSong(song models.Song) error {
	return s.repo.CreateSong(song)
}