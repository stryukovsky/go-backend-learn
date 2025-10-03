package aave

import (
	"fmt"
	"log/slog"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
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

func (h *AaveHandler) parseSupplyEvents(chainId string, event *PoolSupplyIterator) ([]trade.AaveEvent, error) {
	result := make([]trade.AaveEvent, 0)
	for event.Next() {
		err := event.Error()
		if err != nil {
			return nil, err
		}
		timestamp, err := cache.GetCachedBlockTimestamp(h.pool.client, h.rdb, event.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}
		item := trade.NewAaveEvent(
			chainId,
			"supply",
			event.Event.OnBehalfOf,
			event.Event.Reserve,
			event.Event.Amount,
			*timestamp,
			event.Event.Raw.TxHash.Hex(),
		)
		result = append(result, item)
	}
	return result, nil
}

func (h *AaveHandler) parseWithdrawEvents(chainId string, event *PoolWithdrawIterator) ([]trade.AaveEvent, error) {
	result := make([]trade.AaveEvent, 0)
	for event.Next() {
		err := event.Error()
		if err != nil {
			return nil, err
		}
		timestamp, err := cache.GetCachedBlockTimestamp(h.pool.client, h.rdb, event.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}
		item := trade.NewAaveEvent(
			chainId,
			"withdraw",
			event.Event.To,
			event.Event.Reserve,
			event.Event.Amount,
			*timestamp,
			event.Event.Raw.TxHash.Hex(),
		)
		result = append(result, item)
	}
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
	supplyEventsRaw, err := h.pool.filterer.FilterSupply(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock},
		[]common.Address{},
		formattedParticipants,
		[]uint16{},
	)
	if err != nil {
		return nil, err
	}
	withdrawEventsRaw, err := h.pool.filterer.FilterWithdraw(&bind.FilterOpts{}, []common.Address{}, []common.Address{}, formattedParticipants)
	if err != nil {
		return nil, err
	}
	defer supplyEventsRaw.Close()
	defer withdrawEventsRaw.Close()
	supplyEvents, err := h.parseSupplyEvents(chainId, supplyEventsRaw)
	if err != nil {
		return nil, err
	}
	withdrawEvents, err := h.parseWithdrawEvents(chainId, withdrawEventsRaw)
	if err != nil {
		return nil, err
	}
	return append(supplyEvents, withdrawEvents...), nil
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
