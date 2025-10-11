package analytics

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/protocols"
	"github.com/stryukovsky/go-backend-learn/trade/protocols/uniswapv3"
	"gorm.io/gorm"
)

const (
	ParallelFactor = 16
)

func fetchInteractionsFromEthJSONRPC(
	chainId string,
	db *gorm.DB,
	startBlock uint64,
	endBlock uint64,
	handler *uniswapv3.UniswapV3PoolHandler,
) error {
	blockchainInteractions, mintedPositions, err := handler.FetchLiquidityInteractions(
		chainId,
		startBlock,
		endBlock,
	)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot fetch blockchain interactions: %s", handler.Name(), err.Error()))
		return err
	}
	if len(mintedPositions) > 0 {
		err = db.Save(mintedPositions).Error
		if err != nil {
			slog.Warn(fmt.Sprintf("[%s] Cannot save minted uniswapv3 positions: %s", handler.Name(), err.Error()))
			return err
		}
	}
	if len(blockchainInteractions) == 0 {
		slog.Warn(fmt.Sprintf("[%s] No blockchain interactions", handler.Name()))
		return nil
	}
	err = db.Save(blockchainInteractions).Error
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot save blockchain interactions: %s", handler.Name(), err.Error()))
		return err
	}
	slog.Info(fmt.Sprintf(
		"[%s] Found %d blockchain interactions where tracked wallets participated",
		handler.Name(),
		len(blockchainInteractions)))
	financialInteractions, err := handler.PopulateWithFinanceInfoConcurrently(blockchainInteractions)
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot fetch financial interactions: %s", handler.Name(), err.Error()))
		return err
	}
	err = db.Save(financialInteractions).Error
	if err != nil {
		slog.Warn(fmt.Sprintf("[%s] Cannot save financial interactions: %s", handler.Name(), err.Error()))
		return err
	}

	return nil
}

func Analyze(blocksCount uint64, poolAddress string, db *gorm.DB, rdb *redis.Client) {
	slog.Info("Starting worker")
	var config trade.AnalyticsWorker
	result := db.First(&config)
	if result.Error != nil {
		slog.Warn("No config")
		return
	}
	startBlock := config.LastBlock
	currentBlock := startBlock

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

	var uniswapV3Pool trade.DeFiPlatform
	err = db.Find(&uniswapV3Pool, &trade.DeFiPlatform{
		ChainId: chainId.String(),
		Type:    trade.UniswapV3,
		Address: poolAddress,
	}).Error
	if err != nil {
		slog.Warn(fmt.Sprintf("Cannot get UniswapV3 pool: %s", err.Error()))
		return
	}
	var uniswapv3Handlers []protocols.DeFiProtocolHandler[trade.UniswapV3Event, trade.UniswapV3Deal]
	uniswapv3Handler, err := uniswapv3.NewUniswapV3PoolHandler(
		uniswapV3Pool,
		client,
		rdb,
		db,
		ParallelFactor,
	)
	if err != nil {
		slog.Warn(fmt.Sprintf("Cannot get uniswapv3 platform handler: %s", err.Error()))
		return
	}
	var casted protocols.DeFiProtocolHandler[trade.UniswapV3Event, trade.UniswapV3Deal]
	casted = uniswapv3Handler
	uniswapv3Handlers = append(uniswapv3Handlers, casted)

	for {
		lastBlockInBlockchain, err := client.BlockNumber(context.Background())
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
		endBlock := min(currentBlock+config.BlocksInterval, lastBlockInBlockchain)
		if endBlock-startBlock < 50 {
			slog.Info("Seems we've reached the top of blockchain. Sleep for 3 minutes")
			time.Sleep(3 * time.Minute)
			continue
		}
		dbTx := db
		err = fetchInteractionsFromEthJSONRPC(
			chainId.String(),
			dbTx,
			currentBlock,
			endBlock,
			uniswapv3Handler,
		)
		if err != nil {
			slog.Info(fmt.Sprintf("Cannot fetch UniswapV3 interactions due to %s", err.Error()))
			return
		} else {
			slog.Info(fmt.Sprintf("Successfully fetched blockchain events so mark worker indexed on block %d", endBlock))
			config.LastBlock = endBlock
			db.Save(&config)
			currentBlock = endBlock + 1
		}
	}
}
