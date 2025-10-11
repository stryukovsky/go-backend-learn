package database

import (
	"github.com/stryukovsky/go-backend-learn/trade"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&trade.Deal{},
		&trade.ERC20Transfer{},
		&trade.Worker{},
		&trade.Token{},
		&trade.TrackedWallet{},
		&trade.Chain{},
		&trade.AaveEvent{},
		&trade.AaveInteraction{},
		&trade.DeFiPlatform{},
		&trade.UniswapV3Event{},
		&trade.UniswapV3Deal{},
		&trade.UniswapV3Position{},
		&trade.AnalyticsWorker{},
	)
	return err
}
