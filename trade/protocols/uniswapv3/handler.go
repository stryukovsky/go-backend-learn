package uniswapv3

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"github.com/stryukovsky/go-backend-learn/trade"
	"gorm.io/gorm"
)

type UniswapV3PoolHandler struct {
	pool           UniswapV3PoolInstance
	rdb            *redis.Client
	db             *gorm.DB
	name           string
	tokenA         trade.Token
	tokenB         trade.Token
	parallelFactor int
}

func (h *UniswapV3PoolHandler) ParallelFactor() int { return h.parallelFactor }

func NewUniswapV3PoolHandler(
	instance trade.DeFiPlatform,
	client *ethclient.Client,
	rdb *redis.Client,
	db *gorm.DB,
	tokenA trade.Token,
	tokenB trade.Token,
	parallelFactor int,
) (*UniswapV3PoolHandler, error) {
	pool, err := NewUniswapV3PoolInstance(client, instance.Address)
	if err != nil {
		return nil, err
	}
	return &UniswapV3PoolHandler{
		pool: *pool,
		rdb: rdb,
		db: db,
		name: fmt.Sprintf("Uniswap V3 Pool %s - %s", tokenA.Symbol, tokenB.Symbol),
		tokenA: tokenA,
		tokenB: tokenB,
		parallelFactor: parallelFactor,
	}, nil
}

func (h *UniswapV3PoolHandler) FetchBlockchainInteractions(
	chainId string,
	participants []string,
	fromBlock uint64,
	toBlock uint64,
) ([]trade.UniswapV3Event, error) {
	formattedParticipants := lo.Map(participants, func(p string, _ int) common.Address {return common.HexToAddress(p)})
	mintEventsIter, err := h.pool.filterer.FilterMint(&bind.FilterOpts{Start: fromBlock, End: &toBlock}, formattedParticipants, nil, nil)
	if err != nil {
		return nil, err
	}
	mintEventsRaw := make([]UniswapV3PoolMint, 0)
	for mintEventsIter.Next() { 
		mintEventsRaw = append(mintEventsRaw, *mintEventsIter.Event)
	}
	mintEvents := make([]trade.UniswapV3Event, len(mintEventsRaw))
	chunkSize := len(mintEventsRaw) / h.ParallelFactor()
	chunks := lo.Chunk(mintEvents, chunkSize)
	for _, chunk := range chunks {

		go func ()  {

			
		}()
	}
	
	 

	


}
