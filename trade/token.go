package trade

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"math/big"

	"github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/eth"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joomcode/errorx"
	"github.com/redis/go-redis/v9"
)

type ERC20 struct {
	W3       *web3.Web3
	Address  string
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

// func (token *ERC20) ListTransfers(
// 	fliter *types.Fliter,
// 	cache *redis.Client,
// 	filterTransfer func(string, string, *big.Int) bool,
// ) ([]ERC20Transfer, error) {
// 	logs, err := token.W3.Eth.GetLogs(fliter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	slog.Info(fmt.Sprintf("[%s] Found %d events. Filter them to find transfers of tracked wallets", token.Symbol, len(logs)))
// 	result := make([]ERC20Transfer, 0, len(logs))
// 	for _, e := range logs {
// 		if len(e.Topics) != 3 {
// 			continue
// 		}
// 		from := common.HexToAddress(e.Topics[1]).Hex()
// 		to := common.HexToAddress(e.Topics[2]).Hex()
// 		amount := common.HexToHash(e.Data).Big()
// 		if filterTransfer(from, to, amount) {
//
// 			blockNumber := common.HexToHash(e.BlockNumber).Big()
// 			timestamp, err := GetCachedBlockTimestamp(token.W3, cache, blockNumber.Uint64())
// 			if err != nil {
// 				return nil, err
// 			}
// 			transfer := ERC20Transfer{
// 				Sender:       from,
// 				Recipient:    to,
// 				Amount:       DBInt{amount},
// 				TokenAddress: token.Contract.Address().Hex(),
// 				Block:        DBInt{blockNumber},
// 				Timestamp:    *timestamp,
// 				TxId:         e.TransactionHash.Hex(),
// 			}
// 			result = append(result, transfer)
// 		}
// 	}
// 	return result, nil
// }

func (token *ERC20) ListTransfersOfParticipants(
	rpcUrl string,
	chainId string,
	participants []string,
	fromBlock uint64,
	toBlock uint64,
	cache *redis.Client,
) ([]ERC20Transfer, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return []ERC20Transfer{}, err
	}
	formattedParticipants := make([]common.Hash, len(participants))
	for i, participant := range participants {
		formattedParticipants[i] = common.HexToHash(participant)
	}
	logsParticipantsSenders, err := 
		client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{common.HexToAddress(token.Address)},
		Topics: [][]common.Hash{
			{common.HexToHash(TransferTopic)},
			formattedParticipants,
			{},
		},
	})
	if err != nil {
		return []ERC20Transfer{}, err
	}
	logsParticipantsRecipients, err := 
		client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{common.HexToAddress(token.Address)},
		Topics: [][]common.Hash{
			{common.HexToHash(TransferTopic)},
			{},
			formattedParticipants,
		},
	})
	if err != nil {
		return []ERC20Transfer{}, err
	}
	allLogs := append(logsParticipantsSenders, logsParticipantsRecipients...)
	result := make([]ERC20Transfer, len(allLogs))
	for i, event := range allLogs {
		sender := common.BytesToAddress(event.Topics[1].Bytes()).Hex()
		recipient := common.BytesToAddress(event.Topics[2].Bytes()).Hex()
		amount := common.BytesToHash(event.Data).Big()
		txId := event.TxHash.Hex()
		block := big.NewInt(int64(event.BlockNumber))
		timestamp, err := GetCachedBlockTimestamp(token.W3, cache, event.BlockNumber)
		if err != nil {
			slog.Warn("Cannot fetch from cache or blockchain info on block %d timestamp: %s", event.BlockNumber, err.Error())	
			continue
		}
		transfer := NewERC20Transfer(token.Address, sender, recipient, amount, block, chainId, timestamp, txId)
		result[i] = transfer
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
	return &ERC20{W3: w3, Address: address, Contract: *contract, Decimals: *decimals, Name: name, Symbol: symbol}, nil
}
