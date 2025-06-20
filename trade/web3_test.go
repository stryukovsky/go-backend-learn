package trade

import (
	"math/big"
	"testing"

	"github.com/chenzhijie/go-web3"
	"github.com/stretchr/testify/assert"
)

func TestCreateContract(t *testing.T) {
	provider, err := web3.NewWeb3("http://localhost:8545")
	if err != nil {
		t.Fatal(err)
	}
	contract, err := CreateContract(provider, "../abi/ERC20.json", "0xdac17f958d2ee523a2206206994597c13d831ec7")
	if err != nil {
		// return nil, err
		t.Fatal(err)
	}
	result, err := contract.Call("decimals")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *big.NewInt(6), *result.(*big.Int))
}
