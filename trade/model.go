package trade

import "math/big"

type Deal struct {
	InputToken  string `json:"inputToken" binding:"required"`
	InputAmount string `json:"inputAmount" binding:"required"`
	OutputToken string `json:"OutputToken" binding:"required"`
}

type ERC20Transfer struct {
	TokenAddress string  `json:"tokenAddress" binding:"required"`
	Sender       string  `json:"sender" binding:"required"`
	Recipient    string  `json:"recipient" binding:"required"`
	Amount       big.Int `json:"amount" binding:"required"`
}
