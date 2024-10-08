package client

import "tx-status-server/sdk/types"

type IClient interface {
	ChainConfig() map[string]types.ChainConfig
	Put(*types.PutReq) (*types.PutResp, error)
	Get(*types.GetReq) (*types.GetResp, error)
}
