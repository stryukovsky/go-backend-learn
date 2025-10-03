package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/cache"
	"gorm.io/gorm"
)

func apiErr(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func BalanceByWallet(ctx *gin.Context, db *gorm.DB, rdb *redis.Client) {
	walletAddress := common.HexToAddress(ctx.Param("wallet")).Hex()
	balance, err := cache.GetCachedBalanceOfWallet(db, rdb, walletAddress)
	if err != nil {
		apiErr(ctx, err)
		return
	}
	slog.Info(fmt.Sprintf("Find balances across all blockchains of %s", walletAddress))
	ctx.JSON(http.StatusOK, balance)
}

func BalanceByWalletAndChain(ctx *gin.Context, db *gorm.DB, rdb *redis.Client) {
	walletAddress := common.HexToAddress(ctx.Param("wallet")).Hex()
	chainId := ctx.Param("chainId")
	result, err := cache.GetCachedBalanceOfWalletOnChain(db, rdb, chainId, walletAddress)
	if err != nil {
		apiErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func GetWalletsOnChain(ctx *gin.Context, db *gorm.DB) {
	chainId := ctx.Param("chainId")
	wallets := []trade.TrackedWallet{}
	err := db.Find(&wallets, trade.TrackedWallet{ChainId: chainId}).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, wallets)
}

func ListWallets(ctx *gin.Context, db *gorm.DB) {
	wallets := []trade.TrackedWallet{}
	err := db.Find(&wallets).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, wallets)
}

func ListChains(ctx *gin.Context, db *gorm.DB) {
	chains := []trade.Chain{}
	err := db.Find(&chains).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, chains)
}

func ListAaveInteractions(ctx *gin.Context, db *gorm.DB) {

	wallet := common.HexToAddress(ctx.Param("wallet")).Hex()
	chainId := ctx.Param("chainId")
	var aaveInteractions []trade.AaveInteraction
	err := db.Preload("BlockchainEvent").Find(
		&aaveInteractions,
		trade.AaveInteraction{BlockchainEvent: trade.AaveEvent{WalletAddress: wallet, ChainId: chainId}},
	).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, aaveInteractions)
}

func ListDealsByWalletAndChain(ctx *gin.Context, db *gorm.DB) {
	wallet := common.HexToAddress(ctx.Param("wallet")).Hex()
	chainId := ctx.Param("chainId")
	dealsAsSender := []trade.Deal{}
	err := db.Preload("BlockchainTransfer").Find(&dealsAsSender, trade.Deal{BlockchainTransfer: trade.ERC20Transfer{Sender: wallet, ChainId: chainId}}).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}

	dealsAsRecipient := []trade.Deal{}
	err = db.Preload("BlockchainTransfer").Find(&dealsAsRecipient, trade.Deal{BlockchainTransfer: trade.ERC20Transfer{Recipient: wallet}}).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}
	result := trade.NewDealsByWallet(wallet, dealsAsRecipient, dealsAsSender)
	ctx.JSON(http.StatusOK, result)
}

func CreateApi(router *gin.Engine, db *gorm.DB, rdb *redis.Client) {
	router.GET("/api/wallets", func(ctx *gin.Context) {
		ListWallets(ctx, db)
	})
	router.GET("/api/chains", func(ctx *gin.Context) {
		ListChains(ctx, db)
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
	router.GET("/api/deals/:chainId/:wallet", func(ctx *gin.Context) {
		ListDealsByWalletAndChain(ctx, db)
	})
	router.GET("/api/aave/:chainId/:wallet", func(ctx *gin.Context) {
		ListAaveInteractions(ctx, db)
	})
}
