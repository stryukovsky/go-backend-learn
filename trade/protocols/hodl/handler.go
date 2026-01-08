package hodl

import (
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/cache"
	"github.com/stryukovsky/go-backend-learn/trade/web3client"
)

type HODLHandler struct {
	token          ERC20
	cm             *cache.CacheManager
	parallelFactor int
}

func NewHODLHandler(client *web3client.MultiURLClient, token trade.Token, cm *cache.CacheManager, parallelFactor int) (*HODLHandler, error) {
	erc20, err := NewERC20(client, token)
	if err != nil {
		return nil, err
	}
	return &HODLHandler{
		token:          *erc20,
		cm:             cm,
		parallelFactor: parallelFactor,
	}, nil
}

func (h *HODLHandler) FetchBlockchainInteractions(
	chainId string,
	participants []string,
	fromBlock uint64,
	toBlock uint64,
) ([]trade.ERC20Transfer, error) {
	formattedParticipants := make([]common.Address, len(participants))
	for i, participant := range participants {
		formattedParticipants[i] = common.HexToAddress(participant)
	}
	transfersParticipantsSenders, err := h.token.filterer.FilterTransfer(&bind.FilterOpts{
		Start: fromBlock,
		End:   &toBlock,
	}, formattedParticipants, []common.Address{})
	if err != nil {
		return []trade.ERC20Transfer{}, err
	}
	transfersParticipantsRecipients, err := h.token.filterer.FilterTransfer(&bind.FilterOpts{
		Start: fromBlock,
		End:   &toBlock,
	}, []common.Address{}, formattedParticipants)
	if err != nil {
		return []trade.ERC20Transfer{}, err
	}
	allTransfers := make([]IERC20Transfer, 0)
	for transfersParticipantsSenders.Next() {
		allTransfers = append(allTransfers, *transfersParticipantsSenders.Event)
	}
	for transfersParticipantsRecipients.Next() {
		allTransfers = append(allTransfers, *transfersParticipantsRecipients.Event)
	}
	if len(allTransfers) == 0 {
		return []trade.ERC20Transfer{}, nil
	}
	slog.Info(fmt.Sprintf("Scanned %d transfers", len(allTransfers)))
	return trade.ParseEVMEvents(h.parallelFactor, h.token.Info.Symbol, chainId, allTransfers, func(task *trade.ParallelEVMParserTask[trade.ERC20Transfer], event IERC20Transfer) {
		sender := event.From
		recipient := event.To
		amount := event.Value
		txId := event.Raw.TxHash
		block := event.Raw.BlockNumber
		timestamp, err := h.cm.GetCachedBlockTimestamp(block)
		if err != nil {
			slog.Warn(fmt.Sprintf("Cannot fetch from cache or blockchain info on block %d timestamp: %s", block, err.Error()))
		} else {
			transfer := trade.NewERC20Transfer(h.token.Info.Address, sender.String(), recipient.String(), amount, block, chainId, timestamp, txId.Hex(), event.Raw.Index)
			task.ValuesCh <- transfer
		}
	},
	)
}

func (h *HODLHandler) ParallelFactor() int { return h.parallelFactor }

func (h *HODLHandler) PopulateWithFinanceInfo(interactions []trade.ERC20Transfer) ([]trade.Deal, error) {
	result := make([]trade.Deal, len(interactions))
	for i, transfer := range interactions {
		closePrice, err := h.cm.GetCachedSymbolPriceAtTime(h.token.Info.Symbol, &transfer.Timestamp)
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
