package ethclient

import (
	"etherum-json-rpc/config"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

var rpcClient *ethclient.Client

func init() {
	var err error
	rpcClient, err = ethclient.Dial(config.GetEthNodeUrl())
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