package main

import (
	"context"
	"etherum-json-rpc/ethclient"
	"etherum-json-rpc/mongodb"
)

var ADDRESS_HASHED = []string{"0x28c6c06298d514db089934071355e5743bf21d60"}
var BLOCK_NUMERS = []int64{17065470, 17065471}
const databaseName = "ethereum-block"

func main() {
    defer ethclient.Close()
    defer mongodb.Close()
    ctx := context.Background()
    ethRunner, err := ethclient.NewEthTransactionGetter(ctx)
    if err != nil {
        panic(err)
    }
    //Implements
    transactions, err := ethRunner.GetJsonTransactionByBlocks(ADDRESS_HASHED[0], BLOCK_NUMERS...)
    if err != nil {
        panic(err)
    }
    // fmt.Printf("%v\n", strings.Join(transactions,","))
    inserter := mongodb.NewInserter(databaseName, ctx)
    if err := inserter.InsertJsonDataTransactions(transactions); err != nil {
        panic(err)
    }
}