package ton

import (
	"context"
	"tx-status-server/appctx"

	"github.com/hibiken/asynq"
)

func HandleTonTxStatusTask(ctx *appctx.Context) asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		return nil
	}
}
