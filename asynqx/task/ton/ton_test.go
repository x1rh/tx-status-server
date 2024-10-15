package ton

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"testing"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)

func TestTon(t *testing.T) {
	_ = ""
	client := liteclient.NewConnectionPool()

	testnetURL := "https://ton-blockchain.github.io/testnet-global.config.json"
	mainnetURL := "https://ton.org/global.config.json"
	_ = mainnetURL

	cfg, err := liteclient.GetConfigFromUrl(context.Background(), testnetURL)
	if err != nil {
		log.Fatalln("get config err: ", err.Error())
		return
	}

	// connect to mainnet lite servers
	err = client.AddConnectionsFromConfig(context.Background(), cfg)
	if err != nil {
		log.Fatalln("connection err: ", err.Error())
		return
	}

	// initialize ton api lite connection wrapper with full proof checks
	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()
	api.SetTrustedBlockFromConfig(cfg)

	master, err := api.CurrentMasterchainInfo(context.Background()) // we fetch block just to trigger chain proof check
	if err != nil {
		log.Fatalln("get masterchain info err: ", err.Error())
		return
	}

	// address on which we are accepting payments

	myWalletAddress := "UQCwxyBrDh2FL3MycXF0-XgInMp1aMbND_1SIggGS2cVjAlV" // my wallet
	// myWalletAddress := "0QAsR34m_SDd4fbk-vJCwUWNK17AVuRf152vtbXV9C9Tizlf" // my wallet
	// myWalletAddress := "EQCSES0TZYqcVkgoguhIb8iMEo4cvaEwmIrU5qbQgnN8fmvP" // tester giver bot

	// treasuryAddress := address.MustParseAddr("EQAYqo4u7VF0fa4DPAebk4g9lBytj2VFny7pzXR0trjtXQaO")
	treasuryAddress := address.MustParseAddr(myWalletAddress)

	fmt.Printf("%+v\n", treasuryAddress)

	acc, err := api.GetAccount(context.Background(), master, treasuryAddress)
	if err != nil {
		log.Fatalln("get masterchain info err: ", err.Error())
		return
	}

	// Cursor of processed transaction, save it to your db
	// We start from last transaction, will not process transactions older than we started from.
	// After each processed transaction, save lt to your db, to continue after restart

	lastProcessedLT := acc.LastTxLT
	fmt.Println("lastProcessedLT", lastProcessedLT)

	// channel with new transactions

	// lt := uint64(26725328000003)
	// txHash := "746AA9D08540713EEC53646655ED216F5CACCC6CDA042D6BD865FE6B8A2C4C1C"

	// transactions := make(chan *tlb.Transaction)

	// it is a blocking call, so we start it asynchronously
	// go api.SubscribeOnTransactions(context.Background(), treasuryAddress, lt-1, transactions)

	log.Println("waiting for transfers...")

	// // USDT master contract addr, but can be any jetton
	// usdt := jetton.NewJettonMasterClient(api, address.MustParseAddr("EQCxE6mUtQJKFnGfaROTKOt1lZbDiiX1kCixRv7Nw2Id_sDs"))
	// // get our jetton wallet address

	// treasuryJettonWallet, err := usdt.GetJettonWalletAtBlock(context.Background(), treasuryAddress, master)
	// if err != nil {
	// 	log.Fatalln("get jetton wallet address err: ", err.Error())
	// 	return
	// }

	fmt.Printf("acc.LastTxLt: %v\n", acc.LastTxLT)
	fmt.Printf("acc.LastTxHash: %v\n", hex.EncodeToString(acc.LastTxHash))

	transactionList, err := api.ListTransactions(context.Background(), treasuryAddress, 15, acc.LastTxLT, acc.LastTxHash)
	if err != nil {
		// In some cases you can get error:
		// lite server error, code XXX: cannot compute block with specified transaction: lt not in db
		// it means that current lite server does not store older data, you can query one with full history
		log.Printf("send err: %s", err.Error())
		return
	}

	for _, tx := range transactionList {
		fmt.Printf("transaction: %+v\n", tx)
		fmt.Printf("base64 tx hash: %s\n", base64.StdEncoding.EncodeToString(tx.Hash))
		fmt.Printf("tx hash: %s\n", hex.EncodeToString(tx.Hash))

		/*
			dv, ok := tx.Description.Description.(tlb.TransactionDescriptionOrdinary)
			if !ok {
				fmt.Println("not a tlb.TransactionDescriptionOrdinary")
			}
			fmt.Printf("description: %+v\n", dv)

			if dv.Aborted {
				fmt.Println("transaction is aborted")
			}

			pv, ok := dv.ComputePhase.Phase.(tlb.ComputePhaseVM)
			if !ok {
				fmt.Println("not a tlb.ComputePhaseVM")
			} else {
				fmt.Printf("exitcode=%d\n", pv.Details.ExitCode)
			}

		*/
	}

	// it can happen due to none of available liteservers know old enough state for our address
	// (when our unprocessed transactions are too old)
	log.Println("something went wrong, transaction listening unexpectedly finished")
}
