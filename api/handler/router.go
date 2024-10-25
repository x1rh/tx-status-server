package handler

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", pingHandler)
	r.GET("/tx/status", getTxStatusHandler)
	r.POST("/tx/status", postTxStatusHandler)
}
