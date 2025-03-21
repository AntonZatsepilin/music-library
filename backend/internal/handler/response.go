package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Error response
// swagger:response errorResponse
type errorResponse struct {
	// Error message
    // Example: invalid request parameters
	Message string `json:"message"`
}

// Status response
// swagger:response statusResponse
type statusResponse struct {
    // Status message
    // Example: Song created successfully
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}