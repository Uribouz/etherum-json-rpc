package main

import (
	"context"
	"etherum-json-rpc/ethclient"
	"etherum-json-rpc/mongodb"
)

var ADDRESS_HASHED = "0x28c6c06298d514db089934071355e5743bf21d60"
var BLOCK_NUMERS = []int64{17065470, 17065471}

func main() {
    _= mongodb.GetDBClient()
    client, err := ethclient.NewEthClient(context.Background())
    if err != nil {
        panic(err)
    }
    client.PrintInfoTransactions(ADDRESS_HASHED, BLOCK_NUMERS...)


}