package trade

import (
	"math/big"
	"strconv"

	"github.com/chenzhijie/go-web3"
	"gorm.io/gorm"
)

func Cycle(db *gorm.DB, id uint) {
	var config Worker
	result := db.First(&config, id)
	if result.Error != nil {
		panic("No worker with id " + strconv.Itoa(int(id)))
	}

	web3, err := web3.NewWeb3(config.BlockchainUrl)
	if err != nil {
		panic(err)
	}

	tokens := make([]ERC20, 0, len(config.Tokens))
	for _, token := range config.Tokens {
		erc20, err := CreateToken(web3, token.Address, token.Symbol)
		if err != nil {
			continue
		}
		tokens = append(tokens, *erc20)
	}

	startFromBlock := config.LastBlock.Int
	endInBlock := new(big.Int).Add(startFromBlock, config.BlocksInterval.Int)
	for _, token := range tokens {
		transfers, err := token.ListTransfers(startFromBlock, endInBlock)
		if err != nil {
			panic(err)
		}
		// deals := make([]Deal, 0, len(transfers))
		for _, transfer := range transfers {
			deal, err := CreateDeal(transfer)
			if err != nil {
				panic(err)
			}
			// deals = append(deals, *deal)
			db.Save(deal)
		}
	}
	config.LastBlock = DBInt{endInBlock}
	db.Save(config)



}
