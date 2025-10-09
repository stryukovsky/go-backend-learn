package hodl

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"sync"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/cache"
)

type HODLHandler struct {
	token          ERC20
	rdb            *redis.Client
	parallelFactor int
}

func NewHODLHandler(client *ethclient.Client, token trade.Token, rdb *redis.Client, parallelFactor int) (*HODLHandler, error) {
	erc20, err := NewERC20(client, token)
	if err != nil {
		return nil, err
	}
	return &HODLHandler{
		token:          *erc20,
		rdb:            rdb,
		parallelFactor: parallelFactor,
	}, nil
}

func (h *HODLHandler) FetchBlockchainInteractions(
	chainId string,
	participants []string,
	fromBlock uint64,
	toBlock uint64,

) ([]trade.ERC20Transfer, error) {
	formattedParticipants := make([]common.Hash, len(participants))
	for i, participant := range participants {
		formattedParticipants[i] = common.HexToHash(participant)
	}
	logsParticipantsSenders, err :=
		h.token.client.FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(fromBlock)),
			ToBlock:   big.NewInt(int64(toBlock)),
			Addresses: []common.Address{common.HexToAddress(h.token.Info.Address)},
			Topics: [][]common.Hash{
				{common.HexToHash(TransferTopic)},
				formattedParticipants,
				{},
			},
		})
	if err != nil {
		return []trade.ERC20Transfer{}, err
	}
	logsParticipantsRecipients, err :=
		h.token.client.FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(fromBlock)),
			ToBlock:   big.NewInt(int64(toBlock)),
			Addresses: []common.Address{common.HexToAddress(h.token.Info.Address)},
			Topics: [][]common.Hash{
				{common.HexToHash(TransferTopic)},
				{},
				formattedParticipants,
			},
		})
	if err != nil {
		return []trade.ERC20Transfer{}, err
	}
	allLogs := append(logsParticipantsSenders, logsParticipantsRecipients...)
	if len(allLogs) == 0 {
		return []trade.ERC20Transfer{}, nil
	}
	slog.Info(fmt.Sprintf("Scanned %d transfers", len(allLogs)))
	logsChunks := lo.Chunk(allLogs, h.ParallelFactor())
	resultCh := make(chan trade.ERC20Transfer)
	result := make([]trade.ERC20Transfer, len(allLogs))

	var wg sync.WaitGroup
	wg.Add(h.ParallelFactor())

	go func() {
		for i := range h.ParallelFactor() {
			go func() {
				defer wg.Done()
				if i >= len(logsChunks) {
					return
				}
				for _, event := range logsChunks[i] {
					sender := common.BytesToAddress(event.Topics[1].Bytes()).Hex()
					recipient := common.BytesToAddress(event.Topics[2].Bytes()).Hex()
					amount := common.BytesToHash(event.Data).Big()
					txId := event.TxHash.Hex()
					block := big.NewInt(int64(event.BlockNumber))
					timestamp, err := cache.GetCachedBlockTimestamp(h.token.client, h.rdb, event.BlockNumber)
					if err != nil {
						slog.Warn(fmt.Sprintf("Cannot fetch from cache or blockchain info on block %d timestamp: %s", event.BlockNumber, err.Error()))
						continue
					}
					transfer := trade.NewERC20Transfer(h.token.Info.Address, sender, recipient, amount, block, chainId, timestamp, txId, event.Index)
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

func (h *HODLHandler) ParallelFactor() int { return h.parallelFactor }

func (h *HODLHandler) PopulateWithFinanceInfo(interactions []trade.ERC20Transfer) ([]trade.Deal, error) {
	result := make([]trade.Deal, len(interactions))
	for i, transfer := range interactions {
		closePrice, err := cache.GetCachedSymbolPriceAtTime(h.rdb, h.token.Info.Symbol, &transfer.Timestamp)
		if err != nil {
			return nil, err
		}
		volumeToken := big.NewRat(1, 1)
		volumeToken = volumeToken.SetFrac(transfer.Amount.Int, new(big.Int).Exp(big.NewInt(10), h.token.Info.Decimals.Int, nil))
		volumeUSD := new(big.Rat).Mul(volumeToken, closePrice)
		deal := trade.Deal{
			Price:              trade.NewDBNumeric(closePrice),
			VolumeUSD:          trade.NewDBNumeric(volumeUSD),
			VolumeTokens:       trade.NewDBNumeric(volumeToken),
			BlockchainTransfer: transfer,
		}
		result[i] = deal
	}
	return result, nil
}

func (h *HODLHandler) Name() string {
	return h.token.Info.Symbol
}
