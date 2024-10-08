package appctx

import (
	"tx-status-server/asynqx/client"
	"tx-status-server/config"
)

var defaultAppCtx *Context

type Context struct {
	Config     config.Config
	TaskClient *client.Client
}

func New(c config.Config) *Context {
	return &Context{
		Config:     c,
		TaskClient: client.New(c),
	}
}

func GetContext() *Context {
	return defaultAppCtx
}
