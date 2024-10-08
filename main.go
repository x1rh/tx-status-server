package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"tx-status-server/api"
	"tx-status-server/appctx"
	"tx-status-server/asynqx/server"
	"tx-status-server/config"
	"tx-status-server/logger"
)

// func randomTask(c *client.Client, app string) {
// 	id := 0
// 	r := rand.New(rand.NewSource(time.Now().Unix()))
// 	ticker := time.NewTicker(time.Millisecond * 100)
// 	maxRetry := 16
// 	for range ticker.C {
// 		x := r.Int()
// 		if x%2 == 0 {
// 			taskInfo, err := c.Enqueue(
// 				task.TypeTask1,
// 				biz1.Task1{
// 					Id:  id,
// 					App: app,
// 				},
// 				asynq.TaskID(fmt.Sprintf("%s:%s:id:%d", app, task.TypeTask1, id)),
// 				asynq.MaxRetry(maxRetry),
// 				asynq.Timeout(3*time.Second),
// 			)
// 			if err != nil {
// 				slog.Error("fail to enqueue task1", slog.Any("err", err))
// 				return
// 			}
// 			slog.Debug("enqueue task1 success", slog.Any("task", taskInfo))
// 		} else {
// 			taskInfo, err := c.Enqueue(
// 				task.TypeTask2,
// 				biz2.Task2{
// 					Id:  id,
// 					App: app,
// 				},
// 				asynq.TaskID(fmt.Sprintf("%s:%s:id:%d", app, task.TypeTask2, id)),
// 				asynq.MaxRetry(maxRetry),
// 				asynq.Timeout(3*time.Second),
// 			)
// 			if err != nil {
// 				slog.Error("fail to enqueue task2", slog.Any("err", err))
// 				return
// 			}
// 			slog.Debug("enqueue task2 success", slog.Any("task", taskInfo))
// 		}
// 		id += 1
// 	}
// }

func main() {
	// var app = flag.String("n", "app-main", "app name")
	// flag.Parse()

	logger.Init(slog.LevelInfo, false)

	c := config.Config{
		RedisConfig: config.RedisConfig{
			Addr:      "0.0.0.0:63790",
			Password:  "",
			Username:  "",
			DB:        0,
			EnableTls: false,
		},
	}
	ctx := appctx.New(c)
	svr, err := server.New(ctx)
	if err != nil {
		slog.Error("fail to run server")
		return
	}

	go svr.Start()
	go api.Run()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	slog.Info("received signal SIGINT or SIGTERM, graceful shutdown now")
}
