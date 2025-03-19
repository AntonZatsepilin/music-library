package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/AntonZatsepilin/music-library.git/internal/repository"
)

var (
    ErrInvalidInput    = errors.New("invalid input")
    ErrExternalService = errors.New("external service error")
)

type SongServiceImpl struct {
    repo       repository.SongRepository
    infoClient *MusicInfoClient
}

func NewSongService(repo repository.SongRepository, infoClient *MusicInfoClient) *SongServiceImpl {
    return &SongServiceImpl{
        repo:       repo,
        infoClient: infoClient,
    }
}


func (s *SongServiceImpl) CreateSong(input models.CreateSongRequest) error {
    detail, err := s.infoClient.GetSongDetail(context.Background(), input.Group, input.Song)
    if err != nil {
        var apiErr *APIError
        if errors.As(err, &apiErr) {
            switch apiErr.StatusCode {
            case http.StatusBadRequest:
                return fmt.Errorf("%w: %s", ErrInvalidInput, apiErr.Error())
            default:
                return fmt.Errorf("%w: %s", ErrExternalService, apiErr.Error())
            }
        }
        return fmt.Errorf("failed to get song details: %w", err)
    }

    song := models.Song{
        Group:       input.Group,
        SongName:    input.Song,
        ReleaseDate: detail.ReleaseDate,
        Text:        detail.Text,
        Link:        detail.Link,
    }

    return s.repo.CreateSong(song)
}