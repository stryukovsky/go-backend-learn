package trade

import (
	"math/big"

	"github.com/chenzhijie/go-web3/eth"
)

type ERC20 struct {
	contract eth.Contract
	decimals uint8
	name     string
}

func (token *ERC20) Transfer(recipient string, amount big.Int) {
  token.contract.Call("transfer", recipient, amount)

}
