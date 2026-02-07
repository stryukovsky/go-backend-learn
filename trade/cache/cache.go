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
	clients := make([]CacheEthJSONRPC, 0)
	for _, url := range ethereumUrls {
		client, err := ethclient.Dial(url)
		if err != nil {
			slog.Warn(fmt.Sprintf("Failed to connect to JSON RPC %s: %v", url, err))
		} else {
			clients = append(clients, CacheEthJSONRPC{Eth: *client, RpcUrl: url})
		}
	}
	if len(clients) == 0 {
		return nil, fmt.Errorf("No available JSON RPC found")
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
	if len(cm.clients) > 1 {
		result := trade.RandomChoice(cm.clients[1:])
		return &result
	} else if len(cm.clients) == 1 {
		return &cm.clients[0]
	} else {
		panic("No eth clients for cache client")
	}
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
		err = db.Preload("BlockchainTransfer").
			Joins("JOIN erc20_transfers ON erc20_transfers.id = deals.blockchain_transfer_id").
			Where("erc20_transfers.recipient = ?", walletAddress).
			Find(&dealsIncome).Error
		if err != nil {
			return nil, err
		}

		var dealsOutcome []trade.Deal
		err = db.Preload("BlockchainTransfer").
			Joins("JOIN erc20_transfers ON erc20_transfers.id = deals.blockchain_transfer_id").
			Where("erc20_transfers.sender = ?", walletAddress).
			Find(&dealsOutcome).Error
		if err != nil {
			return nil, err
		}
		slog.Debug(fmt.Sprintf("Found %d income and %d outcome deals of %s", len(dealsIncome), len(dealsOutcome), walletAddress))

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
	err = db.Preload("BlockchainTransfer").
		Joins("JOIN erc20_transfers ON erc20_transfers.id = deals.blockchain_transfer_id").
		Where("erc20_transfers.recipient = ? AND erc20_transfers.chain_id = ?", walletAddress, chainId).
		Find(&dealsIncome).Error
	if err != nil {
		return nil, err
	}

	dealsOutcome := []trade.Deal{}
	err = db.Preload("BlockchainTransfer").
		Joins("JOIN erc20_transfers ON erc20_transfers.id = deals.blockchain_transfer_id").
		Where("erc20_transfers.sender = ? AND erc20_transfers.chain_id = ?", walletAddress, chainId).
		Find(&dealsOutcome).Error
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

func calculateTokenBalance(income []trade.Deal, outcome []trade.Deal) string {
	result := big.NewRat(0, 1)
	for _, deal := range income {
		result = result.Add(result, deal.VolumeTokens.Rat)
	}
	for _, deal := range outcome {
		result = result.Sub(result, deal.VolumeTokens.Rat)
	}
	balance := result.FloatString(6)
	return balance
}

func (cm *CacheManager) GetCachedTokenBalancesByChain(db *gorm.DB, chainId string) ([]trade.TokenBalanceByChain, error) {
	cacheKey := fmt.Sprintf("tokenBalancesByChain:%s", chainId)
	cachedData, err := cm.Get(cacheKey)

	if err != nil && err != redis.Nil {
		return nil, err
	}

	if "a" == "" {
		var result []trade.TokenBalanceByChain
		err = json.Unmarshal([]byte(cachedData), &result)
		if err != nil {
			return nil, err
		}
		slog.Debug(fmt.Sprintf("[Cache] Token balances for chain %s retrieved from cache", chainId))
		return result, nil
	}

	// Get all tracked wallets for this chain
	var wallets []trade.TrackedWallet
	err = db.Find(&wallets, trade.TrackedWallet{ChainId: chainId}).Error
	if err != nil {
		return nil, err
	}

	if len(wallets) == 0 {
		return []trade.TokenBalanceByChain{}, nil
	}

	// Map to aggregate balances: tokenAddress -> TokenBalanceByChain
	tokenBalancesMap := make(map[string]*trade.TokenBalanceByChain)

	// For each wallet, get deals and aggregate by token
	for _, wallet := range wallets {
		walletAddr := wallet.Address

		// Get income deals
		// In the section where we query deals, replace with:
		var dealsIncome []trade.Deal
		err = db.Preload("BlockchainTransfer").
			Joins("JOIN erc20_transfers ON erc20_transfers.id = deals.blockchain_transfer_id").
			Where("erc20_transfers.recipient = ? AND erc20_transfers.chain_id = ?", walletAddr, chainId).
			Find(&dealsIncome).Error
		if err != nil {
			slog.Warn(fmt.Sprintf("Error getting income deals for wallet %s: %v", walletAddr, err))
			continue
		}

		// Get outcome deals
		var dealsOutcome []trade.Deal
		err = db.Preload("BlockchainTransfer").
			Joins("JOIN erc20_transfers ON erc20_transfers.id = deals.blockchain_transfer_id").
			Where("erc20_transfers.sender = ? AND erc20_transfers.chain_id = ?", walletAddr, chainId).
			Find(&dealsOutcome).Error
		if err != nil {
			slog.Warn(fmt.Sprintf("Error getting outcome deals for wallet %s: %v", walletAddr, err))
			continue
		}

		// Group deals by token for this wallet
		tokenDealsIncome := make(map[string][]trade.Deal)
		tokenDealsOutcome := make(map[string][]trade.Deal)

		for _, deal := range dealsIncome {
			tokenAddr := deal.BlockchainTransfer.TokenAddress
			tokenDealsIncome[tokenAddr] = append(tokenDealsIncome[tokenAddr], deal)
		}

		for _, deal := range dealsOutcome {
			tokenAddr := deal.BlockchainTransfer.TokenAddress
			tokenDealsOutcome[tokenAddr] = append(tokenDealsOutcome[tokenAddr], deal)
		}

		// Calculate balance for each token this wallet has
		allTokens := make(map[string]bool)
		for token := range tokenDealsIncome {
			allTokens[token] = true
		}
		for token := range tokenDealsOutcome {
			allTokens[token] = true
		}

		for tokenAddr := range allTokens {
			income := tokenDealsIncome[tokenAddr]
			outcome := tokenDealsOutcome[tokenAddr]
			balance := calculateTokenBalance(income, outcome)

			// Skip zero balances
			balanceRat, _ := new(big.Rat).SetString(balance)
			if balanceRat.Sign() == 0 {
				continue
			}

			// Get or create token balance entry
			if _, exists := tokenBalancesMap[tokenAddr]; !exists {
				// Get token symbol from database
				var token trade.Token
				err = db.First(&token, trade.Token{ChainId: chainId, Address: tokenAddr}).Error
				symbol := "UNKNOWN"
				if err == nil {
					symbol = token.Symbol
				}

				tokenBalancesMap[tokenAddr] = &trade.TokenBalanceByChain{
					ChainId:      chainId,
					TokenAddress: tokenAddr,
					TokenSymbol:  symbol,
					TotalBalance: "0",
					Wallets:      []trade.WalletBalance{},
				}
			}

			// Add wallet balance
			tokenBalancesMap[tokenAddr].Wallets = append(
				tokenBalancesMap[tokenAddr].Wallets,
				trade.WalletBalance{
					WalletAddress: walletAddr,
					Balance:       balance,
				},
			)

			// Update total balance
			currentTotal, _ := new(big.Rat).SetString(tokenBalancesMap[tokenAddr].TotalBalance)
			walletBalance, _ := new(big.Rat).SetString(balance)
			newTotal := new(big.Rat).Add(currentTotal, walletBalance)
			tokenBalancesMap[tokenAddr].TotalBalance = newTotal.FloatString(6)
		}
	}

	// Convert map to slice
	result := make([]trade.TokenBalanceByChain, 0, len(tokenBalancesMap))
	for _, tokenBalance := range tokenBalancesMap {
		result = append(result, *tokenBalance)
	}

	// Cache the result
	resultJSON, err := json.Marshal(result)
	if err != nil {
		slog.Warn(fmt.Sprintf("[Cache] Failed to marshal token balances for chain %s: %v", chainId, err))
	} else {
		err = cm.SetWithTTL(cacheKey, resultJSON, 10*time.Minute)
		if err != nil {
			slog.Warn(fmt.Sprintf("[Cache] Failed to cache token balances for chain %s: %v", chainId, err))
		}
	}

	slog.Info(fmt.Sprintf("[Cache] Calculated token balances for chain %s: %d tokens", chainId, len(result)))
	return result, nil
}
