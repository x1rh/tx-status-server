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

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("tx-status-server")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./etc")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var c config.Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
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
