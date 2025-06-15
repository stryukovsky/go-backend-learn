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
}
