package ethclient

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
)

type ethTransactionGetter struct {
	ctx context.Context
}
func NewEthTransactionGetter(ctx context.Context) (client *ethTransactionGetter,err error) {
	return &ethTransactionGetter{ctx:ctx},nil
}

func (e ethTransactionGetter) GetJsonTransactionByBlocks(address string, blockNo... int64) ([]string, error) {
	var result []string
    for _, each := range blockNo {
		blockNumber, err := rpcClient.BlockByNumber(e.ctx,  big.NewInt(each))
		if err != nil {
			return nil, fmt.Errorf("cannot do BlockByHash, %v", err)
		}
		data, err := e.GetJsonTransaction(address, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("cannot do GetJsonTransaction, %v", err)
		}
		result = append(result, data...)
	}
	return result, nil
}

func (e ethTransactionGetter) GetJsonTransaction(address string, block *types.Block) ([]string, error) {
	var result []string
	for _, each := range block.Transactions() {
		from, err := types.Sender(types.LatestSignerForChainID(each.ChainId()), each)
		if err != nil {
			return nil, fmt.Errorf("cannot do Sender, %v", err)
		}
		if (!strings.EqualFold(from.String(), address) &&
		 !strings.EqualFold(each.To().String(), address) ) {
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