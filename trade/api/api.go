package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/cache"
	"gorm.io/gorm"
)

func apiErr(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func BalanceByWallet(ctx *gin.Context, db *gorm.DB, cm *cache.CacheManager) {
	walletAddress := common.HexToAddress(ctx.Param("wallet")).Hex()
	balance, err := cm.GetCachedBalanceOfWallet(db, walletAddress)
	if err != nil {
		apiErr(ctx, err)
		return
	}
	slog.Info(fmt.Sprintf("Find balances across all blockchains of %s", walletAddress))
	ctx.JSON(http.StatusOK, balance)
}

func BalanceByWalletAndChain(ctx *gin.Context, db *gorm.DB, cm *cache.CacheManager) {
	walletAddress := common.HexToAddress(ctx.Param("wallet")).Hex()
	chainId := ctx.Param("chainId")
	result, err := cm.GetCachedBalanceOfWalletOnChain(db, chainId, walletAddress)
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
	err := db.Preload("BlockchainEvent").
		Joins("JOIN aave_events ON aave_events.id = aave_interactions.blockchain_event_id").
		Where("aave_events.wallet_address = ? AND aave_events.chain_id = ?", wallet, chainId).
		Find(&aaveInteractions).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, aaveInteractions)
}

func ListCompound3Interactions(ctx *gin.Context, db *gorm.DB) {
	wallet := common.HexToAddress(ctx.Param("wallet")).Hex()
	chainId := ctx.Param("chainId")
	var compoundInteractions []trade.Compound3Interaction
	err := db.Preload("BlockchainEvent").
		Joins("JOIN compound3_events ON compound3_events.id = compound3_interactions.blockchain_event_id").
		Where("compound3_events.wallet_address = ? AND compound3_events.chain_id = ?", wallet, chainId).
		Find(&compoundInteractions).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, compoundInteractions)
}

func GetTokenBalancesByChain(ctx *gin.Context, db *gorm.DB, cm *cache.CacheManager) {
	chainId := ctx.Param("chainId")

	tokenBalances, err := cm.GetCachedTokenBalancesByChain(db, chainId)
	if err != nil {
		apiErr(ctx, err)
		return
	}

	slog.Info(fmt.Sprintf("Retrieved token balances for chain %s: %d tokens", chainId, len(tokenBalances)))
	ctx.JSON(http.StatusOK, tokenBalances)
}

func ListUniswapV3Interactions(ctx *gin.Context, db *gorm.DB) {
	wallet := common.HexToAddress(ctx.Param("wallet")).Hex()
	chainId := ctx.Param("chainId")
	var uniswapv3Interactions []trade.UniswapV3Deal
	err := db.Preload("BlockchainEvent").
		Joins("JOIN uniswap_v3_events ON uniswap_v3_events.id = uniswap_v3_deals.blockchain_event_id").
		Where("uniswap_v3_events.wallet_address = ? AND uniswap_v3_events.chain_id = ?", wallet, chainId).
		Find(&uniswapv3Interactions).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, uniswapv3Interactions)
}

func ListDealsByWalletAndChain(ctx *gin.Context, db *gorm.DB) {
	wallet := common.HexToAddress(ctx.Param("wallet")).Hex()
	chainId := ctx.Param("chainId")

	dealsAsSender := []trade.Deal{}
	err := db.Preload("BlockchainTransfer").
		Joins("JOIN erc20_transfers ON erc20_transfers.id = deals.blockchain_transfer_id").
		Where("erc20_transfers.sender = ? AND erc20_transfers.chain_id = ?", wallet, chainId).
		Find(&dealsAsSender).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}

	dealsAsRecipient := []trade.Deal{}
	err = db.Preload("BlockchainTransfer").
		Joins("JOIN erc20_transfers ON erc20_transfers.id = deals.blockchain_transfer_id").
		Where("erc20_transfers.recipient = ? AND erc20_transfers.chain_id = ?", wallet, chainId).
		Find(&dealsAsRecipient).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}

	result := trade.NewDealsByWallet(wallet, dealsAsRecipient, dealsAsSender)
	ctx.JSON(http.StatusOK, result)
}

func ListTokensByChain(ctx *gin.Context, db *gorm.DB) {
	chainId := ctx.Param("chainId")
	var tokens []trade.Token
	err := db.Where("chain_id = ?", chainId).Find(&tokens).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, tokens)
}

func CreateApi(router *gin.Engine, db *gorm.DB, cm *cache.CacheManager) {
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
		BalanceByWalletAndChain(ctx, db, cm)
	})
	router.GET("/api/balance/wallet/:wallet", func(ctx *gin.Context) {
		BalanceByWallet(ctx, db, cm)
	})
	router.GET("/api/deals/:chainId/:wallet", func(ctx *gin.Context) {
		ListDealsByWalletAndChain(ctx, db)
	})
	router.GET("/api/aave/:chainId/:wallet", func(ctx *gin.Context) {
		ListAaveInteractions(ctx, db)
	})
	router.GET("/api/uniswapv3/:chainId/:wallet", func(ctx *gin.Context) {
		ListUniswapV3Interactions(ctx, db)
	})
	router.GET("/api/compound3/:chainId/:wallet", func(ctx *gin.Context) {
		ListCompound3Interactions(ctx, db)
	})
	router.GET("/api/chain/:chainId/token-balances", func(ctx *gin.Context) {
		GetTokenBalancesByChain(ctx, db, cm)
	})
	router.GET("/api/:chainId/tokens", func(ctx *gin.Context) {
		ListTokensByChain(ctx, db)
	})
}
