package trade

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"log/slog"

	"github.com/chenzhijie/go-web3"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func FetchTransfersFromNode(
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
		transfers, err := token.ListTransfersOfParticipants(participants, startFromBlock, endInBlock, cache)
		if err != nil {
			slog.Warn(fmt.Sprintf("[%s] Cannot fetch transfers: %s", token.Symbol, err.Error()))
			continue
		}
		slog.Info(fmt.Sprintf("[%s] Found %d transfers where tracked wallets participated", token.Symbol, len(transfers)))
		for _, transfer := range transfers {
			deal, err := CreateDeal(cache, transfer)
			if err != nil {
				slog.Warn(fmt.Sprintf("[%s] Cannot create deal object for ERC20 transfer %s: %s", token.Symbol, transfer.TxId, err.Error()))
			} else {
				slog.Info(fmt.Sprintf("[%s] Found deal with volume $ %s", token.Symbol, deal.VolumeUSD.FloatString(5)))
				db.Save(deal)
			}
		}
	}
}

func Cycle(db *gorm.DB, id uint) {
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

	cache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "redis",
		DB:       0,
	})

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

	// criteria to consider wallet info is outdated (i.e. 100 intervals ago was updated)
	criteria := 100 * config.BlocksInterval

	alchemyIsAvailable := strings.TrimSpace(config.AlchemyApiUrl) != ""

	// wallets which are actual enough to fetch transfers from blockchain
	walletsToFetchFromNode := make([]TrackedWallet, 0, len(trackedWallets))
	// minimal block number of every wallet
	var minBlockOfWalletsToFetchFromNode uint64 = math.MaxUint64
	for _, wallet := range trackedWallets {
		startFromBlock := wallet.LastBlock
		if currentBlockchainBlock-startFromBlock > criteria && alchemyIsAvailable {
			transfers, err := AlchemyGetTransfersForAccount(web3, cache, config, wallet)
			if err != nil {
				slog.Warn(fmt.Sprintf("Error when interacting with alchemy: %s", err.Error()))
				continue
			}
			db.CreateInBatches(transfers, len(transfers))
			wallet.LastBlock = currentBlockchainBlock
			db.Save(wallet)
		} else {
			walletsToFetchFromNode = append(walletsToFetchFromNode, wallet)
			minBlockOfWalletsToFetchFromNode = min(minBlockOfWalletsToFetchFromNode, wallet.LastBlock)
		}
	}
	var participants []string
	for _, wallet := range walletsToFetchFromNode {
		slog.Info(fmt.Sprintf("Wallet %s will be updated with transfers fetched from blockchain", wallet.Address))
		participants = append(participants, strings.ToLower(wallet.Address))
	}
	if len(walletsToFetchFromNode) > 0 {
		FetchTransfersFromNode(db,
			cache,
			&config,
			minBlockOfWalletsToFetchFromNode,
			currentBlockchainBlock,
			tokens,
			trackedWallets,
			participants)
	}
}
