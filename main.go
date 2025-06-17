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
	trade.CreateApi(router, db)
	router.Run()
}

// import (
// 	"fmt"
//
// 	"github.com/stryukovsky/go-backend-learn/trade"
// )
// func main() {
// 	fmt.Println(trade.PerfromDeal(&trade.Deal{InputToken: "0xdac17f958d2ee523a2206206994597c13d831ec7", InputAmount: "10", OutputToken: "abc"}))
// }
