package handler

import (
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
		api.GET("/:id/text", h.GetSongText)
		api.POST("", h.CreateSong)
		api.PUT("/:id", h.UpdateSong)
		api.DELETE("/:id", h.DeleteSong)
	}

	return router
}