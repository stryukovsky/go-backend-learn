package aave

import (
	"fmt"
	"log/slog"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/cache"
	"github.com/stryukovsky/go-backend-learn/trade/web3client"
	"gorm.io/gorm"
)

type AaveHandler struct {
	pool           AavePool
	cm             *cache.CacheManager
	db             *gorm.DB
	name           string
	tokens         []trade.Token
	parallelFactor int
}

func (h *AaveHandler) ParallelFactor() int { return h.parallelFactor }

func NewAaveHandler(
	instance trade.DeFiPlatform,
	client *web3client.MultiURLClient,
	rdb *cache.CacheManager,
	tokens []trade.Token,
	parallelFactor int,
) (*AaveHandler, error) {
	pool, err := NewAavePool(client, instance.Address)
	if err != nil {
		return nil, err
	}
	return &AaveHandler{
		pool:           *pool,
		cm:             rdb,
		name:           fmt.Sprintf("Aave on %s", instance.Address),
		tokens:         tokens,
		parallelFactor: parallelFactor,
	}, nil
}

func (h *AaveHandler) parseAaveEvents(chainId string, events []any) ([]trade.AaveEvent, error) {
	return trade.ParseEVMEvents(h.ParallelFactor(),
		h.Name(),
		chainId,
		events,
		func(task trade.ParallelEVMParserTask[trade.AaveEvent],
			generalEvent any,
		) error {
			switch generalEvent := generalEvent.(type) {
			default:
				return fmt.Errorf("[%s] Unexpected event type %s in chunk of Supply Events", h.Name(), generalEvent)
			case PoolSupply:
				var event PoolSupply = generalEvent
				timestamp, err := h.cm.GetCachedBlockTimestamp(event.Raw.BlockNumber)
				if err != nil {
					slog.Warn(fmt.Sprintf("[%s] Failure on parsing Supply event %s", h.Name(), err.Error()))
					return err
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
				task.ValuesCh <- item
			case PoolWithdraw:
				var event PoolWithdraw = generalEvent
				timestamp, err := h.cm.GetCachedBlockTimestamp(event.Raw.BlockNumber)
				if err != nil {
					slog.Warn(fmt.Sprintf("[%s] Failure on parsing Withdraw event %s", h.Name(), err.Error()))
					return err
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
				task.ValuesCh <- item
			}
				return nil
		})
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
	result := make([]trade.AaveInteraction, 0, len(interactions))
	for _, interaction := range interactions {
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

		closePrice, err := h.cm.GetCachedSymbolPriceAtTime(token.Symbol, &interaction.Timestamp)
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
		result = append(result, deal)
	}
	return result, nil
}
func (h *AaveHandler) Name() string { return h.name }
