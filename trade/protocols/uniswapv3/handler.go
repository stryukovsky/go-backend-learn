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
	"github.com/samber/lo"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/cache"
	"gorm.io/gorm"
)

type UniswapV3PoolHandler struct {
	pool            UniswapV3PoolInstance
	positionManager NFPositionManagerInstance
	cm              *cache.CacheManager
	db              *gorm.DB
	name            string
	tokenA          trade.Token
	tokenB          trade.Token
	parallelFactor  int
	chainId         string
}

func (h *UniswapV3PoolHandler) ParallelFactor() int { return h.parallelFactor }

func NewUniswapV3PoolHandler(
	instance trade.DeFiPlatform,
	client *ethclient.Client,
	cm *cache.CacheManager,
	db *gorm.DB,
	parallelFactor int,
) (*UniswapV3PoolHandler, error) {
	pool, err := NewUniswapV3PoolInstance(client, instance.Address)
	if err != nil {
		return nil, err
	}
	nfPositionManager, err := NewNFPositionManagerInstance(client, instance.ExtraContractAddress1)
	if err != nil {
		return nil, err
	}
	tokenAddressA, err := pool.caller.Token0(nil)
	if err != nil {
		return nil, err
	}
	tokenAddressB, err := pool.caller.Token1(nil)
	if err != nil {
		return nil, err
	}
	var tokenA trade.Token
	db.First(&tokenA, trade.Token{ChainId: instance.ChainId, Address: tokenAddressA.Hex()})
	var tokenB trade.Token
	db.First(&tokenB, trade.Token{ChainId: instance.ChainId, Address: tokenAddressB.Hex()})

	return &UniswapV3PoolHandler{
		pool:            *pool,
		positionManager: *nfPositionManager,
		cm:              cm,
		db:              db,
		name:            fmt.Sprintf("Uniswap V3 Pool %s - %s", tokenA.Symbol, tokenB.Symbol),
		tokenA:          tokenA,
		tokenB:          tokenB,
		parallelFactor:  parallelFactor,
		chainId:         instance.ChainId,
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
	timestamp, err := h.cm.GetCachedBlockTimestamp(
		event.Raw.BlockNumber,
	)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot get block timestamp from cache/blockchain: %s", h.Name(), err.Error()))
		return nil, err
	}
	if timestamp.IsZero() {
		slog.Warn(fmt.Sprintf("[%s] Block timestamp is zero", h.Name()))
	}
	result := trade.NewUniswapV3Event(
		h.chainId,
		trade.UniswapV3Mint,
		event.Owner.Hex(),
		h.pool.Address.Hex(),
		event.Amount0,
		event.Amount1,
		lowerPrice,
		upperPrice,
		*timestamp,
		event.Raw.TxHash.Hex(),
		event.Raw.Index,
		event.Raw.BlockNumber,
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
	timestamp, err := h.cm.GetCachedBlockTimestamp(
		event.Raw.BlockNumber,
	)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot get block timestamp from cache/blockchain: %s", h.Name(), err.Error()))
		return nil, err
	}
	if timestamp.IsZero() {
		slog.Warn(fmt.Sprintf("[%s] Block timestamp is zero", h.Name()))
	}
	result := trade.NewUniswapV3Event(
		h.chainId,
		trade.UniswapV3Burn,
		event.Owner.Hex(),
		h.pool.Address.Hex(),
		event.Amount0,
		event.Amount1,
		lowerPrice,
		upperPrice,
		*timestamp,
		event.Raw.TxHash.Hex(),
		event.Raw.Index,
		event.Raw.BlockNumber,
	)
	return &result, nil
}

func (h *UniswapV3PoolHandler) parseSwap(event UniswapV3PoolSwap) (*trade.UniswapV3Event, error) {
	price, err := SqrtPrice2Price(event.SqrtPriceX96, h.tokenA, h.tokenB)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot parse price of swap event: %s", h.Name(), err.Error()))
		return nil, err
	}
	timestamp, err := h.cm.GetCachedBlockTimestamp(
		event.Raw.BlockNumber,
	)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot get block timestamp from cache/blockchain: %s", h.Name(), err.Error()))
		return nil, err
	}
	if timestamp.IsZero() {
		slog.Warn(fmt.Sprintf("[%s] Block timestamp is zero", h.Name()))
	}
	result := trade.NewUniswapV3Event(
		h.chainId,
		trade.UniswapV3Swap,
		event.Sender.Hex(),
		h.pool.Address.Hex(),
		event.Amount0,
		event.Amount1,
		price,
		price,
		*timestamp,
		event.Raw.TxHash.Hex(),
		event.Raw.Index,
		event.Raw.BlockNumber,
	)
	return &result, nil
}

func (h *UniswapV3PoolHandler) parseCollect(event UniswapV3PoolCollect) (*trade.UniswapV3Event, error) {
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
	timestamp, err := h.cm.GetCachedBlockTimestamp(
		event.Raw.BlockNumber,
	)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot get block timestamp from cache/blockchain: %s", h.Name(), err.Error()))
		return nil, err
	}
	if timestamp.IsZero() {
		slog.Warn(fmt.Sprintf("[%s] Block timestamp is zero", h.Name()))
	}
	result := trade.NewUniswapV3Event(
		h.chainId,
		trade.UniswapV3Collect,
		event.Owner.Hex(),
		h.pool.Address.Hex(),
		event.Amount0,
		event.Amount1,
		lowerPrice,
		upperPrice,
		*timestamp,
		event.Raw.TxHash.Hex(),
		event.Raw.Index,
		event.Raw.BlockNumber,
	)
	return &result, nil
}

// We identify liquidity interaction (mint/burn liquidity)
// As unique combination of liquidity amount and corresponding amount of token0 and token1
type LiquidityActionIdentity struct {
	Amount  string
	Amount0 string
	Amount1 string
}

func NewLiquidityActionIdentity(liquidityAmount *big.Int, amount0 *big.Int, amount1 *big.Int) LiquidityActionIdentity {
	return LiquidityActionIdentity{
		Amount:  liquidityAmount.String(),
		Amount0: amount0.String(),
		Amount1: amount1.String(),
	}
}

var addressZero = common.BigToAddress(big.NewInt(0))

// pool events are of
// UniswapV3PoolMint
// UniswapV3PoolBurn
// UniswapV3PoolCollect
// UniswapV3PoolSwap
//
// position manager liquidity events are of
func (h *UniswapV3PoolHandler) parseEvents(
	poolEvents []any,
	pmLiquidityEvents []any,
	actualWalletsBurnedLiquidity map[string]common.Address,
	positions []trade.UniswapV3Position,
) ([]trade.UniswapV3Event, error) {
	actualWalletsMintedLiquidity := make(map[string]common.Address)
	for _, position := range positions {
		actualWalletsMintedLiquidity[position.TokenId] = common.HexToAddress(position.Owner)
	}

	liquidityAdded := make(map[LiquidityActionIdentity]string)
	liquidityRemoved := make(map[LiquidityActionIdentity]string)
	feesCollected := make(map[LiquidityActionIdentity]string)
	for _, liquidityEvent := range pmLiquidityEvents {
		switch casted := liquidityEvent.(type) {
		case INonFungiblePositionsManagerIncreaseLiquidity:
			liquidityAdded[NewLiquidityActionIdentity(casted.Liquidity, casted.Amount0, casted.Amount1)] = casted.TokenId.String()
		case INonFungiblePositionsManagerDecreaseLiquidity:
			liquidityRemoved[NewLiquidityActionIdentity(casted.Liquidity, casted.Amount0, casted.Amount1)] = casted.TokenId.String()
		case INonFungiblePositionsManagerCollect:
			feesCollected[NewLiquidityActionIdentity(big.NewInt(0), casted.Amount0, casted.Amount1)] = casted.TokenId.String()
		}
	}

	chunkSize := len(poolEvents) / h.ParallelFactor()
	if chunkSize == 0 {
		return []trade.UniswapV3Event{}, nil
	}
	chunks := lo.Chunk(poolEvents, chunkSize)
	chunksCount := len(chunks)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(chunksCount)
	resultCh := make(chan trade.UniswapV3Event, h.ParallelFactor())
	for i, chunk := range chunks {
		go func() {
			slog.Debug(fmt.Sprintf("[%s] Starting %d worker", h.Name(), i))
			defer wg.Done()
			for _, uncastedEvent := range chunk {
				select {
				case <-ctx.Done():
					return
				default:
					switch castedEvent := uncastedEvent.(type) {
					case UniswapV3PoolMint:
						parsedEvent, err := h.parseMint(castedEvent)
						if err != nil {
							cancel()
							return
						} else {
							liquidityIdentity := NewLiquidityActionIdentity(castedEvent.Amount, castedEvent.Amount0, castedEvent.Amount1)
							if tokenId, ok := liquidityAdded[liquidityIdentity]; ok {
								if walletAddress, ok := actualWalletsMintedLiquidity[tokenId]; ok {
									slog.Info(fmt.Sprintf(
										"[%s] Wallet %s has minted token %s which corresponds to liquidity event being parsed",
										h.Name(),
										walletAddress.Hex(),
										tokenId,
									))
									parsedEvent.WalletAddress = walletAddress.Hex()
									parsedEvent.PositionTokenId = tokenId
								} else {
									slog.Warn(fmt.Sprintf(
										"[%s] Token with tokenId %s found, but no wallet holding it found", h.Name(), tokenId))
								}
							} else {
								slog.Warn(fmt.Sprintf("[%s] liquidity minted, but no token ID found", h.Name()))
							}
							resultCh <- *parsedEvent
						}
					case UniswapV3PoolBurn:
						parsedEvent, err := h.parseBurn(castedEvent)
						if err != nil {
							cancel()
							return
						} else {
							liquidityIdentity := NewLiquidityActionIdentity(castedEvent.Amount, castedEvent.Amount0, castedEvent.Amount1)
							if tokenId, ok := liquidityRemoved[liquidityIdentity]; ok {
								if walletAddress, ok := actualWalletsBurnedLiquidity[tokenId]; ok {
									slog.Info(fmt.Sprintf(
										"Wallet %s has burned token %s which corresponds to liquidity event being parsed",
										walletAddress.Hex(),
										tokenId,
									))
									parsedEvent.WalletAddress = walletAddress.Hex()
									parsedEvent.PositionTokenId = tokenId
								}
							}
							resultCh <- *parsedEvent
						}
					case UniswapV3PoolSwap:
						parsedEvent, err := h.parseSwap(castedEvent)
						if err != nil {
							cancel()
							return
						} else {
							resultCh <- *parsedEvent
						}
					case UniswapV3PoolCollect:
						parsedEvent, err := h.parseCollect(castedEvent)
						if err != nil {
							cancel()
							return
						} else {
							liquidityIdentity := NewLiquidityActionIdentity(big.NewInt(0), castedEvent.Amount0, castedEvent.Amount1)
							if tokenId, ok := feesCollected[liquidityIdentity]; ok {
								if walletAddress, ok := actualWalletsBurnedLiquidity[tokenId]; ok {
									slog.Info(fmt.Sprintf(
										"Wallet %s has collected fees on liquidity token %s which corresponds to liquidity event being parsed",
										walletAddress.Hex(),
										tokenId,
									))
									parsedEvent.WalletAddress = walletAddress.Hex()
									parsedEvent.PositionTokenId = tokenId
								}
							}
							resultCh <- *parsedEvent
						}
					default:
						slog.Info(fmt.Sprintf("[%s] Skip event of type %s since no parsing implemented for it", h.Name(), uncastedEvent))
						continue
					}
				}
			}
			slog.Debug(fmt.Sprintf("[%s] %d worker finished parsing events", h.Name(), i))
		}()
	}
	results := make([]trade.UniswapV3Event, 0, len(poolEvents))
	go func() {
		wg.Wait()
		close(resultCh)
	}()
	for item := range resultCh {
		select {
		case <-ctx.Done():
			return []trade.UniswapV3Event{}, errors.New("Some worker goroutine encountered error. Details in logs")
		default:
			results = append(results, item)
		}
	}
	return results, nil
}

func (h *UniswapV3PoolHandler) fetchPoolLiquidityEvents(fromBlock uint64, toBlock uint64) ([]any, error) {
	mintEventsIter, err := h.pool.filterer.FilterMint(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock},
		nil,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	burnEventsIter, err := h.pool.filterer.FilterBurn(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock},
		nil,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	swapEventsIter, err := h.pool.filterer.FilterSwap(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock},
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	liquidityPoolEvents := make([]any, 0, 100)
	for mintEventsIter.Next() {
		liquidityPoolEvents = append(liquidityPoolEvents, *mintEventsIter.Event)
	}
	for burnEventsIter.Next() {
		liquidityPoolEvents = append(liquidityPoolEvents, *burnEventsIter.Event)
	}
	for swapEventsIter.Next() {
		liquidityPoolEvents = append(liquidityPoolEvents, *swapEventsIter.Event)
	}
	return liquidityPoolEvents, nil
}

func (h *UniswapV3PoolHandler) fetchPositionsManagerLiquidityEvents(fromBlock uint64, toBlock uint64) ([]any, error) {
	// Parse INonFungiblePositionsManagerIncreaseLiquidity event
	liquidityAddedIter, err := h.positionManager.filterer.FilterIncreaseLiquidity(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock},
		nil,
	)
	if err != nil {
		return nil, err
	}

	// Parse INonFungiblePositionsManagerDecreaseLiquidity event
	liquidityRemovedIter, err := h.positionManager.filterer.FilterDecreaseLiquidity(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock},
		nil,
	)
	if err != nil {
		return nil, err
	}

	feesCollected, err := h.positionManager.filterer.FilterCollect(&bind.FilterOpts{Start: fromBlock, End: &toBlock}, nil)
	if err != nil {
		return nil, err
	}

	// Parse INonFungiblePositionsManagerIncreaseLiquidity event
	liquidityPositionManagerEvents := make([]any, 0, 100)
	for liquidityAddedIter.Next() {
		liquidityPositionManagerEvents = append(liquidityPositionManagerEvents, *liquidityAddedIter.Event)
	}
	for liquidityRemovedIter.Next() {
		liquidityPositionManagerEvents = append(liquidityPositionManagerEvents, *liquidityRemovedIter.Event)
	}
	for feesCollected.Next() {
		liquidityPositionManagerEvents = append(liquidityPositionManagerEvents, *feesCollected.Event)
	}
	return liquidityPositionManagerEvents, nil
}

func (h *UniswapV3PoolHandler) fetchERC721TransferEvents(fromBlock uint64, toBlock uint64) ([]INonFungiblePositionsManagerTransfer, error) {
	erc721TransferEventsIter, err := h.positionManager.filterer.FilterTransfer(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock},
		nil,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	transferEvents := make([]INonFungiblePositionsManagerTransfer, 0, 100)
	for erc721TransferEventsIter.Next() {
		transferEvents = append(transferEvents, *erc721TransferEventsIter.Event)
	}
	return transferEvents, nil
}

func (h *UniswapV3PoolHandler) FetchLiquidityInteractions(
	chainId string,
	fromBlock uint64,
	toBlock uint64,
) ([]trade.UniswapV3Event, []trade.UniswapV3Position, error) {
	// Parse ERC721 Transfer events
	transferEvents, err := h.fetchERC721TransferEvents(fromBlock, toBlock)
	if err != nil {
		return nil, nil, err
	}
	mintedPositions := make([]trade.UniswapV3Position, 0, len(transferEvents))
	positionBurnedEvents := make(map[string]common.Address)
	var alreadyMintedPositions []trade.UniswapV3Position
	err = h.db.Find(&alreadyMintedPositions, trade.UniswapV3Position{
		ChainId:                 h.chainId,
		UniswapPositionsManager: h.positionManager.Address.Hex(),
	}).Error
	if err != nil {
		return nil, nil, err
	}
	for _, event := range transferEvents {
		if event.From == addressZero {
			slog.Debug(fmt.Sprintf(
				"[%s] Found new position minted to %s with id %s",
				h.Name(),
				event.To.Hex(),
				event.TokenId.String(),
			))
			mintedPositions = append(mintedPositions, trade.NewUniswapV3Position(
				h.chainId,
				h.positionManager.Address,
				event.TokenId,
				event.To,
			))
		}
		if event.To == addressZero {
			positionBurnedEvents[event.TokenId.String()] = event.From
		}
	}
	allPositionsAvailableAtTheMoment := append(alreadyMintedPositions, mintedPositions...)
	liquidityPoolEvents, err := h.fetchPoolLiquidityEvents(fromBlock, toBlock)
	if err != nil {
		return nil, nil, err
	}
	if len(liquidityPoolEvents) == 0 {
		slog.Warn(fmt.Sprintf("[%s] no events in block range %d - %d", h.Name(), fromBlock, toBlock))
		return make([]trade.UniswapV3Event, 0), mintedPositions, nil
	} else {
		slog.Info(fmt.Sprintf("[%s] found %d events in block range %d - %d", h.Name(), len(liquidityPoolEvents), fromBlock, toBlock))
	}

	liquidityPositionManagerEvents, err := h.fetchPositionsManagerLiquidityEvents(fromBlock, toBlock)
	if err != nil {
		return nil, nil, err
	}

	result, err := h.parseEvents(liquidityPoolEvents, liquidityPositionManagerEvents, positionBurnedEvents, allPositionsAvailableAtTheMoment)
	if err != nil {
		return nil, nil, err
	}
	return result, mintedPositions, nil
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
	events, _, err := h.FetchLiquidityInteractions(chainId, fromBlock, toBlock)
	if err != nil {
		return nil, err
	}
	// FIXME: filter and also return created positions
	return events, nil
}

func (h *UniswapV3PoolHandler) humanVolumeOfToken(amount *big.Int, token *trade.Token, dealTime *time.Time) (*big.Rat, *big.Rat, *big.Rat, error) {
	closePrice, err := h.cm.GetCachedSymbolPriceAtTime(token.Symbol, dealTime)
	if err != nil {
		return nil, nil, nil, err
	}

	decimalsMultiplier := new(big.Int).Exp(big.NewInt(10), token.Decimals.Int, nil)
	volumeToken := new(big.Rat).SetFrac(amount, decimalsMultiplier)
	volumeUSD := new(big.Rat).Mul(volumeToken, closePrice)
	return volumeUSD, volumeToken, closePrice, nil
}

func (h *UniswapV3PoolHandler) PopulateWithFinanceInfoConcurrently(interactions []trade.UniswapV3Event) ([]trade.UniswapV3Deal, error) {
	chunkSize := len(interactions) / h.ParallelFactor()
	if chunkSize == 0 {
		return []trade.UniswapV3Deal{}, nil
	}
	resultCh := make(chan trade.UniswapV3Deal)
	chunks := lo.Chunk(interactions, chunkSize)
	var wg sync.WaitGroup
	wg.Add(len(chunks))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for _, chunk := range chunks {
		go func() {
			defer wg.Done()
			for _, interaction := range chunk {
				select {
				case <-ctx.Done():
					return
				default:
					volumeAInUSD, volumeA, priceAInUSD, err := h.humanVolumeOfToken(
						interaction.AmountTokenA.Int,
						&h.tokenA,
						&interaction.Timestamp,
					)
					if err != nil {
						slog.Warn(fmt.Sprintf(
							"[%s] Error on token %s volume and price calculation in USD: %s",
							h.Name(),
							h.tokenA.Symbol,
							err.Error(),
						))
						cancel()
						return
					}

					volumeBInUSD, volumeB, priceBInUSD, err := h.humanVolumeOfToken(interaction.AmountTokenB.Int, &h.tokenB, &interaction.Timestamp)
					if err != nil {
						slog.Warn(fmt.Sprintf("[%s] Error on token %s volume and price calculation in USD: %s", h.Name(), h.tokenB.Symbol, err.Error()))
						cancel()
						return
					}
					volumeTotalUSD := new(big.Rat).Add(volumeAInUSD, volumeBInUSD)
					deal := trade.NewUniswapV3Deal(
						h.tokenA.Symbol,
						h.tokenB.Symbol,
						priceAInUSD,
						priceBInUSD,
						volumeAInUSD,
						volumeBInUSD,
						volumeA,
						volumeB,
						volumeTotalUSD,
						interaction,
					)
					resultCh <- deal
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	result := make([]trade.UniswapV3Deal, len(interactions))
	i := 0
	for deal := range resultCh {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			result[i] = deal
			i++
		}
	}
	return result, nil
}

// two methods exist because we possibly can reach out of binance api limits if asking it too frequent
func (h *UniswapV3PoolHandler) PopulateWithFinanceInfo(interactions []trade.UniswapV3Event) ([]trade.UniswapV3Deal, error) {
	result := make([]trade.UniswapV3Deal, len(interactions))
	for i, interaction := range interactions {
		volumeAInUSD, volumeA, priceAInUSD, err := h.humanVolumeOfToken(interaction.AmountTokenA.Int, &h.tokenA, &interaction.Timestamp)
		if err != nil {
			return nil, err
		}

		volumeBInUSD, volumeB, priceBInUSD, err := h.humanVolumeOfToken(interaction.AmountTokenB.Int, &h.tokenB, &interaction.Timestamp)
		if err != nil {
			return nil, err
		}
		volumeTotalUSD := new(big.Rat).Add(volumeAInUSD, volumeBInUSD)
		deal := trade.NewUniswapV3Deal(
			h.tokenA.Symbol,
			h.tokenB.Symbol,
			priceAInUSD,
			priceBInUSD,
			volumeAInUSD,
			volumeBInUSD,
			volumeA,
			volumeB,
			volumeTotalUSD,
			interaction,
		)
		result[i] = deal
	}
	return result, nil
}
