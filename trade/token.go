package trade

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
)

type ERC20 struct {
	client *ethclient.Client
	caller *IERC20Caller
	Info   Token
}

func (token *ERC20) BalanceOf(recipient string) (*big.Int, error) {
	balance, err := token.caller.BalanceOf(&bind.CallOpts{}, common.HexToAddress(recipient))
	if err != nil {
		return nil, err
	}
	return balance, nil
}

var TransferTopic string = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

const ParallelFactor = 16

func (token *ERC20) ListTransfersOfParticipants(
	chainId string,
	participants []string,
	fromBlock uint64,
	toBlock uint64,
	cache *redis.Client,
) ([]ERC20Transfer, error) {
	formattedParticipants := make([]common.Hash, len(participants))
	for i, participant := range participants {
		formattedParticipants[i] = common.HexToHash(participant)
	}
	logsParticipantsSenders, err :=
		token.client.FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(fromBlock)),
			ToBlock:   big.NewInt(int64(toBlock)),
			Addresses: []common.Address{common.HexToAddress(token.Info.Address)},
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
		token.client.FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(fromBlock)),
			ToBlock:   big.NewInt(int64(toBlock)),
			Addresses: []common.Address{common.HexToAddress(token.Info.Address)},
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
	slog.Info(fmt.Sprintf("Scanned %d transfers", len(allLogs)))
	logsChunks := lo.Chunk(allLogs, ParallelFactor)
	resultCh := make(chan ERC20Transfer)
	result := make([]ERC20Transfer, len(allLogs))

	var wg sync.WaitGroup
	wg.Add(ParallelFactor)

	go func() {
		for i := range ParallelFactor {
			go func() {
				defer wg.Done()
				for _, event := range logsChunks[i] {
					sender := common.BytesToAddress(event.Topics[1].Bytes()).Hex()
					recipient := common.BytesToAddress(event.Topics[2].Bytes()).Hex()
					amount := common.BytesToHash(event.Data).Big()
					txId := event.TxHash.Hex()
					block := big.NewInt(int64(event.BlockNumber))
					timestamp, err := GetCachedBlockTimestamp(token.client, cache, event.BlockNumber)
					if err != nil {
						slog.Warn(fmt.Sprintf("Cannot fetch from cache or blockchain info on block %d timestamp: %s", event.BlockNumber, err.Error()))
						continue
					}
					transfer := NewERC20Transfer(token.Info.Address, sender, recipient, amount, block, chainId, timestamp, txId)
					resultCh <- transfer
				}
			}()
		}
		wg.Wait()
		close(resultCh)
	}()
	i := 0
	for transfer := range resultCh {
		result[i] = transfer
		i++
	}

	return result, nil
}

var (
	BadDecimalsValue error = errors.New("Bad decimals of ERC20 contract")
	BadNameValue     error = errors.New("Bad name of ERC20 contract")
	BadSymbolValue   error = errors.New("Bad symbol of ERC20 contract")
)

func NewERC20(client *ethclient.Client, token Token) (*ERC20, error) {
	caller, err := NewIERC20Caller(common.HexToAddress(token.Address), client)
	if err != nil {
		return nil, err
	}
	return &ERC20{client: client, caller: caller, Info: token}, nil
}
