package trade

import (
	"errors"
	"math/big"

	"github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/eth"
)

type ERC20 struct {
	contract eth.Contract
	decimals big.Int
	name     string
}

var (
	BadBalance error = errors.New("Bad balance of ERC20 contract")
)

func (token *ERC20) BalanceOf(recipient string) (*big.Int, error) {
	rawBalance, err := token.contract.Call("balanceOf", recipient)
	if err != nil {
		return nil, err
	}
	var balance *big.Int
	var success bool
	if balance, success = rawBalance.(*big.Int); !success {
		return nil, BadBalance
	}
	return balance, nil
}

var (
	BadDecimalsValue error = errors.New("Bad decimals of ERC20 contract")
	BadNameValue     error = errors.New("Bad name of ERC20 contract")
)

func CreateToken(w3 *web3.Web3, address string) (*ERC20, error) {
	contract, err := CreateContract(w3, "../abi/ERC20.json", address)
	if err != nil {
		return nil, err
	}
	decimalsRaw, err := contract.Call("decimals")
	if err != nil {
		return nil, err
	}
	var decimals *big.Int
	var success bool
	if decimals, success = decimalsRaw.(*big.Int); !success {
		return nil, BadDecimalsValue
	}
	nameRaw, err := contract.Call("name")
	if err != nil {
		return nil, err
	}
	var name string
	if name, success = nameRaw.(string); !success {
		return nil, BadNameValue
	}
	return &ERC20{contract: *contract, decimals: *decimals, name: name}, nil
}
