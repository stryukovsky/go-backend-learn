package trade

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BalanceAcrossAllChains struct {
	Address string `json:"address" binding:"required"`
	Balance string `json:"balance" binding:"required"`
}

type BalanceOnChain struct {
	ChainId string `json:"chainId" binding:"required"`
	Address string `json:"address" binding:"required"`
	Balance string `json:"balance" binding:"required"`
}

func AddDeal(ctx *gin.Context, db *gorm.DB) {
	var deal Deal
	if err := ctx.ShouldBindJSON(&deal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tx := db.Create(&deal)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "cannot add"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

func GetDeal(ctx *gin.Context, db *gorm.DB) {
	id := ctx.Param("id")
	var deal Deal
	if err := db.First(&deal, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "deal not found"})
		return
	}
	ctx.JSON(http.StatusOK, deal)
}

func ListDeals(ctx *gin.Context, db *gorm.DB) {
	var items []Deal
	if tx := db.Find(&items); tx.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "cannot list"})
	} else {
		ctx.JSON(http.StatusOK, items)
	}
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

const cacheKeyPrefixBalanceAcrossAllChains = "balanceAcrossAllChains:"

func BalanceByWallet(ctx *gin.Context, db *gorm.DB, rdb *redis.Client) {
	walletAddress := common.HexToAddress(ctx.Param("wallet")).Hex()
	cacheKey := cacheKeyPrefixBalanceAcrossAllChains + walletAddress
	cachedBalance, err := rdb.Get(context.Background(), cacheKey).Result()
	if err == nil && cachedBalance != "" {
		var balanceAcrossAllChains BalanceAcrossAllChains
		json.Unmarshal([]byte(cachedBalance), &balanceAcrossAllChains)
		ctx.JSON(http.StatusOK, balanceAcrossAllChains)
		return
	}

	slog.Info(fmt.Sprintf("Find balances across all blockchains of %s", walletAddress))

	var dealsIncome []Deal
	countIncome := 0
	err = db.Preload("BlockchainTransfer").Where("blockchain_transfer.recipient = ?", walletAddress).First(&dealsIncome, &countIncome).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Cannot fetch income deals for %s: %s", walletAddress, err.Error)})
		return
	}

	var dealsOutcome []Deal
	countOutcome := 0
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Cannot fetch outcome deals for %s: %s", walletAddress, err.Error)})
		return
	}
	slog.Info(fmt.Sprintf("Found %d income and %d outcome deals of %s", len(dealsIncome), countOutcome, walletAddress))

	balance := calculateBalance(dealsIncome, dealsOutcome)
	cachedData, _ := json.Marshal(BalanceAcrossAllChains{Address: walletAddress, Balance: balance})
	rdb.Set(ctx, cacheKey, cachedData, 5*time.Minute)
	ctx.JSON(http.StatusOK, BalanceAcrossAllChains{Address: walletAddress, Balance: balance})
}

const cacheKeyPrefixBalanceOfWalletOnChain = "balanceOnChain:"

func BalanceByWalletAndChain(ctx *gin.Context, db *gorm.DB, rdb *redis.Client) {
	walletAddress := common.HexToAddress(ctx.Param("wallet")).Hex()
	chainId := ctx.Param("chainId")

	key := cacheKeyPrefixBalanceOfWalletOnChain + chainId + ":" + walletAddress
	cached, err := rdb.Get(context.Background(), key).Result()
	if err != nil && err != redis.Nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get balance from cache due to internal error " + err.Error()})
		return
	}

	if cached != "" {
		var result BalanceOnChain
		deserialized := json.Unmarshal([]byte(cached), &result)
		ctx.JSON(http.StatusOK, deserialized)
		return
	}

	dealsIncome := []Deal{}
	dealsOutcome := []Deal{}
	err = db.
		Preload("BlockchainTransfer").
		Find(&dealsIncome, Deal{BlockchainTransfer: ERC20Transfer{Recipient: walletAddress, ChainId: chainId}}).
		Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Cannot fetch income deals for %s: %s", walletAddress, err.Error),
		})
		return
	}
	err = db.
		Preload("BlockchainTransfer").
		Find(&dealsOutcome, Deal{BlockchainTransfer: ERC20Transfer{Sender: walletAddress, ChainId: chainId}}).
		Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Cannot fetch income deals for %s: %s", walletAddress, err.Error())})
		return
	}

	balance := big.NewRat(0, 1)
	for _, deal := range dealsIncome {
		balance = balance.Add(balance, deal.VolumeUSD.Rat)
	}
	for _, deal := range dealsOutcome {
		balance = balance.Sub(balance, deal.VolumeUSD.Rat)
	}
	result := BalanceOnChain{
		ChainId: chainId,
		Address: walletAddress,
		Balance: balance.FloatString(2),
	}
	err = rdb.Set(context.Background(), key, result, 15*time.Minute).Err()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Cannot set to cache balance", err.Error())})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func GetWalletsOnChain(ctx *gin.Context, db *gorm.DB) {
	chainId := ctx.Param("chainId")
	wallets := []TrackedWallet{}
	err := db.Find(&wallets, TrackedWallet{ChainId: chainId}).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find wallets"})
		return
	}
	ctx.JSON(http.StatusOK, wallets)
}

func ListWallets(ctx *gin.Context, db *gorm.DB) {
	wallets := []TrackedWallet{}
	err := db.Find(&wallets).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list wallets"})
		return
	}
	ctx.JSON(http.StatusOK, wallets)
}

func CreateApi(router *gin.Engine, db *gorm.DB, rdb *redis.Client) {
	router.POST("/api/deal", func(ctx *gin.Context) {
		AddDeal(ctx, db)
	})
	router.GET("/api/deal", func(ctx *gin.Context) {
		ListDeals(ctx, db)
	})
	router.GET("/api/deal/:id", func(ctx *gin.Context) {
		GetDeal(ctx, db)
	})
	router.GET("/api/chain/:chainId/wallets", func(ctx *gin.Context) {
		GetWalletsOnChain(ctx, db)
	})
	router.GET("/api/balance/chainAndWallet/:chainId/:wallet", func(ctx *gin.Context) {
		BalanceByWalletAndChain(ctx, db, rdb)
	})
	router.GET("/api/balance/wallet/:wallet", func(ctx *gin.Context) {
		BalanceByWallet(ctx, db, rdb)
	})
}
