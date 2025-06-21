package trade

import (
	"fmt"
	"testing"

	"github.com/chenzhijie/go-web3"
)

func TestEvents(t *testing.T) {

	w3, err := web3.NewWeb3("http://localhost:8545")
	if err != nil {
		t.Fatal(err)
	}

	events, err := EventTransfer(w3, "0xdac17f958d2ee523a2206206994597c13d831ec7")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(events)
}
