package ton

import (
	"context"
	"encoding/json"
	"log/slog"
	"tx-status-server/appctx"

	"github.com/hibiken/asynq"
	"github.com/xssnick/tonutils-go/tlb"
)

type TonTxStatusTask struct {
	Id      int    `json:"id"`
	TxType  int    `json:"txType"` // NOTICE: mainnet:0 or testnet:1 | TODO: maybe use workchain_id
	Lt      uint64 `json:"lt"`
	TxHash  string `json:"txHash"`
	Address string `json:"address"`
	Status  int    `json:"status"`
}

func HandleTonTxStatusTask(ctx *appctx.Context) asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		var t TonTxStatusTask
		err := json.Unmarshal(task.Payload(), &t)
		if err != nil {
			slog.Error("fail to unmarshal task", slog.Any("err", err))
			return err
		}

		var tx *tlb.Transaction
		switch t.TxType {
		case 0:
			tx, err = DefaultMainnetGetTransaction(t.Lt, t.Address, t.TxHash)
			if err != nil {
				slog.Error("fail to get tx on mainnet", slog.Any("tx hash", t.TxHash))
				return err
			}
		case 1:
			tx, err = DefaultTestnetGetTransaction(t.Lt, t.Address, t.TxHash)
			if err != nil {
				slog.Error("fail to get tx on testnet", slog.Any("tx hash", t.TxHash))
				return err
			}
		}

		// todo: handle tx
		slog.Info("handle tx", slog.Any("tx", tx))
		return nil
	}
}
