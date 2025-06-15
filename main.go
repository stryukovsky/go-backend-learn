package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stryukovsky/go-backend-learn/trade"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main() {
	router := gin.Default()
	db, err := gorm.Open(postgres.Open("postgresql://user:pass@localhost:5432/db"), &gorm.Config{})
	if err != nil {
		panic("Cannot start db connection" + err.Error())
	}
	db.AutoMigrate(&trade.Deal{})
	router.POST("/api/deal", func(ctx *gin.Context) {
		trade.AddDeal(ctx, db)
	})

	router.Run()
}
