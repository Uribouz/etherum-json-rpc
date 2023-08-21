package ethclient

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type subscriber struct {
	chDataIn	chan types.Log
	ChDataOut  	chan common.Hash
	ethereum.Subscription
}

func NewHashSubscriber(ctx context.Context, addresses ...string) (client subscriber, err error) {
	commonAddresses := make([]common.Address, len(addresses))
	for i, each := range addresses {
		commonAddresses[i] = common.HexToAddress(each)
	}

	chDataIn := make(chan types.Log)
	sub, err := rpcClient.SubscribeFilterLogs(ctx, ethereum.FilterQuery{ Addresses:commonAddresses }, chDataIn)
	if err != nil {
		return subscriber{}, fmt.Errorf("cannot do SubscribeFilterLogs, %v", err);
	}
	return subscriber{
		chDataIn: 			chDataIn,
		ChDataOut: 			make(chan common.Hash),
		Subscription: 		sub,
	}, nil
}

func (e *subscriber) Subscribe() {
	for {
        select {
        case err := <-e.Err():
            log.Fatal(err)
        case data := <-e.chDataIn:
            log.Printf("receive data: %v\n", data)
			e.ChDataOut <- data.BlockHash
        }
    }
}
