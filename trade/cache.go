package trade

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/chenzhijie/go-web3"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func GetBlockTimestamp(w3 *web3.Web3, rdb *redis.Client, block uint64) (*time.Time, error) {
	blockIdentifierStr := fmt.Sprintf("%d", block)
	timestampString, err := rdb.Get(ctx, blockIdentifierStr).Result()
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}

		blockHeader, err := w3.Eth.GetBlockHeaderByNumber(big.NewInt(int64(block)), false)
		if err != nil {
			return nil, err
		}
		blockTimestamp := blockHeader.Time
		rdb.Set(ctx, fmt.Sprintf("%d", blockIdentifierStr), fmt.Sprintf("%d", blockTimestamp), time.Hour*3)
		result := time.Unix(int64(blockTimestamp), 0)
		return &result, nil
	}
	timestamp, err := strconv.ParseInt(timestampString, 10, 64)
	if err != nil {

		return nil, err
	}
	result := time.Unix(timestamp, 0)
	return &result, nil
}
