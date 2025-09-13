package trade

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestGetEvents(t *testing.T) {
	fmt.Println("Test eth_getLogs()")
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}
	result, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(23353060),
		ToBlock: big.NewInt(23353465),
		Addresses: []common.Address{common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")},
		Topics: [][]common.Hash{
			{common.HexToHash(TransferTopic)},
			{},
			{common.HexToHash("0xc7bBeC68d12a0d1830360F8Ec58fA599bA1b0e9b")},
			{},
		},
	})
	fmt.Println(common.HexToHash("0xc7bBeC68d12a0d1830360F8Ec58fA599bA1b0e9b").Hex())
	if err != nil {
		panic(err)
	}
	fmt.Println("Logs:", len(result))
	for _, entry := range result {
		fmt.Println(entry.Topics)

	}

}
