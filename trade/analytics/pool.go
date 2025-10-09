package analytics

import (
	"context"
	"fmt"
	"strconv"

	"log/slog"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgconn"
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
	config *trade.Worker,
	startFromBlock uint64,
	currentBlockchainBlock uint64,
	handler *uniswapv3.UniswapV3PoolHandler,
) error {
	endInBlock := min(startFromBlock+config.BlocksInterval, currentBlockchainBlock)
	blockchainInteractions, err := handler.FetchLiquidityInteractions(
		chainId,
		startFromBlock,
		endInBlock,
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
		return nil
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
	return nil
}

func Analyze(startBlock uint64, blocksCount uint64, poolAddress string, db *gorm.DB, rdb *redis.Client, id uint) {
	lastBlockToAnalyze := startBlock + blocksCount
	currentBlock := startBlock
	slog.Info("Starting worker")
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

	for currentBlock <= lastBlockToAnalyze {
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
		endBlock := min(currentBlock+config.BlocksInterval, lastBlockInBlockchain, lastBlockToAnalyze)
		dbTx := db
		err = fetchInteractionsFromEthJSONRPC(
			chainId.String(),
			dbTx,
			&config,
			currentBlock,
			endBlock,
			uniswapv3Handler,
		)
		if err != nil {
			slog.Info(fmt.Sprintf("Cannot fetch UniswapV3 interactions due to %s", err.Error()))
			return
		} else {
			slog.Info(fmt.Sprintf("Successfully fetched blockchain events so mark wallets as indexed on block %d", endBlock))
			currentBlock = endBlock + 1
		}
	}
}
