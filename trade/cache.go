package trade

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
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	BadRationalValue error = errors.New("Bad rational value stored in cache")
)

var ctx = context.Background()

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "redis",
		DB:       0,
	})
}

func GetCachedBlockTimestamp(client *ethclient.Client, rdb *redis.Client, block uint64) (*time.Time, error) {
	blockIdentifierStr := fmt.Sprintf("block:%d", block)
	timestampString, err := rdb.Get(ctx, blockIdentifierStr).Result()
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
		slog.Info(fmt.Sprintf("Block %s is new, fetching its date from blockchain", blockIdentifierStr))

		blockHeader, err := client.BlockByNumber(context.Background(), big.NewInt(int64(block)))
		if err != nil {
			return nil, err
		}
		blockTimestamp := blockHeader.Time()
		blockTimestampString := fmt.Sprintf("%d", blockTimestamp)
		err = rdb.Set(ctx, blockIdentifierStr, blockTimestampString, time.Hour*3).Err()
		if err != nil {
			slog.Warn(fmt.Sprintf("Cannot update value in cache %s=%s", blockIdentifierStr, blockTimestampString))
			return nil, err
		}
		slog.Info(fmt.Sprintf("Written to cache pair %s = %s", blockIdentifierStr, blockTimestampString))
		result := time.Unix(int64(blockTimestamp), 0)
		return &result, nil
	}
	slog.Info(fmt.Sprintf("Block %d is already cached with unix timestamp %s", block, timestampString))
	timestamp, err := strconv.ParseInt(timestampString, 10, 64)
	if err != nil {
		return nil, err
	}
	result := time.Unix(timestamp, 0)
	return &result, nil
}

func GetCachedSymbolPriceAtTime(rdb *redis.Client, symbol string, instant *time.Time) (*big.Rat, error) {
	instantString := fmt.Sprintf("%d", instant.UnixMilli())
	identifierStr := fmt.Sprintf("quote_%s_%s", symbol, instantString)
	quoteString, err := rdb.Get(ctx, identifierStr).Result()
	if err != nil {
		if err == redis.Nil {
			price, err := GetClosePrice(symbol, instant)
			if err != nil {
				slog.Warn(fmt.Sprintf("Cannot get price for symbol %s at instant %s: %s", symbol, instantString, err.Error()))
				return nil, err
			}
			err = rdb.Set(ctx, identifierStr, price.String(), time.Hour*3).Err()
			if err != nil {
				slog.Warn(fmt.Sprintf("Cannot update in cache price of symbol %s at instant %s ms: %s", symbol, instantString, err.Error()))
				return nil, err
			}
			return price, nil
		}
		return nil, err
	}

	slog.Info(fmt.Sprintf("Key already cached %s = %s", identifierStr, quoteString))
	quote, success := new(big.Rat).SetString(quoteString)
	if !success {
		return nil, BadRationalValue
	}
	return quote, nil
}

func calculateBalance(income []Deal, outcome []Deal) string {
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

func GetCachedBalanceOfWallet(db *gorm.DB, rdb *redis.Client, walletAddress string) (*BalanceAcrossAllChains, error) {
	cacheKey := fmt.Sprintf("balanceAcrossAllChains:%s", walletAddress)
	cachedBalance, err := rdb.Get(context.Background(), cacheKey).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if cachedBalance == "" {
		var dealsIncome []Deal
		countIncome := 0
		err = db.Preload("BlockchainTransfer").Where("blockchain_transfer.recipient = ?", walletAddress).First(&dealsIncome, &countIncome).Error
		if err != nil {
			return nil, err
		}

		var dealsOutcome []Deal
		countOutcome := 0
		err = db.Preload("BlockchainTransfer").Where("blockchain_transfer.sender = ?", walletAddress).First(&dealsIncome, &countOutcome).Error
		if err != nil {
			return nil, err
		}
		slog.Info(fmt.Sprintf("Found %d income and %d outcome deals of %s", len(dealsIncome), countOutcome, walletAddress))

		balance := calculateBalance(dealsIncome, dealsOutcome)
		cachedData, _ := json.Marshal(BalanceAcrossAllChains{Address: walletAddress, Balance: balance})
		rdb.Set(ctx, cacheKey, cachedData, 5*time.Minute)
		return NewBalanceAcrossAllChains(walletAddress, balance), nil

	} else {
		var balanceAcrossAllChains BalanceAcrossAllChains
		err = json.Unmarshal([]byte(cachedBalance), &balanceAcrossAllChains)
		if err != nil {
			return nil, err
		}
		return &balanceAcrossAllChains, nil
	}
}

func GetCachedBalanceOfWalletOnChain(db *gorm.DB, rdb *redis.Client, chainId string, walletAddress string) (*BalanceOnChain, error) {
	key := fmt.Sprintf("BalanceOnChain:%s:%s", chainId, walletAddress)
	cached, err := rdb.Get(context.Background(), key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if cached != "" {
		var result BalanceOnChain
		err = json.Unmarshal([]byte(cached), &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}

	dealsIncome := []Deal{}
	dealsOutcome := []Deal{}
	err = db.
		Preload("BlockchainTransfer").
		Find(&dealsIncome, Deal{BlockchainTransfer: ERC20Transfer{Recipient: walletAddress, ChainId: chainId}}).
		Error
	if err != nil {
		return nil, err
	}
	err = db.
		Preload("BlockchainTransfer").
		Find(&dealsOutcome, Deal{BlockchainTransfer: ERC20Transfer{Sender: walletAddress, ChainId: chainId}}).
		Error
	if err != nil {
		return nil, err
	}
	balance := calculateBalance(dealsIncome, dealsOutcome)

	result := NewBalanceOnChain(chainId, walletAddress, balance)
	err = rdb.Set(context.Background(), key, result, 15*time.Minute).Err()
	if err != nil {
		return nil, err
	}
	return result, nil
}
