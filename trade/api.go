package trade

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func apiErr(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

type BalanceAcrossAllChains struct {
	Address string `json:"address" binding:"required"`
	Balance string `json:"balance" binding:"required"`
}

func NewBalanceAcrossAllChains(address string, balance string) *BalanceAcrossAllChains {
	return &BalanceAcrossAllChains{
		Address: address,
		Balance: balance,
	}
}

type BalanceOnChain struct {
	ChainId string `json:"chainId" binding:"required"`
	Address string `json:"address" binding:"required"`
	Balance string `json:"balance" binding:"required"`
}

func (b *BalanceOnChain) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}

func NewBalanceOnChain(chainId string, address string, balance string) *BalanceOnChain {
	return &BalanceOnChain{
		Address: address,
		Balance: balance,
		ChainId: chainId,
	}
}

type DealsByWallet struct {
	Address  string `json:"address" binding:"required"`
	DealsIn  []Deal `json:"dealsIn" binding:"required"`
	DealsOut []Deal `json:"dealsOut" binding:"required"`
}

func NewDealsByWallet(wallet string, dealsIn []Deal, dealsOut []Deal) *DealsByWallet {
	return &DealsByWallet{
		Address:  wallet,
		DealsIn:  dealsIn,
		DealsOut: dealsOut,
	}
}

func BalanceByWallet(ctx *gin.Context, db *gorm.DB, rdb *redis.Client) {
	walletAddress := common.HexToAddress(ctx.Param("wallet")).Hex()
	balance, err := GetCachedBalanceOfWallet(db, rdb, walletAddress)
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
	result, err := GetCachedBalanceOfWalletOnChain(db, rdb, chainId, walletAddress)
	if err != nil {
		apiErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func GetWalletsOnChain(ctx *gin.Context, db *gorm.DB) {
	chainId := ctx.Param("chainId")
	wallets := []TrackedWallet{}
	err := db.Find(&wallets, TrackedWallet{ChainId: chainId}).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, wallets)
}

func ListWallets(ctx *gin.Context, db *gorm.DB) {
	wallets := []TrackedWallet{}
	err := db.Find(&wallets).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, wallets)
}

func ListChains(ctx *gin.Context, db *gorm.DB) {
	chains := []Chain{}
	err := db.Find(&chains).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, chains)
}

func ListDealsByWalletAndChain(ctx *gin.Context, db *gorm.DB) {
	wallet := common.HexToAddress(ctx.Param("wallet")).Hex()
	chainId := ctx.Param("chainId")
	dealsAsSender := []Deal{}
	err := db.Preload("BlockchainTransfer").Find(&dealsAsSender, Deal{BlockchainTransfer: ERC20Transfer{Sender: wallet, ChainId: chainId}}).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}

	dealsAsRecipient := []Deal{}
	err = db.Preload("BlockchainTransfer").Find(&dealsAsRecipient, Deal{BlockchainTransfer: ERC20Transfer{Recipient: wallet}}).Error
	if err != nil {
		apiErr(ctx, err)
		return
	}
	result := NewDealsByWallet(wallet, dealsAsRecipient, dealsAsSender)
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
}
