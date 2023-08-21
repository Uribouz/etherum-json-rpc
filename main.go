package main

import (
	"etherum-json-rpc/core"
	"etherum-json-rpc/ethclient"
	"etherum-json-rpc/mongodb"
)

var ADDRESS_HASHED = "0x28c6c06298d514db089934071355e5743bf21d60"
var BLOCK_NUMERS = []int64{17065470, 17065471}

func main() {
    defer ethclient.Close()
    defer mongodb.Close()
    core.DoReadAndInsertTransaction(ADDRESS_HASHED, BLOCK_NUMERS...)
}