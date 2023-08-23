package core

import (
	"context"
	"etherum-json-rpc/config"
	"etherum-json-rpc/ethclient"
	"etherum-json-rpc/mongodb"
	"log"
)

//DoReadAndInsertTransaction usage example
// var ADDRESS_HASHED = "0x28c6c06298d514db089934071355e5743bf21d60"
// var BLOCK_NUMERS = []int64{17065470, 17065471}
// core.DoReadAndInsertTransaction(ctx, ADDRESS_HASHED, BLOCK_NUMERS...)

func DoReadAndInsertTransaction(ctx context.Context, address string, blockNo ...int64) {
	ethRunner, err := ethclient.NewEthTransactionGetter(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//Implements
	transactions, err := ethRunner.GetJsonTransactionByBlocks(address, blockNo...)
	if err != nil {
		log.Fatal(err)
	}
	inserter := mongodb.NewInserter(config.GetDatabaseName(), ctx)
	if err := inserter.InsertJsonDataTransactions(transactions); err != nil {
		log.Fatal(err)
	}
}