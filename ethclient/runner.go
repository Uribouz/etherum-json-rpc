package ethclient

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

var rpcClient *ethclient.Client
//TDOO: use wss client instead.
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