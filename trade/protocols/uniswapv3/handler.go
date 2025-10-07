package uniswapv3

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/cache"
	"gorm.io/gorm"
)

type UniswapV3PoolHandler struct {
	pool           UniswapV3PoolInstance
	rdb            *redis.Client
	db             *gorm.DB
	name           string
	tokenA         trade.Token
	tokenB         trade.Token
	parallelFactor int
	chainId        string
}

func (h *UniswapV3PoolHandler) ParallelFactor() int { return h.parallelFactor }

func NewUniswapV3PoolHandler(
	instance trade.DeFiPlatform,
	client *ethclient.Client,
	rdb *redis.Client,
	db *gorm.DB,
	tokenA trade.Token,
	tokenB trade.Token,
	parallelFactor int,
) (*UniswapV3PoolHandler, error) {
	pool, err := NewUniswapV3PoolInstance(client, instance.Address)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolHandler{
		pool:           *pool,
		rdb:            rdb,
		db:             db,
		name:           fmt.Sprintf("Uniswap V3 Pool %s - %s", tokenA.Symbol, tokenB.Symbol),
		tokenA:         tokenA,
		tokenB:         tokenB,
		parallelFactor: parallelFactor,
		chainId:        instance.ChainId,
	}, nil
}

const (
	Tick2PriceBase float64 = 1.0001
)

func Adjustment2HumanPrice(adjustedPrice *big.Rat, token0 trade.Token, token1 trade.Token) *big.Rat {
	// adjustment can help us return back to human price, its 10 ** (decimals0 - decimals1)
	adjustment := math.Pow10(int(token0.Decimals.Int64())) / math.Pow10(int(token1.Decimals.Int64()))
	humanPrice := new(big.Rat).Mul(adjustedPrice, new(big.Rat).SetFloat64(adjustment))
	return humanPrice
}

func Tick2Price(tick *big.Int, token0 trade.Token, token1 trade.Token) (*big.Rat, error) {
	if tick.Sign() < 0 {
		return nil, errors.New("Tick must be greater than or equal to zero")
	}
	if !tick.IsInt64() {
		return nil, errors.New("Tick must be representable as an int64")
	}
	power := float64(tick.Int64())
	adjustedPrice := new(big.Rat).SetFloat64(math.Pow(Tick2PriceBase, power))
	return Adjustment2HumanPrice(adjustedPrice, token0, token1), nil
}

func SqrtPrice2Price(source *big.Int, token0 trade.Token, token1 trade.Token) (*big.Rat, error) {
	sqrtPriceX96 := new(big.Rat).SetInt(source)
	pow96 := new(big.Rat).SetFloat64(math.Pow(2, -96))
	// multiply by 2 ** (-96)
	sqrtPrice := new(big.Rat).Mul(sqrtPriceX96, pow96)
	// i.e. simply take value from the square root
	adjustedPrice := new(big.Rat).Mul(sqrtPrice, sqrtPrice)
	return Adjustment2HumanPrice(adjustedPrice, token0, token1), nil
}

func (h *UniswapV3PoolHandler) parseMint(event UniswapV3PoolMint) (*trade.UniswapV3Event, error) {

	lowerPrice, err := Tick2Price(event.TickLower, h.tokenA, h.tokenB)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot parse lower price of mint event: %s", h.Name(), err.Error()))
		return nil, err
	}
	upperPrice, err := Tick2Price(event.TickUpper, h.tokenA, h.tokenB)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot parse upper price of mint event: %s", h.Name(), err.Error()))
		return nil, err
	}
	timestamp, err := cache.GetCachedBlockTimestamp(
		h.pool.client,
		h.rdb,
		event.Raw.BlockNumber,
	)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot get block timestamp from cache/blockchain: %s", h.Name(), err.Error()))
		return nil, err
	}
	result := trade.NewUniswapV3Event(
		h.chainId,
		trade.UniswapV3Mint,
		event.Owner.Hex(),
		h.pool.Address.Hex(),
		event.Amount0,
		event.Amount1,
		upperPrice,
		lowerPrice,
		*timestamp,
		event.Raw.TxHash.Hex(),
	)
	return &result, nil
}

func (h *UniswapV3PoolHandler) parseBurn(event UniswapV3PoolBurn) (*trade.UniswapV3Event, error) {
	lowerPrice, err := Tick2Price(event.TickLower, h.tokenA, h.tokenB)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot parse lower price of burn event: %s", h.Name(), err.Error()))
		return nil, err
	}
	upperPrice, err := Tick2Price(event.TickUpper, h.tokenA, h.tokenB)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot parse upper price of burn event: %s", h.Name(), err.Error()))
		return nil, err
	}
	timestamp, err := cache.GetCachedBlockTimestamp(
		h.pool.client,
		h.rdb,
		event.Raw.BlockNumber,
	)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot get block timestamp from cache/blockchain: %s", h.Name(), err.Error()))
		return nil, err
	}
	result := trade.NewUniswapV3Event(
		h.chainId,
		trade.UniswapV3Burn,
		event.Owner.Hex(),
		h.pool.Address.Hex(),
		event.Amount0,
		event.Amount1,
		upperPrice,
		lowerPrice,
		*timestamp,
		event.Raw.TxHash.Hex(),
	)
	return &result, nil
}

func (h *UniswapV3PoolHandler) parseSwap(event UniswapV3PoolSwap) (*trade.UniswapV3Event, error) {
	price, err := SqrtPrice2Price(event.SqrtPriceX96, h.tokenA, h.tokenB)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot parse price of burn event: %s", h.Name(), err.Error()))
		return nil, err
	}
	timestamp, err := cache.GetCachedBlockTimestamp(
		h.pool.client,
		h.rdb,
		event.Raw.BlockNumber,
	)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot get block timestamp from cache/blockchain: %s", h.Name(), err.Error()))
		return nil, err
	}
	result := trade.NewUniswapV3Event(
		h.chainId,
		trade.UniswapV3Burn,
		event.Sender.Hex(),
		h.pool.Address.Hex(),
		event.Amount0,
		event.Amount1,
		price,
		price,
		*timestamp,
		event.Raw.TxHash.Hex(),
	)
	return &result, nil
}

// events are of
// UniswapV3PoolMint
func (h *UniswapV3PoolHandler) parseEvents(events []any) ([]trade.UniswapV3Event, error) {
	chunkSize := len(events) / h.ParallelFactor()
	chunks := lo.Chunk(events, chunkSize)
	chunksCount := len(chunks)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(chunksCount)
	resultCh := make(chan trade.UniswapV3Event)
	// func Chunk[T any, Slice ~[]T](collection Slice, size int) []Slice {
	for _, chunk := range chunks {
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				for _, uncastedEvent := range chunk {
					switch castedEvent := uncastedEvent.(type) {
					case UniswapV3PoolMint:
						parsedEvent, err := h.parseMint(castedEvent)
						if err != nil {
							cancel()
							return
						}
						resultCh <- *parsedEvent
					case UniswapV3PoolBurn:
						parsedEvent, err := h.parseBurn(castedEvent)
						if err != nil {
							cancel()
							return
						}
						resultCh <- *parsedEvent
					case UniswapV3PoolSwap:
						parsedEvent, err := h.parseSwap(castedEvent)
						if err != nil {
							cancel()
							return
						}
						resultCh <- *parsedEvent
					default:
						slog.Info(fmt.Sprintf("[%s] Skip event of type %s since no parsing implemented for it", h.Name(), uncastedEvent))
						continue
					}
				}
			}
		}()
	}
	results := make([]trade.UniswapV3Event, len(events))
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			return
		default:
			i := 0
			for item := range resultCh {
				results[i] = item
				i++
			}
		}
	}()
	wg.Wait()
	return results, nil
}

func (h *UniswapV3PoolHandler) Name() string {
	return h.name
}

func (h *UniswapV3PoolHandler) FetchBlockchainInteractions(
	chainId string,
	participants []string,
	fromBlock uint64,
	toBlock uint64,
) ([]trade.UniswapV3Event, error) {
	formattedParticipants := lo.Map(participants, func(p string, _ int) common.Address { return common.HexToAddress(p) })
	mintEventsIter, err := h.pool.filterer.FilterMint(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock},
		formattedParticipants,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	burnEventsIter, err := h.pool.filterer.FilterBurn(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock},
		formattedParticipants,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	swapEventsIter, err := h.pool.filterer.FilterSwap(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock},
		nil,
		formattedParticipants,
	)
	if err != nil {
		return nil, err
	}
	events := make([]any, 0, 100)
	for mintEventsIter.Next() {
		events = append(events, *mintEventsIter.Event)
	}
	for burnEventsIter.Next() {
		events = append(events, *burnEventsIter.Event)
	}
	for swapEventsIter.Next() {
		events = append(events, *swapEventsIter.Event)
	}

	result, err := h.parseEvents(events)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (h *UniswapV3PoolHandler) humanVolumeOfToken(amount *big.Int, token *trade.Token, dealTime *time.Time) (*big.Rat, *big.Rat, error) {
	closePrice, err := cache.GetCachedSymbolPriceAtTime(h.rdb, token.Symbol, dealTime)
	if err != nil {
		return nil, nil, err
	}

	decimalsMultiplier := new(big.Int).Exp(big.NewInt(10), token.Decimals.Int, nil)
	volumeToken := new(big.Rat).SetFrac(amount, decimalsMultiplier)
	volumeUSD := new(big.Rat).Mul(volumeToken, closePrice)
	return volumeUSD, closePrice, nil
}

func (h *UniswapV3PoolHandler) PopulateWithFinanceInfo(interactions []trade.UniswapV3Event) ([]trade.UniswapV3Deal, error) {
	result := make([]trade.UniswapV3Deal, len(interactions))
	for i, interaction := range interactions {
		volumeAInUSD, priceAInUSD, err := h.humanVolumeOfToken(interaction.AmountTokenA.Int, &h.tokenA, &interaction.Timestamp)
		if err != nil {
			return nil, err
		}
		volumeBInUSD, priceBInUSD, err := h.humanVolumeOfToken(interaction.AmountTokenB.Int, &h.tokenB, &interaction.Timestamp)
		if err != nil {
			return nil, err
		}
		volumeTotalUSD := new(big.Rat).Add(volumeAInUSD, volumeBInUSD)
		deal := trade.NewUniswapV3Deal(priceAInUSD, priceBInUSD, volumeAInUSD, volumeBInUSD, volumeTotalUSD, interaction)
		result[i] = deal
	}
	return result, nil
}
