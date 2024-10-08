package types

type ChainConfig struct {
	ChainKind int // 1:ethereum|2:solana|3:ton
	ChainName string
	ChainId   int // 0 if not exists
	URL       string
}
