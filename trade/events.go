package trade

import (
	"math/big"

	"github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/types"
	"github.com/ethereum/go-ethereum/common"
)

func EventTransfer(w3 *web3.Web3, contract string) ([]ERC20Transfer, error) {
	contractAddress := common.HexToAddress(contract)
	fliter := &types.Fliter{Address: contractAddress, Topics: []string{"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"}}
	logs, err := w3.Eth.GetLogs(fliter)
	if err != nil {
		return nil, err
	}
	result := make([]ERC20Transfer, 0, len(logs))
	for _, e := range logs {
		if len(e.Topics) != 3 {
			continue
		}
		from := common.HexToAddress(e.Topics[1])
		to := common.HexToAddress(e.Topics[1])
		amount := big.NewInt(0)
		amount.SetString(e.Data, 16)
		transfer := ERC20Transfer{
			Sender:       from.Hex(),
			Recipient:    to.Hex(),
			Amount:       *amount,
			TokenAddress: contractAddress.Hex(),
		}
		result = append(result, transfer)
	}
	return result, nil
}
