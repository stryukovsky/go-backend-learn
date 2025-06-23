package main

import (

	// "github.com/gin-gonic/gin"
	// "math/big"

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
	db.AutoMigrate(&trade.Deal{}, &trade.ERC20Transfer{}, &trade.Worker{})
	// worker := trade.Worker{
	// 	BlockchainUrl:  "http://localhost:8545",
	// 	LastBlock:      trade.DBInt{Int: big.NewInt(22761436)},
	// 	BlocksInterval: trade.DBInt{Int: big.NewInt(100)},
	// }
	// db.Create(&worker)
	// db.Model(&worker).Association("Tokens").Append([]trade.Token{{ChainId: "1", Address: "0xdac17f958d2ee523a2206206994597c13d831ec7", Symbol: "USDT"}})
	trade.Cycle(db, 6)
	// router.Run()
}
