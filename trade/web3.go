package trade

import (
	"os"
	"github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/eth"
)


func CreateContract(w3 *web3.Web3, abiFilename string, address string) (*eth.Contract, error) {
	rawAbi, err := os.ReadFile(abiFilename)	
	if err != nil {
		return nil, err
	}
	abi := string(rawAbi)
	return w3.Eth.NewContract(abi, address)
}


