package main

import (
	"github.com/stryukovsky/go-backend-learn/trade"
	"gorm.io/gorm"
)

func Fixture(db *gorm.DB) {
	// db.Create(&trade.Worker{BlockchainUrl: "http://localhost:8545", BlocksInterval: 1000})
	// db.Create(&trade.TrackedWallet{ChainId: "1", Address: "0x8EB8a3b98659Cce290402893d0123abb75E3ab28", LastBlock: 23152597})
	// db.Create(&trade.Token{ChainId: "1", Address: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", Symbol: "ETH"})
	db.Create(&trade.Chain{Name: "Ethereum mainnet", ChainId: "1"})
}
