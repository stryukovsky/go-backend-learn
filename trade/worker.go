package trade

import (
	"fmt"
	"math"
	"strconv"

	"log/slog"

	"github.com/chenzhijie/go-web3"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func FetchTransfersFromNode(
	rpcUrl string,
	db *gorm.DB,
	cache *redis.Client,
	config *Worker,
	startFromBlock uint64,
	currentBlockchainBlock uint64,
	tokens []ERC20,
	trackedWallets []TrackedWallet,
	participants []string) {
	endInBlock := min(startFromBlock+config.BlocksInterval, currentBlockchainBlock)
	slog.Info(fmt.Sprintf("Interacting with %d tokens. Find events from block %d to %d", len(tokens), startFromBlock, endInBlock))
	for _, token := range tokens {
		transfers, err := token.ListTransfersOfParticipants(rpcUrl, participants, startFromBlock, endInBlock, cache)
		if err != nil {
			slog.Warn(fmt.Sprintf("[%s] Cannot fetch transfers: %s", token.Symbol, err.Error()))
			continue
		}
		slog.Info(fmt.Sprintf("[%s] Found %d transfers where tracked wallets participated", token.Symbol, len(transfers)))
		for _, transfer := range transfers {
			deal, err := CreateDeal(cache, transfer, token)
			if err != nil {
				slog.Warn(fmt.Sprintf("[%s] Cannot create deal object for ERC20 transfer %s: %s", token.Symbol, transfer.TxId, err.Error()))
			} else {
				slog.Info(fmt.Sprintf("[%s] Found deal with volume $ %s", token.Symbol, deal.VolumeUSD.FloatString(5)))
				db.Save(deal)
			}
		}
	}
}

func Cycle(db *gorm.DB, cache *redis.Client, id uint) {
	var config Worker
	result := db.First(&config, id)
	if result.Error != nil {
		slog.Warn("No config with id " + strconv.Itoa(int(id)))
		return
	}

	web3, err := web3.NewWeb3(config.BlockchainUrl)
	if err != nil {
		slog.Warn(fmt.Sprintf("Cannot connect to blockchain: %e", err))
		return
	}

	chainId, err := web3.Eth.ChainID()
	if err != nil {
		slog.Warn("Cannot fetch chain id")
		return
	}

	var trackedWallets []TrackedWallet
	err = db.Find(&trackedWallets, &TrackedWallet{ChainId: chainId.String()}).Error
	if err != nil {
		slog.Warn("Failed to get tracked wallets")
		return
	}

	currentBlockchainBlock, err := web3.Eth.GetBlockNumber()
	if err != nil {
		slog.Warn(fmt.Sprintf("Cannot get last blockchain block: %s", err.Error()))
		return
	}

	var tokensFromDB []Token
	err = db.Find(&tokensFromDB, &Token{ChainId: chainId.String()}).Error
	if err != nil {
		slog.Warn(fmt.Sprintf("Cannot get tokens of config: %s", err.Error()))
		return
	}

	tokens := make([]ERC20, 0, len(tokensFromDB))
	for _, token := range tokensFromDB {
		erc20, err := CreateERC20(web3, token.Address, token.Symbol)
		if err != nil {
			slog.Warn(fmt.Sprintf("Cannot create token %s: %e", token.Address, err))
			continue
		}
		tokens = append(tokens, *erc20)
	}

	var participants []string
	var minBlockOfWalletsToFetchFromNode uint64 = math.MaxUint64
	for _, wallet := range trackedWallets {
		// criteriaWalletIsOutDated := wallet.LastBlock < currentBlockchainBlock-config.BlocksInterval
		// if criteriaWalletIsOutDated {
			slog.Info(fmt.Sprintf("Wallet %s will be updated with transfers fetched from blockchain", wallet.Address))
			participants = append(participants, wallet.Address)
		// } 
		minBlockOfWalletsToFetchFromNode = min(wallet.LastBlock, minBlockOfWalletsToFetchFromNode)
	}
	if len(participants) > 0 {
		FetchTransfersFromNode(config.BlockchainUrl, db,
			cache,
			&config,
			minBlockOfWalletsToFetchFromNode,
			currentBlockchainBlock,
			tokens,
			trackedWallets,
			participants)
	}
}
