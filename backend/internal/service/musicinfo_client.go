package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/sirupsen/logrus"
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

func (c *MusicInfoClient) GetSongDetail(group, song string) (*models.SongDetail, error) {

    logrus.WithFields(logrus.Fields{
        "group": group,
        "song":  song,
        "url":   c.baseURL,
    }).Debug("Data request from external API")

    url := fmt.Sprintf("%s/info?group=%s&song=%s", c.baseURL, group, song)
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return nil, err
    }

    logrus.Debug("Sending a request to the external API")

    resp, err := c.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    logrus.WithField("status", resp.StatusCode).Debug("Received a response from the external API")

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