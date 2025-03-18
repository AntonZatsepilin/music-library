package handler

import (
	"github.com/AntonZatsepilin/music-library.git/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateSong(c *gin.Context) {
	var inputSong models.Song

	if err := c.BindJSON(&inputSong); err != nil {
		newErrorResponse(c, 400, err.Error())
		return
	}

	if err := h.services.CreateSong(inputSong); err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(200, statusResponse{"Song created successfully"})

}