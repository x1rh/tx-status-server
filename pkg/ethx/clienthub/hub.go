package clienthub

import (
	"errors"
	"tx-status-server/constants"
	"tx-status-server/pkg/ethx/chain"
	"tx-status-server/pkg/ethx/client"
	"tx-status-server/pkg/ethx/constanst"
)

type ClientHub struct {
	ChainConfig map[int]*chain.Config  // map[chainId]ChainConfig
	Clients     map[int]*client.Client // map[chainId]Client
}

func (h *ClientHub) GetChainIdByName(chainName string) (int, error) {
	switch chainName {
	case constanst.ChainNameEthereum:
		return constants.ChainIdEthereum, nil
	case constanst.ChainNameEthereumSepolia:
		return constants.ChainIdEthereumSepolia, nil
	default:
		return 0, errors.New("invalid chainName")
	}
}

func (h *ClientHub) WithChainID(chainId int) (*client.Client, error) {
	c, found := h.Clients[chainId]
	if !found {
		return nil, errors.New("invalid chainId")
	}
	return c, nil
}

func (h *ClientHub) MustWithChainID(chainId int) *client.Client {
	c, _ := h.WithChainID(chainId)
	return c
}
