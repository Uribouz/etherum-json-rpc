package main

import (
	"context"
	"etherum-json-rpc/ethclient"
	"etherum-json-rpc/mongodb"
	"fmt"
)

var ADDRESS_HASHED = "0x28c6c06298d514db089934071355e5743bf21d60"
var BLOCK_NUMERS = []int64{17065470, 17065471}

func main() {
    defer ethclient.Close()
    defer mongodb.Close()
    ethRunner, err := ethclient.NewEthRunner(context.Background())
    if err != nil {
        panic(err)
    }
    //Implements
    transactions, err := ethRunner.GetJsonTransactionByBlocks(ADDRESS_HASHED, BLOCK_NUMERS...)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%v\n", transactions)
    _= mongodb.GetDBClient()
}