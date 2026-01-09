package worker

import (
	"fmt"
	"log/slog"
	"math"
	"strconv"

	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/cache"
	"github.com/stryukovsky/go-backend-learn/trade/protocols"
	"github.com/stryukovsky/go-backend-learn/trade/protocols/aave"
	"github.com/stryukovsky/go-backend-learn/trade/protocols/compound3"
	"github.com/stryukovsky/go-backend-learn/trade/protocols/hodl"
	"github.com/stryukovsky/go-backend-learn/trade/protocols/uniswapv3"
	"github.com/stryukovsky/go-backend-learn/trade/web3client"
	"gorm.io/gorm"
)

const (
	ParallelFactor = 16
)


func Cycle(db *gorm.DB, cm *cache.CacheManager, id uint) {
	slog.Info("Starting worker")
	var config trade.Worker
	result := db.First(&config, id)
	if result.Error != nil {
		slog.Warn("No config with id " + strconv.Itoa(int(id)))
		return
	}

	client, err := web3client.NewMultiURLClient(config.BlockchainUrlsForEvents)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to connect to Ethereum node: %s", err.Error()))
		return
	}

	chainId, err := client.ChainID()
	if err != nil {
		slog.Warn(fmt.Sprintf("Cannot fetch chain id: %s", err.Error()))
		return
	}

	var trackedWallets []trade.TrackedWallet
	err = db.Find(&trackedWallets, &trade.TrackedWallet{ChainId: chainId.String()}).Error
	if err != nil {
		slog.Warn("Failed to get tracked wallets")
		return
	}

	currentBlockchainBlock, err := client.BlockNumber()
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
	var participants []string
	var minBlockOfWalletsToFetchFromNode uint64 = math.MaxUint64
	for _, wallet := range trackedWallets {
		slog.Info(fmt.Sprintf("Wallet %s will be updated with transfers fetched from blockchain", wallet.Address))
		participants = append(participants, wallet.Address)
		minBlockOfWalletsToFetchFromNode = min(wallet.LastBlock, minBlockOfWalletsToFetchFromNode)
	}
	if minBlockOfWalletsToFetchFromNode == math.MaxUint64 {
		slog.Warn("Cannot determine where to start indexing. Maybe there is no tracked wallets?")
		return
	}
	startBlock := minBlockOfWalletsToFetchFromNode
	endBlock := min(startBlock+config.BlocksInterval, currentBlockchainBlock)

	var erc20Handlers []protocols.DeFiProtocolHandler[trade.ERC20Transfer, trade.Deal]
	for _, token := range tokensFromDB {
		erc20, err := hodl.NewHODLHandler(client, token, cm, ParallelFactor)
		if err != nil {
			slog.Warn(fmt.Sprintf("Cannot create token %s: %e", token.Address, err))
			continue
		}
		var casted protocols.DeFiProtocolHandler[trade.ERC20Transfer, trade.Deal]
		casted = erc20
		erc20Handlers = append(erc20Handlers, casted)
	}

	var aaveInstances []trade.DeFiPlatform
	err = db.Find(&aaveInstances, &trade.DeFiPlatform{ChainId: chainId.String(), Type: trade.Aave}).Error
	if err != nil {
		slog.Warn(fmt.Sprintf("Cannot get aave platform instances: %s", err.Error()))
		return
	}
	var aaveHandlers []protocols.DeFiProtocolHandler[trade.AaveEvent, trade.AaveInteraction]
	for _, aaveInstance := range aaveInstances {
		var tokens []trade.Token
		db.Find(&tokens, trade.Token{ChainId: aaveInstance.ChainId})
		aaveHandler, err := aave.NewAaveHandler(aaveInstance, client, cm, tokens, ParallelFactor)
		if err != nil {
			slog.Warn(fmt.Sprintf("Cannot get aave platform handler: %s", err.Error()))
			continue
		}
		var casted protocols.DeFiProtocolHandler[trade.AaveEvent, trade.AaveInteraction]
		casted = aaveHandler
		aaveHandlers = append(aaveHandlers, casted)
	}

	var compoundInstances []trade.DeFiPlatform
	err = db.Find(&compoundInstances, &trade.DeFiPlatform{ChainId: chainId.String(), Type: trade.Compound3}).Error
	if err != nil {
		slog.Warn(fmt.Sprintf("Cannot get compound platform instances: %s", err.Error()))
		return
	}
	var compoundHandlers []protocols.DeFiProtocolHandler[trade.Compound3Event, trade.Compound3Interaction]
	for _, compoundInstance := range compoundInstances {
		var tokens []trade.Token
		db.Find(&tokens, trade.Token{ChainId: compoundInstance.ChainId})
		compoundHandler, err := compound3.NewCompound3Handler(compoundInstance, client, cm, tokens, ParallelFactor)
		if err != nil {
			slog.Warn(fmt.Sprintf("Cannot get compound platform handler: %s", err.Error()))
			continue
		}
		var casted protocols.DeFiProtocolHandler[trade.Compound3Event, trade.Compound3Interaction]
		casted = compoundHandler
		compoundHandlers = append(compoundHandlers, casted)
	}

	var uniswapV3Pools []trade.DeFiPlatform
	err = db.Find(&uniswapV3Pools, &trade.DeFiPlatform{ChainId: chainId.String(), Type: trade.UniswapV3}).Error
	if err != nil {
		slog.Warn(fmt.Sprintf("Cannot get UniswapV3 pools: %s", err.Error()))
		return
	}
	var uniswapv3Handlers []protocols.DeFiProtocolHandler[trade.UniswapV3Event, trade.UniswapV3Deal]
	for _, uniswapv3Instance := range uniswapV3Pools {
		uniswapv3Handler, err := uniswapv3.NewUniswapV3PoolHandler(
			uniswapv3Instance,
			client,
			cm,
			db,
			ParallelFactor,
		)
		if err != nil {
			slog.Warn(fmt.Sprintf("Cannot get uniswapv3 platform handler: %s", err.Error()))
			continue
		}
		var casted protocols.DeFiProtocolHandler[trade.UniswapV3Event, trade.UniswapV3Deal]
		casted = uniswapv3Handler
		uniswapv3Handlers = append(uniswapv3Handlers, casted)
	}
	if len(participants) == 0 {
		return
	}

	// Environment is ready to setup
	env := NewFetchEnvironment(chainId.String(), db, trackedWallets, participants, erc20Handlers, aaveHandlers, compoundHandlers, uniswapv3Handlers)
	env.Fetch(startBlock, endBlock)
}
