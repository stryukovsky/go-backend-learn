package trade

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/chenzhijie/go-web3"
)

func TestPriceFetching(t *testing.T) {
	w3, err := web3.NewWeb3("http://localhost:8545")
	if err != nil {
		t.Fatal(err)
	}

	token, err := CreateToken(w3, "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", "ETH")

	transfers, err := token.ListTransfers(big.NewInt(1), big.NewInt(1))
	if err != nil {
		t.Fatal(err)
	}

	firstTransfer := transfers[0]
	deal, err := CreateDeal(firstTransfer)
	if err != nil {
		t.Fatal(err)
	}

	dealJSON, err := json.Marshal(deal)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(dealJSON))
}
