package hodl

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/web3client"
)

type ERC20CallerWithURL struct {
	Caller *IERC20Caller
	Url    string
}

func (c *ERC20CallerWithURL) URL() string { return c.Url }

type MultiURLERC20Caller struct {
	callers []*ERC20CallerWithURL
}

func (c *MultiURLERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	return trade.RetryEthCall(
		func() []*ERC20CallerWithURL { return c.callers },
		func(caller *ERC20CallerWithURL) (*big.Int, error) { return caller.Caller.BalanceOf(opts, account) },
	)
}

type ERC20FiltererWithURL struct {
	filterer *IERC20Filterer
	url      string
}

type MultiURLERC20Filterer struct {
	filterers []*ERC20FiltererWithURL
}

func (f *ERC20FiltererWithURL) URL() string { return f.url }

func (m *MultiURLERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {
	return trade.RetryEthCall(
		func() []*ERC20FiltererWithURL { return m.filterers },
		func(filterer *ERC20FiltererWithURL) (*IERC20TransferIterator, error) {
			return filterer.filterer.FilterTransfer(opts, from, to)
		})
}

type ERC20 struct {
	client   *web3client.MultiURLClient
	caller   *MultiURLERC20Caller
	filterer *MultiURLERC20Filterer
	Info     trade.Token
}

func (token *ERC20) BalanceOf(recipient string) (*big.Int, error) {
	balance, err := token.caller.BalanceOf(&bind.CallOpts{}, common.HexToAddress(recipient))
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func NewERC20(client *web3client.MultiURLClient, token trade.Token) (*ERC20, error) {
	callers := make([]*ERC20CallerWithURL, client.Length())
	for i, clientWithURL := range client.Iter() {
		caller, err := NewIERC20Caller(common.HexToAddress(token.Address), clientWithURL.Client)
		if err != nil {
			return nil, fmt.Errorf("%s. URL of provider is %s", err.Error(), clientWithURL.Url)
		}
		callers[i] = &ERC20CallerWithURL{Caller: caller, Url: clientWithURL.Url}
	}

	filterers := make([]*ERC20FiltererWithURL, client.Length())
	for i, clientWithURL := range client.Iter() {
		filterer, err := NewIERC20Filterer(common.HexToAddress(token.Address), clientWithURL.Client)
		if err != nil {
			return nil, fmt.Errorf("%s. URL of provider is %s", err.Error(), clientWithURL.Url)
		}
		filterers[i] = &ERC20FiltererWithURL{filterer: filterer, url: clientWithURL.Url}
	}
	return &ERC20{client: client, caller: &MultiURLERC20Caller{callers}, filterer: &MultiURLERC20Filterer{filterers}, Info: token}, nil
}
