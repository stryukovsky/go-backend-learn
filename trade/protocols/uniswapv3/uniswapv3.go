package uniswapv3

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UniswapV3PoolInstance struct {
	client   *ethclient.Client
	caller   *UniswapV3PoolCaller
	filterer *UniswapV3PoolFilterer
	Address  common.Address
}

type NFPositionManagerInstance struct {
	client   *ethclient.Client
	caller   *INonFungiblePositionsManagerCaller
	filterer *INonFungiblePositionsManagerFilterer
	Address  common.Address
}

func NewUniswapV3PoolInstance(client *ethclient.Client, address string) (*UniswapV3PoolInstance, error) {
	checksumAddr := common.HexToAddress(address)
	caller, err := NewUniswapV3PoolCaller(checksumAddr, client)
	if err != nil {
		return nil, err
	}
	filterer, err := NewUniswapV3PoolFilterer(checksumAddr, client)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolInstance{
		client:   client,
		caller:   caller,
		filterer: filterer,
		Address:  checksumAddr,
	}, nil
}

func NewNFPositionManagerInstance(client *ethclient.Client, address string) (*NFPositionManagerInstance, error) {
	checksumAddr := common.HexToAddress(address)
	caller, err := NewINonFungiblePositionsManagerCaller(checksumAddr, client)
	if err != nil {
		return nil, err
	}
	filterer, err := NewINonFungiblePositionsManagerFilterer(checksumAddr, client)
	if err != nil {
		return nil, err
	}
	return &NFPositionManagerInstance{
		client:   client,
		caller:   caller,
		filterer: filterer,
		Address:  checksumAddr,
	}, nil
}
