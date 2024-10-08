package types

type PutReq struct {
	ChainKind int
	ChainName string
	Tx        string
}

type PutResp struct {
}

type GetReq struct {
	ChainKind int
	ChainName string
	Tx        string
}

type GetResp struct {
}
