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

// @Summary Get Transaction Status
// @Description Get the status of a transaction based on chain name and transaction hash
// @Tags transaction
// @Accept json
// @Produce json
// @Param chainName query string true "Chain Name" Enums(ethereum, solana, ton)
// @Param tx query string true "Transaction Hash"
// @Success 200 {object} types.GetResp
// @Failure 400 {object} map[string]string
// @Router /tx/status [get]
func getTxStatusHandler(c *gin.Context) {
	chainName := c.Query("chainName")
	txHash := c.Query("tx")

	if chainName == "" || txHash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "chainName and txHash are required"})
		return
	}

	// 处理逻辑
	slog.Infof("Getting status for chain: %s, txHash: %s", chainName, txHash)
}

// @Summary Create Transaction Status
// @Description Create a new transaction status for a specific chain
// @Tags transaction
// @Accept json
// @Produce json
// @Param request body types.PutReq true "Transaction Request"
// @Success 201 {object} types.PutResp
// @Failure 400 {object} map[string]string
// @Router /tx/status [post]
func postTxStatusHandler(c *gin.Context) {
	appctx := appctx.GetContext()
	var req types.PutReq
	if err := c.ShouldBindJSON(&req); err != nil {
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
}
