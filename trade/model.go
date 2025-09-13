package trade

import (
	"database/sql/driver"
	"fmt"
	"math/big"
	"time"

	"gorm.io/gorm"
)

type DBNumeric struct {
	*big.Rat
}

func (br *DBNumeric) Scan(value any) error {
	if value == nil {
		br.Rat = nil
		return nil
	}

	decimalStr, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan BigRat: expected string, got %T", value)
	}

	if br.Rat == nil {
		br.Rat = new(big.Rat)
	}

	// Convert from PostgreSQL numeric string to big.Rat
	_, ok = br.Rat.SetString(decimalStr)
	if !ok {
		return fmt.Errorf("failed to parse BigRat from decimal string: %s", decimalStr)
	}

	return nil
}

func (br DBNumeric) Value() (driver.Value, error) {
	if br.Rat == nil {
		return nil, nil
	}
	// Convert to string representation that PostgreSQL numeric can handle
	return br.Rat.FloatString(20), nil
}

func (DBNumeric) GormDataType() string {
	return "numeric" // PostgreSQL numeric type
}

// DBInt wrapper for GORM
type DBInt struct {
	*big.Int
}

// Scan implements sql.Scanner interface
func (b *DBInt) Scan(value any) error {
	if value == nil {
		b.Int = nil
		return nil
	}

	// PostgreSQL returns numeric types as string
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan BigInt: expected string, got %T", value)
	}

	if b.Int == nil {
		b.Int = new(big.Int)
	}

	// Parse the string representation
	_, success := b.Int.SetString(str, 10)
	if !success {
		return fmt.Errorf("failed to parse BigInt from string: %s", str)
	}

	return nil
}

// Value implements driver.Valuer interface
func (b DBInt) Value() (driver.Value, error) {
	if b.Int == nil {
		return nil, nil
	}
	return b.Int.String(), nil
}

// GormDataType declares the database type
func (DBInt) GormDataType() string {
	return "numeric" // PostgreSQL numeric type
}

type Deal struct {
	gorm.Model
	Price                DBNumeric `json:"price" binding:"required"`
	VolumeTokens         DBNumeric `json:"volumeTokens" binding:"required"`
	VolumeUSD            DBNumeric `json:"volumeUSD" binding:"required"`
	BlockchainTransferID int
	BlockchainTransfer   ERC20Transfer `json:"blockchainTransfer" binding:"required"`
}

type ERC20Transfer struct {
	gorm.Model
	TokenAddress string    `json:"tokenAddress" binding:"required"`
	Sender       string    `json:"sender" binding:"required"`
	Recipient    string    `json:"recipient" binding:"required"`
	Amount       DBInt     `json:"amount" binding:"required" gorm:"type:numeric(78,0)"`
	Block        DBInt     `json:"blockNumber" binding:"required"`
	ChainId      string    `json:"chainId" binding:"required"`
	Timestamp    time.Time `json:"timestamp" binding:"required"`
	TxId         string    `json:"txId" gorm:"uniqueIndex" binding:"required"`
}

func NewERC20Transfer(address string, sender string, recipient string, amount *big.Int, block *big.Int, chainId string, timestamp *time.Time, txId string) ERC20Transfer {
	return ERC20Transfer{
		TokenAddress: address,
		Sender:       sender,
		Recipient:    recipient,
		Amount:       DBInt{amount},
		Block:        DBInt{block},
		ChainId:      chainId,
		Timestamp:    *timestamp,
		TxId:         txId,
	}
}

type Worker struct {
	gorm.Model
	BlockchainUrl  string `json:"blockchainUrl" binding:"required"`
	BlocksInterval uint64 `json:"blocksInterval" binding:"required"`
}

type Token struct {
	gorm.Model
	ChainId  string `json:"chainId" binding:"required" gorm:"uniqueIndex:idx_token_uniqueness"`
	Symbol   string `json:"symbol" binding:"required" gorm:"uniqueIndex:idx_token_uniqueness"`
	Address  string `json:"address" binding:"required" gorm:"uniqueIndex:idx_token_uniqueness"`
	Decimals DBInt  `json:"decimals" binding:"required"`
}

type TrackedWallet struct {
	gorm.Model
	Address   string `json:"address" binding:"required" gorm:"uniqueIndex:idx_wallet_uniqueness"`
	ChainId   string `json:"chainId" binding:"required" gorm:"uniqueIndex:idx_wallet_uniqueness"`
	LastBlock uint64 `json:"lastBlock" binding:"required"`
}

type Balance struct {
}
