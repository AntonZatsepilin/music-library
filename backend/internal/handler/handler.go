package handler

import (
	"github.com/AntonZatsepilin/music-library.git/internal/service"
	"github.com/gin-gonic/gin"
)


type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}


func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/songs")
	{
		api.GET("", h.GetSongs)
		api.POST("", h.CreateSong)
		api.GET("/:id", h.GetSongById)
		api.PUT("/:id", h.UpdateSongById)
		api.DELETE("/:id", h.DeleteSongById)
		api.GET("/:id/lyrics", h.GetSongLyrics)
		api.GET("/generate", h.GenerateFakeSongs)
	}

	return router
}