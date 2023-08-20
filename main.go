package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ADDRESS_HASHED = "0x28c6c06298d514db089934071355e5743bf21d60"
var BLOCK_NUMERS = []int64{17065470, 17065471}

func main() {
    const url = "https://rpc.ankr.com/eth"  // url string
    
    rpcClient,err := ethclient.Dial(url)
    if err != nil {
        panic(err)
    }
    ctx := context.Background()
    printInfoTransactions := func (blockNo int64) {
        blockNum := big.NewInt(blockNo)
        blocks, err := rpcClient.BlockByNumber(ctx, blockNum)
        if err != nil {
            fmt.Printf("cannot do BlockByHash, %v", err)
            return
        }
        trans := blocks.Transactions()
        for _, each := range trans {
            from, err := types.Sender(types.LatestSignerForChainID(each.ChainId()), each)   
            if err != nil {
                fmt.Printf("cannot do Sender, %v", err)
                return
            }
            if strings.EqualFold(from.String(),ADDRESS_HASHED) {
                fmt.Printf("Hash: %v, From: %v, To: %v, Value: %v\n", each.Hash(), from.String(), each.To(), each.Value())
            }
            // each.UnmarshalJSON()
        }
    }
    for _, each := range BLOCK_NUMERS {
        printInfoTransactions(each)
    }
}