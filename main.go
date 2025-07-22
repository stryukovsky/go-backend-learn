package main

import (

	// "github.com/gin-gonic/gin"

	"fmt"

	"github.com/stryukovsky/go-backend-learn/trade"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// router := gin.Default()
	db, err := gorm.Open(postgres.Open("postgresql://user:pass@localhost:5432/db"), &gorm.Config{})
	if err != nil {
		panic("Cannot start db connection" + err.Error())
	}
	db.AutoMigrate(&trade.Deal{}, &trade.ERC20Transfer{}, &trade.Worker{}, &trade.Token{}, &trade.TrackedWallet{})
	// worker := trade.Worker{
	// 	BlockchainUrl:  "http://localhost:8545",
	// 	LastBlock:      22761436,
	// 	BlocksInterval: 100,
	// }
	// db.Create(&worker)
	// db.Create(&trade.Token{ChainId: "1", Address: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", Symbol: "ETH"})
	// db.Create(&trade.TrackedWallet{ChainId: "1", Address: "0x8EB8a3b98659Cce290402893d0123abb75E3ab28"})

	// trade.Cycle(db, 1)
	dealsIncome := []trade.Deal{}
	err = db.Joins("ERC20Transfer").Find(&dealsIncome).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(dealsIncome))

	// router.Run()
}
