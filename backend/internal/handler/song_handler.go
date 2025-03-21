package handler

import (
	"net/http"
	"strconv"

	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

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

func (h *Handler) GetSongs(c *gin.Context) {
    var filter models.SongFilter
    if err := c.ShouldBindQuery(&filter); err != nil {
        newErrorResponse(c, http.StatusBadRequest, "invalid filter parameters")
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