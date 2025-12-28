package trade

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type DBNumeric struct {
	*big.Rat
}

func NewDBNumeric(value *big.Rat) DBNumeric {
	return DBNumeric{value}
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

func (br DBNumeric) MarshalJSON() ([]byte, error) {
	return []byte("\"" + br.Rat.FloatString(5) + "\""), nil
}

func (DBNumeric) GormDataType() string {
	return "numeric" // PostgreSQL numeric type
}

// DBInt wrapper for GORM
type DBInt struct {
	*big.Int
}

func NewDBInt(value *big.Int) DBInt {
	return DBInt{value}
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

type AaveInteraction struct {
	gorm.Model
	Price             DBNumeric `json:"price" binding:"required"`
	VolumeTokens      DBNumeric `json:"volumeTokens" binding:"required"`
	VolumeUSD         DBNumeric `json:"volumeUSD" binding:"required"`
	BlockchainEventID int
	BlockchainEvent   AaveEvent `json:"blockchainEvent" binding:"required"`
}

type AaveEvent struct {
	gorm.Model
	ChainId       string    `json:"chainId" binding:"required" gorm:"uniqueIndex:aave_idx_event_uniqueness"`
	Direction     string    `json:"direction" binding:"required"`
	WalletAddress string    `json:"walletAddress" binding:"required"`
	TokenAddress  string    `json:"tokenAddress" binding:"required"`
	Amount        DBInt     `json:"amount" binding:"required"`
	Timestamp     time.Time `json:"timestamp" binding:"required"`
	TxId          string    `json:"txId" binding:"required" gorm:"uniqueIndex:aave_idx_event_uniqueness"`
	LogIndex      uint      `json:"logIndex" binding:"required" gorm:"uniqueIndex:aave_idx_event_uniqueness"`
}

func NewAaveEvent(
	chainId string,
	direction string,
	walletAddress common.Address,
	tokenAddress common.Address,
	amount *big.Int,
	timestamp time.Time,
	txId string,
	logIndex uint,
) AaveEvent {
	return AaveEvent{
		ChainId:       chainId,
		Direction:     direction,
		WalletAddress: walletAddress.Hex(),
		TokenAddress:  tokenAddress.Hex(),
		Amount:        DBInt{amount},
		Timestamp:     timestamp,
		TxId:          txId,
		LogIndex:      logIndex,
	}
}

type Compound3Interaction struct {
	gorm.Model
	Price             DBNumeric `json:"price" binding:"required"`
	VolumeTokens      DBNumeric `json:"volumeTokens" binding:"required"`
	VolumeUSD         DBNumeric `json:"volumeUSD" binding:"required"`
	BlockchainEventID int
	BlockchainEvent   Compound3Event `json:"blockchainEvent" binding:"required"`
}

type Compound3Event struct {
	gorm.Model
	ChainId       string    `json:"chainId" binding:"required" gorm:"uniqueIndex:aave_idx_event_uniqueness"`
	Direction     string    `json:"direction" binding:"required"`
	WalletAddress string    `json:"walletAddress" binding:"required"`
	TokenAddress  string    `json:"tokenAddress" binding:"required"`
	Amount        DBInt     `json:"amount" binding:"required"`
	Timestamp     time.Time `json:"timestamp" binding:"required"`
	TxId          string    `json:"txId" binding:"required" gorm:"uniqueIndex:aave_idx_event_uniqueness"`
	LogIndex      uint      `json:"logIndex" binding:"required" gorm:"uniqueIndex:aave_idx_event_uniqueness"`
}

func NewCompound3Event(
	chainId string,
	direction string,
	walletAddress common.Address,
	tokenAddress common.Address,
	amount *big.Int,
	timestamp time.Time,
	txId string,
	logIndex uint,
) Compound3Event {
	return Compound3Event{
		ChainId:       chainId,
		Direction:     direction,
		WalletAddress: walletAddress.Hex(),
		TokenAddress:  tokenAddress.Hex(),
		Amount:        DBInt{amount},
		Timestamp:     timestamp,
		TxId:          txId,
		LogIndex:      logIndex,
	}
}

const (
	UniswapV3Swap    = "Swap"
	UniswapV3Mint    = "Mint"
	UniswapV3Burn    = "Burn"
	UniswapV3Collect = "Collect"
)

type UniswapV3Event struct {
	gorm.Model
	ChainId       string `json:"chainId" binding:"required" gorm:"uniqueIndex:uniswap_v3_idx_event_uniqueness"`
	Type          string `json:"type" binding:"required"`
	WalletAddress string `json:"walletAddress" binding:"required"`
	PoolAddress   string `json:"poolAddress" binding:"required"`
	AmountTokenA  DBInt  `json:"amountTokenA" binding:"required"`
	AmountTokenB  DBInt  `json:"amountTokenB" binding:"required"`
	// for swaps PriceLower == PriceUpper
	PriceLower      DBNumeric `json:"priceLower" binding:"required"`
	PriceUpper      DBNumeric `json:"priceUpper" binding:"required"`
	Timestamp       time.Time `json:"timestamp" binding:"required"`
	PositionTokenId string    `json:"tokenId" binding:"required"`
	TxId            string    `json:"txId" binding:"required" gorm:"uniqueIndex:uniswap_v3_idx_event_uniqueness"`
	LogIndex        uint      `json:"logIndex" binding:"required" gorm:"uniqueIndex:uniswap_v3_idx_event_uniqueness"`
	BlockNumber     uint64    `json:"blockNumber" binding:"required"`
}

func NewUniswapV3Event(
	chainId string,
	eventType string,
	walletAddress string,
	poolAddress string,
	amountTokenA *big.Int,
	amountTokenB *big.Int,
	priceLower *big.Rat,
	priceUpper *big.Rat,
	timestamp time.Time,
	txId string,
	logIndex uint,
	blockNumber uint64,
) UniswapV3Event {
	return UniswapV3Event{
		ChainId:         chainId,
		Type:            eventType,
		WalletAddress:   walletAddress,
		PoolAddress:     poolAddress,
		AmountTokenA:    NewDBInt(amountTokenA),
		AmountTokenB:    NewDBInt(amountTokenB),
		PriceLower:      NewDBNumeric(priceLower),
		PriceUpper:      NewDBNumeric(priceUpper),
		Timestamp:       timestamp,
		PositionTokenId: "",
		TxId:            txId,
		LogIndex:        logIndex,
		BlockNumber:     blockNumber,
	}
}

type UniswapV3Deal struct {
	gorm.Model
	SymbolA            string    `json:"symbolA" binding:"required"`
	SymbolB            string    `json:"symbolB" binding:"required"`
	PriceTokenA        DBNumeric `json:"priceTokenA" binding:"required"`
	PriceTokenB        DBNumeric `json:"priceTokenB" binding:"required"`
	VolumeTokensAInUSD DBNumeric `json:"volumeTokensAInUSD" binding:"required"`
	VolumeTokensBInUSD DBNumeric `json:"volumeTokensBInUSD" binding:"required"`
	VolumeTokensA      DBNumeric `json:"volumeTokensA" binding:"required"`
	VolumeTokensB      DBNumeric `json:"volumeTokensB" binding:"required"`
	VolumeTotalUSD     DBNumeric `json:"volumeTotalUSD" binding:"required"`
	BlockchainEventID  int
	BlockchainEvent    UniswapV3Event `json:"blockchainEvent" binding:"required"`
}

func NewUniswapV3Deal(
	tickerA string,
	tickerB string,
	priceTokenA *big.Rat,
	priceTokenB *big.Rat,
	volumeTokensAInUSD *big.Rat,
	volumeTokensBInUSD *big.Rat,
	volumeTokensA *big.Rat,
	volumeTokensB *big.Rat,
	volumeTotalUSD *big.Rat,
	blockchainEvent UniswapV3Event,
) UniswapV3Deal {
	return UniswapV3Deal{
		SymbolA:            tickerA,
		SymbolB:            tickerB,
		PriceTokenA:        NewDBNumeric(priceTokenA),
		PriceTokenB:        NewDBNumeric(priceTokenB),
		VolumeTokensAInUSD: NewDBNumeric(volumeTokensAInUSD),
		VolumeTokensBInUSD: NewDBNumeric(volumeTokensBInUSD),
		VolumeTokensA:      NewDBNumeric(volumeTokensA),
		VolumeTokensB:      NewDBNumeric(volumeTokensB),
		VolumeTotalUSD:     NewDBNumeric(volumeTotalUSD),
		BlockchainEvent:    blockchainEvent,
	}
}

// Historical info on any position minted UniswapV3
type UniswapV3Position struct {
	ChainId                 string `json:"chainId" binding:"required" gorm:"uniqueIndex:uniswap_v3_position_uniqueness"`
	UniswapPositionsManager string `json:"uniswapPositionsManager" binding:"required" gorm:"uniqueIndex:uniswap_v3_position_uniqueness"`
	TokenId                 string `json:"tokenId" binding:"required" gorm:"uniqueIndex:uniswap_v3_position_uniqueness"`
	Owner                   string `json:"owner" binding:"required"`
}

func NewUniswapV3Position(chainId string, uniswapPositionsManager common.Address, tokenId *big.Int, owner common.Address) UniswapV3Position {
	return UniswapV3Position{
		ChainId:                 chainId,
		UniswapPositionsManager: uniswapPositionsManager.Hex(),
		TokenId:                 tokenId.String(),
		Owner:                   owner.Hex(),
	}
}

type Deal struct {
	gorm.Model
	Price                DBNumeric `json:"price" binding:"required"`
	VolumeTokens         DBNumeric `json:"volumeTokens" binding:"required"`
	VolumeUSD            DBNumeric `json:"volumeUSD" binding:"required"`
	BlockchainTransferID int
	BlockchainTransfer   ERC20Transfer `json:"blockchainTransfer" binding:"required"`
}

type Chain struct {
	gorm.Model
	Name    string `json:"name" binding:"required"`
	ChainId string `json:"ChainId" binding:"required" gorm:"uniqueIndex"`
}

// TODO: Block from DBInt to uint64
type ERC20Transfer struct {
	gorm.Model
	TokenAddress string    `json:"tokenAddress" binding:"required"`
	Sender       string    `json:"sender" binding:"required"`
	Recipient    string    `json:"recipient" binding:"required"`
	Amount       DBInt     `json:"amount" binding:"required"`
	Block        DBInt     `json:"blockNumber" binding:"required"`
	ChainId      string    `json:"chainId" binding:"required" gorm:"uniqueIndex:erc20_idx_event_uniqueness"`
	Timestamp    time.Time `json:"timestamp" binding:"required"`
	TxId         string    `json:"txId" binding:"required" gorm:"uniqueIndex:erc20_idx_event_uniqueness"`
	LogIndex     uint      `json:"logIndex" binding:"required" gorm:"uniqueIndex:erc20_idx_event_uniqueness"`
}

func NewERC20Transfer(
	address string,
	sender string,
	recipient string,
	amount *big.Int,
	block uint64,
	chainId string,
	timestamp *time.Time,
	txId string,
	logIndex uint,
) ERC20Transfer {
	return ERC20Transfer{
		TokenAddress: address,
		Sender:       sender,
		Recipient:    recipient,
		Amount:       DBInt{amount},
		Block:        DBInt{big.NewInt(int64(block))},
		ChainId:      chainId,
		Timestamp:    *timestamp,
		TxId:         txId,
		LogIndex:     logIndex,
	}
}

type Worker struct {
	gorm.Model
	BlockchainUrls pq.StringArray `json:"blockchainUrl" binding:"required" gorm:"type:text[]"`
	BlocksInterval uint64         `json:"blocksInterval" binding:"required"`
}

type AnalyticsWorker struct {
	gorm.Model
	BlockchainUrls pq.StringArray `json:"blockchainUrl" binding:"required" gorm:"type:text[]"`
	BlocksInterval uint64         `json:"blocksInterval" binding:"required"`
	LastBlock      uint64         `json:"lastBlock" binding:"required"`
}

type Token struct {
	gorm.Model
	ChainId  string `json:"chainId" binding:"required" gorm:"uniqueIndex:idx_token_uniqueness"`
	Symbol   string `json:"symbol" binding:"required" gorm:"uniqueIndex:idx_token_uniqueness"`
	Address  string `json:"address" binding:"required" gorm:"uniqueIndex:idx_token_uniqueness"`
	Decimals DBInt  `json:"decimals" binding:"required"`
}

const (
	Aave      = "Aave"
	Compound3 = "Compound3"
	UniswapV3 = "UniswapV3"
)

type DeFiPlatform struct {
	gorm.Model
	ChainId               string `json:"chainId" binding:"required" gorm:"uniqueIndex:idx_platform_uniqueness"`
	Address               string `json:"address" binding:"required" gorm:"uniqueIndex:idx_platform_uniqueness"`
	ExtraContractAddress1 string `json:"extraContractAddress1" binding:"required"`
	Type                  string `json:"type" binding:"required"`
}

type TrackedWallet struct {
	gorm.Model
	Address   string `json:"address" binding:"required" gorm:"uniqueIndex:idx_wallet_uniqueness"`
	ChainId   string `json:"chainId" binding:"required" gorm:"uniqueIndex:idx_wallet_uniqueness"`
	LastBlock uint64 `json:"lastBlock" binding:"required"`
}

type BalanceAcrossAllChains struct {
	Address string `json:"address" binding:"required"`
	Balance string `json:"balance" binding:"required"`
}

func NewBalanceAcrossAllChains(address string, balance string) *BalanceAcrossAllChains {
	return &BalanceAcrossAllChains{
		Address: address,
		Balance: balance,
	}
}

type BalanceOnChain struct {
	ChainId string `json:"chainId" binding:"required"`
	Address string `json:"address" binding:"required"`
	Balance string `json:"balance" binding:"required"`
}

func (b *BalanceOnChain) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}

func NewBalanceOnChain(chainId string, address string, balance string) *BalanceOnChain {
	return &BalanceOnChain{
		Address: address,
		Balance: balance,
		ChainId: chainId,
	}
}

type DealsByWallet struct {
	Address  string `json:"address" binding:"required"`
	DealsIn  []Deal `json:"dealsIn" binding:"required"`
	DealsOut []Deal `json:"dealsOut" binding:"required"`
}

func NewDealsByWallet(wallet string, dealsIn []Deal, dealsOut []Deal) *DealsByWallet {
	return &DealsByWallet{
		Address:  wallet,
		DealsIn:  dealsIn,
		DealsOut: dealsOut,
	}
}
