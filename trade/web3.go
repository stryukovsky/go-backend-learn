package trade

import (
	"fmt"
	"os"

	"github.com/IBM/fp-go/either"
	E "github.com/IBM/fp-go/either"
	F "github.com/IBM/fp-go/function"

	"github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/eth"
)

func CreateToken(w3 *web3.Web3, address string) either.Either[error, *eth.Contract] {
	read := E.Eitherize1(os.ReadFile)
	fromByte := E.Eitherize1(func(raw []byte) (string, error) { return string(raw), nil })
	createSC := E.Eitherize1(func(abi string) (*eth.Contract, error) {
		return eth.NewContract(abi, address)
	})
	return F.Pipe3("abi/ERC20.json", read, E.Chain(fromByte), E.Chain(createSC))
}

func TokenDecimals(token *eth.Contract) either.Either[error, int] {
	call := E.Eitherize1(func(method string) (any, error) { return token.Call(method) })
	convert := E.Eitherize1(func(value any) (int, error) {
		result, ok := value.(int)
		if ok {
			return result, nil
		} else {
			return 0, fmt.Errorf("Bad number value for decimals")
		}
	})
	result := F.Flow2(call, E.Chain(convert))
	return result("decimals")
}

func CreateProvider(provider string) either.Either[error, *web3.Web3] {
	web3 := E.Eitherize1(web3.NewWeb3)
	return F.Pipe1(provider, web3)
}

func PerfromDeal(deal *Deal) either.Either[error, int] {
	createToken := E.Chain(F.Bind2nd(CreateToken, deal.InputToken))
	decimals := E.Chain(TokenDecimals)
	result := F.Pipe3("http://localhost:8545", CreateProvider, createToken, decimals)
	return result
}

