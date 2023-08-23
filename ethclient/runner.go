package ethclient

import (
	"etherum-json-rpc/config"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

var rpcClient *ethclient.Client
var once sync.Once
func Init() {
	once.Do(func() {
		var err error
		rpcClient, err = ethclient.Dial(config.GetEthNodeUrl())
		if err != nil {
			panic(err)
		}
		fmt.Println("EthClient: Connection established.")
	})
}

func Close() {
	if rpcClient != nil {
		rpcClient.Close()
	}
}