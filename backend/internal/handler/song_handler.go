package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CreateSong godoc
// @Summary Create new song
// @Description Create new song with metadata
// @Tags songs
// @Accept json
// @Produce json
// @Param input body models.CreateSongRequest true "Song data"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /songs [post]
func (h *Handler) CreateSong(c *gin.Context) {
	logrus.Debug("Received a request to create a song")

	var inputSong models.CreateSongRequest

	if err := c.BindJSON(&inputSong); err != nil {
		logrus.WithError(err).Warn("Invalid request format")
		newErrorResponse(c, 400, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
        "group": inputSong.Group,
        "song":  inputSong.Song,
    }).Info("An attempt at a song creation")

	if err := h.services.SongService.CreateSong(inputSong); err != nil {
		logrus.WithError(err).Error("Song creation error")
		newErrorResponse(c, 500, err.Error())
		return
	}

	logrus.Info("Song created successfully")
	c.JSON(200, statusResponse{"Song created successfully"})

}

// DeleteSongById godoc
// @Summary Delete song
// @Description Delete song by ID
// @Tags songs
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /songs/{id} [delete]
func (h *Handler) DeleteSongById(c *gin.Context) {
	logrus.Debug("Received a request to delete a song")
	
	songId, err := getSongId(c)

	if err != nil {
		return
	}

	if err := h.services.SongService.DeleteSongById(songId); err != nil {
		logrus.WithError(err).Error("Song deletion error")
		newErrorResponse(c, 500, err.Error())
		return
	}

	logrus.Info("Song deleted successfully")
	c.JSON(200, statusResponse{"Song deleted successfully"})
}

// UpdateSongById godoc
// @Summary Update song
// @Description Update existing song details
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param input body models.UpdateSongRequest true "Update data"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /songs/{id} [put]
func (h *Handler) UpdateSongById(c *gin.Context) {
	logrus.Debug("Received a request to update a song")

	songId, err := getSongId(c)

	if err != nil {
		return
	}

	var inputSong models.UpdateSongRequest

	if err := c.BindJSON(&inputSong); err != nil {
		logrus.WithError(err).Warn("Invalid request format")
		newErrorResponse(c, 400, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"group": inputSong.Group,
		"song":  inputSong.Song,
	}).Info("An attempt at a song update")

	if err := h.services.SongService.UpdateSongById(songId, inputSong); err != nil {
		logrus.WithError(err).Error("Song update error")
		newErrorResponse(c, 500, err.Error())
		return
	}

	logrus.Info("Song updated successfully")
	c.JSON(200, statusResponse{"Song updated successfully"})
}

// GetSongById godoc
// @Summary Get song by ID
// @Description Get song details by ID
// @Tags songs
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} models.Song
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /songs/{id} [get]
func (h *Handler) GetSongById(c *gin.Context) {
	logrus.Debug("Received a request to get a song")

	songId, err := getSongId(c)

	if err != nil {
		return
	}

	song, err := h.services.SongService.GetSongById(songId)

	if err != nil {
		logrus.WithError(err).Error("Song retrieval error")
		newErrorResponse(c, 500, err.Error())
		return
	}

	logrus.Info("Song retrieved successfully")
	c.JSON(200, song)
}

// GetSongLyrics godoc
// @Summary Get song lyrics
// @Description Get paginated song lyrics verses
// @Tags lyrics
// @Produce json
// @Param id path int true "Song ID"
// @Param page query int false "Page number" default(1) minimum(1)
// @Param limit query int false "Items per page" default(10) minimum(1) maximum(100)
// @Success 200 {object} models.LyricResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /songs/{id}/lyrics [get]
func (h *Handler) GetSongLyrics(c *gin.Context) {
	logrus.Debug("Received a request to get a song lyrics")

	songId, err := getSongId(c)

	if err != nil {
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
    if err != nil || page < 1 {
        newErrorResponse(c, http.StatusBadRequest, "invalid page number")
        return
    }

    limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
    if err != nil || limit < 1 {
        newErrorResponse(c, http.StatusBadRequest, "invalid limit value")
        return
    }

	verses, total, err := h.services.SongService.GetSongLyrics(songId, page, limit)
    if err != nil {
        logrus.WithError(err).Error("Lyrics retrieval error")
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

	response := models.LyricResponse{
        Verses: verses,
        Total:  total,
        Page:   page,
        Limit:  limit,
    }

    logrus.Info("Lyrics retrieved successfully")
    c.JSON(http.StatusOK, response)
}

// GetSongs godoc
// @Summary Get songs list
// @Description Get filtered and paginated list of songs
// @Tags songs
// @Produce json
// @Param group query string false "Filter by group name"
// @Param song query string false "Filter by song name"
// @Param releaseDate query string false "Filter by release date (YYYY-MM-DD)"
// @Param text query string false "Search in lyrics"
// @Param link query string false "Filter by link"
// @Param sort_by query string false "Sort field (group|song|releaseDate|text|link)"
// @Param sort_order query string false "Sort order (ASC|DESC)"
// @Param page query int false "Page number" default(1) minimum(1)
// @Param limit query int false "Items per page" default(10) minimum(1) maximum(100)
// @Success 200 {object} models.SongsResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /songs [get]
func (h *Handler) GetSongs(c *gin.Context) {
    var filter models.SongFilter
    if err := c.ShouldBindQuery(&filter); err != nil {
        newErrorResponse(c, http.StatusBadRequest, "invalid filter parameters")
        return
    }

    if filter.SortBy != "" {
        allowedFields := map[string]bool{
            "group":       true,
            "song":       true,
            "releaseDate": true,
            "text":       true,
            "link":       true,
            "":           true,
        }
        if !allowedFields[filter.SortBy] {
            newErrorResponse(c, http.StatusBadRequest, "invalid sort_by parameter")
            return
        }
    }

    if filter.SortOrder != "" {
        order := strings.ToUpper(filter.SortOrder)
        if order != "ASC" && order != "DESC" && order != "" {
            newErrorResponse(c, http.StatusBadRequest, "sort_order must be ASC or DESC")
            return
        }
    }

    page, err := strconv.Atoi(c.DefaultQuery("page", "1"))	
    if err != nil || page < 1 {
        newErrorResponse(c, http.StatusBadRequest, "invalid page number")
        return
    }

    limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
    if err != nil || limit < 1 || limit > 100 {
        newErrorResponse(c, http.StatusBadRequest, "limit must be between 1 and 100")
        return
    }

    songs, total, err := h.services.SongService.GetSongs(filter, page, limit)
    if err != nil {
        newErrorResponse(c, http.StatusInternalServerError, "failed to get songs")
        return
    }

    response := models.SongsResponse{
        Data:  songs,
        Total: total,
        Page:  page,
        Limit: limit,
    }

    c.JSON(http.StatusOK, response)
}


// GenerateFakeSongs godoc
// @Summary Generate fake songs
// @Description Generate test songs with random data
// @Tags songs
// @Accept json
// @Produce json
// @Param count query int false "Number of songs to generate" default(1) minimum(1) maximum(100)
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /songs/generate [get]
func (h *Handler) GenerateFakeSongs(c *gin.Context) {
	logrus.Debug("Received a request to generate fake songs")

	count, err := strconv.Atoi(c.DefaultQuery("count", "1"))
	if err != nil || count < 1 {
		newErrorResponse(c, http.StatusBadRequest, "invalid count value")
		return
	}

	if err := h.services.SongService.GenerateFakeSongs(count); err != nil {
		logrus.WithError(err).Error("Failed to generate fake songs")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Info("Fake songs generated successfully")
	c.JSON(http.StatusOK, statusResponse{"Fake songs generated successfully"})
}