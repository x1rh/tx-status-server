package client

import (
	"tx-status-server/pkg/ethx/chain"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	ChainConfig chain.Config
	Client      *ethclient.Client
}
