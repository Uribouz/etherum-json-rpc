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

func init() {
	var err error
	rpcClient, err = ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("EthClient: Connection established.")
}

func Close() {
	if rpcClient != nil {
		rpcClient.Close()
	}
}

type ethRunner struct {
	ctx context.Context
}
func NewEthRunner(ctx context.Context) (client *ethRunner,err error) {
	return &ethRunner{ctx:ctx},nil
}

func (e ethRunner)PrintInfoTransactions(address string, blockNo... int64) {
    for _, each := range blockNo {
		e.PrintInfoTransaction(address, each)
	}
}

func (e ethRunner)PrintInfoTransaction(address string, blockNo int64) {
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
	}
}

func (e ethRunner) GetJsonTransactionByBlocks(address string, blockNo... int64) ([]string, error) {
	var result []string
    for _, each := range blockNo {
		data, err := e.GetJsonTransaction(address, each)
		if err != nil {
			return nil, fmt.Errorf("cannot do GetJsonTransaction, %v", err)
		}
		result = append(result, data...)
	}
	return result, nil
}

func (e ethRunner) GetJsonTransaction(address string, blockNo int64) ([]string, error) {
	blockNum := big.NewInt(blockNo)
	blocks, err := rpcClient.BlockByNumber(e.ctx, blockNum)
	if err != nil {
		return nil, fmt.Errorf("cannot do BlockByHash, %v", err)
	}
	trans := blocks.Transactions()
	var result []string
	for _, each := range trans {
		from, err := types.Sender(types.LatestSignerForChainID(each.ChainId()), each)
		if err != nil {
			return nil, fmt.Errorf("cannot do Sender, %v", err)
		}
		if !strings.EqualFold(from.String(), address) {
			continue
		}
		data, err := each.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("cannot do MarshalJSON, %v", err) 
		}
		result = append(result, string(data))	
	}
	return result, nil
}