package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getSongId(c *gin.Context) (int, error) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, "invalid song id")
        return 0, errors.New("invalid song id")
    }

    return id, nil
}