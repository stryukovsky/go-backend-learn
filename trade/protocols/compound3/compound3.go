package compound3

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/web3client"
)

// --- Caller wrappers ---

type CometCallerWithURL struct {
	Caller *CometCaller
	Url    string
}

func (c *CometCallerWithURL) URL() string { return c.Url }

type MultiURLCometCaller struct {
	callers []*CometCallerWithURL
}

// Helper to get BaseToken from any working provider
func (m *MultiURLCometCaller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	return trade.RetryEthCall(
		func() []*CometCallerWithURL { return m.callers },
		func(caller *CometCallerWithURL) (common.Address, error) {
			return caller.Caller.BaseToken(opts)
		})
}

// --- Filterer wrappers ---

type CometFiltererWithURL struct {
	filterer *CometFilterer
	url      string
}

func (f *CometFiltererWithURL) URL() string { return f.url }

type MultiURLCometFilterer struct {
	filterers []*CometFiltererWithURL
}

func (m *MultiURLCometFilterer) FilterSupply(
	opts *bind.FilterOpts,
	from []common.Address,
	dst []common.Address,
) (*CometSupplyIterator, error) {
	return trade.RetryEthCall(
		func() []*CometFiltererWithURL { return m.filterers },
		func(filterer *CometFiltererWithURL) (*CometSupplyIterator, error) {
			return filterer.filterer.FilterSupply(opts, from, dst)
		})
}

func (m *MultiURLCometFilterer) FilterSupplyCollateral(
	opts *bind.FilterOpts,
	from []common.Address,
	dst []common.Address,
	asset []common.Address,
) (*CometSupplyCollateralIterator, error) {
	return trade.RetryEthCall(
		func() []*CometFiltererWithURL { return m.filterers },
		func(filterer *CometFiltererWithURL) (*CometSupplyCollateralIterator, error) {
			return filterer.filterer.FilterSupplyCollateral(opts, from, dst, asset)
		})
}

func (m *MultiURLCometFilterer) FilterWithdraw(
	opts *bind.FilterOpts,
	src []common.Address,
	to []common.Address,
) (*CometWithdrawIterator, error) {
	return trade.RetryEthCall(
		func() []*CometFiltererWithURL { return m.filterers },
		func(filterer *CometFiltererWithURL) (*CometWithdrawIterator, error) {
			return filterer.filterer.FilterWithdraw(opts, src, to)
		})
}

func (m *MultiURLCometFilterer) FilterWithdrawCollateral(
	opts *bind.FilterOpts,
	src []common.Address,
	to []common.Address,
	asset []common.Address,
) (*CometWithdrawCollateralIterator, error) {
	return trade.RetryEthCall(
		func() []*CometFiltererWithURL { return m.filterers },
		func(filterer *CometFiltererWithURL) (*CometWithdrawCollateralIterator, error) {
			return filterer.filterer.FilterWithdrawCollateral(opts, src, to, asset)
		})
}

// --- Main Compound3 struct with multi-URL support ---

type Compound3 struct {
	client      *web3client.MultiURLClient
	caller      *MultiURLCometCaller
	filterer    *MultiURLCometFilterer
	Address     common.Address
	MainAsset   common.Address
}

func NewCompound3(client *web3client.MultiURLClient, address string) (*Compound3, error) {
	checksumAddr := common.HexToAddress(address)

	// Build callers
	callers := make([]*CometCallerWithURL, client.Length())
	for i, clientWithURL := range client.Iter() {
		caller, err := NewCometCaller(checksumAddr, clientWithURL.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to create CometCaller: %w. URL: %s", err, clientWithURL.Url)
		}
		callers[i] = &CometCallerWithURL{Caller: caller, Url: clientWithURL.Url}
	}

	// Build filterers
	filterers := make([]*CometFiltererWithURL, client.Length())
	for i, clientWithURL := range client.Iter() {
		filterer, err := NewCometFilterer(checksumAddr, clientWithURL.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to create CometFilterer: %w. URL: %s", err, clientWithURL.Url)
		}
		filterers[i] = &CometFiltererWithURL{filterer: filterer, url: clientWithURL.Url}
	}

	// Create multi-url wrappers
	multiCaller := &MultiURLCometCaller{callers: callers}
	multiFilterer := &MultiURLCometFilterer{filterers: filterers}

	// Fetch base token (MainAsset) using retry logic
	mainAsset, err := multiCaller.BaseToken(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch BaseToken for Comet at %s: %w", address, err)
	}

	return &Compound3{
		client:      client,
		caller:      multiCaller,
		filterer:    multiFilterer,
		Address:     checksumAddr,
		MainAsset:   mainAsset,
	}, nil
}
