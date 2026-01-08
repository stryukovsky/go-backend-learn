package uniswapv3

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/web3client"
)

// =============== Uniswap V3 Pool ===============

type UniswapV3PoolCallerWithURL struct {
	Caller *UniswapV3PoolCaller
	Url    string
}

func (c *UniswapV3PoolCallerWithURL) URL() string { return c.Url }
func (m *MultiURLUniswapV3PoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	return trade.RetryEthCall(
		func() []*UniswapV3PoolCallerWithURL { return m.callers },
		func(f *UniswapV3PoolCallerWithURL) (common.Address, error) {
			return f.Caller.Token0(opts)
		})
}

func (m *MultiURLUniswapV3PoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	return trade.RetryEthCall(
		func() []*UniswapV3PoolCallerWithURL { return m.callers },
		func(f *UniswapV3PoolCallerWithURL) (common.Address, error) {
			return f.Caller.Token1(opts)
		})
}

type MultiURLUniswapV3PoolCaller struct {
	callers []*UniswapV3PoolCallerWithURL
}

type UniswapV3PoolFiltererWithURL struct {
	filterer *UniswapV3PoolFilterer
	url      string
}

func (f *UniswapV3PoolFiltererWithURL) URL() string { return f.url }

type MultiURLUniswapV3PoolFilterer struct {
	filterers []*UniswapV3PoolFiltererWithURL
}

func (m *MultiURLUniswapV3PoolFilterer) FilterMint(
	opts *bind.FilterOpts,
	owner []common.Address,
	tickLower []*big.Int,
	tickUpper []*big.Int,
) (*UniswapV3PoolMintIterator, error) {
	return trade.RetryEthCall(
		func() []*UniswapV3PoolFiltererWithURL { return m.filterers },
		func(f *UniswapV3PoolFiltererWithURL) (*UniswapV3PoolMintIterator, error) {
			return f.filterer.FilterMint(opts, owner, tickLower, tickUpper)
		})
}

func (m *MultiURLUniswapV3PoolFilterer) FilterBurn(
	opts *bind.FilterOpts,
	owner []common.Address,
	tickLower []*big.Int,
	tickUpper []*big.Int,
) (*UniswapV3PoolBurnIterator, error) {
	return trade.RetryEthCall(
		func() []*UniswapV3PoolFiltererWithURL { return m.filterers },
		func(f *UniswapV3PoolFiltererWithURL) (*UniswapV3PoolBurnIterator, error) {
			return f.filterer.FilterBurn(opts, owner, tickLower, tickUpper)
		})
}

func (m *MultiURLUniswapV3PoolFilterer) FilterSwap(
	opts *bind.FilterOpts,
	sender []common.Address,
	recipient []common.Address,
) (*UniswapV3PoolSwapIterator, error) {
	return trade.RetryEthCall(
		func() []*UniswapV3PoolFiltererWithURL { return m.filterers },
		func(f *UniswapV3PoolFiltererWithURL) (*UniswapV3PoolSwapIterator, error) {
			return f.filterer.FilterSwap(opts, sender, recipient)
		})
}

type UniswapV3PoolInstance struct {
	client   *web3client.MultiURLClient
	caller   *MultiURLUniswapV3PoolCaller
	filterer *MultiURLUniswapV3PoolFilterer
	Address  common.Address
}

func NewUniswapV3PoolInstance(client *web3client.MultiURLClient, address string) (*UniswapV3PoolInstance, error) {
	checksumAddr := common.HexToAddress(address)

	callers := make([]*UniswapV3PoolCallerWithURL, client.Length())
	for i, clientWithURL := range client.Iter() {
		caller, err := NewUniswapV3PoolCaller(checksumAddr, clientWithURL.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to create UniswapV3PoolCaller: %w. URL: %s", err, clientWithURL.Url)
		}
		callers[i] = &UniswapV3PoolCallerWithURL{Caller: caller, Url: clientWithURL.Url}
	}

	filterers := make([]*UniswapV3PoolFiltererWithURL, client.Length())
	for i, clientWithURL := range client.Iter() {
		filterer, err := NewUniswapV3PoolFilterer(checksumAddr, clientWithURL.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to create UniswapV3PoolFilterer: %w. URL: %s", err, clientWithURL.Url)
		}
		filterers[i] = &UniswapV3PoolFiltererWithURL{filterer: filterer, url: clientWithURL.Url}
	}

	return &UniswapV3PoolInstance{
		client:   client,
		caller:   &MultiURLUniswapV3PoolCaller{callers},
		filterer: &MultiURLUniswapV3PoolFilterer{filterers},
		Address:  checksumAddr,
	}, nil
}

// =============== Non-Fungible Position Manager ===============

type NFPositionManagerCallerWithURL struct {
	Caller *INonFungiblePositionsManagerCaller
	Url    string
}

func (c *NFPositionManagerCallerWithURL) URL() string { return c.Url }

type MultiURLNFPositionManagerCaller struct {
	callers []*NFPositionManagerCallerWithURL
}

type NFPositionManagerFiltererWithURL struct {
	filterer *INonFungiblePositionsManagerFilterer
	url      string
}

func (f *NFPositionManagerFiltererWithURL) URL() string { return f.url }

type MultiURLNFPositionManagerFilterer struct {
	filterers []*NFPositionManagerFiltererWithURL
}

func (m *MultiURLNFPositionManagerFilterer) FilterIncreaseLiquidity(
	opts *bind.FilterOpts,
	tokenId []*big.Int,
) (*INonFungiblePositionsManagerIncreaseLiquidityIterator, error) {
	return trade.RetryEthCall(
		func() []*NFPositionManagerFiltererWithURL { return m.filterers },
		func(f *NFPositionManagerFiltererWithURL) (*INonFungiblePositionsManagerIncreaseLiquidityIterator, error) {
			return f.filterer.FilterIncreaseLiquidity(opts, tokenId)
		})
}

func (m *MultiURLNFPositionManagerFilterer) FilterDecreaseLiquidity(
	opts *bind.FilterOpts,
	tokenId []*big.Int,
) (*INonFungiblePositionsManagerDecreaseLiquidityIterator, error) {
	return trade.RetryEthCall(
		func() []*NFPositionManagerFiltererWithURL { return m.filterers },
		func(f *NFPositionManagerFiltererWithURL) (*INonFungiblePositionsManagerDecreaseLiquidityIterator, error) {
			return f.filterer.FilterDecreaseLiquidity(opts, tokenId)
		})
}

func (m *MultiURLNFPositionManagerFilterer) FilterCollect(
	opts *bind.FilterOpts,
	tokenId []*big.Int,
) (*INonFungiblePositionsManagerCollectIterator, error) {
	return trade.RetryEthCall(
		func() []*NFPositionManagerFiltererWithURL { return m.filterers },
		func(f *NFPositionManagerFiltererWithURL) (*INonFungiblePositionsManagerCollectIterator, error) {
			return f.filterer.FilterCollect(opts, tokenId)
		})
}

func (m *MultiURLNFPositionManagerFilterer) FilterTransfer(
	opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int,
) (*INonFungiblePositionsManagerTransferIterator, error) {
	return trade.RetryEthCall(
		func() []*NFPositionManagerFiltererWithURL { return m.filterers },
		func(f *NFPositionManagerFiltererWithURL) (*INonFungiblePositionsManagerTransferIterator, error) {
			return f.filterer.FilterTransfer(opts, from, to, tokenId)
		})
}

type NFPositionManagerInstance struct {
	client   *web3client.MultiURLClient
	caller   *MultiURLNFPositionManagerCaller
	filterer *MultiURLNFPositionManagerFilterer
	Address  common.Address
}

func NewNFPositionManagerInstance(client *web3client.MultiURLClient, address string) (*NFPositionManagerInstance, error) {
	checksumAddr := common.HexToAddress(address)

	callers := make([]*NFPositionManagerCallerWithURL, client.Length())
	for i, clientWithURL := range client.Iter() {
		caller, err := NewINonFungiblePositionsManagerCaller(checksumAddr, clientWithURL.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to create INonFungiblePositionsManagerCaller: %w. URL: %s", err, clientWithURL.Url)
		}
		callers[i] = &NFPositionManagerCallerWithURL{Caller: caller, Url: clientWithURL.Url}
	}

	filterers := make([]*NFPositionManagerFiltererWithURL, client.Length())
	for i, clientWithURL := range client.Iter() {
		filterer, err := NewINonFungiblePositionsManagerFilterer(checksumAddr, clientWithURL.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to create INonFungiblePositionsManagerFilterer: %w. URL: %s", err, clientWithURL.Url)
		}
		filterers[i] = &NFPositionManagerFiltererWithURL{filterer: filterer, url: clientWithURL.Url}
	}

	return &NFPositionManagerInstance{
		client:   client,
		caller:   &MultiURLNFPositionManagerCaller{callers},
		filterer: &MultiURLNFPositionManagerFilterer{filterers},
		Address:  checksumAddr,
	}, nil
}
