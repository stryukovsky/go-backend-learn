package trade

import (
	"fmt"
	"log/slog"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BalanceAcrossAllChains struct {
	Address string `json:"address" binding:"required"`
	Balance string `json:"balance" binding:"required"`
}

func AddDeal(ctx *gin.Context, db *gorm.DB) {
	var deal Deal
	if err := ctx.ShouldBindJSON(&deal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

// TODO: add caching
func BalanceByWallet(ctx *gin.Context, db *gorm.DB) {
	// make address checksum-format
	walletAddress := common.HexToAddress(ctx.Param("wallet")).Hex()
	slog.Info(fmt.Sprintf("Find balances across all blockchains of %s", walletAddress))

	dealsIncome := []Deal{}
	err := db.Preload("BlockchainTransfer").Find(&dealsIncome, Deal{BlockchainTransfer: ERC20Transfer{Recipient: walletAddress}}).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Cannot fetch income deals for %s: %s", walletAddress, err.Error)})
		return
	}
	dealsOutcome := []Deal{}
	err = db.Preload("BlockchainTransfer").Find(&dealsOutcome, Deal{BlockchainTransfer: ERC20Transfer{Recipient: walletAddress}}).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Cannot fetch outcome deals for %s: %s", walletAddress, err.Error)})
		return
	}
	slog.Info(fmt.Sprintf("Found %d income and %d outcome deals of %s", len(dealsIncome), len(dealsOutcome), walletAddress))

	balance := calculateBalance(dealsIncome, dealsOutcome)
	ctx.JSON(http.StatusOK, BalanceAcrossAllChains{Address: walletAddress, Balance: balance})
}

// TODO: add caching
func BalanceByWalletAndChain(ctx *gin.Context, db *gorm.DB) {
	walletAddress := common.HexToAddress(ctx.Param("wallet")).Hex()
	chainId := ctx.Param("chainId")
	dealsIncome := []Deal{}
	dealsOutcome := []Deal{}
	err := db.Preload("BlockchainTransfer").Find(dealsIncome, Deal{BlockchainTransfer: ERC20Transfer{Recipient: walletAddress, ChainId: chainId}})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Cannot fetch income deals for %s: %s", walletAddress, err.Error)})
		return
	}
	err = db.Preload("BlockchainTransfer").Find(dealsOutcome, Deal{BlockchainTransfer: ERC20Transfer{Sender: walletAddress, ChainId: chainId}})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Cannot fetch income deals for %s: %s", walletAddress, err.Error)})
		return
	}
}

func GetWalletsOnChain(ctx *gin.Context, db *gorm.DB) {
	
}

func CreateApi(router *gin.Engine, db *gorm.DB) {
	router.POST("/api/deal", func(ctx *gin.Context) {
		AddDeal(ctx, db)
	})
	router.GET("/api/deal", func(ctx *gin.Context) {
		ListDeals(ctx, db)
	})
	router.GET("/api/deal/:id", func(ctx *gin.Context) {
		GetDeal(ctx, db)
	})
}
