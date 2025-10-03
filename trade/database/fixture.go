package database

import (
	"math/big"

	"github.com/stryukovsky/go-backend-learn/trade"
	"gorm.io/gorm"
)

func Fixture(db *gorm.DB) {
	db.Create(&trade.Worker{BlockchainUrl: "http://localhost:8545", BlocksInterval: 1000})
	db.Create(&trade.TrackedWallet{ChainId: "1", Address: "0x8EB8a3b98659Cce290402893d0123abb75E3ab28", LastBlock: 23152597})
	db.Create(
		&trade.Token{
			ChainId:  "1",
			Address:  "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
			Symbol:   "ETH",
			Decimals: trade.NewDBInt(big.NewInt(18)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "1",
			Address:  "0x1aBaEA1f7C830bD89Acc67eC4af516284b1bC33c",
			Symbol:   "EUR",
			Decimals: trade.NewDBInt(big.NewInt(6)),
		})
	db.Create(&trade.Chain{Name: "Ethereum mainnet", ChainId: "1"})
	db.Create(&trade.DeFiPlatform{Type: trade.Aave, ChainId: "1", Address: "0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"})
	db.Create(&trade.TrackedWallet{ChainId: "1", Address: "0x49ecd0F2De4868E5130fdC2C45D4d97444B7c269", LastBlock: 23495633})
}
