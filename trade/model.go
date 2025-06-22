package trade

import (
	"math/big"
	"time"
)

type Deal struct {
	Price              big.Rat       `json:"price" binding:"required"`
	VolumeTokens       big.Rat       `json:"volumeTokens" binding:"required"`
	VolumeUSD          big.Rat       `json:"volumeUSD" binding:"required"`
	BlockchainTransfer ERC20Transfer `json:"blockchainTransfer" binding:"required"`
}

type ERC20Transfer struct {
	TokenAddress string    `json:"tokenAddress" binding:"required"`
	Name         string    `json:"name" binding:"required"`
	Symbol       string    `json:"symbol" binding:"required"`
	Decimals     big.Int   `json:"decimals" binding:"required"`
	Sender       string    `json:"sender" binding:"required"`
	Recipient    string    `json:"recipient" binding:"required"`
	Amount       big.Int   `json:"amount" binding:"required"`
	Block        big.Int   `json:"blockNumber" binding:"required"`
	Timestamp    time.Time `json:"timestamp" binding:"required"`
}
