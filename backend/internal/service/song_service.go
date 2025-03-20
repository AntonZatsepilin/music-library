package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/AntonZatsepilin/music-library.git/internal/repository"
	"github.com/bxcodec/faker/v3"
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

func (s *SongServiceImpl) GenerateFakeSongs(count int) error {
    rand.Seed(time.Now().UnixNano())
    
    for i := 0; i < count; i++ {
        year := 1990 + rand.Intn(34)
        month := rand.Intn(12) + 1
        day := rand.Intn(28) + 1
        
        song := models.Song{
            Group:       fmt.Sprintf("%s %s", faker.FirstName(), faker.LastName()),
            SongName:    faker.Word(),
            ReleaseDate: fmt.Sprintf("%d-%02d-%02d", year, month, day),
            Text:        faker.Paragraph(),
            Link:        fmt.Sprintf("https://example.com/%s", faker.UUIDDigit()),
        }
        
        if err := s.repo.CreateSong(song); err != nil {
            return fmt.Errorf("failed to generate song: %w", err)
        }
    }
    return nil
}