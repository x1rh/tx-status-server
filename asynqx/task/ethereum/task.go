package ethereum

import (
	"context"
	"encoding/json"
	"log/slog"
	"tx-status-server/appctx"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hibiken/asynq"
)

type EthTxStatusQueryTask struct {
	Id      int    `json:"id"`
	ChainId int    `json:"chainId"`
	TxHash  string `json:"txHash"`
	Status  int    `json:"status"`
}

func HandleEthereumTxStatusTask(_appctx *appctx.Context) asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		var t EthTxStatusQueryTask
		err := json.Unmarshal(task.Payload(), &t)
		if err != nil {
			slog.Error("fail to unmarshal task", slog.Any("err", err))
			return err
		}

		receipt, err := _appctx.EthClientHub.MustWithChainID(t.ChainId).TransactionReceipt(ctx, common.HexToHash(t.TxHash))
		if err != nil {
			// transaction maybe still pending
			slog.Error("fail to check receipt", slog.Any("err", err))
			return err
		}

		if receipt.Status == 0 {
			t.Status = 2 // fail
		} else if receipt.Status == 1 {
			t.Status = 1 // success
		}
		slog.Debug(
			"tx detail",
			slog.Any("status", receipt.Status),
			slog.Any("tx hash", receipt.TxHash.Hex()),
			slog.Any("gas used", receipt.GasUsed),
		)
		return nil
	}
}
