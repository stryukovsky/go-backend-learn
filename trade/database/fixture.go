package database

import (
	"math/big"

	"github.com/stryukovsky/go-backend-learn/trade"
	"gorm.io/gorm"
)

func Fixture(db *gorm.DB) {
	url := "http://localhost:8545"
	db.Create(&trade.Worker{BlockchainUrls: []*string{&url}, BlocksInterval: 1000})
	db.Create(&trade.AnalyticsWorker{BlockchainUrls: []*string{&url}, BlocksInterval: 1000, LastBlock: 12369651})
	db.Create(&trade.TrackedWallet{ChainId: "1", Address: "0x8EB8a3b98659Cce290402893d0123abb75E3ab28", LastBlock: 23152597})
	db.Create(&trade.TrackedWallet{ChainId: "1", Address: "0x49ecd0F2De4868E5130fdC2C45D4d97444B7c269", LastBlock: 23495633})
	db.Create(&trade.TrackedWallet{ChainId: "1", Address: "0x66a9893cC07D91D95644AEDD05D03f95e1dBA8Af", LastBlock: 23153501})

	db.Create(
		&trade.Token{
			ChainId:  "1",
			Address:  "0xdAC17F958D2ee523a2206206994597C13D831ec7",
			Symbol:   "USDT",
			Decimals: trade.NewDBInt(big.NewInt(6)),
		})
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
	db.Create(
		&trade.Token{
			ChainId:  "1",
			Address:  "0xB8c77482e45F1F44dE1745F52C74426C631bDD52",
			Symbol:   "BNB",
			Decimals: trade.NewDBInt(big.NewInt(18)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "1",
			Address:  "0x514910771AF9Ca656af840dff83E8264EcF986CA",
			Symbol:   "LINK",
			Decimals: trade.NewDBInt(big.NewInt(18)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "1",
			Address:  "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
			Symbol:   "UNI",
			Decimals: trade.NewDBInt(big.NewInt(18)),
		})
	db.Create(&trade.Chain{Name: "Ethereum mainnet", ChainId: "1"})
	db.Create(&trade.DeFiPlatform{Type: trade.Aave, ChainId: "1", Address: "0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"})

	// Uniswap V3 Pools for Ethereum Mainnet
	// Fee tiers: 100 (0.01%), 500 (0.05%), 3000 (0.3%), 10000 (1%)
	// Stablecoin pairs typically use 100 or 500 basis points
	// Standard pairs typically use 3000 basis points (0.3%)

	// ========================================
	// STABLECOIN POOLS (Lower fees: 0.01% or 0.05%)
	// ========================================

	nonFungiblePositionsManager := "0xC36442b4a4522E871399CD717aBDD847Ab11FE88"

	// USDC/USDT - 0.01% fee (most liquid stablecoin pair)
	db.Create(&trade.DeFiPlatform{
		Type:                  trade.UniswapV3,
		ChainId:               "1",
		Address:               "0x3416cf6c708da44db2624d63ea0aaef7113527c6",
		ExtraContractAddress1: nonFungiblePositionsManager,
	})

	// EURC/USDC - 0.05% fee (Euro stablecoin to USD stablecoin)
	db.Create(&trade.DeFiPlatform{
		Type:                  trade.UniswapV3,
		ChainId:               "1",
		Address:               "0x95dbb3c7546f22bce375900abfdd64a4e5bd73d6",
		ExtraContractAddress1: nonFungiblePositionsManager,
	})

	// USDT/EURC - 0.05% fee (alternative stablecoin pair)
	// Note: Check if this pool exists with sufficient liquidity
	// You may need to query the factory contract to confirm

	// ========================================
	// WETH PAIRS (Standard 0.3% fee)
	// ========================================

	// USDC/WETH - 0.3% fee (highest liquidity USDC/ETH pool)
	db.Create(&trade.DeFiPlatform{
		Type:                  trade.UniswapV3,
		ChainId:               "1",
		Address:               "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640",
		ExtraContractAddress1: nonFungiblePositionsManager,
	})

	// Alternative: USDC/WETH - 0.3% fee (secondary pool)
	db.Create(&trade.DeFiPlatform{
		Type:                  trade.UniswapV3,
		ChainId:               "1",
		Address:               "0x8ad599c3a0ff1de082011efddc58f1908eb6e6d8",
		ExtraContractAddress1: nonFungiblePositionsManager,
	})

	// USDT/WETH - 0.3% fee (high liquidity USDT/ETH pool)
	db.Create(&trade.DeFiPlatform{
		Type:                  trade.UniswapV3,
		ChainId:               "1",
		Address:               "0x4e68ccd3e89f51c3074ca5072bbac773960dfa36",
		ExtraContractAddress1: nonFungiblePositionsManager,
	})

	// Alternative: USDT/WETH - 0.3% fee (secondary pool)
	db.Create(&trade.DeFiPlatform{
		Type:                  trade.UniswapV3,
		ChainId:               "1",
		Address:               "0x11b815efb8f581194ae79006d24e0d814b7697f6",
		ExtraContractAddress1: nonFungiblePositionsManager,
	})

	// LINK/WETH - 0.3% fee (high liquidity Chainlink pool)
	db.Create(&trade.DeFiPlatform{
		Type:                  trade.UniswapV3,
		ChainId:               "1",
		Address:               "0xa6cc3c2531fdaa6ae1a3ca84c2855806728693e8",
		ExtraContractAddress1: nonFungiblePositionsManager,
	})

	// UNI/WETH - 0.3% fee (highest liquidity Uniswap token pool)
	db.Create(&trade.DeFiPlatform{
		Type:                  trade.UniswapV3,
		ChainId:               "1",
		Address:               "0x1d42064fc4beb5f8aaf85f4617ae8b3b5b8bd801",
		ExtraContractAddress1: nonFungiblePositionsManager,
	})

	// ========================================
	// BNB PAIRS (1% fee due to lower liquidity)
	// ========================================

	// BNB/WETH - 1% fee (most liquid BNB pool on Ethereum)
	db.Create(&trade.DeFiPlatform{
		Type:                  trade.UniswapV3,
		ChainId:               "1",
		Address:               "0x9e7809c21ba130c1a51c112928ea6474d9a9ae3c",
		ExtraContractAddress1: nonFungiblePositionsManager,
	})

	// Alternative: BNB/WETH - 0.3% fee (lower liquidity)
	// Only use if you need lower fee tier despite lower liquidity
	// db.Create(&trade.DeFiPlatform{
	// 	Type:                  trade.UniswapV3,
	// 	ChainId:               "1",
	// 	Address:               "0x881b8d0b1ad9d1b1db918342b064d10afb9eaa69",
	// 	ExtraContractAddress1: nonFungiblePositionsManager,
	// })

	// ========================================
	// CROSS-STABLECOIN PAIRS TO CONSIDER
	// ========================================

	// USDC/USDT already covered above with 0.01% fee

	// For EURC pairs with USDT, you should query the Uniswap V3 Factory:
	// factory.getPool(token0, token1, fee)
	// Address: 0x1F98431c8aD98523631AE4a59f267346ea31F984
	//
	// Example fees to try:
	// - 100 (0.01%) for tight stablecoin pairs
	// - 500 (0.05%) for correlated pairs
	//
	// Token addresses for reference:
	// USDT:  0xdAC17F958D2ee523a2206206994597C13D831ec7
	// USDC:  0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48
	// EURC:  0x1aBaEA1f7C830bD89Acc67eC4af516284b1bC33c
	// WETH:  0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
	// BNB:   0xB8c77482e45F1F44dE1745F52C74426C631bDD52
	// LINK:  0x514910771AF9Ca656af840dff83E8264EcF986CA
	// UNI:   0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984

	// ========================================
	// ADDITIONAL PAIRS YOU MAY WANT
	// ========================================

	// LINK/USDC, LINK/USDT, UNI/USDC, UNI/USDT, BNB/USDC, BNB/USDT
	// These can be found by querying the factory contract
	// Most will likely use 0.3% fee tier (3000 basis points)
	//
	// To find them programmatically, use:
	// web3.Call(factoryContract, "getPool", [token0Address, token1Address, feeAmount])
	//
	// Where feeAmount is typically:
	// - 100 for tight stablecoin pairs (0.01%)
	// - 500 for correlated assets (0.05%)
	// - 3000 for standard pairs (0.3%)
	// - 10000 for exotic pairs (1%)
}
