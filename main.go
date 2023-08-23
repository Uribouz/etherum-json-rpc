package main

import (
	"context"
	"etherum-json-rpc/adapter"
	"etherum-json-rpc/config"
	"etherum-json-rpc/core"
	"etherum-json-rpc/ethclient"
	"etherum-json-rpc/mockdata"
	"etherum-json-rpc/mongodb"
	"log"
)


func main() {

    //Initialize
    config.Init()
    ethclient.Init()
    defer ethclient.Close()
    mongodb.Init()
    defer mongodb.Close()

    //Get Datasource
    dataSource := mockdata.NewFileReader()
    data, err := adapter.JsonDataToAddresses(dataSource)
    if err != nil {
        log.Fatal(err);
    }

    //Running
    ctx := context.Background()
    core.DoMultipleSubscribeAddress(ctx, data...)
}