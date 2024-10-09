package ton

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
)

func TestTon(t *testing.T) {
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

	transactions := make(chan *tlb.Transaction)

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
	}

	fmt.Println("======================")

	// listen for new transactions from channel

	_ = transactions

	for tx := range transactions {
		// only internal messages can increase the balance
		// if tx.IO.In != nil && tx.IO.In.MsgType == tlb.MsgTypeInternal {
		// 	ti := tx.IO.In.AsInternal()
		// 	src := ti.SrcAddr

		// 	// verify that event sender is our jetton wallet
		// 	if ti.SrcAddr.Equals(treasuryJettonWallet.Address()) {
		// 		var transfer jetton.TransferNotification
		// 		if err = tlb.LoadFromCell(&transfer, ti.Body.BeginParse()); err == nil {
		// 			// convert decimals to 6 for USDT (it can be fetched from jetton details too), default is 9
		// 			amt := tlb.MustFromNano(transfer.Amount.Nano(), 6)

		// 			// reassign sender to real jetton sender instead of its jetton wallet contract
		// 			src = transfer.Sender
		// 			log.Println("received", amt.String(), "USDT from", src.String())
		// 		}
		// 	}

		// 	// show received ton amount
		// 	log.Println("received", ti.Amount.String(), "TON from", src.String())
		// }

		// update last processed lt and save it in db

		// lastProcessedLT = tx.LT
		fmt.Printf("%+v\n", tx)
	}

	// it can happen due to none of available liteservers know old enough state for our address
	// (when our unprocessed transactions are too old)
	log.Println("something went wrong, transaction listening unexpectedly finished")
}
