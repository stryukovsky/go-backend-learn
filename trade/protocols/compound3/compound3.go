package compound3

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Compound3 struct {
	client   *ethclient.Client
	caller   *CometCaller
	filterer *CometFilterer
	Address  common.Address
	MainAsset common.Address
}

func NewCompound3(client *ethclient.Client, address string) (*Compound3, error) {
	checksumAddr := common.HexToAddress(address)
	caller, err := NewCometCaller(checksumAddr, client)
	if err != nil {
		return nil, err
	}
	mainAsset, err := caller.BaseToken(&bind.CallOpts{});
	if err != nil {
		return nil, err
	}
	filterer, err := NewCometFilterer(checksumAddr, client)
	if err != nil {
		return nil, err
	}
	return &Compound3{
		caller:   caller,
		filterer: filterer,
		client:   client,
		Address:  checksumAddr,
		MainAsset: mainAsset,
	}, nil
}
