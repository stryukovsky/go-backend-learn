package worker

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/cache"
	"github.com/stryukovsky/go-backend-learn/trade/protocols"
	"github.com/stryukovsky/go-backend-learn/trade/protocols/aave"
	"github.com/stryukovsky/go-backend-learn/trade/protocols/compound3"
	"github.com/stryukovsky/go-backend-learn/trade/protocols/hodl"
	"github.com/stryukovsky/go-backend-learn/trade/protocols/uniswapv3"
	"gorm.io/gorm"
)

const (
	ParallelFactor = 16
)

func fetchInteractionsFromEthJSONRPC[A any, B any](
	chainId string,
	db *gorm.DB,
	startBlock uint64,
	endBlock uint64,
	handlers []protocols.DeFiProtocolHandler[A, B],
	participants []string,
) error {
	for _, handler := range handlers {
		blockchainInteractions, err := handler.FetchBlockchainInteractions(
			chainId,
			participants,
			startBlock,
			endBlock,
		)
		if len(blockchainInteractions) > 0 {
			for _, blockchainInteraction := range blockchainInteractions {
				err = db.Create(&blockchainInteraction).Error
				if err != nil {
					slog.Warn(fmt.Sprintf("[%s] Cannot save blockchain interaction: %s", handler.Name(), err.Error()))
					if _, ok := err.(*pgconn.PgError); ok {
						// ignore error if duplicate
						err = nil
					} else {
						return err
					}
				}
			}
		}
		if err != nil {
			slog.Warn(fmt.Sprintf("[%s] Cannot fetch blockchain interactions: %s", handler.Name(), err.Error()))
			return err
		}
		if len(blockchainInteractions) == 0 {
			slog.Warn(fmt.Sprintf("[%s] No blockchain interactions found", handler.Name()))
			continue
		}
		slog.Info(fmt.Sprintf(
			"[%s] Found %d blockchain interactions where tracked wallets participated",
			handler.Name(),
			len(blockchainInteractions)))
		financialInteractions, err := handler.PopulateWithFinanceInfo(blockchainInteractions)
		if err != nil {
			slog.Warn(fmt.Sprintf("[%s] Cannot fetch financial interactions: %s", handler.Name(), err.Error()))
			return err
		}
		for _, financialInteraction := range financialInteractions {
			err = db.Create(&financialInteraction).Error
			if err != nil {
				slog.Warn(fmt.Sprintf("[%s] Cannot save financial interaction: %s", handler.Name(), err.Error()))
				if _, ok := err.(*pgconn.PgError); ok {
					// ignore error if duplicate
					err = nil
				} else {
					return err
				}
			}
		}
	}
	return nil
}

func Cycle(db *gorm.DB, cm *cache.CacheManager, id uint) {
	slog.Info("Starting worker")
	var config trade.Worker
	result := db.First(&config, id)
	if result.Error != nil {
		slog.Warn("No config with id " + strconv.Itoa(int(id)))
		return
	}

	client, err := ethclient.Dial(config.BlockchainUrls[0])
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to connect to Ethereum node: %s", err.Error()))
		return
	}

	chainId, err := client.ChainID(context.Background())
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
	if len(participants) > 0 {
		// tx := db.Begin()
		tx := db
		err = fetchInteractionsFromEthJSONRPC(
			chainId.String(),
			tx,
			startBlock,
			endBlock,
			erc20Handlers,
			participants,
		)
		if err != nil {
			slog.Info(fmt.Sprintf("Cannot fetch ERC20 transfers due to %s", err.Error()))
		}

		err = fetchInteractionsFromEthJSONRPC(
			chainId.String(),
			tx,
			startBlock,
			endBlock,
			aaveHandlers,
			participants,
		)
		if err != nil {
			slog.Info(fmt.Sprintf("Cannot fetch Aave interactions due to %s", err.Error()))
		}

		err = fetchInteractionsFromEthJSONRPC(
			chainId.String(),
			tx,
			startBlock,
			endBlock,
			compoundHandlers,
			participants,
		)
		if err != nil {
			slog.Info(fmt.Sprintf("Cannot fetch Compound3 interactions due to %s", err.Error()))
		}

		err = fetchInteractionsFromEthJSONRPC(
			chainId.String(),
			tx,
			startBlock,
			endBlock,
			uniswapv3Handlers,
			participants,
		)
		if err != nil {
			slog.Info(fmt.Sprintf("Cannot fetch UniswapV3 interactions due to %s", err.Error()))
		}

		if err == nil {
			for i := range trackedWallets {
				trackedWallets[i].LastBlock = endBlock
			}
			slog.Info(fmt.Sprintf("Successfully fetched blockchain events so mark wallets as indexed on block %d", endBlock))
			db.Save(trackedWallets)
		}
	}
}
