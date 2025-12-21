package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/binance"
	"gorm.io/gorm"
)

var BadRationalValue error = errors.New("Bad rational value stored in cache")

type CacheEthJSONRPC struct {
	RpcUrl string
	Eth    ethclient.Client
}

type CacheManager struct {
	clients []CacheEthJSONRPC
	rdb     redis.Client
}

func NewCacheManager(ethereumUrls pq.StringArray, redisAddr, redisPassword string, redisDb int) (*CacheManager, error) {
	clients := make([]CacheEthJSONRPC, len(ethereumUrls))
	for i, url := range ethereumUrls {
		client, err := ethclient.Dial(url)
		if err != nil {
			return nil, err
		}
		clients[i] = CacheEthJSONRPC{Eth: *client, RpcUrl: url}
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDb,
	})
	return &CacheManager{clients, *rdb}, nil
}

func (cm *CacheManager) GetBasicClient() *ethclient.Client {
	return &cm.clients[0].Eth
}

func (cm *CacheManager) GetReadonlyClient() *CacheEthJSONRPC {
	result := trade.RandomChoice(cm.clients[1:])
	return &result
}

func (cm *CacheManager) Set(key string, value any) error {
	return cm.rdb.Set(ctx, key, value, 0).Err()
}

func (cm *CacheManager) SetWithTTL(key string, value any, ttl time.Duration) error {
	return cm.rdb.Set(ctx, key, value, ttl).Err()
}

func (cm *CacheManager) Get(key string) (string, error) {
	return cm.rdb.Get(ctx, key).Result()
}

var ctx = context.Background()

func (cm *CacheManager) GetCachedBlockTimestamp(block uint64) (*time.Time, error) {
	blockIdentifierStr := fmt.Sprintf("block:%d", block)
	timestampString, err := cm.Get(blockIdentifierStr)
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
		slog.Debug(fmt.Sprintf("[Cache] Block %s is new, fetching its date from blockchain", blockIdentifierStr))

		client := cm.GetReadonlyClient()
		blockHeader, err := client.Eth.HeaderByNumber(context.Background(), big.NewInt(int64(block)))
		if err != nil {
			slog.Warn(fmt.Sprintf("[Cache] Cannot get block header for %d block using %s: %s", block, client.RpcUrl, err.Error()))
			return nil, err
		}
		time.Sleep(time.Second * 1)
		blockTimestamp := blockHeader.Time
		if blockTimestamp <= 0 {
			return nil, fmt.Errorf("[Cache] Invalid timestamp. Timestamp: %d", blockTimestamp)
		}
		blockTimestampString := fmt.Sprintf("%d", blockTimestamp)
		err = cm.Set(blockIdentifierStr, blockTimestampString)
		if err != nil {
			slog.Warn(fmt.Sprintf("[Cache] Cannot update value in cache %s=%s", blockIdentifierStr, blockTimestampString))
			return nil, err
		}
		slog.Debug(fmt.Sprintf("[Cache] Written to cache pair %s = %s", blockIdentifierStr, blockTimestampString))
		result := time.Unix(int64(blockTimestamp), 0)
		return &result, nil
	}
	slog.Debug(fmt.Sprintf("[Cache] Block %d is already cached with unix timestamp %s", block, timestampString))
	timestamp, err := strconv.ParseInt(timestampString, 10, 64)
	if err != nil {
		return nil, err
	}
	result := time.Unix(timestamp, 0)
	return &result, nil
}

func (cm *CacheManager) GetCachedSymbolPriceAtTime(symbol string, instant *time.Time) (*big.Rat, error) {
	truncated := instant.Truncate(5 * time.Minute)
	instantString := fmt.Sprintf("%d", truncated.UnixMilli())
	identifierStr := fmt.Sprintf("quote:%s:%s", symbol, instantString)
	quoteString, err := cm.Get(identifierStr)
	if err != nil {
		if err == redis.Nil {
			price, err := binance.GetClosePrice(symbol, &truncated)
			if err != nil {
				slog.Warn(fmt.Sprintf("[Cache] Cannot get price for symbol %s at instant %s: %s", symbol, instantString, err.Error()))
				return nil, err
			}
			err = cm.Set(identifierStr, price.String())
			if err != nil {
				slog.Warn(fmt.Sprintf("[Cache] Cannot update in cache price of symbol %s at instant %s ms: %s", symbol, instantString, err.Error()))
				return nil, err
			}
			return price, nil
		}
		return nil, err
	}

	slog.Debug(fmt.Sprintf("[Cache] Key already cached %s = %s", identifierStr, quoteString))
	quote, success := new(big.Rat).SetString(quoteString)
	if !success {
		return nil, BadRationalValue
	}
	return quote, nil
}

func calculateBalance(income []trade.Deal, outcome []trade.Deal) string {
	result := big.NewRat(0, 1)
	for _, deal := range income {
		result = result.Add(result, deal.VolumeUSD.Rat)
	}
	for _, deal := range outcome {
		result = result.Sub(result, deal.VolumeUSD.Rat)
	}

	balance := result.FloatString(2)
	return balance
}

func (cm *CacheManager) GetCachedBalanceOfWallet(db *gorm.DB, walletAddress string) (*trade.BalanceAcrossAllChains, error) {
	cacheKey := fmt.Sprintf("balanceAcrossAllChains:%s", walletAddress)
	cachedBalance, err := cm.Get(cacheKey)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if cachedBalance == "" {
		var dealsIncome []trade.Deal
		countIncome := 0
		err = db.Preload("BlockchainTransfer").Where("blockchain_transfer.recipient = ?", walletAddress).First(&dealsIncome, &countIncome).Error
		if err != nil {
			return nil, err
		}

		var dealsOutcome []trade.Deal
		countOutcome := 0
		err = db.Preload("BlockchainTransfer").Where("blockchain_transfer.sender = ?", walletAddress).First(&dealsIncome, &countOutcome).Error
		if err != nil {
			return nil, err
		}
		slog.Debug(fmt.Sprintf("Found %d income and %d outcome deals of %s", len(dealsIncome), countOutcome, walletAddress))

		balance := calculateBalance(dealsIncome, dealsOutcome)
		cachedData, _ := json.Marshal(trade.BalanceAcrossAllChains{Address: walletAddress, Balance: balance})
		cm.SetWithTTL(cacheKey, cachedData, 5*time.Minute)
		return trade.NewBalanceAcrossAllChains(walletAddress, balance), nil

	} else {
		var balanceAcrossAllChains trade.BalanceAcrossAllChains
		err = json.Unmarshal([]byte(cachedBalance), &balanceAcrossAllChains)
		if err != nil {
			return nil, err
		}
		return &balanceAcrossAllChains, nil
	}
}

func (cm *CacheManager) GetCachedBalanceOfWalletOnChain(db *gorm.DB, chainId string, walletAddress string) (*trade.BalanceOnChain, error) {
	key := fmt.Sprintf("BalanceOnChain:%s:%s", chainId, walletAddress)
	cached, err := cm.Get(key)
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if cached != "" {
		var result trade.BalanceOnChain
		err = json.Unmarshal([]byte(cached), &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}

	dealsIncome := []trade.Deal{}
	dealsOutcome := []trade.Deal{}
	err = db.
		Preload("BlockchainTransfer").
		Find(&dealsIncome, trade.Deal{BlockchainTransfer: trade.ERC20Transfer{Recipient: walletAddress, ChainId: chainId}}).
		Error
	if err != nil {
		return nil, err
	}
	err = db.
		Preload("BlockchainTransfer").
		Find(&dealsOutcome, trade.Deal{BlockchainTransfer: trade.ERC20Transfer{Sender: walletAddress, ChainId: chainId}}).
		Error
	if err != nil {
		return nil, err
	}
	balance := calculateBalance(dealsIncome, dealsOutcome)

	result := trade.NewBalanceOnChain(chainId, walletAddress, balance)
	err = cm.SetWithTTL(key, result, 15*time.Minute)
	if err != nil {
		return nil, err
	}
	return result, nil
}
