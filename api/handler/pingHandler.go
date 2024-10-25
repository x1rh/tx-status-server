package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Ping
// @Description Ping the server to check if it's alive
// @Tags ping
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
