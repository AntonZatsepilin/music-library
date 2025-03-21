package service

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/AntonZatsepilin/music-library.git/internal/repository"
	"github.com/bxcodec/faker/v3"
	"github.com/sirupsen/logrus"
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
    logrus.WithFields(logrus.Fields{
        "group": input.Group,
        "song":  input.Song,
    }).Debug("Data request from external API")
    detail, err := s.infoClient.GetSongDetail(input.Group, input.Song)
    if err != nil {
        logrus.WithError(err).Error("API error")
        return fmt.Errorf("API error: %w", err)
        }

    song := models.Song{
        Group:       input.Group,
        SongName:    input.Song,
        ReleaseDate: detail.ReleaseDate,
        Text:        detail.Text,
        Link:        detail.Link,
    }
    logrus.Debug("Saving a song to the database")
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

func (s *SongServiceImpl) DeleteSongById(id int) error {
    return s.repo.DeleteSongById(id)
}

func (s *SongServiceImpl) UpdateSongById(id int, input models.UpdateSongRequest) error {
    return s.repo.UpdateSongById(id, input)
}

func (s *SongServiceImpl) GetSongById(id int) (models.Song, error) {
    return s.repo.GetSongById(id)
}

func (s *SongServiceImpl) GetSongLyrics(songId int, page, limit int) ([]string, int, error) {

    song, err := s.repo.GetSongById(songId)
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "songId": songId,
            "error":  err,
        }).Error("Failed to get song for lyrics")
        return nil, 0, fmt.Errorf("song not found: %w", err)
    }

    verses := strings.Split(song.Text, "\n\n")
    totalVerses := len(verses)

    start := (page - 1) * limit
    if start >= totalVerses {
        return []string{}, totalVerses, nil
    }

    end := start + limit
    if end > totalVerses {
        end = totalVerses
    }

    logrus.WithFields(logrus.Fields{
        "songId":      songId,
        "totalVerses": totalVerses,
        "page":        page,
        "limit":       limit,
    }).Debug("Successfully processed lyrics request")

    return verses[start:end], totalVerses, nil
}

func (s *SongServiceImpl) GetSongs(filter models.SongFilter, page, limit int) ([]models.Song, int, error) {
    if page < 1 {
        page = 1
    }
    if limit < 1 {
        limit = 10
    } else if limit > 100 {
        limit = 100
    }
    
    return s.repo.GetSongs(filter, page, limit)
}