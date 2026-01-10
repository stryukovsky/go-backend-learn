package database

import (
	"encoding/json"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lib/pq"
	"github.com/samber/lo"
	"github.com/stryukovsky/go-backend-learn/trade"
	"gorm.io/gorm"
)

func Arbitrum(db *gorm.DB) {
	blockchainUrlsForCache := make(pq.StringArray, 0)
	blockchainUrlsForEvents := make(pq.StringArray, 0)
	blockchainUrlsForEvents = append(blockchainUrlsForEvents, "https://arb1.arbitrum.io/rpc")

	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://1rpc.io/arb")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://arbitrum-one-rpc.publicnode.com")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://arbitrum-one-public.nodies.app")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://arb-mainnet.g.alchemy.com/v2/demo")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://arbitrum.public.blockpi.network/v1/rpc/public")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://arbitrum-one.public.blastapi.io")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://arbitrum-one-rpc.publicnode.com")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "wss://arbitrum-one-rpc.publicnode.com")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://arbitrum.meowrpc.com")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://api.zan.top/arb-one")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://arbitrum.drpc.org")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://api.stateless.solutions/arbitrum-one/v1/demo")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://arbitrum.gateway.tenderly.co")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://endpoints.omniatech.io/v1/arbitrum/one/public")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://arb1.lava.build")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://arbitrum.api.onfinality.io/public")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://arb-one-mainnet.gateway.tatum.io")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://rpc.poolz.finance/arbitrum")
	blockchainUrlsForCache = append(blockchainUrlsForCache, "https://arb-one.api.pocket.network")

	db.Create(&trade.Worker{BlockchainUrlsForEvents: blockchainUrlsForEvents, BlockchainUrlsForCacheManager: blockchainUrlsForCache, BlocksInterval: 50000})
	db.Create(
		&trade.Token{
			ChainId:  "42161",
			Address:  "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9",
			Symbol:   "USDT",
			Decimals: trade.NewDBInt(big.NewInt(6)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "42161",
			Address:  "0x82aF49447D8a07e3bd95BD0d56f35241523fBab1",
			Symbol:   "WETH",
			Decimals: trade.NewDBInt(big.NewInt(18)),
		})
	db.Create(
		// Note: No direct EUR stablecoin equivalent found on Arbitrum
		&trade.Token{
			ChainId:  "42161",
			Address:  "0xaf88d065e77c8cC2239327C5EDb3A432268e5831",
			Symbol:   "USDC",
			Decimals: trade.NewDBInt(big.NewInt(6)),
		})
	// Note: BNB is not natively available on Arbitrum (it's a BSC token)
	db.Create(
		&trade.Token{
			ChainId:  "42161",
			Address:  "0xf97f4df75117a78c1A5a0DBb814Af92458539FB4",
			Symbol:   "LINK",
			Decimals: trade.NewDBInt(big.NewInt(18)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "42161",
			Address:  "0xFa7F8980b0f1E64A2062791cc3b0871572f1F7f0",
			Symbol:   "UNI",
			Decimals: trade.NewDBInt(big.NewInt(18)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "42161",
			Address:  "0x2f2a2543B76A4166549F7aaB2e75Bef0aefC5B0f",
			Symbol:   "WBTC",
			Decimals: trade.NewDBInt(big.NewInt(8)),
		})
	db.Create(&trade.Chain{Name: "Arbitrum One", ChainId: "42161"})
	db.Create(&trade.DeFiPlatform{Type: trade.Aave, ChainId: "42161", Address: "0x794a61358D6845594F94dc1DB02A252b5b4814aD"})
	db.Create(&trade.DeFiPlatform{Type: trade.Compound3, ChainId: "42161", Address: "0x9c4ec768c28520B50860ea7a15bd7213a9fF58bf"})
}

func Binance(db *gorm.DB) {
	// BSC Chain ID is 56
	db.Create(
		&trade.Token{
			ChainId:  "56",
			Address:  "0x55d398326f99059fF775485246999027B3197955",
			Symbol:   "USDT",
			Decimals: trade.NewDBInt(big.NewInt(6)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "56",
			Address:  "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
			Symbol:   "BNB",
			Decimals: trade.NewDBInt(big.NewInt(18)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "56",
			Address:  "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",
			Symbol:   "BUSD",
			Decimals: trade.NewDBInt(big.NewInt(6)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "56",
			Address:  "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d",
			Symbol:   "USDC",
			Decimals: trade.NewDBInt(big.NewInt(6)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "56",
			Address:  "0x2170Ed0880ac9A755fd29B2688956BD959F933F8",
			Symbol:   "ETH",
			Decimals: trade.NewDBInt(big.NewInt(18)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "56",
			Address:  "0xF8A0BF9cF54Bb92F17374d9e9A321E6a111a51bD",
			Symbol:   "LINK",
			Decimals: trade.NewDBInt(big.NewInt(18)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "56",
			Address:  "0xBf5140A22578168FD562DCcF235E5D43A02ce9B1",
			Symbol:   "UNI",
			Decimals: trade.NewDBInt(big.NewInt(18)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "56",
			Address:  "0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c",
			Symbol:   "BTC",
			Decimals: trade.NewDBInt(big.NewInt(8)),
		})
	db.Create(&trade.Chain{Name: "BNB Smart Chain", ChainId: "56"})
	db.Create(&trade.DeFiPlatform{Type: trade.Aave, ChainId: "56", Address: "0x6807dc923806fE8Fd134338EABCA509979a7e0cB"})
}

func Base(db *gorm.DB) {
	// Base Chain ID is 8453
	db.Create(
		&trade.Token{
			ChainId:  "8453",
			Address:  "0xfde4C96c8593536E31F229EA8f37b2ADa2699bb2",
			Symbol:   "USDT",
			Decimals: trade.NewDBInt(big.NewInt(6)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "8453",
			Address:  "0x4200000000000000000000000000000000000006",
			Symbol:   "ETH",
			Decimals: trade.NewDBInt(big.NewInt(18)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "8453",
			Address:  "0x833589fcd6edb6e08f4c7c32d4f71b54bda02913",
			Symbol:   "USDC",
			Decimals: trade.NewDBInt(big.NewInt(6)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "8453",
			Address:  "0x1d4bB5576050E69a95d8A726dF4fF23125130EB5",
			Symbol:   "LINK",
			Decimals: trade.NewDBInt(big.NewInt(18)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "8453",
			Address:  "0x90dF94286a30E52a776ed76Ec97d9dAC6C745936",
			Symbol:   "UNI",
			Decimals: trade.NewDBInt(big.NewInt(18)),
		})
	db.Create(
		&trade.Token{
			ChainId:  "8453",
			Address:  "0x0555E30da8f98308EdB960aa94C0Db47230d2B9c",
			Symbol:   "BTC",
			Decimals: trade.NewDBInt(big.NewInt(8)),
		})
	db.Create(&trade.Chain{Name: "Base", ChainId: "8453"})
	db.Create(&trade.DeFiPlatform{Type: trade.Aave, ChainId: "8453", Address: "0xA4CF4740E5F564D2BFC7382352c61F202D894287"})
	db.Create(&trade.DeFiPlatform{Type: trade.Compound3, ChainId: "8453", Address: "0xBA12222222228d8Ba445958a75a0704d566BF2C8"})
}

func Wallets(db *gorm.DB) {
	addresses := []string{}
	addressesSource, err := os.ReadFile("wallets.json")
	if err != nil {
		panic("Cannot read wallets file")
	}
	err = json.Unmarshal(addressesSource, &addresses)
	if err != nil {
		panic("Cannot read wallets file: file exists but malformed")
	}
	lo.ForEach(addresses, func(item string, index int) {
		db.Create(&trade.TrackedWallet{ChainId: "42161", Address: common.HexToAddress(item).String(), LastBlock: 210147506})
		db.Create(&trade.TrackedWallet{ChainId: "56", Address: item, LastBlock: 33484103})
		db.Create(&trade.TrackedWallet{ChainId: "8453", Address: item, LastBlock: 20163035})
	})
}

func Ethereum(db *gorm.DB) {
	url := "https://eth-mainnet.public.blastapi.io"
	blockchainUrls := make(pq.StringArray, 1)
	blockchainUrls[0] = url

	db.Create(&trade.Worker{BlockchainUrlsForCacheManager: blockchainUrls, BlocksInterval: 1000})
	db.Create(&trade.AnalyticsWorker{BlockchainUrls: blockchainUrls, BlocksInterval: 1000, LastBlock: 12369651})

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
	db.Create(
		&trade.Token{
			ChainId:  "1",
			Address:  "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
			Symbol:   "BTC",
			Decimals: trade.NewDBInt(big.NewInt(8)),
		})
	db.Create(&trade.Chain{Name: "Ethereum mainnet", ChainId: "1"})
	db.Create(&trade.DeFiPlatform{Type: trade.Aave, ChainId: "1", Address: "0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"})
	db.Create(&trade.DeFiPlatform{Type: trade.Compound3, ChainId: "1", Address: "0xc3d688B66703497DAA19211EEdff47f25384cdc3"})

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
}
