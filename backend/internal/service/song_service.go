package service

import "github.com/AntonZatsepilin/music-library.git/internal/repository"

type SongService struct {
	repo repository.SongRepository
}

func NewSongService(repo repository.SongRepository) *SongService {
	return &SongService{repo: repo}
}