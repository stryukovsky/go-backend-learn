package trade

import (
	"fmt"
	"strconv"

	"log/slog"

	"github.com/chenzhijie/go-web3"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Cycle(db *gorm.DB, id uint) {
	var config Worker
	result := db.First(&config, id)
	if result.Error != nil {
		slog.Warn("No config with id " + strconv.Itoa(int(id)))
		return
	}

	web3, err := web3.NewWeb3(config.BlockchainUrl)
	if err != nil {
		slog.Warn(fmt.Sprintf("Cannot connect to blockchain: %e", err))
		return
	}

	cache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "redis",
		DB:       0,
	})

	currentBlockchainBlock, err := web3.Eth.GetBlockNumber()
	if err != nil {
		slog.Warn(fmt.Sprintf("Cannot get last blockchain block: %e", err))
		return
	}

	var tokensFromDB []Token
	err = db.Model(&config).Association("Tokens").Find(&tokensFromDB)
	if err != nil {
		slog.Warn(fmt.Sprintf("Cannot get tokens of config: %e", err))
		return
	}

	tokens := make([]ERC20, 0, len(tokensFromDB))
	for _, token := range tokensFromDB {
		erc20, err := CreateERC20(web3, token.Address, token.Symbol)
		if err != nil {
			slog.Warn(fmt.Sprintf("Cannot create token %s: %e", token.Address, err))
			continue
		}
		tokens = append(tokens, *erc20)
	}

	startFromBlock := config.LastBlock
	endInBlock := min(startFromBlock+config.BlocksInterval, currentBlockchainBlock)
	slog.Info(fmt.Sprintf("Interacting with %d tokens. Find events from block %d to %d", len(tokens), startFromBlock, endInBlock))
	for _, token := range tokens {
		transfers, err := token.ListTransfers(startFromBlock, endInBlock, cache)
		if err != nil {
			slog.Warn(fmt.Sprintf("[%s] Cannot fetch transfers: %s", token.Symbol, err.Error()))
			continue
		}
		slog.Info(fmt.Sprintf("[%s] Found %d transfers", token.Symbol, len(transfers)))
		for _, transfer := range transfers {
			deal, err := CreateDeal(cache, transfer)
			if err != nil {
				slog.Warn(fmt.Sprintf("[%s] Cannot create deal object for ERC20 transfer %s: %s", token.Symbol, transfer.TxId, err.Error()))
			} else {
				slog.Info(fmt.Sprintf("[%s] Found deal with volume $ %s", token.Symbol, deal.VolumeUSD.FloatString(5)))
				db.Save(deal)
			}
		}
	}
	slog.Info(fmt.Sprintf("Worker last block updated to %d", endInBlock))
	config.LastBlock = endInBlock
	db.Save(config)

}
