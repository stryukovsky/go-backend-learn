package worker

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"log/slog"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/protocols"
	"github.com/stryukovsky/go-backend-learn/trade/protocols/hodl"
	"gorm.io/gorm"
)

func  fetchTransfersFromEthJSONRPC[A any, B any](
	chainId string,
	db *gorm.DB,
	cache *redis.Client,
	config *trade.Worker,
	startFromBlock uint64,
	currentBlockchainBlock uint64,
	handlers []protocols.DeFiProtocolHandler[A, B],
	trackedWallets []trade.TrackedWallet,
	participants []string) {
	endInBlock := min(startFromBlock+config.BlocksInterval, currentBlockchainBlock)
	slog.Info(fmt.Sprintf("Interacting with %d tokens. Find events from block %d to %d", len(handlers), startFromBlock, endInBlock))
	for _, handler := range handlers {
		interactions, err := handler.FetchBlockchainInteractions(
			chainId,
			participants,
			startFromBlock,
			endInBlock,
		)
		if err != nil {
			slog.Warn(fmt.Sprintf("[%s] Cannot fetch blockchain interactions: %s", handler.Name(), err.Error()))
			continue
		}
		slog.Info(fmt.Sprintf("[%s] Found %d blockchain interactions where tracked wallets participated", handler.Name(), len(interactions)))
		financialInteractions, err := handler.PopulateWithFinanceInfo(interactions)
		if err != nil {
			slog.Warn(fmt.Sprintf("[%s] Cannot fetch financial interactions: %s", handler.Name(), err.Error()))
			continue
		}
		err = db.Create(financialInteractions).Error
		if err != nil {
			slog.Warn(fmt.Sprintf("[%s] Cannot save financial interactions: %s", handler.Name(), err.Error()))
			continue
		}
	}
}

func Cycle(db *gorm.DB, rdb *redis.Client, id uint) {
	var config trade.Worker
	result := db.First(&config, id)
	if result.Error != nil {
		slog.Warn("No config with id " + strconv.Itoa(int(id)))
		return
	}

	client, err := ethclient.Dial(config.BlockchainUrl)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to connect to Ethereum node: %s", err.Error()))
		return
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		slog.Warn("Cannot fetch chain id")
		return
	}

	var trackedWallets []trade.TrackedWallet
	err = db.Find(&trackedWallets, &trade.TrackedWallet{ChainId: chainId.String()}).Error
	if err != nil {
		slog.Warn("Failed to get tracked wallets")
		return
	}

	currentBlockchainBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		slog.Warn(fmt.Sprintf("Cannot get last blockchain block: %s", err.Error()))
		return
	}

	var tokensFromDB []trade.Token
	err = db.Find(&tokensFromDB, &trade.Token{ChainId: chainId.String()}).Error
	if err != nil {
		slog.Warn(fmt.Sprintf("Cannot get tokens of config: %s", err.Error()))
		return
	}

	var defiProtocolHandlers []protocols.DeFiProtocolHandler[any, any]
	for _, token := range tokensFromDB {
		erc20, err := hodl.NewHODLHandler(client, token, rdb)
		if err != nil {
			slog.Warn(fmt.Sprintf("Cannot create token %s: %e", token.Address, err))
			continue
		}
		var casted protocols.DeFiProtocolHandler[trade.ERC20Transfer, trade.Deal]
		casted = erc20
		defiProtocolHandlers = append(defiProtocolHandlers, casted)
	}

	var participants []string
	var minBlockOfWalletsToFetchFromNode uint64 = math.MaxUint64
	for _, wallet := range trackedWallets {
		slog.Info(fmt.Sprintf("Wallet %s will be updated with transfers fetched from blockchain", wallet.Address))
		participants = append(participants, wallet.Address)
		minBlockOfWalletsToFetchFromNode = min(wallet.LastBlock, minBlockOfWalletsToFetchFromNode)
	}
	if len(participants) > 0 {
		fetchTransfersFromEthJSONRPC(
			chainId.String(),
			db,
			rdb,
			&config,
			minBlockOfWalletsToFetchFromNode,
			currentBlockchainBlock,
			defiProtocolHandlers,
			trackedWallets,
			participants)
	}
}
