package trade

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"strconv"
	"time"

	"github.com/chenzhijie/go-web3"
	"github.com/redis/go-redis/v9"
)

var (
	BadRationalValue error = errors.New("Bad rational value stored in cache")
)

var ctx = context.Background()

func GetCachedBlockTimestamp(w3 *web3.Web3, rdb *redis.Client, block uint64) (*time.Time, error) {
	blockIdentifierStr := fmt.Sprintf("block:%d", block)
	timestampString, err := rdb.Get(ctx, blockIdentifierStr).Result()
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
		slog.Info(fmt.Sprintf("Block %s is new, fetching its date from blockchain", blockIdentifierStr))

		blockHeader, err := w3.Eth.GetBlockHeaderByNumber(big.NewInt(int64(block)), false)
		if err != nil {
			return nil, err
		}
		blockTimestamp := blockHeader.Time
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
	quote , success:= new(big.Rat).SetString(quoteString)
	if !success {
		return nil, BadRationalValue
	}
	return quote, nil

	

}
