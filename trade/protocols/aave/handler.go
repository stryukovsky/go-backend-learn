package aave

import "github.com/redis/go-redis/v9"

type AaveHandler struct {

	pool AavePool
	rdb *redis.Client

}

func NewAaveHandler(address string)
