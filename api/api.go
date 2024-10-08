package api

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

func Run() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/tx/status", func(c *gin.Context) {
	})

	r.POST("tx/status", func(c *gin.Context) {
		appctx := appctx.GetContext()
		var req types.PutReq
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		switch req.ChainKind {
		case constants.ChainKindEthereum:
			switch req.ChainName {
			case constants.ChainNameEthereum:
				slog.Info("handle ethereum tx status")
			case constants.ChainNameEthereumSepolia:
				slog.Info("handle ethereum sepolia tx status")
			default:
				slog.Error("req.ChainName not implement", "chainName", req.ChainName)
			}
		case constants.ChainKindSolana:
			slog.Info("handle solana tx status")
		case constants.ChainKindTon:
			_, err := appctx.TaskClient.Enqueue(
				task.TypeTxStatusTon,
				asynq.TaskID(fmt.Sprintf("%s:%s", req.ChainName, req.Tx)),
				asynq.MaxRetry(32),
				asynq.Timeout(10*time.Second),
			)
			if err != nil {
				slog.Error("fail to enqueue ton tx status query", slog.Any("err", err))
			}
		default:
			slog.Error("req.ChainKind not implement", "chainKind", req.ChainKind)
		}
	})

	r.Run()
}
