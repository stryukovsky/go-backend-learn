package trade

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
