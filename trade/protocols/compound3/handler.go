package compound3

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
	client *web3client.MultiURLClient,
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

	return trade.ParseEVMEvents(
		h.ParallelFactor(),
		h.Name(), chainId, eventsRaw, func(task trade.ParallelEVMParserTask[trade.Compound3Event], generalEvent any) error {
			switch generalEvent := generalEvent.(type) {
			default:
				slog.Info(fmt.Sprintf("[%s] Unexpected event type %s in chunk of Events", h.Name(), generalEvent))
			case CometSupply:
				var event CometSupply = generalEvent
				timestamp, err := h.cm.GetCachedBlockTimestamp(event.Raw.BlockNumber)
				if err != nil {
					slog.Warn(fmt.Sprintf("[%s] Failure on parsing Supply event %s", h.Name(), err.Error()))
					return err
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
				task.ValuesCh <- item

			case CometSupplyCollateral:
				var event CometSupplyCollateral = generalEvent
				timestamp, err := h.cm.GetCachedBlockTimestamp(event.Raw.BlockNumber)
				if err != nil {
					slog.Warn(fmt.Sprintf("[%s] Failure on parsing Supply event %s", h.Name(), err.Error()))
					return err
				}
				item := trade.NewCompound3Event(
					chainId,
					"supply_collateral",
					event.Dst,
					event.Asset,
					event.Amount,
					*timestamp,
					event.Raw.TxHash.Hex(),
					event.Raw.Index,
				)
				task.ValuesCh <- item

			case CometWithdrawCollateral:
				var event CometWithdrawCollateral = generalEvent
				timestamp, err := h.cm.GetCachedBlockTimestamp(event.Raw.BlockNumber)
				if err != nil {
					slog.Warn(fmt.Sprintf("[%s] Failure on parsing Withdraw event %s", h.Name(), err.Error()))
					return err
				}
				item := trade.NewCompound3Event(
					chainId,
					"withdraw_collateral",
					event.To,
					event.Asset,
					event.Amount,
					*timestamp,
					event.Raw.TxHash.Hex(),
					event.Raw.Index,
				)
				task.ValuesCh <- item

			case CometWithdraw:
				var event CometWithdraw = generalEvent
				timestamp, err := h.cm.GetCachedBlockTimestamp(event.Raw.BlockNumber)
				if err != nil {
					slog.Warn(fmt.Sprintf("[%s] Failure on parsing Withdraw event %s", h.Name(), err.Error()))
					return err
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
				task.ValuesCh <- item
			}
			return nil
		})
}

func (h *Compound3Handler) PopulateWithFinanceInfo(interactions []trade.Compound3Event) ([]trade.Compound3Interaction, error) {
	result := make([]trade.Compound3Interaction, 0, len(interactions))
	for _, interaction := range interactions {
		tokenAddress := common.HexToAddress(interaction.TokenAddress)
		token := trade.Token{}
		for _, t := range h.tokens {
			if strings.EqualFold(t.Address, tokenAddress.Hex()) {
				token = t
			}
		}

		if len(token.Address) == 0 {
			slog.Warn(fmt.Sprintf("Found compound interaction with unknown token address %s", tokenAddress))
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
		result = append(result, deal)
	}
	return result, nil
}
func (h *Compound3Handler) Name() string { return h.name }
