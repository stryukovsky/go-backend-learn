package trade

import (
	"os"
	"strings"
		E "github.com/IBM/fp-go/either"
	F "github.com/IBM/fp-go/function"

	"github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/eth"
)

func CreateToken(w3 *web3.Web3, address string) (*eth.Contract, error) {
	rawABI := mo.TupleToResult(os.ReadFile("abi/ERC20.json")).ToEither()
	// abi := rawABI.MapRight())
}
