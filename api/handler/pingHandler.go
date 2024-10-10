package handler 

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
	"tx-status-server/appctx"
	"tx-status-server/asynqx/task"
	"tx-status-server/constants"
	"tx-status-server/sdk/types"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
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