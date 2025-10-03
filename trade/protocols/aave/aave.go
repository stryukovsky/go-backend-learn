package aave

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AavePool struct {
	client   *ethclient.Client
	caller   *PoolCaller
	filterer *PoolFilterer
	Address  common.Address
}

func NewAavePool(client *ethclient.Client, address string) (*AavePool, error) {
	checksumAddr := common.HexToAddress(address)
	caller, err := NewPoolCaller(checksumAddr, client)
	if err != nil {
		return nil, err
	}
	filterer, err := NewPoolFilterer(checksumAddr, client)
	if err != nil {
		return nil, err
	}
	return &AavePool{
		client:   client,
		caller:   caller,
		filterer: filterer,
		Address:  checksumAddr,
	}, nil
}

