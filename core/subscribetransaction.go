package core

import (
	"context"
	"etherum-json-rpc/chunker"
	"etherum-json-rpc/config"
	"etherum-json-rpc/ethclient"
	"etherum-json-rpc/mongodb"
	"fmt"
	"log"
	"strings"
	"sync"
)

func DoMultipleSubscribeAddress(ctx context.Context, address... string) error {
	chunkAddresses := chunker.Chunk(config.GetWorkerTotalNum(), address)
	var wg sync.WaitGroup
	for i, each :=range chunkAddresses {
		wg.Add(1)
		go func(idx int, data []string) {
			defer wg.Done()
			log.Printf("Start worker no. %d, with addresses: %s\n", idx, strings.Join(data,","))
			err := DoSubscribeAddress(ctx, data...)
			if err != nil {
				log.Printf("worker no. %d stopped, got an error: %s\n", idx, err.Error())
			}
		}(i, each)
	}
	wg.Wait()
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

	inserter := mongodb.NewInserter(config.GetDatabaseName(), ctx)
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