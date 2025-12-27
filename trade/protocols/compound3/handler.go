package compound3

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
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/cache"
	"gorm.io/gorm"
)

type Compound3Handler struct {
	compoundCometContract Compound3
	cm                    *cache.CacheManager
	db                    *gorm.DB
	name                  string
	tokens                []trade.Token
	parallelFactor        int
}

func (h *Compound3Handler) ParallelFactor() int { return h.parallelFactor }

func NewCompound3Handler(
	instance trade.DeFiPlatform,
	client *ethclient.Client,
	rdb *cache.CacheManager,
	tokens []trade.Token,
	parallelFactor int,
) (*Compound3Handler, error) {
	compoundComet, err := NewCompound3(client, instance.Address)
	if err != nil {
		return nil, err
	}
	return &Compound3Handler{
		compoundCometContract: *compoundComet,
		cm:                    rdb,
		name:                  fmt.Sprintf("Compound3 on %s", instance.Address),
		tokens:                tokens,
		parallelFactor:        parallelFactor,
	}, nil
}

func (h *Compound3Handler) parseCompound3Events(chainId string, events []any) ([]trade.Compound3Event, error) {
	if len(events) == 0 {
		return []trade.Compound3Event{}, nil
	}
	eventChunks := trade.Chunks(events, h.ParallelFactor())
	var wg sync.WaitGroup
	wg.Add(len(eventChunks))
	valuesCh := make(chan trade.Compound3Event, h.ParallelFactor())
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
						slog.Info(fmt.Sprintf("[%s] Unexpected event type %s in chunk of Events", h.Name(), generalEvent))
					case CometSupply:
						var event CometSupply = generalEvent
						timestamp, err := h.cm.GetCachedBlockTimestamp(event.Raw.BlockNumber)
						if err != nil {
							slog.Warn(fmt.Sprintf("[%s] Failure on parsing Supply event %s", h.Name(), err.Error()))
							wg.Done()
							cancel()
						}
						item := trade.NewCompound3Event(
							chainId,
							"supply",
							event.Dst,
							h.compoundCometContract.MainAsset,
							event.Amount,
							*timestamp,
							event.Raw.TxHash.Hex(),
							event.Raw.Index,
						)
						valuesCh <- item

					case CometSupplyCollateral:
						var event CometSupplyCollateral = generalEvent
						timestamp, err := h.cm.GetCachedBlockTimestamp(event.Raw.BlockNumber)
						if err != nil {
							slog.Warn(fmt.Sprintf("[%s] Failure on parsing Supply event %s", h.Name(), err.Error()))
							wg.Done()
							cancel()
						}
						item := trade.NewCompound3Event(
							chainId,
							"supply",
							event.Dst,
							event.Asset,
							event.Amount,
							*timestamp,
							event.Raw.TxHash.Hex(),
							event.Raw.Index,
						)
						valuesCh <- item

					case CometWithdrawCollateral:
						var event CometWithdrawCollateral = generalEvent
						timestamp, err := h.cm.GetCachedBlockTimestamp(event.Raw.BlockNumber)
						if err != nil {
							slog.Warn(fmt.Sprintf("[%s] Failure on parsing Withdraw event %s", h.Name(), err.Error()))
							wg.Done()
							cancel()
						}
						item := trade.NewCompound3Event(
							chainId,
							"withdraw",
							event.To,
							event.Asset,
							event.Amount,
							*timestamp,
							event.Raw.TxHash.Hex(),
							event.Raw.Index,
						)
						valuesCh <- item

					case CometWithdraw:
						var event CometWithdraw = generalEvent
						timestamp, err := h.cm.GetCachedBlockTimestamp(event.Raw.BlockNumber)
						if err != nil {
							slog.Warn(fmt.Sprintf("[%s] Failure on parsing Withdraw event %s", h.Name(), err.Error()))
							wg.Done()
							cancel()
						}
						item := trade.NewCompound3Event(
							chainId,
							"withdraw",
							event.To,
							h.compoundCometContract.MainAsset,
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
	// go func() {
	// 	wg.Wait()
	// 	close(valuesCh)
	// }()
	result := make([]trade.Compound3Event, 0, len(events))
	for item := range valuesCh {
		result = append(result, item)
	}
	cancel()
	return result, nil
}

func (h *Compound3Handler) FetchBlockchainInteractions(
	chainId string,
	participants []string,
	fromBlock uint64,
	toBlock uint64,
) ([]trade.Compound3Event, error) {
	formattedParticipants := make([]common.Address, len(participants))
	for i, p := range participants {
		formattedParticipants[i] = common.HexToAddress(p)
	}
	supplyEventsIter, err := h.compoundCometContract.filterer.FilterSupply(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock},
		[]common.Address{},
		formattedParticipants,
	)
	if err != nil {
		return nil, err
	}
	defer supplyEventsIter.Close()

	collateralSupplyEventsIter, err := h.compoundCometContract.filterer.FilterSupplyCollateral(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock},
		[]common.Address{},
		formattedParticipants,
		[]common.Address{},
	)
	if err != nil {
		return nil, err
	}
	defer collateralSupplyEventsIter.Close()

	withdrawEventsIter, err := h.compoundCometContract.filterer.FilterWithdraw(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock}, []common.Address{}, formattedParticipants)
	if err != nil {
		return nil, err
	}
	defer withdrawEventsIter.Close()

	collateralWithdrawEventsIter, err := h.compoundCometContract.filterer.FilterWithdrawCollateral(
		&bind.FilterOpts{Start: fromBlock, End: &toBlock}, []common.Address{}, formattedParticipants, []common.Address{})
	if err != nil {
		return nil, err
	}
	defer collateralWithdrawEventsIter.Close()
	// any is because go do not support generic methods, we have two types for each event: Supply and Withdraw
	eventsRaw := make([]any, 0)
	for supplyEventsIter.Next() {
		if err = supplyEventsIter.Error(); err != nil {
			return nil, err
		}
		eventsRaw = append(eventsRaw, *supplyEventsIter.Event)
	}
	for collateralSupplyEventsIter.Next() {
		if err = collateralSupplyEventsIter.Error(); err != nil {
			return nil, err
		}
		eventsRaw = append(eventsRaw, *collateralSupplyEventsIter.Event)
	}
	for withdrawEventsIter.Next() {
		err := withdrawEventsIter.Error()
		if err != nil {
			return nil, err
		}
		eventsRaw = append(eventsRaw, *withdrawEventsIter.Event)
	}
	for collateralWithdrawEventsIter.Next() {
		err := collateralWithdrawEventsIter.Error()
		if err != nil {
			return nil, err
		}
		eventsRaw = append(eventsRaw, *collateralWithdrawEventsIter.Event)
	}
	if len(eventsRaw) == 0 {
		slog.Warn(fmt.Sprintf("[%s] no events in block range %d - %d", h.Name(), fromBlock, toBlock))
		return make([]trade.Compound3Event, 0), nil
	} else {
		slog.Info(fmt.Sprintf("[%s] found %d events in block range %d - %d", h.Name(), len(eventsRaw), fromBlock, toBlock))
	}
	events, err := h.parseCompound3Events(chainId, eventsRaw)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (h *Compound3Handler) PopulateWithFinanceInfo(interactions []trade.Compound3Event) ([]trade.Compound3Interaction, error) {
	result := make([]trade.Compound3Interaction, len(interactions))
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

		closePrice, err := h.cm.GetCachedSymbolPriceAtTime(token.Symbol, &interaction.Timestamp)
		if err != nil {
			return nil, err
		}

		volumeToken := big.NewRat(1, 1)
		decimalsMultiplier := new(big.Int).Exp(big.NewInt(10), token.Decimals.Int, nil)
		volumeToken = volumeToken.SetFrac(interaction.Amount.Int, decimalsMultiplier)

		volumeUSD := new(big.Rat).Mul(volumeToken, closePrice)
		deal := trade.Compound3Interaction{
			Price:           trade.NewDBNumeric(closePrice),
			VolumeTokens:    trade.NewDBNumeric(volumeToken),
			VolumeUSD:       trade.NewDBNumeric(volumeUSD),
			BlockchainEvent: interaction,
		}
		result[i] = deal
	}
	return result, nil
}
func (h *Compound3Handler) Name() string { return h.name }
