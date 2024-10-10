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


func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", pingHandler)
	r.GET("/tx/status", getTxStatusHandler)
	r.POST("/tx/status", postTxStatusHandler)
}

