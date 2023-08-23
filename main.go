package main

import (
	"context"
	"etherum-json-rpc/adapter"
	"etherum-json-rpc/core"
	"etherum-json-rpc/ethclient"
	"etherum-json-rpc/mockdata"
	"etherum-json-rpc/mongodb"
	"log"
)

// var ADDRESS_HASHED = "0x28c6c06298d514db089934071355e5743bf21d60"
var BLOCK_NUMERS = []int64{17065470, 17065471}

func main() {
    defer ethclient.Close()
    defer mongodb.Close()
    dataSource := mockdata.NewFileReader()
    data, err := adapter.JsonDataToAddresses(dataSource)
    if err != nil {
        log.Fatal(err);
    }
    ctx := context.Background()
    // core.DoReadAndInsertTransaction(ctx, data[0], BLOCK_NUMERS...)
    core.DoMultipleSubscribeAddress(ctx, data...)
}