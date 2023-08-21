package core

import (
	"context"
	"etherum-json-rpc/ethclient"
	"etherum-json-rpc/mongodb"
	"log"
)

func DoSubscribeAddress(address string) {
	ctx := context.Background()
    subscriber, err := ethclient.NewHashSubscriber(ctx, address)
	if err != nil {
		log.Fatal(err)
	}
    go subscriber.Subscribe()
	defer subscriber.Unsubscribe()
	ethRunner, err := ethclient.NewEthTransactionGetter(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//Implements
	for data := range subscriber.ChDataOut {
		transactions, err := ethRunner.GetJsonTransactionByHash(address, data)
		if err != nil {
			log.Fatal(err)
		}
		inserter := mongodb.NewInserter(DATABASE_NAME, ctx)
		if err := inserter.InsertJsonDataTransactions(transactions); err != nil {
			log.Fatal(err)
		}
	}
}