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

	api := router.Group("/song")
	{
		api.POST("", h.CreateSong)
		api.GET("/:id", h.GetSongById)
		api.PUT("/:id", h.UpdateSongById)
		api.DELETE("/:id", h.DeleteSongById)
	}

	return router
}