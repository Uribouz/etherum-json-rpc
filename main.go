package main

import (
	"context"
	"etherum-json-rpc/ethclient"
	"etherum-json-rpc/mongodb"
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
    _= mongodb.GetDBClient()
    ethRunner.PrintInfoTransactions(ADDRESS_HASHED, BLOCK_NUMERS...)
}