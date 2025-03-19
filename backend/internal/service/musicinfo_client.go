package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/AntonZatsepilin/music-library.git/internal/models"
)

type MusicInfoClient struct {
    baseURL string
    client  *http.Client
}

func NewMusicInfoClient(baseURL string) *MusicInfoClient {
    return &MusicInfoClient{
        baseURL: baseURL,
        client: &http.Client{
            Timeout: 10 * time.Second,
        },
    }
}

type APIError struct {
    StatusCode int
    Body       string
}

func (e *APIError) Error() string {
    return fmt.Sprintf("API request failed: status %d, body %s", e.StatusCode, e.Body)
}

func (c *MusicInfoClient) GetSongDetail(ctx context.Context, group, song string) (*models.SongDetail, error) {
    url := fmt.Sprintf("%s/info?group=%s&song=%s", c.baseURL, group, song)
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        return nil, err
    }

    resp, err := c.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return nil, &APIError{
            StatusCode: resp.StatusCode,
            Body:       string(body),
        }
    }

    var detail models.SongDetail
    if err := json.NewDecoder(resp.Body).Decode(&detail); err != nil {
        return nil, err
    }

    return &detail, nil
}