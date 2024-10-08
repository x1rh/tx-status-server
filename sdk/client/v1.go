package client

import (
	"tx-status-server/sdk/types"
)

type ClientV1 struct {
	URL string
}

func (c *ClientV1) ChainConfig() map[string]types.ChainConfig {
	return nil
}

func (c *ClientV1) Put(req *types.PutReq) (*types.PutResp, error) {
	return nil, nil
}

func (c *ClientV1) Get(req *types.GetReq) (*types.GetResp, error) {
	return nil, nil
}

func NewV1() *ClientV1 {
	return &ClientV1{}
}
