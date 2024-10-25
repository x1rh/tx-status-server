package config

import (
	"github.com/x1rh/web3go/ethx/chain"
)

type Config struct {
	RedisConfig RedisConfig
	ChainConfig []chain.Config
}

type RedisConfig struct {
	Addr      string
	Password  string
	Username  string
	DB        int
	EnableTls bool
}
