package core

import (
	"context"
	"etherum-json-rpc/ethclient"
	"etherum-json-rpc/mongodb"
	"fmt"
)

func DoMultipleSubsctibeAddress(parentCtx context.Context, address... string) error {
	for i:=0; i<= WORKER_TOTAL_NUM; i++ {

	}
	return nil
}

func DoSubscribeAddress(parentCtx context.Context, address... string) error {
	ctx, ctxCancel := context.WithCancel(parentCtx)
	defer ctxCancel()
    subscriber, err := ethclient.NewHashSubscriber(ctx, address...)
	if err != nil {
		return fmt.Errorf("cannot NewHashSubscriber, %v", err)
	}
    go subscriber.Subscribe()
	defer subscriber.Unsubscribe()
	ethRunner, err := ethclient.NewEthTransactionGetter(ctx)
	if err != nil {
		return fmt.Errorf("cannot NewEthTransactionGetter, %v", err)
	}

	inserter := mongodb.NewInserter(DATABASE_NAME, ctx)
	for data := range subscriber.ChDataOut {
		transactions, err := ethRunner.GetJsonTransactionByHash(data.Address.String(), data.Hash)
		if err != nil {
			return fmt.Errorf("cannot GetJsonTransactionByHash, %v", err)
		}
		if err := inserter.InsertJsonDataTransactions(transactions); err != nil {
			return fmt.Errorf("cannot InsertJsonDataTransactions, %v", err)
		}
	}
	return nil
}