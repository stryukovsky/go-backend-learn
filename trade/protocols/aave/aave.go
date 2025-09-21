package aave

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stryukovsky/go-backend-learn/trade"
)

type AavePool struct {
	client  *ethclient.Client
	caller  *PoolCaller
	filterer  *PoolFilterer
	Address common.Address
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
		client:  client,
		caller:  caller,
		filterer: filterer,
		Address: checksumAddr,
	}, nil
}

func parseSupplyEvents(event *PoolSupplyIterator) ([]trade.AaveEvent, error){
	result := make([]trade.AaveEvent, 5)
	for event.Next() {
		err := event.Error()
		if err != nil {
			return nil, err
		}
		item := trade.NewAaveEvent("supply", event.Event.OnBehalfOf, event.Event.Reserve.Big())
		result = append(result, item)
	}
	return result, nil
}

func parseWithdrawEvents(event *PoolWithdrawIterator) ([]trade.AaveEvent, error){
	result := make([]trade.AaveEvent, 5)
	for event.Next() {
		err := event.Error()
		if err != nil {
			return nil, err
		}
		item := trade.NewAaveEvent("withdraw", event.Event.To, event.Event.Reserve.Big())
		result = append(result, item)
	}
	return result, nil
}

func (a *AavePool) GetInvestmentsByParticipants(participants []string) ([]trade.AaveEvent, error){
	formattedParticipants := make([]common.Address, len(participants))
	for i, p := range participants {
		formattedParticipants[i] = common.HexToAddress(p)
	}
	supplyEventsRaw, err := a.filterer.FilterSupply(&bind.FilterOpts{}, []common.Address{}, formattedParticipants, []uint16{})
	if err != nil {
		return nil, err
	}
	withdrawEventsRaw, err := a.filterer.FilterWithdraw(&bind.FilterOpts{}, []common.Address{}, []common.Address{}, formattedParticipants)
	if err != nil {
		return nil, err
	}
	defer supplyEventsRaw.Close()
	defer withdrawEventsRaw.Close()
	supplyEvents, err := parseSupplyEvents(supplyEventsRaw)
	if err != nil {
		return nil, err
	}
	withdrawEvents, err := parseWithdrawEvents(withdrawEventsRaw)
	if err != nil {
		return nil, err
	}
	return append(supplyEvents, withdrawEvents...), nil
}
