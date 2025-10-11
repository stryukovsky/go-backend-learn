package aave

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/cache"
	"gorm.io/gorm"
)

type AaveHandler struct {
	pool           AavePool
	rdb            *redis.Client
	db             *gorm.DB
	name           string
	tokens         []trade.Token
	parallelFactor int
}

func (h *AaveHandler) ParallelFactor() int { return h.parallelFactor }

func NewAaveHandler(
	instance trade.DeFiPlatform,
	client *ethclient.Client,
	rdb *redis.Client,
	tokens []trade.Token,
	parallelFactor int,
) (*AaveHandler, error) {
	pool, err := NewAavePool(client, instance.Address)
	if err != nil {
		return nil, err
	}
	return &AaveHandler{
		pool:           *pool,
		rdb:            rdb,
		name:           fmt.Sprintf("Aave on %s", instance.Address),
		tokens:         tokens,
		parallelFactor: parallelFactor,
	}, nil
}

func (h *AaveHandler) parseAaveEvents(chainId string, events []any) ([]trade.AaveEvent, error) {
	chunkSize := len(events) / h.ParallelFactor()
	if chunkSize == 0 {
		return []trade.AaveEvent{}, nil
	}
	eventChunks := lo.Chunk(events, chunkSize)
	var wg sync.WaitGroup
	wg.Add(h.ParallelFactor())
	valuesCh := make(chan trade.AaveEvent, h.ParallelFactor())
	ctx, cancel := context.WithCancel(context.Background())
	defer ctx.Done()
	for i, chunk := range eventChunks {
		go func() {
			slog.Debug(fmt.Sprintf("[%s] Parsing %d-th chunk of Supply Events", h.Name(), i+1))
			for _, generalEvent := range chunk {
				select {
				case <-ctx.Done():
					return
				default:
					switch generalEvent := generalEvent.(type) {
					default:
						slog.Info(fmt.Sprintf("[%s] Unexpected event type %s in chunk of Supply Events", h.Name(), generalEvent))
					case PoolSupply:
						var event PoolSupply = generalEvent
						timestamp, err := cache.GetCachedBlockTimestamp(h.pool.client, h.rdb, event.Raw.BlockNumber)
						if err != nil {
							slog.Warn(fmt.Sprintf("[%s] Failure on parsing Supply event %s", h.Name(), err.Error()))
							wg.Done()
							cancel()
						}
						item := trade.NewAaveEvent(
							chainId,
							"supply",
							event.OnBehalfOf,
							event.Reserve,
							event.Amount,
							*timestamp,
							event.Raw.TxHash.Hex(),
							event.Raw.Index,
						)
						valuesCh <- item
					case PoolWithdraw:
						var event PoolWithdraw = generalEvent
						timestamp, err := cache.GetCachedBlockTimestamp(h.pool.client, h.rdb, event.Raw.BlockNumber)
						if err != nil {
							slog.Warn(fmt.Sprintf("[%s] Failure on parsing Withdraw event %s", h.Name(), err.Error()))
							wg.Done()
							cancel()
						}
						item := trade.NewAaveEvent(
							chainId,
							"withdraw",
							event.To,
							event.Reserve,
							event.Amount,
							*timestamp,
							event.Raw.TxHash.Hex(),
							event.Raw.Index,
						)
						valuesCh <- item
					}
				}
				wg.Done()
			}
		}()
	}
	wg.Wait()
	result := make([]trade.AaveEvent, len(events))
	i := 0
	for item := range valuesCh {
		result[i] = item
		i++
	}
	cancel()
	return result, nil
}

func (h *AaveHandler) FetchBlockchainInteractions(
	chainId string,
	participants []string,
	fromBlock uint64,
	toBlock uint64,
) ([]trade.AaveEvent, error) {
	formattedParticipants := make([]common.Address, len(participants))
	for i, p := range participants {
		formattedParticipants[i] = common.HexToAddress(p)
	}
	supplyEventsIter, err := h.pool.filterer.FilterSupply(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock},
		[]common.Address{},
		formattedParticipants,
		[]uint16{},
	)
	if err != nil {
		return nil, err
	}
	defer supplyEventsIter.Close()
	withdrawEventsIter, err := h.pool.filterer.FilterWithdraw(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock}, []common.Address{}, []common.Address{}, formattedParticipants)
	if err != nil {
		return nil, err
	}
	defer withdrawEventsIter.Close()

	// any is because go do not support generic methods, we have two types for each event: Supply and Withdraw
	eventsRaw := make([]any, 0)
	for supplyEventsIter.Next() {
		if err = supplyEventsIter.Error(); err != nil {
			return nil, err
		}
		eventsRaw = append(eventsRaw, *supplyEventsIter.Event)
	}
	for withdrawEventsIter.Next() {
		err := withdrawEventsIter.Error()
		if err != nil {
			return nil, err
		}
		eventsRaw = append(eventsRaw, *withdrawEventsIter.Event)
	}
	if len(eventsRaw) == 0 {
		slog.Warn(fmt.Sprintf("[%s] no events in block range %d - %d", h.Name(), fromBlock, toBlock))
		return make([]trade.AaveEvent, 0), nil
	} else {
		slog.Info(fmt.Sprintf("[%s] found %d events in block range %d - %d", h.Name(), len(eventsRaw), fromBlock, toBlock))
	}
	events, err := h.parseAaveEvents(chainId, eventsRaw)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (h *AaveHandler) PopulateWithFinanceInfo(interactions []trade.AaveEvent) ([]trade.AaveInteraction, error) {
	result := make([]trade.AaveInteraction, len(interactions))
	for i, interaction := range interactions {
		tokenAddress := common.HexToAddress(interaction.TokenAddress)
		token := trade.Token{}
		for _, t := range h.tokens {
			if strings.EqualFold(t.Address, tokenAddress.Hex()) {
				token = t
			}
		}

		if len(token.Address) == 0 {
			slog.Warn(fmt.Sprintf("Found aave interaction with unknown token address %s", tokenAddress))
			continue
		}

		closePrice, err := cache.GetCachedSymbolPriceAtTime(h.rdb, token.Symbol, &interaction.Timestamp)
		if err != nil {
			return nil, err
		}

		volumeToken := big.NewRat(1, 1)
		decimalsMultiplier := new(big.Int).Exp(big.NewInt(10), token.Decimals.Int, nil)
		volumeToken = volumeToken.SetFrac(interaction.Amount.Int, decimalsMultiplier)

		volumeUSD := new(big.Rat).Mul(volumeToken, closePrice)
		deal := trade.AaveInteraction{
			Price:           trade.NewDBNumeric(closePrice),
			VolumeTokens:    trade.NewDBNumeric(volumeToken),
			VolumeUSD:       trade.NewDBNumeric(volumeUSD),
			BlockchainEvent: interaction,
		}
		result[i] = deal
	}
	return result, nil
}
func (h *AaveHandler) Name() string { return h.name }
