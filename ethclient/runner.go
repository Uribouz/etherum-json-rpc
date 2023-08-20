package ethclient

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var rpcClient *ethclient.Client
const url = "https://rpc.ankr.com/eth"  // ethereum serviceNode

type ethClient struct {
	ctx context.Context
}

func init() {
	var err error
	rpcClient, err = ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
}

func NewEthClient(ctx context.Context) (client *ethClient,err error) {
	return &ethClient{ctx:ctx},nil
}

func (e ethClient)PrintInfoTransactions(address string, blockNo... int64) {
    for _, each := range blockNo {
		e.PrintInfoTransaction(address, each)
	}
}

func (e ethClient)PrintInfoTransaction(address string, blockNo int64) {
	blockNum := big.NewInt(blockNo)
	blocks, err := rpcClient.BlockByNumber(e.ctx, blockNum)
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
		if strings.EqualFold(from.String(), address) {
			fmt.Printf("Hash: %v, From: %v, To: %v, Value: %v\n", each.Hash(), from.String(), each.To(), each.Value())
		}
		// each.UnmarshalJSON()
	}
}