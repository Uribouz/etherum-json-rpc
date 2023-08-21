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
		data, err := e.GetJsonTransaction(address, each)
		if err != nil {
			return nil, fmt.Errorf("cannot do GetJsonTransaction, %v", err)
		}
		result = append(result, data...)
	}
	return result, nil
}

func (e ethTransactionGetter) GetJsonTransaction(address string, blockNo int64) ([]string, error) {
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