package core

import (
	"context"
	"etherum-json-rpc/ethclient"
	"etherum-json-rpc/mongodb"
	"log"
)


func DoReadAndInsertTransaction(address string, blockNo ...int64) {
	ctx := context.Background()
	ethRunner, err := ethclient.NewEthTransactionGetter(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//Implements
	transactions, err := ethRunner.GetJsonTransactionByBlocks(address, blockNo...)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%v\n", strings.Join(transactions,","))
	inserter := mongodb.NewInserter(DATABASE_NAME, ctx)
	if err := inserter.InsertJsonDataTransactions(transactions); err != nil {
		log.Fatal(err)
	}
}