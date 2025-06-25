package trade

import (
	"errors"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/eth"
	"github.com/chenzhijie/go-web3/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joomcode/errorx"
	"github.com/redis/go-redis/v9"
)

type ERC20 struct {
	W3       *web3.Web3
	Contract eth.Contract
	Decimals big.Int
	Name     string
	Symbol   string
}

var (
	BadBalance error = errors.New("Bad balance of ERC20 contract")
)

func (token *ERC20) BalanceOf(recipient string) (*big.Int, error) {
	rawBalance, err := token.Contract.Call("balanceOf", common.HexToAddress(recipient))
	if err != nil {
		return nil, errorx.Decorate(err, "Cannot call balanceOf")
	}
	var balance *big.Int
	var success bool
	if balance, success = rawBalance.(*big.Int); !success {
		return nil, BadBalance
	}
	return balance, nil
}

var TransferTopic string = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

func (token *ERC20) ListTransfers(fromBlock uint64, toBlock uint64, cache *redis.Client) ([]ERC20Transfer, error) {
	fromBlockHex := fmt.Sprintf("0x%x", fromBlock)
	toBlockHex := fmt.Sprintf("0x%x", toBlock)
	fliter := &types.Fliter{Address: token.Contract.Address(),
		FromBlock: fromBlockHex,
		ToBlock:   toBlockHex,
		Topics:    []string{TransferTopic}}
	logs, err := token.W3.Eth.GetLogs(fliter)
	if err != nil {
		return nil, err
	}
	slog.Info(fmt.Sprintf("[%s] Found %d events between %d and %d blocks", token.Symbol, len(logs), fromBlock, toBlock))
	result := make([]ERC20Transfer, 0, len(logs))
	for _, e := range logs {
		if len(e.Topics) != 3 {
			continue
		}
		from := common.HexToAddress(e.Topics[1])
		to := common.HexToAddress(e.Topics[1])
		amount := common.HexToHash(e.Data).Big()
		blockNumber := common.HexToHash(e.BlockNumber).Big()
		timestamp, err := GetCachedBlockTimestamp(token.W3, cache, blockNumber.Uint64())
		if err != nil {
			return nil, err
		}
		transfer := ERC20Transfer{
			Sender:       from.Hex(),
			Recipient:    to.Hex(),
			Amount:       DBInt{amount},
			TokenAddress: token.Contract.Address().Hex(),
			Block:        DBInt{blockNumber},
			Timestamp:    *timestamp,
			Decimals:     DBInt{&token.Decimals},
			Symbol:       token.Symbol,
			TxId:         e.TransactionHash.Hex(),
		}
		result = append(result, transfer)
	}
	return result, nil
}

var (
	BadDecimalsValue error = errors.New("Bad decimals of ERC20 contract")
	BadNameValue     error = errors.New("Bad name of ERC20 contract")
	BadSymbolValue   error = errors.New("Bad symbol of ERC20 contract")
)

func CreateERC20(w3 *web3.Web3, address string, symbol string) (*ERC20, error) {
	contract, err := CreateContract(w3, "abi/ERC20.json", address)
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
	return &ERC20{W3: w3, Contract: *contract, Decimals: *decimals, Name: name, Symbol: symbol}, nil
}
