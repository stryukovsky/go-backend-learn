package database

import (
	"math/big"

	"github.com/stryukovsky/go-backend-learn/trade"
	"gorm.io/gorm"
)

func Fixture(db *gorm.DB) {
	db.Create(&trade.Worker{BlockchainUrl: "http://localhost:8545", BlocksInterval: 1000})
	db.Create(&trade.TrackedWallet{ChainId: "1", Address: "0x8EB8a3b98659Cce290402893d0123abb75E3ab28", LastBlock: 23152597})
	db.Create(&trade.TrackedWallet{ChainId: "1", Address: "0x49ecd0F2De4868E5130fdC2C45D4d97444B7c269", LastBlock: 23495633})
	db.Create(&trade.TrackedWallet{ChainId: "1", Address: "0x66a9893cC07D91D95644AEDD05D03f95e1dBA8Af", LastBlock: 23153501})
	
	db.Create(
		&trade.Token{
			ChainId:  "1",
			Address:  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
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
	db.Create(
		&trade.Token{
			ChainId:  "1",
			Address:  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			Symbol:   "USDC",
			Decimals: trade.NewDBInt(big.NewInt(6)),
		})
	db.Create(&trade.Chain{Name: "Ethereum mainnet", ChainId: "1"})
	db.Create(&trade.DeFiPlatform{Type: trade.Aave, ChainId: "1", Address: "0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"})
	db.Create(&trade.DeFiPlatform{Type: trade.UniswapV3, ChainId: "1", Address: "0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640"})
}
