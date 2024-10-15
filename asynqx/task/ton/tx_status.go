package ton

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
)

func DefaultMainnetGetTransaction(lt uint64, address, txHash string) (*tlb.Transaction, error) {
	mainnet := "https://ton.org/global.config.json"
	return GetTransaction(mainnet, lt, address, txHash)
}

func DefaultTestnetGetTransaction(lt uint64, address, txHash string) (*tlb.Transaction, error) {
	testnet := "https://ton-blockchain.github.io/testnet-global.config.json"
	return GetTransaction(testnet, lt, address, txHash)
}

func GetTransaction(url string, lt uint64, walletAddress, txHash string) (*tlb.Transaction, error) {
	client := liteclient.NewConnectionPool()
	cfg, err := liteclient.GetConfigFromUrl(context.Background(), url)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get config")
	}

	err = client.AddConnectionsFromConfig(context.Background(), cfg)
	if err != nil {
		return nil, errors.Wrap(err, "connection error")
	}

	// initialize ton api lite connection wrapper with full proof checks
	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()
	api.SetTrustedBlockFromConfig(cfg)

	wa := address.MustParseAddr(walletAddress)
	txb, err := hex.DecodeString(txHash)
	if err != nil {
		return nil, errors.Wrap(err, "fail to decode tx hash to []byte")
	}

	// list 5 transaction to find the specific tx
	transactionList, err := api.ListTransactions(context.Background(), wa, 5, lt, txb)
	if err != nil {
		return nil, errors.Wrap(err, "fail to list transactions")
	}

	for _, tx := range transactionList {
		hexTxHash := hex.EncodeToString(tx.Hash)
		fmt.Printf("tx hash: %s\n", hexTxHash)
		if strings.EqualFold(hexTxHash, txHash) {
			return tx, nil
		}
	}
	return nil, fmt.Errorf("fail to find the tx: %s", txHash)
}
