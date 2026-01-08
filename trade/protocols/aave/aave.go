package aave

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/web3client"
)

type AaveCallerWithURL struct {
	Caller *PoolCaller
	Url    string
}

func (c *AaveCallerWithURL) URL() string { return c.Url }

type MultiURLAaveCaller struct {
	callers []*AaveCallerWithURL
}

// Add any caller methods you need here (e.g., getReserveData, getUserAccountData, etc.)
// For now, since your use case seems focused on events, we'll leave it minimal unless needed.

// --- Filterer wrappers ---

type AaveFiltererWithURL struct {
	filterer *PoolFilterer
	url      string
}

func (f *AaveFiltererWithURL) URL() string { return f.url }

type MultiURLAaveFilterer struct {
	filterers []*AaveFiltererWithURL
}

func (m *MultiURLAaveFilterer) FilterSupply(
	opts *bind.FilterOpts,
	reserve []common.Address,
	onBehalfOf []common.Address,
	referralCode []uint16,
) (*PoolSupplyIterator, error) {
	return trade.RetryEthCall(
		func() []*AaveFiltererWithURL { return m.filterers },
		func(filterer *AaveFiltererWithURL) (*PoolSupplyIterator, error) {
			return filterer.filterer.FilterSupply(opts, reserve, onBehalfOf, referralCode)
		})
}

func (m *MultiURLAaveFilterer) FilterWithdraw(
	opts *bind.FilterOpts,
	reserve []common.Address,
	to []common.Address,
	repayFromAToken []common.Address,
) (*PoolWithdrawIterator, error) {
	return trade.RetryEthCall(
		func() []*AaveFiltererWithURL { return m.filterers },
		func(filterer *AaveFiltererWithURL) (*PoolWithdrawIterator, error) {
			return filterer.filterer.FilterWithdraw(opts, reserve, to, repayFromAToken)
		})
}

// --- Main AavePool struct using multi-url clients ---

type AavePool struct {
	client   *web3client.MultiURLClient
	caller   *MultiURLAaveCaller
	filterer *MultiURLAaveFilterer
	Address  common.Address
}

func NewAavePool(client *web3client.MultiURLClient, address string) (*AavePool, error) {
	checksumAddr := common.HexToAddress(address)

	// Build multi-url callers
	callers := make([]*AaveCallerWithURL, client.Length())
	for i, clientWithURL := range client.Iter() {
		caller, err := NewPoolCaller(checksumAddr, clientWithURL.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to create PoolCaller: %w. URL: %s", err, clientWithURL.Url)
		}
		callers[i] = &AaveCallerWithURL{Caller: caller, Url: clientWithURL.Url}
	}

	// Build multi-url filterers
	filterers := make([]*AaveFiltererWithURL, client.Length())
	for i, clientWithURL := range client.Iter() {
		filterer, err := NewPoolFilterer(checksumAddr, clientWithURL.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to create PoolFilterer: %w. URL: %s", err, clientWithURL.Url)
		}
		filterers[i] = &AaveFiltererWithURL{filterer: filterer, url: clientWithURL.Url}
	}

	return &AavePool{
		client:   client,
		caller:   &MultiURLAaveCaller{callers},
		filterer: &MultiURLAaveFilterer{filterers},
		Address:  checksumAddr,
	}, nil
}
