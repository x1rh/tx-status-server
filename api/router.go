package api

import (
	"tx-status-server/api/handler"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	handler.RegisterRoutes(r)
	r.Run()
}
