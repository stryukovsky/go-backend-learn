package hodl

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stryukovsky/go-backend-learn/trade"
)

type ERC20 struct {
	client   *ethclient.Client
	caller   *IERC20Caller
	filterer *IERC20Filterer
	Info     trade.Token
}

func (token *ERC20) BalanceOf(recipient string) (*big.Int, error) {
	balance, err := token.caller.BalanceOf(&bind.CallOpts{}, common.HexToAddress(recipient))
	if err != nil {
		return nil, err
	}
	return balance, nil
}

var TransferTopic string = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

var (
	BadDecimalsValue error = errors.New("Bad decimals of ERC20 contract")
	BadNameValue     error = errors.New("Bad name of ERC20 contract")
	BadSymbolValue   error = errors.New("Bad symbol of ERC20 contract")
)

func NewERC20(client *ethclient.Client, token trade.Token) (*ERC20, error) {
	caller, err := NewIERC20Caller(common.HexToAddress(token.Address), client)
	if err != nil {
		return nil, err
	}

	filterer, err := NewIERC20Filterer(common.HexToAddress(token.Address), client)
	if err != nil {
		return nil, err
	}
	return &ERC20{client: client, caller: caller, filterer: filterer, Info: token}, nil
}
