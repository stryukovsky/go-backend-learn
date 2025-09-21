// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DataTypesCollateralConfig is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCollateralConfig struct {
	Ltv                  uint16
	LiquidationThreshold uint16
	LiquidationBonus     uint16
}

// DataTypesEModeCategoryBaseConfiguration is an auto generated low-level Go binding around an user-defined struct.
type DataTypesEModeCategoryBaseConfiguration struct {
	Ltv                  uint16
	LiquidationThreshold uint16
	LiquidationBonus     uint16
	Label                string
}

// DataTypesEModeCategoryLegacy is an auto generated low-level Go binding around an user-defined struct.
type DataTypesEModeCategoryLegacy struct {
	Ltv                  uint16
	LiquidationThreshold uint16
	LiquidationBonus     uint16
	PriceSource          common.Address
	Label                string
}

// DataTypesReserveConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveConfigurationMap struct {
	Data *big.Int
}

// DataTypesReserveDataLegacy is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveDataLegacy struct {
	Configuration               DataTypesReserveConfigurationMap
	LiquidityIndex              *big.Int
	CurrentLiquidityRate        *big.Int
	VariableBorrowIndex         *big.Int
	CurrentVariableBorrowRate   *big.Int
	CurrentStableBorrowRate     *big.Int
	LastUpdateTimestamp         *big.Int
	Id                          uint16
	ATokenAddress               common.Address
	StableDebtTokenAddress      common.Address
	VariableDebtTokenAddress    common.Address
	InterestRateStrategyAddress common.Address
	AccruedToTreasury           *big.Int
	Unbacked                    *big.Int
	IsolationModeTotalDebt      *big.Int
}

// DataTypesUserConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesUserConfigurationMap struct {
	Data *big.Int
}

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"contractIReserveInterestRateStrategy\",\"name\":\"interestRateStrategy_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssetNotListed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotAToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPoolAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPoolConfigurator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPositionManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotUmbrella\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EModeCategoryReserved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddressesProvider\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotValid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumDataTypes.InterestRateMode\",\"name\":\"interestRateMode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountCovered\",\"type\":\"uint256\"}],\"name\":\"DeficitCovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"debtAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountCreated\",\"type\":\"uint256\"}],\"name\":\"DeficitCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumDataTypes.InterestRateMode\",\"name\":\"interestRateMode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDebt\",\"type\":\"uint256\"}],\"name\":\"IsolationModeTotalDebtUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"debtAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtToCover\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidatedCollateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"receiveAToken\",\"type\":\"bool\"}],\"name\":\"LiquidationCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountMinted\",\"type\":\"uint256\"}],\"name\":\"MintedToTreasury\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"positionManager\",\"type\":\"address\"}],\"name\":\"PositionManagerApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"positionManager\",\"type\":\"address\"}],\"name\":\"PositionManagerRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"repayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useATokens\",\"type\":\"bool\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"ReserveDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"Supply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"categoryId\",\"type\":\"uint8\"}],\"name\":\"UserEModeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLASHLOAN_PREMIUM_TOTAL\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLASHLOAN_PREMIUM_TO_PROTOCOL\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUMBER_RESERVES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_INTEREST_RATE_STRATEGY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UMBRELLA\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"positionManager\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approve\",\"type\":\"bool\"}],\"name\":\"approvePositionManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ltv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationBonus\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"internalType\":\"structDataTypes.EModeCategoryBaseConfiguration\",\"name\":\"category\",\"type\":\"tuple\"}],\"name\":\"configureEModeCategory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"borrowableBitmap\",\"type\":\"uint128\"}],\"name\":\"configureEModeCategoryBorrowableBitmap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"collateralBitmap\",\"type\":\"uint128\"}],\"name\":\"configureEModeCategoryCollateralBitmap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"dropReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"eliminateReserveDeficit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scaledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledBalanceFromBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledBalanceToBefore\",\"type\":\"uint256\"}],\"name\":\"finalizeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"interestRateModes\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"flashLoanSimple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBorrowLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getEModeCategoryBorrowableBitmap\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getEModeCategoryCollateralBitmap\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getEModeCategoryCollateralConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ltv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationBonus\",\"type\":\"uint16\"}],\"internalType\":\"structDataTypes.CollateralConfig\",\"name\":\"res\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getEModeCategoryData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ltv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationBonus\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceSource\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"internalType\":\"structDataTypes.EModeCategoryLegacy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getEModeCategoryLabel\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEModeLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFlashLoanLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getLiquidationGracePeriod\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidationLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveAToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"}],\"name\":\"getReserveAddressById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentLiquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentVariableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentStableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accruedToTreasury\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"unbacked\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"isolationModeTotalDebt\",\"type\":\"uint128\"}],\"internalType\":\"structDataTypes.ReserveDataLegacy\",\"name\":\"res\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveDeficit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedVariableDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveVariableDebtToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupplyLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserEMode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getVirtualUnderlyingBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtAddress\",\"type\":\"address\"}],\"name\":\"initReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"positionManager\",\"type\":\"address\"}],\"name\":\"isApprovedPositionManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtToCover\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"receiveAToken\",\"type\":\"bool\"}],\"name\":\"liquidationCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"mintToTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"renouncePositionManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"}],\"name\":\"repayWithATokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"permitV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"permitR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"permitS\",\"type\":\"bytes32\"}],\"name\":\"repayWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"resetIsolationModeTotalDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"}],\"name\":\"setConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"until\",\"type\":\"uint40\"}],\"name\":\"setLiquidationGracePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"categoryId\",\"type\":\"uint8\"}],\"name\":\"setUserEMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"categoryId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"setUserEModeOnBehalfOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useAsCollateral\",\"type\":\"bool\"}],\"name\":\"setUserUseReserveAsCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useAsCollateral\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"setUserUseReserveAsCollateralOnBehalfOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"supply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"permitV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"permitR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"permitS\",\"type\":\"bytes32\"}],\"name\":\"supplyWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"syncIndexesState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"syncRatesState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"flashLoanPremium\",\"type\":\"uint128\"}],\"name\":\"updateFlashloanPremium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Pool *PoolCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Pool *PoolSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _Pool.Contract.ADDRESSESPROVIDER(&_Pool.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Pool *PoolCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _Pool.Contract.ADDRESSESPROVIDER(&_Pool.CallOpts)
}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint128)
func (_Pool *PoolCaller) FLASHLOANPREMIUMTOTAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "FLASHLOAN_PREMIUM_TOTAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint128)
func (_Pool *PoolSession) FLASHLOANPREMIUMTOTAL() (*big.Int, error) {
	return _Pool.Contract.FLASHLOANPREMIUMTOTAL(&_Pool.CallOpts)
}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint128)
func (_Pool *PoolCallerSession) FLASHLOANPREMIUMTOTAL() (*big.Int, error) {
	return _Pool.Contract.FLASHLOANPREMIUMTOTAL(&_Pool.CallOpts)
}

// FLASHLOANPREMIUMTOPROTOCOL is a free data retrieval call binding the contract method 0x6a99c036.
//
// Solidity: function FLASHLOAN_PREMIUM_TO_PROTOCOL() view returns(uint128)
func (_Pool *PoolCaller) FLASHLOANPREMIUMTOPROTOCOL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "FLASHLOAN_PREMIUM_TO_PROTOCOL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLASHLOANPREMIUMTOPROTOCOL is a free data retrieval call binding the contract method 0x6a99c036.
//
// Solidity: function FLASHLOAN_PREMIUM_TO_PROTOCOL() view returns(uint128)
func (_Pool *PoolSession) FLASHLOANPREMIUMTOPROTOCOL() (*big.Int, error) {
	return _Pool.Contract.FLASHLOANPREMIUMTOPROTOCOL(&_Pool.CallOpts)
}

// FLASHLOANPREMIUMTOPROTOCOL is a free data retrieval call binding the contract method 0x6a99c036.
//
// Solidity: function FLASHLOAN_PREMIUM_TO_PROTOCOL() view returns(uint128)
func (_Pool *PoolCallerSession) FLASHLOANPREMIUMTOPROTOCOL() (*big.Int, error) {
	return _Pool.Contract.FLASHLOANPREMIUMTOPROTOCOL(&_Pool.CallOpts)
}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint16)
func (_Pool *PoolCaller) MAXNUMBERRESERVES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "MAX_NUMBER_RESERVES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint16)
func (_Pool *PoolSession) MAXNUMBERRESERVES() (uint16, error) {
	return _Pool.Contract.MAXNUMBERRESERVES(&_Pool.CallOpts)
}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint16)
func (_Pool *PoolCallerSession) MAXNUMBERRESERVES() (uint16, error) {
	return _Pool.Contract.MAXNUMBERRESERVES(&_Pool.CallOpts)
}

// POOLREVISION is a free data retrieval call binding the contract method 0x0148170e.
//
// Solidity: function POOL_REVISION() view returns(uint256)
func (_Pool *PoolCaller) POOLREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "POOL_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// POOLREVISION is a free data retrieval call binding the contract method 0x0148170e.
//
// Solidity: function POOL_REVISION() view returns(uint256)
func (_Pool *PoolSession) POOLREVISION() (*big.Int, error) {
	return _Pool.Contract.POOLREVISION(&_Pool.CallOpts)
}

// POOLREVISION is a free data retrieval call binding the contract method 0x0148170e.
//
// Solidity: function POOL_REVISION() view returns(uint256)
func (_Pool *PoolCallerSession) POOLREVISION() (*big.Int, error) {
	return _Pool.Contract.POOLREVISION(&_Pool.CallOpts)
}

// RESERVEINTERESTRATESTRATEGY is a free data retrieval call binding the contract method 0x1b8feb0e.
//
// Solidity: function RESERVE_INTEREST_RATE_STRATEGY() view returns(address)
func (_Pool *PoolCaller) RESERVEINTERESTRATESTRATEGY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "RESERVE_INTEREST_RATE_STRATEGY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RESERVEINTERESTRATESTRATEGY is a free data retrieval call binding the contract method 0x1b8feb0e.
//
// Solidity: function RESERVE_INTEREST_RATE_STRATEGY() view returns(address)
func (_Pool *PoolSession) RESERVEINTERESTRATESTRATEGY() (common.Address, error) {
	return _Pool.Contract.RESERVEINTERESTRATESTRATEGY(&_Pool.CallOpts)
}

// RESERVEINTERESTRATESTRATEGY is a free data retrieval call binding the contract method 0x1b8feb0e.
//
// Solidity: function RESERVE_INTEREST_RATE_STRATEGY() view returns(address)
func (_Pool *PoolCallerSession) RESERVEINTERESTRATESTRATEGY() (common.Address, error) {
	return _Pool.Contract.RESERVEINTERESTRATESTRATEGY(&_Pool.CallOpts)
}

// UMBRELLA is a free data retrieval call binding the contract method 0x71459c15.
//
// Solidity: function UMBRELLA() view returns(bytes32)
func (_Pool *PoolCaller) UMBRELLA(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "UMBRELLA")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UMBRELLA is a free data retrieval call binding the contract method 0x71459c15.
//
// Solidity: function UMBRELLA() view returns(bytes32)
func (_Pool *PoolSession) UMBRELLA() ([32]byte, error) {
	return _Pool.Contract.UMBRELLA(&_Pool.CallOpts)
}

// UMBRELLA is a free data retrieval call binding the contract method 0x71459c15.
//
// Solidity: function UMBRELLA() view returns(bytes32)
func (_Pool *PoolCallerSession) UMBRELLA() ([32]byte, error) {
	return _Pool.Contract.UMBRELLA(&_Pool.CallOpts)
}

// GetBorrowLogic is a free data retrieval call binding the contract method 0x2be29fa7.
//
// Solidity: function getBorrowLogic() pure returns(address)
func (_Pool *PoolCaller) GetBorrowLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getBorrowLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBorrowLogic is a free data retrieval call binding the contract method 0x2be29fa7.
//
// Solidity: function getBorrowLogic() pure returns(address)
func (_Pool *PoolSession) GetBorrowLogic() (common.Address, error) {
	return _Pool.Contract.GetBorrowLogic(&_Pool.CallOpts)
}

// GetBorrowLogic is a free data retrieval call binding the contract method 0x2be29fa7.
//
// Solidity: function getBorrowLogic() pure returns(address)
func (_Pool *PoolCallerSession) GetBorrowLogic() (common.Address, error) {
	return _Pool.Contract.GetBorrowLogic(&_Pool.CallOpts)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_Pool *PoolCaller) GetConfiguration(opts *bind.CallOpts, asset common.Address) (DataTypesReserveConfigurationMap, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getConfiguration", asset)

	if err != nil {
		return *new(DataTypesReserveConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveConfigurationMap)).(*DataTypesReserveConfigurationMap)

	return out0, err

}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_Pool *PoolSession) GetConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _Pool.Contract.GetConfiguration(&_Pool.CallOpts, asset)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_Pool *PoolCallerSession) GetConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _Pool.Contract.GetConfiguration(&_Pool.CallOpts, asset)
}

// GetEModeCategoryBorrowableBitmap is a free data retrieval call binding the contract method 0x903a2c71.
//
// Solidity: function getEModeCategoryBorrowableBitmap(uint8 id) view returns(uint128)
func (_Pool *PoolCaller) GetEModeCategoryBorrowableBitmap(opts *bind.CallOpts, id uint8) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getEModeCategoryBorrowableBitmap", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEModeCategoryBorrowableBitmap is a free data retrieval call binding the contract method 0x903a2c71.
//
// Solidity: function getEModeCategoryBorrowableBitmap(uint8 id) view returns(uint128)
func (_Pool *PoolSession) GetEModeCategoryBorrowableBitmap(id uint8) (*big.Int, error) {
	return _Pool.Contract.GetEModeCategoryBorrowableBitmap(&_Pool.CallOpts, id)
}

// GetEModeCategoryBorrowableBitmap is a free data retrieval call binding the contract method 0x903a2c71.
//
// Solidity: function getEModeCategoryBorrowableBitmap(uint8 id) view returns(uint128)
func (_Pool *PoolCallerSession) GetEModeCategoryBorrowableBitmap(id uint8) (*big.Int, error) {
	return _Pool.Contract.GetEModeCategoryBorrowableBitmap(&_Pool.CallOpts, id)
}

// GetEModeCategoryCollateralBitmap is a free data retrieval call binding the contract method 0xb0771dba.
//
// Solidity: function getEModeCategoryCollateralBitmap(uint8 id) view returns(uint128)
func (_Pool *PoolCaller) GetEModeCategoryCollateralBitmap(opts *bind.CallOpts, id uint8) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getEModeCategoryCollateralBitmap", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEModeCategoryCollateralBitmap is a free data retrieval call binding the contract method 0xb0771dba.
//
// Solidity: function getEModeCategoryCollateralBitmap(uint8 id) view returns(uint128)
func (_Pool *PoolSession) GetEModeCategoryCollateralBitmap(id uint8) (*big.Int, error) {
	return _Pool.Contract.GetEModeCategoryCollateralBitmap(&_Pool.CallOpts, id)
}

// GetEModeCategoryCollateralBitmap is a free data retrieval call binding the contract method 0xb0771dba.
//
// Solidity: function getEModeCategoryCollateralBitmap(uint8 id) view returns(uint128)
func (_Pool *PoolCallerSession) GetEModeCategoryCollateralBitmap(id uint8) (*big.Int, error) {
	return _Pool.Contract.GetEModeCategoryCollateralBitmap(&_Pool.CallOpts, id)
}

// GetEModeCategoryCollateralConfig is a free data retrieval call binding the contract method 0xb286f467.
//
// Solidity: function getEModeCategoryCollateralConfig(uint8 id) view returns((uint16,uint16,uint16) res)
func (_Pool *PoolCaller) GetEModeCategoryCollateralConfig(opts *bind.CallOpts, id uint8) (DataTypesCollateralConfig, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getEModeCategoryCollateralConfig", id)

	if err != nil {
		return *new(DataTypesCollateralConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesCollateralConfig)).(*DataTypesCollateralConfig)

	return out0, err

}

// GetEModeCategoryCollateralConfig is a free data retrieval call binding the contract method 0xb286f467.
//
// Solidity: function getEModeCategoryCollateralConfig(uint8 id) view returns((uint16,uint16,uint16) res)
func (_Pool *PoolSession) GetEModeCategoryCollateralConfig(id uint8) (DataTypesCollateralConfig, error) {
	return _Pool.Contract.GetEModeCategoryCollateralConfig(&_Pool.CallOpts, id)
}

// GetEModeCategoryCollateralConfig is a free data retrieval call binding the contract method 0xb286f467.
//
// Solidity: function getEModeCategoryCollateralConfig(uint8 id) view returns((uint16,uint16,uint16) res)
func (_Pool *PoolCallerSession) GetEModeCategoryCollateralConfig(id uint8) (DataTypesCollateralConfig, error) {
	return _Pool.Contract.GetEModeCategoryCollateralConfig(&_Pool.CallOpts, id)
}

// GetEModeCategoryData is a free data retrieval call binding the contract method 0x6c6f6ae1.
//
// Solidity: function getEModeCategoryData(uint8 id) view returns((uint16,uint16,uint16,address,string))
func (_Pool *PoolCaller) GetEModeCategoryData(opts *bind.CallOpts, id uint8) (DataTypesEModeCategoryLegacy, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getEModeCategoryData", id)

	if err != nil {
		return *new(DataTypesEModeCategoryLegacy), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesEModeCategoryLegacy)).(*DataTypesEModeCategoryLegacy)

	return out0, err

}

// GetEModeCategoryData is a free data retrieval call binding the contract method 0x6c6f6ae1.
//
// Solidity: function getEModeCategoryData(uint8 id) view returns((uint16,uint16,uint16,address,string))
func (_Pool *PoolSession) GetEModeCategoryData(id uint8) (DataTypesEModeCategoryLegacy, error) {
	return _Pool.Contract.GetEModeCategoryData(&_Pool.CallOpts, id)
}

// GetEModeCategoryData is a free data retrieval call binding the contract method 0x6c6f6ae1.
//
// Solidity: function getEModeCategoryData(uint8 id) view returns((uint16,uint16,uint16,address,string))
func (_Pool *PoolCallerSession) GetEModeCategoryData(id uint8) (DataTypesEModeCategoryLegacy, error) {
	return _Pool.Contract.GetEModeCategoryData(&_Pool.CallOpts, id)
}

// GetEModeCategoryLabel is a free data retrieval call binding the contract method 0x2083e183.
//
// Solidity: function getEModeCategoryLabel(uint8 id) view returns(string)
func (_Pool *PoolCaller) GetEModeCategoryLabel(opts *bind.CallOpts, id uint8) (string, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getEModeCategoryLabel", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEModeCategoryLabel is a free data retrieval call binding the contract method 0x2083e183.
//
// Solidity: function getEModeCategoryLabel(uint8 id) view returns(string)
func (_Pool *PoolSession) GetEModeCategoryLabel(id uint8) (string, error) {
	return _Pool.Contract.GetEModeCategoryLabel(&_Pool.CallOpts, id)
}

// GetEModeCategoryLabel is a free data retrieval call binding the contract method 0x2083e183.
//
// Solidity: function getEModeCategoryLabel(uint8 id) view returns(string)
func (_Pool *PoolCallerSession) GetEModeCategoryLabel(id uint8) (string, error) {
	return _Pool.Contract.GetEModeCategoryLabel(&_Pool.CallOpts, id)
}

// GetEModeLogic is a free data retrieval call binding the contract method 0xf32b9a73.
//
// Solidity: function getEModeLogic() pure returns(address)
func (_Pool *PoolCaller) GetEModeLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getEModeLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEModeLogic is a free data retrieval call binding the contract method 0xf32b9a73.
//
// Solidity: function getEModeLogic() pure returns(address)
func (_Pool *PoolSession) GetEModeLogic() (common.Address, error) {
	return _Pool.Contract.GetEModeLogic(&_Pool.CallOpts)
}

// GetEModeLogic is a free data retrieval call binding the contract method 0xf32b9a73.
//
// Solidity: function getEModeLogic() pure returns(address)
func (_Pool *PoolCallerSession) GetEModeLogic() (common.Address, error) {
	return _Pool.Contract.GetEModeLogic(&_Pool.CallOpts)
}

// GetFlashLoanLogic is a free data retrieval call binding the contract method 0x348fde0f.
//
// Solidity: function getFlashLoanLogic() pure returns(address)
func (_Pool *PoolCaller) GetFlashLoanLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getFlashLoanLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFlashLoanLogic is a free data retrieval call binding the contract method 0x348fde0f.
//
// Solidity: function getFlashLoanLogic() pure returns(address)
func (_Pool *PoolSession) GetFlashLoanLogic() (common.Address, error) {
	return _Pool.Contract.GetFlashLoanLogic(&_Pool.CallOpts)
}

// GetFlashLoanLogic is a free data retrieval call binding the contract method 0x348fde0f.
//
// Solidity: function getFlashLoanLogic() pure returns(address)
func (_Pool *PoolCallerSession) GetFlashLoanLogic() (common.Address, error) {
	return _Pool.Contract.GetFlashLoanLogic(&_Pool.CallOpts)
}

// GetLiquidationGracePeriod is a free data retrieval call binding the contract method 0x5c9a8b18.
//
// Solidity: function getLiquidationGracePeriod(address asset) view returns(uint40)
func (_Pool *PoolCaller) GetLiquidationGracePeriod(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getLiquidationGracePeriod", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidationGracePeriod is a free data retrieval call binding the contract method 0x5c9a8b18.
//
// Solidity: function getLiquidationGracePeriod(address asset) view returns(uint40)
func (_Pool *PoolSession) GetLiquidationGracePeriod(asset common.Address) (*big.Int, error) {
	return _Pool.Contract.GetLiquidationGracePeriod(&_Pool.CallOpts, asset)
}

// GetLiquidationGracePeriod is a free data retrieval call binding the contract method 0x5c9a8b18.
//
// Solidity: function getLiquidationGracePeriod(address asset) view returns(uint40)
func (_Pool *PoolCallerSession) GetLiquidationGracePeriod(asset common.Address) (*big.Int, error) {
	return _Pool.Contract.GetLiquidationGracePeriod(&_Pool.CallOpts, asset)
}

// GetLiquidationLogic is a free data retrieval call binding the contract method 0x911a3413.
//
// Solidity: function getLiquidationLogic() pure returns(address)
func (_Pool *PoolCaller) GetLiquidationLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getLiquidationLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLiquidationLogic is a free data retrieval call binding the contract method 0x911a3413.
//
// Solidity: function getLiquidationLogic() pure returns(address)
func (_Pool *PoolSession) GetLiquidationLogic() (common.Address, error) {
	return _Pool.Contract.GetLiquidationLogic(&_Pool.CallOpts)
}

// GetLiquidationLogic is a free data retrieval call binding the contract method 0x911a3413.
//
// Solidity: function getLiquidationLogic() pure returns(address)
func (_Pool *PoolCallerSession) GetLiquidationLogic() (common.Address, error) {
	return _Pool.Contract.GetLiquidationLogic(&_Pool.CallOpts)
}

// GetPoolLogic is a free data retrieval call binding the contract method 0xd3350155.
//
// Solidity: function getPoolLogic() pure returns(address)
func (_Pool *PoolCaller) GetPoolLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getPoolLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolLogic is a free data retrieval call binding the contract method 0xd3350155.
//
// Solidity: function getPoolLogic() pure returns(address)
func (_Pool *PoolSession) GetPoolLogic() (common.Address, error) {
	return _Pool.Contract.GetPoolLogic(&_Pool.CallOpts)
}

// GetPoolLogic is a free data retrieval call binding the contract method 0xd3350155.
//
// Solidity: function getPoolLogic() pure returns(address)
func (_Pool *PoolCallerSession) GetPoolLogic() (common.Address, error) {
	return _Pool.Contract.GetPoolLogic(&_Pool.CallOpts)
}

// GetReserveAToken is a free data retrieval call binding the contract method 0xcff027d9.
//
// Solidity: function getReserveAToken(address asset) view returns(address)
func (_Pool *PoolCaller) GetReserveAToken(opts *bind.CallOpts, asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getReserveAToken", asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReserveAToken is a free data retrieval call binding the contract method 0xcff027d9.
//
// Solidity: function getReserveAToken(address asset) view returns(address)
func (_Pool *PoolSession) GetReserveAToken(asset common.Address) (common.Address, error) {
	return _Pool.Contract.GetReserveAToken(&_Pool.CallOpts, asset)
}

// GetReserveAToken is a free data retrieval call binding the contract method 0xcff027d9.
//
// Solidity: function getReserveAToken(address asset) view returns(address)
func (_Pool *PoolCallerSession) GetReserveAToken(asset common.Address) (common.Address, error) {
	return _Pool.Contract.GetReserveAToken(&_Pool.CallOpts, asset)
}

// GetReserveAddressById is a free data retrieval call binding the contract method 0x52751797.
//
// Solidity: function getReserveAddressById(uint16 id) view returns(address)
func (_Pool *PoolCaller) GetReserveAddressById(opts *bind.CallOpts, id uint16) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getReserveAddressById", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReserveAddressById is a free data retrieval call binding the contract method 0x52751797.
//
// Solidity: function getReserveAddressById(uint16 id) view returns(address)
func (_Pool *PoolSession) GetReserveAddressById(id uint16) (common.Address, error) {
	return _Pool.Contract.GetReserveAddressById(&_Pool.CallOpts, id)
}

// GetReserveAddressById is a free data retrieval call binding the contract method 0x52751797.
//
// Solidity: function getReserveAddressById(uint16 id) view returns(address)
func (_Pool *PoolCallerSession) GetReserveAddressById(id uint16) (common.Address, error) {
	return _Pool.Contract.GetReserveAddressById(&_Pool.CallOpts, id)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128) res)
func (_Pool *PoolCaller) GetReserveData(opts *bind.CallOpts, asset common.Address) (DataTypesReserveDataLegacy, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getReserveData", asset)

	if err != nil {
		return *new(DataTypesReserveDataLegacy), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveDataLegacy)).(*DataTypesReserveDataLegacy)

	return out0, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128) res)
func (_Pool *PoolSession) GetReserveData(asset common.Address) (DataTypesReserveDataLegacy, error) {
	return _Pool.Contract.GetReserveData(&_Pool.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128) res)
func (_Pool *PoolCallerSession) GetReserveData(asset common.Address) (DataTypesReserveDataLegacy, error) {
	return _Pool.Contract.GetReserveData(&_Pool.CallOpts, asset)
}

// GetReserveDeficit is a free data retrieval call binding the contract method 0xc952485d.
//
// Solidity: function getReserveDeficit(address asset) view returns(uint256)
func (_Pool *PoolCaller) GetReserveDeficit(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getReserveDeficit", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveDeficit is a free data retrieval call binding the contract method 0xc952485d.
//
// Solidity: function getReserveDeficit(address asset) view returns(uint256)
func (_Pool *PoolSession) GetReserveDeficit(asset common.Address) (*big.Int, error) {
	return _Pool.Contract.GetReserveDeficit(&_Pool.CallOpts, asset)
}

// GetReserveDeficit is a free data retrieval call binding the contract method 0xc952485d.
//
// Solidity: function getReserveDeficit(address asset) view returns(uint256)
func (_Pool *PoolCallerSession) GetReserveDeficit(asset common.Address) (*big.Int, error) {
	return _Pool.Contract.GetReserveDeficit(&_Pool.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_Pool *PoolCaller) GetReserveNormalizedIncome(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getReserveNormalizedIncome", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_Pool *PoolSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _Pool.Contract.GetReserveNormalizedIncome(&_Pool.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_Pool *PoolCallerSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _Pool.Contract.GetReserveNormalizedIncome(&_Pool.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_Pool *PoolCaller) GetReserveNormalizedVariableDebt(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getReserveNormalizedVariableDebt", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_Pool *PoolSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _Pool.Contract.GetReserveNormalizedVariableDebt(&_Pool.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_Pool *PoolCallerSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _Pool.Contract.GetReserveNormalizedVariableDebt(&_Pool.CallOpts, asset)
}

// GetReserveVariableDebtToken is a free data retrieval call binding the contract method 0x365090a0.
//
// Solidity: function getReserveVariableDebtToken(address asset) view returns(address)
func (_Pool *PoolCaller) GetReserveVariableDebtToken(opts *bind.CallOpts, asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getReserveVariableDebtToken", asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReserveVariableDebtToken is a free data retrieval call binding the contract method 0x365090a0.
//
// Solidity: function getReserveVariableDebtToken(address asset) view returns(address)
func (_Pool *PoolSession) GetReserveVariableDebtToken(asset common.Address) (common.Address, error) {
	return _Pool.Contract.GetReserveVariableDebtToken(&_Pool.CallOpts, asset)
}

// GetReserveVariableDebtToken is a free data retrieval call binding the contract method 0x365090a0.
//
// Solidity: function getReserveVariableDebtToken(address asset) view returns(address)
func (_Pool *PoolCallerSession) GetReserveVariableDebtToken(asset common.Address) (common.Address, error) {
	return _Pool.Contract.GetReserveVariableDebtToken(&_Pool.CallOpts, asset)
}

// GetReservesCount is a free data retrieval call binding the contract method 0x72218d04.
//
// Solidity: function getReservesCount() view returns(uint256)
func (_Pool *PoolCaller) GetReservesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getReservesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReservesCount is a free data retrieval call binding the contract method 0x72218d04.
//
// Solidity: function getReservesCount() view returns(uint256)
func (_Pool *PoolSession) GetReservesCount() (*big.Int, error) {
	return _Pool.Contract.GetReservesCount(&_Pool.CallOpts)
}

// GetReservesCount is a free data retrieval call binding the contract method 0x72218d04.
//
// Solidity: function getReservesCount() view returns(uint256)
func (_Pool *PoolCallerSession) GetReservesCount() (*big.Int, error) {
	return _Pool.Contract.GetReservesCount(&_Pool.CallOpts)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_Pool *PoolCaller) GetReservesList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getReservesList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_Pool *PoolSession) GetReservesList() ([]common.Address, error) {
	return _Pool.Contract.GetReservesList(&_Pool.CallOpts)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_Pool *PoolCallerSession) GetReservesList() ([]common.Address, error) {
	return _Pool.Contract.GetReservesList(&_Pool.CallOpts)
}

// GetSupplyLogic is a free data retrieval call binding the contract method 0x870e7744.
//
// Solidity: function getSupplyLogic() pure returns(address)
func (_Pool *PoolCaller) GetSupplyLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getSupplyLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSupplyLogic is a free data retrieval call binding the contract method 0x870e7744.
//
// Solidity: function getSupplyLogic() pure returns(address)
func (_Pool *PoolSession) GetSupplyLogic() (common.Address, error) {
	return _Pool.Contract.GetSupplyLogic(&_Pool.CallOpts)
}

// GetSupplyLogic is a free data retrieval call binding the contract method 0x870e7744.
//
// Solidity: function getSupplyLogic() pure returns(address)
func (_Pool *PoolCallerSession) GetSupplyLogic() (common.Address, error) {
	return _Pool.Contract.GetSupplyLogic(&_Pool.CallOpts)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_Pool *PoolCaller) GetUserAccountData(opts *bind.CallOpts, user common.Address) (struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	AvailableBorrowsBase        *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getUserAccountData", user)

	outstruct := new(struct {
		TotalCollateralBase         *big.Int
		TotalDebtBase               *big.Int
		AvailableBorrowsBase        *big.Int
		CurrentLiquidationThreshold *big.Int
		Ltv                         *big.Int
		HealthFactor                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCollateralBase = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalDebtBase = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrowsBase = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentLiquidationThreshold = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HealthFactor = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_Pool *PoolSession) GetUserAccountData(user common.Address) (struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	AvailableBorrowsBase        *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _Pool.Contract.GetUserAccountData(&_Pool.CallOpts, user)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_Pool *PoolCallerSession) GetUserAccountData(user common.Address) (struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	AvailableBorrowsBase        *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _Pool.Contract.GetUserAccountData(&_Pool.CallOpts, user)
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_Pool *PoolCaller) GetUserConfiguration(opts *bind.CallOpts, user common.Address) (DataTypesUserConfigurationMap, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getUserConfiguration", user)

	if err != nil {
		return *new(DataTypesUserConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesUserConfigurationMap)).(*DataTypesUserConfigurationMap)

	return out0, err

}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_Pool *PoolSession) GetUserConfiguration(user common.Address) (DataTypesUserConfigurationMap, error) {
	return _Pool.Contract.GetUserConfiguration(&_Pool.CallOpts, user)
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_Pool *PoolCallerSession) GetUserConfiguration(user common.Address) (DataTypesUserConfigurationMap, error) {
	return _Pool.Contract.GetUserConfiguration(&_Pool.CallOpts, user)
}

// GetUserEMode is a free data retrieval call binding the contract method 0xeddf1b79.
//
// Solidity: function getUserEMode(address user) view returns(uint256)
func (_Pool *PoolCaller) GetUserEMode(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getUserEMode", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserEMode is a free data retrieval call binding the contract method 0xeddf1b79.
//
// Solidity: function getUserEMode(address user) view returns(uint256)
func (_Pool *PoolSession) GetUserEMode(user common.Address) (*big.Int, error) {
	return _Pool.Contract.GetUserEMode(&_Pool.CallOpts, user)
}

// GetUserEMode is a free data retrieval call binding the contract method 0xeddf1b79.
//
// Solidity: function getUserEMode(address user) view returns(uint256)
func (_Pool *PoolCallerSession) GetUserEMode(user common.Address) (*big.Int, error) {
	return _Pool.Contract.GetUserEMode(&_Pool.CallOpts, user)
}

// GetVirtualUnderlyingBalance is a free data retrieval call binding the contract method 0x6fb07f96.
//
// Solidity: function getVirtualUnderlyingBalance(address asset) view returns(uint128)
func (_Pool *PoolCaller) GetVirtualUnderlyingBalance(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getVirtualUnderlyingBalance", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualUnderlyingBalance is a free data retrieval call binding the contract method 0x6fb07f96.
//
// Solidity: function getVirtualUnderlyingBalance(address asset) view returns(uint128)
func (_Pool *PoolSession) GetVirtualUnderlyingBalance(asset common.Address) (*big.Int, error) {
	return _Pool.Contract.GetVirtualUnderlyingBalance(&_Pool.CallOpts, asset)
}

// GetVirtualUnderlyingBalance is a free data retrieval call binding the contract method 0x6fb07f96.
//
// Solidity: function getVirtualUnderlyingBalance(address asset) view returns(uint128)
func (_Pool *PoolCallerSession) GetVirtualUnderlyingBalance(asset common.Address) (*big.Int, error) {
	return _Pool.Contract.GetVirtualUnderlyingBalance(&_Pool.CallOpts, asset)
}

// IsApprovedPositionManager is a free data retrieval call binding the contract method 0xf9c2bd87.
//
// Solidity: function isApprovedPositionManager(address user, address positionManager) view returns(bool)
func (_Pool *PoolCaller) IsApprovedPositionManager(opts *bind.CallOpts, user common.Address, positionManager common.Address) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "isApprovedPositionManager", user, positionManager)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedPositionManager is a free data retrieval call binding the contract method 0xf9c2bd87.
//
// Solidity: function isApprovedPositionManager(address user, address positionManager) view returns(bool)
func (_Pool *PoolSession) IsApprovedPositionManager(user common.Address, positionManager common.Address) (bool, error) {
	return _Pool.Contract.IsApprovedPositionManager(&_Pool.CallOpts, user, positionManager)
}

// IsApprovedPositionManager is a free data retrieval call binding the contract method 0xf9c2bd87.
//
// Solidity: function isApprovedPositionManager(address user, address positionManager) view returns(bool)
func (_Pool *PoolCallerSession) IsApprovedPositionManager(user common.Address, positionManager common.Address) (bool, error) {
	return _Pool.Contract.IsApprovedPositionManager(&_Pool.CallOpts, user, positionManager)
}

// ApprovePositionManager is a paid mutator transaction binding the contract method 0xb8caa7c5.
//
// Solidity: function approvePositionManager(address positionManager, bool approve) returns()
func (_Pool *PoolTransactor) ApprovePositionManager(opts *bind.TransactOpts, positionManager common.Address, approve bool) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "approvePositionManager", positionManager, approve)
}

// ApprovePositionManager is a paid mutator transaction binding the contract method 0xb8caa7c5.
//
// Solidity: function approvePositionManager(address positionManager, bool approve) returns()
func (_Pool *PoolSession) ApprovePositionManager(positionManager common.Address, approve bool) (*types.Transaction, error) {
	return _Pool.Contract.ApprovePositionManager(&_Pool.TransactOpts, positionManager, approve)
}

// ApprovePositionManager is a paid mutator transaction binding the contract method 0xb8caa7c5.
//
// Solidity: function approvePositionManager(address positionManager, bool approve) returns()
func (_Pool *PoolTransactorSession) ApprovePositionManager(positionManager common.Address, approve bool) (*types.Transaction, error) {
	return _Pool.Contract.ApprovePositionManager(&_Pool.TransactOpts, positionManager, approve)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_Pool *PoolTransactor) Borrow(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "borrow", asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_Pool *PoolSession) Borrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Borrow(&_Pool.TransactOpts, asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_Pool *PoolTransactorSession) Borrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Borrow(&_Pool.TransactOpts, asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// ConfigureEModeCategory is a paid mutator transaction binding the contract method 0x7b75d7f4.
//
// Solidity: function configureEModeCategory(uint8 id, (uint16,uint16,uint16,string) category) returns()
func (_Pool *PoolTransactor) ConfigureEModeCategory(opts *bind.TransactOpts, id uint8, category DataTypesEModeCategoryBaseConfiguration) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "configureEModeCategory", id, category)
}

// ConfigureEModeCategory is a paid mutator transaction binding the contract method 0x7b75d7f4.
//
// Solidity: function configureEModeCategory(uint8 id, (uint16,uint16,uint16,string) category) returns()
func (_Pool *PoolSession) ConfigureEModeCategory(id uint8, category DataTypesEModeCategoryBaseConfiguration) (*types.Transaction, error) {
	return _Pool.Contract.ConfigureEModeCategory(&_Pool.TransactOpts, id, category)
}

// ConfigureEModeCategory is a paid mutator transaction binding the contract method 0x7b75d7f4.
//
// Solidity: function configureEModeCategory(uint8 id, (uint16,uint16,uint16,string) category) returns()
func (_Pool *PoolTransactorSession) ConfigureEModeCategory(id uint8, category DataTypesEModeCategoryBaseConfiguration) (*types.Transaction, error) {
	return _Pool.Contract.ConfigureEModeCategory(&_Pool.TransactOpts, id, category)
}

// ConfigureEModeCategoryBorrowableBitmap is a paid mutator transaction binding the contract method 0xff72158a.
//
// Solidity: function configureEModeCategoryBorrowableBitmap(uint8 id, uint128 borrowableBitmap) returns()
func (_Pool *PoolTransactor) ConfigureEModeCategoryBorrowableBitmap(opts *bind.TransactOpts, id uint8, borrowableBitmap *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "configureEModeCategoryBorrowableBitmap", id, borrowableBitmap)
}

// ConfigureEModeCategoryBorrowableBitmap is a paid mutator transaction binding the contract method 0xff72158a.
//
// Solidity: function configureEModeCategoryBorrowableBitmap(uint8 id, uint128 borrowableBitmap) returns()
func (_Pool *PoolSession) ConfigureEModeCategoryBorrowableBitmap(id uint8, borrowableBitmap *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ConfigureEModeCategoryBorrowableBitmap(&_Pool.TransactOpts, id, borrowableBitmap)
}

// ConfigureEModeCategoryBorrowableBitmap is a paid mutator transaction binding the contract method 0xff72158a.
//
// Solidity: function configureEModeCategoryBorrowableBitmap(uint8 id, uint128 borrowableBitmap) returns()
func (_Pool *PoolTransactorSession) ConfigureEModeCategoryBorrowableBitmap(id uint8, borrowableBitmap *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ConfigureEModeCategoryBorrowableBitmap(&_Pool.TransactOpts, id, borrowableBitmap)
}

// ConfigureEModeCategoryCollateralBitmap is a paid mutator transaction binding the contract method 0x92380ecb.
//
// Solidity: function configureEModeCategoryCollateralBitmap(uint8 id, uint128 collateralBitmap) returns()
func (_Pool *PoolTransactor) ConfigureEModeCategoryCollateralBitmap(opts *bind.TransactOpts, id uint8, collateralBitmap *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "configureEModeCategoryCollateralBitmap", id, collateralBitmap)
}

// ConfigureEModeCategoryCollateralBitmap is a paid mutator transaction binding the contract method 0x92380ecb.
//
// Solidity: function configureEModeCategoryCollateralBitmap(uint8 id, uint128 collateralBitmap) returns()
func (_Pool *PoolSession) ConfigureEModeCategoryCollateralBitmap(id uint8, collateralBitmap *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ConfigureEModeCategoryCollateralBitmap(&_Pool.TransactOpts, id, collateralBitmap)
}

// ConfigureEModeCategoryCollateralBitmap is a paid mutator transaction binding the contract method 0x92380ecb.
//
// Solidity: function configureEModeCategoryCollateralBitmap(uint8 id, uint128 collateralBitmap) returns()
func (_Pool *PoolTransactorSession) ConfigureEModeCategoryCollateralBitmap(id uint8, collateralBitmap *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ConfigureEModeCategoryCollateralBitmap(&_Pool.TransactOpts, id, collateralBitmap)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_Pool *PoolTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "deposit", asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_Pool *PoolSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Pool.Contract.Deposit(&_Pool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_Pool *PoolTransactorSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Pool.Contract.Deposit(&_Pool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// DropReserve is a paid mutator transaction binding the contract method 0x63c9b860.
//
// Solidity: function dropReserve(address asset) returns()
func (_Pool *PoolTransactor) DropReserve(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "dropReserve", asset)
}

// DropReserve is a paid mutator transaction binding the contract method 0x63c9b860.
//
// Solidity: function dropReserve(address asset) returns()
func (_Pool *PoolSession) DropReserve(asset common.Address) (*types.Transaction, error) {
	return _Pool.Contract.DropReserve(&_Pool.TransactOpts, asset)
}

// DropReserve is a paid mutator transaction binding the contract method 0x63c9b860.
//
// Solidity: function dropReserve(address asset) returns()
func (_Pool *PoolTransactorSession) DropReserve(asset common.Address) (*types.Transaction, error) {
	return _Pool.Contract.DropReserve(&_Pool.TransactOpts, asset)
}

// EliminateReserveDeficit is a paid mutator transaction binding the contract method 0xa1d2f3c4.
//
// Solidity: function eliminateReserveDeficit(address asset, uint256 amount) returns(uint256)
func (_Pool *PoolTransactor) EliminateReserveDeficit(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "eliminateReserveDeficit", asset, amount)
}

// EliminateReserveDeficit is a paid mutator transaction binding the contract method 0xa1d2f3c4.
//
// Solidity: function eliminateReserveDeficit(address asset, uint256 amount) returns(uint256)
func (_Pool *PoolSession) EliminateReserveDeficit(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.EliminateReserveDeficit(&_Pool.TransactOpts, asset, amount)
}

// EliminateReserveDeficit is a paid mutator transaction binding the contract method 0xa1d2f3c4.
//
// Solidity: function eliminateReserveDeficit(address asset, uint256 amount) returns(uint256)
func (_Pool *PoolTransactorSession) EliminateReserveDeficit(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.EliminateReserveDeficit(&_Pool.TransactOpts, asset, amount)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 scaledAmount, uint256 scaledBalanceFromBefore, uint256 scaledBalanceToBefore) returns()
func (_Pool *PoolTransactor) FinalizeTransfer(opts *bind.TransactOpts, asset common.Address, from common.Address, to common.Address, scaledAmount *big.Int, scaledBalanceFromBefore *big.Int, scaledBalanceToBefore *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "finalizeTransfer", asset, from, to, scaledAmount, scaledBalanceFromBefore, scaledBalanceToBefore)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 scaledAmount, uint256 scaledBalanceFromBefore, uint256 scaledBalanceToBefore) returns()
func (_Pool *PoolSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, scaledAmount *big.Int, scaledBalanceFromBefore *big.Int, scaledBalanceToBefore *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.FinalizeTransfer(&_Pool.TransactOpts, asset, from, to, scaledAmount, scaledBalanceFromBefore, scaledBalanceToBefore)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 scaledAmount, uint256 scaledBalanceFromBefore, uint256 scaledBalanceToBefore) returns()
func (_Pool *PoolTransactorSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, scaledAmount *big.Int, scaledBalanceFromBefore *big.Int, scaledBalanceToBefore *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.FinalizeTransfer(&_Pool.TransactOpts, asset, from, to, scaledAmount, scaledBalanceFromBefore, scaledBalanceToBefore)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] interestRateModes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_Pool *PoolTransactor) FlashLoan(opts *bind.TransactOpts, receiverAddress common.Address, assets []common.Address, amounts []*big.Int, interestRateModes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "flashLoan", receiverAddress, assets, amounts, interestRateModes, onBehalfOf, params, referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] interestRateModes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_Pool *PoolSession) FlashLoan(receiverAddress common.Address, assets []common.Address, amounts []*big.Int, interestRateModes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _Pool.Contract.FlashLoan(&_Pool.TransactOpts, receiverAddress, assets, amounts, interestRateModes, onBehalfOf, params, referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] interestRateModes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_Pool *PoolTransactorSession) FlashLoan(receiverAddress common.Address, assets []common.Address, amounts []*big.Int, interestRateModes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _Pool.Contract.FlashLoan(&_Pool.TransactOpts, receiverAddress, assets, amounts, interestRateModes, onBehalfOf, params, referralCode)
}

// FlashLoanSimple is a paid mutator transaction binding the contract method 0x42b0b77c.
//
// Solidity: function flashLoanSimple(address receiverAddress, address asset, uint256 amount, bytes params, uint16 referralCode) returns()
func (_Pool *PoolTransactor) FlashLoanSimple(opts *bind.TransactOpts, receiverAddress common.Address, asset common.Address, amount *big.Int, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "flashLoanSimple", receiverAddress, asset, amount, params, referralCode)
}

// FlashLoanSimple is a paid mutator transaction binding the contract method 0x42b0b77c.
//
// Solidity: function flashLoanSimple(address receiverAddress, address asset, uint256 amount, bytes params, uint16 referralCode) returns()
func (_Pool *PoolSession) FlashLoanSimple(receiverAddress common.Address, asset common.Address, amount *big.Int, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _Pool.Contract.FlashLoanSimple(&_Pool.TransactOpts, receiverAddress, asset, amount, params, referralCode)
}

// FlashLoanSimple is a paid mutator transaction binding the contract method 0x42b0b77c.
//
// Solidity: function flashLoanSimple(address receiverAddress, address asset, uint256 amount, bytes params, uint16 referralCode) returns()
func (_Pool *PoolTransactorSession) FlashLoanSimple(receiverAddress common.Address, asset common.Address, amount *big.Int, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _Pool.Contract.FlashLoanSimple(&_Pool.TransactOpts, receiverAddress, asset, amount, params, referralCode)
}

// InitReserve is a paid mutator transaction binding the contract method 0x932f12c8.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address variableDebtAddress) returns()
func (_Pool *PoolTransactor) InitReserve(opts *bind.TransactOpts, asset common.Address, aTokenAddress common.Address, variableDebtAddress common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "initReserve", asset, aTokenAddress, variableDebtAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x932f12c8.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address variableDebtAddress) returns()
func (_Pool *PoolSession) InitReserve(asset common.Address, aTokenAddress common.Address, variableDebtAddress common.Address) (*types.Transaction, error) {
	return _Pool.Contract.InitReserve(&_Pool.TransactOpts, asset, aTokenAddress, variableDebtAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x932f12c8.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address variableDebtAddress) returns()
func (_Pool *PoolTransactorSession) InitReserve(asset common.Address, aTokenAddress common.Address, variableDebtAddress common.Address) (*types.Transaction, error) {
	return _Pool.Contract.InitReserve(&_Pool.TransactOpts, asset, aTokenAddress, variableDebtAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_Pool *PoolTransactor) Initialize(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "initialize", provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_Pool *PoolSession) Initialize(provider common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts, provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_Pool *PoolTransactorSession) Initialize(provider common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts, provider)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address borrower, uint256 debtToCover, bool receiveAToken) returns()
func (_Pool *PoolTransactor) LiquidationCall(opts *bind.TransactOpts, collateralAsset common.Address, debtAsset common.Address, borrower common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "liquidationCall", collateralAsset, debtAsset, borrower, debtToCover, receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address borrower, uint256 debtToCover, bool receiveAToken) returns()
func (_Pool *PoolSession) LiquidationCall(collateralAsset common.Address, debtAsset common.Address, borrower common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _Pool.Contract.LiquidationCall(&_Pool.TransactOpts, collateralAsset, debtAsset, borrower, debtToCover, receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address borrower, uint256 debtToCover, bool receiveAToken) returns()
func (_Pool *PoolTransactorSession) LiquidationCall(collateralAsset common.Address, debtAsset common.Address, borrower common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _Pool.Contract.LiquidationCall(&_Pool.TransactOpts, collateralAsset, debtAsset, borrower, debtToCover, receiveAToken)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x9cd19996.
//
// Solidity: function mintToTreasury(address[] assets) returns()
func (_Pool *PoolTransactor) MintToTreasury(opts *bind.TransactOpts, assets []common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "mintToTreasury", assets)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x9cd19996.
//
// Solidity: function mintToTreasury(address[] assets) returns()
func (_Pool *PoolSession) MintToTreasury(assets []common.Address) (*types.Transaction, error) {
	return _Pool.Contract.MintToTreasury(&_Pool.TransactOpts, assets)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x9cd19996.
//
// Solidity: function mintToTreasury(address[] assets) returns()
func (_Pool *PoolTransactorSession) MintToTreasury(assets []common.Address) (*types.Transaction, error) {
	return _Pool.Contract.MintToTreasury(&_Pool.TransactOpts, assets)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Pool *PoolTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Pool *PoolSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Pool.Contract.Multicall(&_Pool.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Pool *PoolTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Pool.Contract.Multicall(&_Pool.TransactOpts, data)
}

// RenouncePositionManagerRole is a paid mutator transaction binding the contract method 0xfea149a6.
//
// Solidity: function renouncePositionManagerRole(address user) returns()
func (_Pool *PoolTransactor) RenouncePositionManagerRole(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "renouncePositionManagerRole", user)
}

// RenouncePositionManagerRole is a paid mutator transaction binding the contract method 0xfea149a6.
//
// Solidity: function renouncePositionManagerRole(address user) returns()
func (_Pool *PoolSession) RenouncePositionManagerRole(user common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RenouncePositionManagerRole(&_Pool.TransactOpts, user)
}

// RenouncePositionManagerRole is a paid mutator transaction binding the contract method 0xfea149a6.
//
// Solidity: function renouncePositionManagerRole(address user) returns()
func (_Pool *PoolTransactorSession) RenouncePositionManagerRole(user common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RenouncePositionManagerRole(&_Pool.TransactOpts, user)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (_Pool *PoolTransactor) Repay(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "repay", asset, amount, interestRateMode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (_Pool *PoolSession) Repay(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Repay(&_Pool.TransactOpts, asset, amount, interestRateMode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (_Pool *PoolTransactorSession) Repay(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Repay(&_Pool.TransactOpts, asset, amount, interestRateMode, onBehalfOf)
}

// RepayWithATokens is a paid mutator transaction binding the contract method 0x2dad97d4.
//
// Solidity: function repayWithATokens(address asset, uint256 amount, uint256 interestRateMode) returns(uint256)
func (_Pool *PoolTransactor) RepayWithATokens(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "repayWithATokens", asset, amount, interestRateMode)
}

// RepayWithATokens is a paid mutator transaction binding the contract method 0x2dad97d4.
//
// Solidity: function repayWithATokens(address asset, uint256 amount, uint256 interestRateMode) returns(uint256)
func (_Pool *PoolSession) RepayWithATokens(asset common.Address, amount *big.Int, interestRateMode *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RepayWithATokens(&_Pool.TransactOpts, asset, amount, interestRateMode)
}

// RepayWithATokens is a paid mutator transaction binding the contract method 0x2dad97d4.
//
// Solidity: function repayWithATokens(address asset, uint256 amount, uint256 interestRateMode) returns(uint256)
func (_Pool *PoolTransactorSession) RepayWithATokens(asset common.Address, amount *big.Int, interestRateMode *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RepayWithATokens(&_Pool.TransactOpts, asset, amount, interestRateMode)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xee3e210b.
//
// Solidity: function repayWithPermit(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns(uint256)
func (_Pool *PoolTransactor) RepayWithPermit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "repayWithPermit", asset, amount, interestRateMode, onBehalfOf, deadline, permitV, permitR, permitS)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xee3e210b.
//
// Solidity: function repayWithPermit(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns(uint256)
func (_Pool *PoolSession) RepayWithPermit(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _Pool.Contract.RepayWithPermit(&_Pool.TransactOpts, asset, amount, interestRateMode, onBehalfOf, deadline, permitV, permitR, permitS)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xee3e210b.
//
// Solidity: function repayWithPermit(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns(uint256)
func (_Pool *PoolTransactorSession) RepayWithPermit(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _Pool.Contract.RepayWithPermit(&_Pool.TransactOpts, asset, amount, interestRateMode, onBehalfOf, deadline, permitV, permitR, permitS)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_Pool *PoolTransactor) RescueTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "rescueTokens", token, to, amount)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_Pool *PoolSession) RescueTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RescueTokens(&_Pool.TransactOpts, token, to, amount)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_Pool *PoolTransactorSession) RescueTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RescueTokens(&_Pool.TransactOpts, token, to, amount)
}

// ResetIsolationModeTotalDebt is a paid mutator transaction binding the contract method 0xe43e88a1.
//
// Solidity: function resetIsolationModeTotalDebt(address asset) returns()
func (_Pool *PoolTransactor) ResetIsolationModeTotalDebt(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "resetIsolationModeTotalDebt", asset)
}

// ResetIsolationModeTotalDebt is a paid mutator transaction binding the contract method 0xe43e88a1.
//
// Solidity: function resetIsolationModeTotalDebt(address asset) returns()
func (_Pool *PoolSession) ResetIsolationModeTotalDebt(asset common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ResetIsolationModeTotalDebt(&_Pool.TransactOpts, asset)
}

// ResetIsolationModeTotalDebt is a paid mutator transaction binding the contract method 0xe43e88a1.
//
// Solidity: function resetIsolationModeTotalDebt(address asset) returns()
func (_Pool *PoolTransactorSession) ResetIsolationModeTotalDebt(asset common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ResetIsolationModeTotalDebt(&_Pool.TransactOpts, asset)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xf51e435b.
//
// Solidity: function setConfiguration(address asset, (uint256) configuration) returns()
func (_Pool *PoolTransactor) SetConfiguration(opts *bind.TransactOpts, asset common.Address, configuration DataTypesReserveConfigurationMap) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setConfiguration", asset, configuration)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xf51e435b.
//
// Solidity: function setConfiguration(address asset, (uint256) configuration) returns()
func (_Pool *PoolSession) SetConfiguration(asset common.Address, configuration DataTypesReserveConfigurationMap) (*types.Transaction, error) {
	return _Pool.Contract.SetConfiguration(&_Pool.TransactOpts, asset, configuration)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xf51e435b.
//
// Solidity: function setConfiguration(address asset, (uint256) configuration) returns()
func (_Pool *PoolTransactorSession) SetConfiguration(asset common.Address, configuration DataTypesReserveConfigurationMap) (*types.Transaction, error) {
	return _Pool.Contract.SetConfiguration(&_Pool.TransactOpts, asset, configuration)
}

// SetLiquidationGracePeriod is a paid mutator transaction binding the contract method 0xb1a99e26.
//
// Solidity: function setLiquidationGracePeriod(address asset, uint40 until) returns()
func (_Pool *PoolTransactor) SetLiquidationGracePeriod(opts *bind.TransactOpts, asset common.Address, until *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setLiquidationGracePeriod", asset, until)
}

// SetLiquidationGracePeriod is a paid mutator transaction binding the contract method 0xb1a99e26.
//
// Solidity: function setLiquidationGracePeriod(address asset, uint40 until) returns()
func (_Pool *PoolSession) SetLiquidationGracePeriod(asset common.Address, until *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetLiquidationGracePeriod(&_Pool.TransactOpts, asset, until)
}

// SetLiquidationGracePeriod is a paid mutator transaction binding the contract method 0xb1a99e26.
//
// Solidity: function setLiquidationGracePeriod(address asset, uint40 until) returns()
func (_Pool *PoolTransactorSession) SetLiquidationGracePeriod(asset common.Address, until *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetLiquidationGracePeriod(&_Pool.TransactOpts, asset, until)
}

// SetUserEMode is a paid mutator transaction binding the contract method 0x28530a47.
//
// Solidity: function setUserEMode(uint8 categoryId) returns()
func (_Pool *PoolTransactor) SetUserEMode(opts *bind.TransactOpts, categoryId uint8) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setUserEMode", categoryId)
}

// SetUserEMode is a paid mutator transaction binding the contract method 0x28530a47.
//
// Solidity: function setUserEMode(uint8 categoryId) returns()
func (_Pool *PoolSession) SetUserEMode(categoryId uint8) (*types.Transaction, error) {
	return _Pool.Contract.SetUserEMode(&_Pool.TransactOpts, categoryId)
}

// SetUserEMode is a paid mutator transaction binding the contract method 0x28530a47.
//
// Solidity: function setUserEMode(uint8 categoryId) returns()
func (_Pool *PoolTransactorSession) SetUserEMode(categoryId uint8) (*types.Transaction, error) {
	return _Pool.Contract.SetUserEMode(&_Pool.TransactOpts, categoryId)
}

// SetUserEModeOnBehalfOf is a paid mutator transaction binding the contract method 0x4ba06814.
//
// Solidity: function setUserEModeOnBehalfOf(uint8 categoryId, address onBehalfOf) returns()
func (_Pool *PoolTransactor) SetUserEModeOnBehalfOf(opts *bind.TransactOpts, categoryId uint8, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setUserEModeOnBehalfOf", categoryId, onBehalfOf)
}

// SetUserEModeOnBehalfOf is a paid mutator transaction binding the contract method 0x4ba06814.
//
// Solidity: function setUserEModeOnBehalfOf(uint8 categoryId, address onBehalfOf) returns()
func (_Pool *PoolSession) SetUserEModeOnBehalfOf(categoryId uint8, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetUserEModeOnBehalfOf(&_Pool.TransactOpts, categoryId, onBehalfOf)
}

// SetUserEModeOnBehalfOf is a paid mutator transaction binding the contract method 0x4ba06814.
//
// Solidity: function setUserEModeOnBehalfOf(uint8 categoryId, address onBehalfOf) returns()
func (_Pool *PoolTransactorSession) SetUserEModeOnBehalfOf(categoryId uint8, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetUserEModeOnBehalfOf(&_Pool.TransactOpts, categoryId, onBehalfOf)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_Pool *PoolTransactor) SetUserUseReserveAsCollateral(opts *bind.TransactOpts, asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setUserUseReserveAsCollateral", asset, useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_Pool *PoolSession) SetUserUseReserveAsCollateral(asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _Pool.Contract.SetUserUseReserveAsCollateral(&_Pool.TransactOpts, asset, useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_Pool *PoolTransactorSession) SetUserUseReserveAsCollateral(asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _Pool.Contract.SetUserUseReserveAsCollateral(&_Pool.TransactOpts, asset, useAsCollateral)
}

// SetUserUseReserveAsCollateralOnBehalfOf is a paid mutator transaction binding the contract method 0x972b35fa.
//
// Solidity: function setUserUseReserveAsCollateralOnBehalfOf(address asset, bool useAsCollateral, address onBehalfOf) returns()
func (_Pool *PoolTransactor) SetUserUseReserveAsCollateralOnBehalfOf(opts *bind.TransactOpts, asset common.Address, useAsCollateral bool, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setUserUseReserveAsCollateralOnBehalfOf", asset, useAsCollateral, onBehalfOf)
}

// SetUserUseReserveAsCollateralOnBehalfOf is a paid mutator transaction binding the contract method 0x972b35fa.
//
// Solidity: function setUserUseReserveAsCollateralOnBehalfOf(address asset, bool useAsCollateral, address onBehalfOf) returns()
func (_Pool *PoolSession) SetUserUseReserveAsCollateralOnBehalfOf(asset common.Address, useAsCollateral bool, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetUserUseReserveAsCollateralOnBehalfOf(&_Pool.TransactOpts, asset, useAsCollateral, onBehalfOf)
}

// SetUserUseReserveAsCollateralOnBehalfOf is a paid mutator transaction binding the contract method 0x972b35fa.
//
// Solidity: function setUserUseReserveAsCollateralOnBehalfOf(address asset, bool useAsCollateral, address onBehalfOf) returns()
func (_Pool *PoolTransactorSession) SetUserUseReserveAsCollateralOnBehalfOf(asset common.Address, useAsCollateral bool, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetUserUseReserveAsCollateralOnBehalfOf(&_Pool.TransactOpts, asset, useAsCollateral, onBehalfOf)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_Pool *PoolTransactor) Supply(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "supply", asset, amount, onBehalfOf, referralCode)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_Pool *PoolSession) Supply(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Pool.Contract.Supply(&_Pool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_Pool *PoolTransactorSession) Supply(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Pool.Contract.Supply(&_Pool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// SupplyWithPermit is a paid mutator transaction binding the contract method 0x02c205f0.
//
// Solidity: function supplyWithPermit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_Pool *PoolTransactor) SupplyWithPermit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "supplyWithPermit", asset, amount, onBehalfOf, referralCode, deadline, permitV, permitR, permitS)
}

// SupplyWithPermit is a paid mutator transaction binding the contract method 0x02c205f0.
//
// Solidity: function supplyWithPermit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_Pool *PoolSession) SupplyWithPermit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _Pool.Contract.SupplyWithPermit(&_Pool.TransactOpts, asset, amount, onBehalfOf, referralCode, deadline, permitV, permitR, permitS)
}

// SupplyWithPermit is a paid mutator transaction binding the contract method 0x02c205f0.
//
// Solidity: function supplyWithPermit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_Pool *PoolTransactorSession) SupplyWithPermit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _Pool.Contract.SupplyWithPermit(&_Pool.TransactOpts, asset, amount, onBehalfOf, referralCode, deadline, permitV, permitR, permitS)
}

// SyncIndexesState is a paid mutator transaction binding the contract method 0xab2b51f6.
//
// Solidity: function syncIndexesState(address asset) returns()
func (_Pool *PoolTransactor) SyncIndexesState(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "syncIndexesState", asset)
}

// SyncIndexesState is a paid mutator transaction binding the contract method 0xab2b51f6.
//
// Solidity: function syncIndexesState(address asset) returns()
func (_Pool *PoolSession) SyncIndexesState(asset common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SyncIndexesState(&_Pool.TransactOpts, asset)
}

// SyncIndexesState is a paid mutator transaction binding the contract method 0xab2b51f6.
//
// Solidity: function syncIndexesState(address asset) returns()
func (_Pool *PoolTransactorSession) SyncIndexesState(asset common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SyncIndexesState(&_Pool.TransactOpts, asset)
}

// SyncRatesState is a paid mutator transaction binding the contract method 0x98c7da4e.
//
// Solidity: function syncRatesState(address asset) returns()
func (_Pool *PoolTransactor) SyncRatesState(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "syncRatesState", asset)
}

// SyncRatesState is a paid mutator transaction binding the contract method 0x98c7da4e.
//
// Solidity: function syncRatesState(address asset) returns()
func (_Pool *PoolSession) SyncRatesState(asset common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SyncRatesState(&_Pool.TransactOpts, asset)
}

// SyncRatesState is a paid mutator transaction binding the contract method 0x98c7da4e.
//
// Solidity: function syncRatesState(address asset) returns()
func (_Pool *PoolTransactorSession) SyncRatesState(asset common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SyncRatesState(&_Pool.TransactOpts, asset)
}

// UpdateFlashloanPremium is a paid mutator transaction binding the contract method 0x9c1d5f00.
//
// Solidity: function updateFlashloanPremium(uint128 flashLoanPremium) returns()
func (_Pool *PoolTransactor) UpdateFlashloanPremium(opts *bind.TransactOpts, flashLoanPremium *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "updateFlashloanPremium", flashLoanPremium)
}

// UpdateFlashloanPremium is a paid mutator transaction binding the contract method 0x9c1d5f00.
//
// Solidity: function updateFlashloanPremium(uint128 flashLoanPremium) returns()
func (_Pool *PoolSession) UpdateFlashloanPremium(flashLoanPremium *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UpdateFlashloanPremium(&_Pool.TransactOpts, flashLoanPremium)
}

// UpdateFlashloanPremium is a paid mutator transaction binding the contract method 0x9c1d5f00.
//
// Solidity: function updateFlashloanPremium(uint128 flashLoanPremium) returns()
func (_Pool *PoolTransactorSession) UpdateFlashloanPremium(flashLoanPremium *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UpdateFlashloanPremium(&_Pool.TransactOpts, flashLoanPremium)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_Pool *PoolTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "withdraw", asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_Pool *PoolSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Withdraw(&_Pool.TransactOpts, asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_Pool *PoolTransactorSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Withdraw(&_Pool.TransactOpts, asset, amount, to)
}

// PoolBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Pool contract.
type PoolBorrowIterator struct {
	Event *PoolBorrow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBorrow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolBorrow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBorrow represents a Borrow event raised by the Pool contract.
type PoolBorrow struct {
	Reserve          common.Address
	User             common.Address
	OnBehalfOf       common.Address
	Amount           *big.Int
	InterestRateMode uint8
	BorrowRate       *big.Int
	ReferralCode     uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xb3d084820fb1a9decffb176436bd02558d15fac9b0ddfed8c465bc7359d7dce0.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint8 interestRateMode, uint256 borrowRate, uint16 indexed referralCode)
func (_Pool *PoolFilterer) FilterBorrow(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (*PoolBorrowIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &PoolBorrowIterator{contract: _Pool.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xb3d084820fb1a9decffb176436bd02558d15fac9b0ddfed8c465bc7359d7dce0.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint8 interestRateMode, uint256 borrowRate, uint16 indexed referralCode)
func (_Pool *PoolFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *PoolBorrow, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBorrow)
				if err := _Pool.contract.UnpackLog(event, "Borrow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBorrow is a log parse operation binding the contract event 0xb3d084820fb1a9decffb176436bd02558d15fac9b0ddfed8c465bc7359d7dce0.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint8 interestRateMode, uint256 borrowRate, uint16 indexed referralCode)
func (_Pool *PoolFilterer) ParseBorrow(log types.Log) (*PoolBorrow, error) {
	event := new(PoolBorrow)
	if err := _Pool.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolDeficitCoveredIterator is returned from FilterDeficitCovered and is used to iterate over the raw logs and unpacked data for DeficitCovered events raised by the Pool contract.
type PoolDeficitCoveredIterator struct {
	Event *PoolDeficitCovered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolDeficitCoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolDeficitCovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolDeficitCovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolDeficitCoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolDeficitCoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolDeficitCovered represents a DeficitCovered event raised by the Pool contract.
type PoolDeficitCovered struct {
	Reserve       common.Address
	Caller        common.Address
	AmountCovered *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeficitCovered is a free log retrieval operation binding the contract event 0x84b203e49f1a4b553088061534231969a68ad1c81be192205e96d23a206cb26a.
//
// Solidity: event DeficitCovered(address indexed reserve, address caller, uint256 amountCovered)
func (_Pool *PoolFilterer) FilterDeficitCovered(opts *bind.FilterOpts, reserve []common.Address) (*PoolDeficitCoveredIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "DeficitCovered", reserveRule)
	if err != nil {
		return nil, err
	}
	return &PoolDeficitCoveredIterator{contract: _Pool.contract, event: "DeficitCovered", logs: logs, sub: sub}, nil
}

// WatchDeficitCovered is a free log subscription operation binding the contract event 0x84b203e49f1a4b553088061534231969a68ad1c81be192205e96d23a206cb26a.
//
// Solidity: event DeficitCovered(address indexed reserve, address caller, uint256 amountCovered)
func (_Pool *PoolFilterer) WatchDeficitCovered(opts *bind.WatchOpts, sink chan<- *PoolDeficitCovered, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "DeficitCovered", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolDeficitCovered)
				if err := _Pool.contract.UnpackLog(event, "DeficitCovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeficitCovered is a log parse operation binding the contract event 0x84b203e49f1a4b553088061534231969a68ad1c81be192205e96d23a206cb26a.
//
// Solidity: event DeficitCovered(address indexed reserve, address caller, uint256 amountCovered)
func (_Pool *PoolFilterer) ParseDeficitCovered(log types.Log) (*PoolDeficitCovered, error) {
	event := new(PoolDeficitCovered)
	if err := _Pool.contract.UnpackLog(event, "DeficitCovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolDeficitCreatedIterator is returned from FilterDeficitCreated and is used to iterate over the raw logs and unpacked data for DeficitCreated events raised by the Pool contract.
type PoolDeficitCreatedIterator struct {
	Event *PoolDeficitCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolDeficitCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolDeficitCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolDeficitCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolDeficitCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolDeficitCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolDeficitCreated represents a DeficitCreated event raised by the Pool contract.
type PoolDeficitCreated struct {
	User          common.Address
	DebtAsset     common.Address
	AmountCreated *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeficitCreated is a free log retrieval operation binding the contract event 0x2bccfb3fad376d59d7accf970515eb77b2f27b082c90ed0fb15583dd5a942699.
//
// Solidity: event DeficitCreated(address indexed user, address indexed debtAsset, uint256 amountCreated)
func (_Pool *PoolFilterer) FilterDeficitCreated(opts *bind.FilterOpts, user []common.Address, debtAsset []common.Address) (*PoolDeficitCreatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "DeficitCreated", userRule, debtAssetRule)
	if err != nil {
		return nil, err
	}
	return &PoolDeficitCreatedIterator{contract: _Pool.contract, event: "DeficitCreated", logs: logs, sub: sub}, nil
}

// WatchDeficitCreated is a free log subscription operation binding the contract event 0x2bccfb3fad376d59d7accf970515eb77b2f27b082c90ed0fb15583dd5a942699.
//
// Solidity: event DeficitCreated(address indexed user, address indexed debtAsset, uint256 amountCreated)
func (_Pool *PoolFilterer) WatchDeficitCreated(opts *bind.WatchOpts, sink chan<- *PoolDeficitCreated, user []common.Address, debtAsset []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "DeficitCreated", userRule, debtAssetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolDeficitCreated)
				if err := _Pool.contract.UnpackLog(event, "DeficitCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeficitCreated is a log parse operation binding the contract event 0x2bccfb3fad376d59d7accf970515eb77b2f27b082c90ed0fb15583dd5a942699.
//
// Solidity: event DeficitCreated(address indexed user, address indexed debtAsset, uint256 amountCreated)
func (_Pool *PoolFilterer) ParseDeficitCreated(log types.Log) (*PoolDeficitCreated, error) {
	event := new(PoolDeficitCreated)
	if err := _Pool.contract.UnpackLog(event, "DeficitCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the Pool contract.
type PoolFlashLoanIterator struct {
	Event *PoolFlashLoan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFlashLoan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolFlashLoan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFlashLoan represents a FlashLoan event raised by the Pool contract.
type PoolFlashLoan struct {
	Target           common.Address
	Initiator        common.Address
	Asset            common.Address
	Amount           *big.Int
	InterestRateMode uint8
	Premium          *big.Int
	ReferralCode     uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0xefefaba5e921573100900a3ad9cf29f222d995fb3b6045797eaea7521bd8d6f0.
//
// Solidity: event FlashLoan(address indexed target, address initiator, address indexed asset, uint256 amount, uint8 interestRateMode, uint256 premium, uint16 indexed referralCode)
func (_Pool *PoolFilterer) FilterFlashLoan(opts *bind.FilterOpts, target []common.Address, asset []common.Address, referralCode []uint16) (*PoolFlashLoanIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "FlashLoan", targetRule, assetRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &PoolFlashLoanIterator{contract: _Pool.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0xefefaba5e921573100900a3ad9cf29f222d995fb3b6045797eaea7521bd8d6f0.
//
// Solidity: event FlashLoan(address indexed target, address initiator, address indexed asset, uint256 amount, uint8 interestRateMode, uint256 premium, uint16 indexed referralCode)
func (_Pool *PoolFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *PoolFlashLoan, target []common.Address, asset []common.Address, referralCode []uint16) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "FlashLoan", targetRule, assetRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFlashLoan)
				if err := _Pool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlashLoan is a log parse operation binding the contract event 0xefefaba5e921573100900a3ad9cf29f222d995fb3b6045797eaea7521bd8d6f0.
//
// Solidity: event FlashLoan(address indexed target, address initiator, address indexed asset, uint256 amount, uint8 interestRateMode, uint256 premium, uint16 indexed referralCode)
func (_Pool *PoolFilterer) ParseFlashLoan(log types.Log) (*PoolFlashLoan, error) {
	event := new(PoolFlashLoan)
	if err := _Pool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolIsolationModeTotalDebtUpdatedIterator is returned from FilterIsolationModeTotalDebtUpdated and is used to iterate over the raw logs and unpacked data for IsolationModeTotalDebtUpdated events raised by the Pool contract.
type PoolIsolationModeTotalDebtUpdatedIterator struct {
	Event *PoolIsolationModeTotalDebtUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolIsolationModeTotalDebtUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolIsolationModeTotalDebtUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolIsolationModeTotalDebtUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolIsolationModeTotalDebtUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolIsolationModeTotalDebtUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolIsolationModeTotalDebtUpdated represents a IsolationModeTotalDebtUpdated event raised by the Pool contract.
type PoolIsolationModeTotalDebtUpdated struct {
	Asset     common.Address
	TotalDebt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIsolationModeTotalDebtUpdated is a free log retrieval operation binding the contract event 0xaef84d3b40895fd58c561f3998000f0583abb992a52fbdc99ace8e8de4d676a5.
//
// Solidity: event IsolationModeTotalDebtUpdated(address indexed asset, uint256 totalDebt)
func (_Pool *PoolFilterer) FilterIsolationModeTotalDebtUpdated(opts *bind.FilterOpts, asset []common.Address) (*PoolIsolationModeTotalDebtUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "IsolationModeTotalDebtUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return &PoolIsolationModeTotalDebtUpdatedIterator{contract: _Pool.contract, event: "IsolationModeTotalDebtUpdated", logs: logs, sub: sub}, nil
}

// WatchIsolationModeTotalDebtUpdated is a free log subscription operation binding the contract event 0xaef84d3b40895fd58c561f3998000f0583abb992a52fbdc99ace8e8de4d676a5.
//
// Solidity: event IsolationModeTotalDebtUpdated(address indexed asset, uint256 totalDebt)
func (_Pool *PoolFilterer) WatchIsolationModeTotalDebtUpdated(opts *bind.WatchOpts, sink chan<- *PoolIsolationModeTotalDebtUpdated, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "IsolationModeTotalDebtUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolIsolationModeTotalDebtUpdated)
				if err := _Pool.contract.UnpackLog(event, "IsolationModeTotalDebtUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIsolationModeTotalDebtUpdated is a log parse operation binding the contract event 0xaef84d3b40895fd58c561f3998000f0583abb992a52fbdc99ace8e8de4d676a5.
//
// Solidity: event IsolationModeTotalDebtUpdated(address indexed asset, uint256 totalDebt)
func (_Pool *PoolFilterer) ParseIsolationModeTotalDebtUpdated(log types.Log) (*PoolIsolationModeTotalDebtUpdated, error) {
	event := new(PoolIsolationModeTotalDebtUpdated)
	if err := _Pool.contract.UnpackLog(event, "IsolationModeTotalDebtUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolLiquidationCallIterator is returned from FilterLiquidationCall and is used to iterate over the raw logs and unpacked data for LiquidationCall events raised by the Pool contract.
type PoolLiquidationCallIterator struct {
	Event *PoolLiquidationCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolLiquidationCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolLiquidationCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolLiquidationCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolLiquidationCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolLiquidationCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolLiquidationCall represents a LiquidationCall event raised by the Pool contract.
type PoolLiquidationCall struct {
	CollateralAsset            common.Address
	DebtAsset                  common.Address
	User                       common.Address
	DebtToCover                *big.Int
	LiquidatedCollateralAmount *big.Int
	Liquidator                 common.Address
	ReceiveAToken              bool
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterLiquidationCall is a free log retrieval operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_Pool *PoolFilterer) FilterLiquidationCall(opts *bind.FilterOpts, collateralAsset []common.Address, debtAsset []common.Address, user []common.Address) (*PoolLiquidationCallIterator, error) {

	var collateralAssetRule []interface{}
	for _, collateralAssetItem := range collateralAsset {
		collateralAssetRule = append(collateralAssetRule, collateralAssetItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "LiquidationCall", collateralAssetRule, debtAssetRule, userRule)
	if err != nil {
		return nil, err
	}
	return &PoolLiquidationCallIterator{contract: _Pool.contract, event: "LiquidationCall", logs: logs, sub: sub}, nil
}

// WatchLiquidationCall is a free log subscription operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_Pool *PoolFilterer) WatchLiquidationCall(opts *bind.WatchOpts, sink chan<- *PoolLiquidationCall, collateralAsset []common.Address, debtAsset []common.Address, user []common.Address) (event.Subscription, error) {

	var collateralAssetRule []interface{}
	for _, collateralAssetItem := range collateralAsset {
		collateralAssetRule = append(collateralAssetRule, collateralAssetItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "LiquidationCall", collateralAssetRule, debtAssetRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolLiquidationCall)
				if err := _Pool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidationCall is a log parse operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_Pool *PoolFilterer) ParseLiquidationCall(log types.Log) (*PoolLiquidationCall, error) {
	event := new(PoolLiquidationCall)
	if err := _Pool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolMintedToTreasuryIterator is returned from FilterMintedToTreasury and is used to iterate over the raw logs and unpacked data for MintedToTreasury events raised by the Pool contract.
type PoolMintedToTreasuryIterator struct {
	Event *PoolMintedToTreasury // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolMintedToTreasuryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMintedToTreasury)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolMintedToTreasury)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolMintedToTreasuryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMintedToTreasuryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMintedToTreasury represents a MintedToTreasury event raised by the Pool contract.
type PoolMintedToTreasury struct {
	Reserve      common.Address
	AmountMinted *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintedToTreasury is a free log retrieval operation binding the contract event 0xbfa21aa5d5f9a1f0120a95e7c0749f389863cbdbfff531aa7339077a5bc919de.
//
// Solidity: event MintedToTreasury(address indexed reserve, uint256 amountMinted)
func (_Pool *PoolFilterer) FilterMintedToTreasury(opts *bind.FilterOpts, reserve []common.Address) (*PoolMintedToTreasuryIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "MintedToTreasury", reserveRule)
	if err != nil {
		return nil, err
	}
	return &PoolMintedToTreasuryIterator{contract: _Pool.contract, event: "MintedToTreasury", logs: logs, sub: sub}, nil
}

// WatchMintedToTreasury is a free log subscription operation binding the contract event 0xbfa21aa5d5f9a1f0120a95e7c0749f389863cbdbfff531aa7339077a5bc919de.
//
// Solidity: event MintedToTreasury(address indexed reserve, uint256 amountMinted)
func (_Pool *PoolFilterer) WatchMintedToTreasury(opts *bind.WatchOpts, sink chan<- *PoolMintedToTreasury, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "MintedToTreasury", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMintedToTreasury)
				if err := _Pool.contract.UnpackLog(event, "MintedToTreasury", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintedToTreasury is a log parse operation binding the contract event 0xbfa21aa5d5f9a1f0120a95e7c0749f389863cbdbfff531aa7339077a5bc919de.
//
// Solidity: event MintedToTreasury(address indexed reserve, uint256 amountMinted)
func (_Pool *PoolFilterer) ParseMintedToTreasury(log types.Log) (*PoolMintedToTreasury, error) {
	event := new(PoolMintedToTreasury)
	if err := _Pool.contract.UnpackLog(event, "MintedToTreasury", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolPositionManagerApprovedIterator is returned from FilterPositionManagerApproved and is used to iterate over the raw logs and unpacked data for PositionManagerApproved events raised by the Pool contract.
type PoolPositionManagerApprovedIterator struct {
	Event *PoolPositionManagerApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolPositionManagerApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolPositionManagerApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolPositionManagerApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolPositionManagerApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolPositionManagerApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolPositionManagerApproved represents a PositionManagerApproved event raised by the Pool contract.
type PoolPositionManagerApproved struct {
	User            common.Address
	PositionManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPositionManagerApproved is a free log retrieval operation binding the contract event 0x540e692f36c2fa13e7583c4deeffd91ce6bc04f91e7d84f295d9d858372875fc.
//
// Solidity: event PositionManagerApproved(address indexed user, address indexed positionManager)
func (_Pool *PoolFilterer) FilterPositionManagerApproved(opts *bind.FilterOpts, user []common.Address, positionManager []common.Address) (*PoolPositionManagerApprovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var positionManagerRule []interface{}
	for _, positionManagerItem := range positionManager {
		positionManagerRule = append(positionManagerRule, positionManagerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "PositionManagerApproved", userRule, positionManagerRule)
	if err != nil {
		return nil, err
	}
	return &PoolPositionManagerApprovedIterator{contract: _Pool.contract, event: "PositionManagerApproved", logs: logs, sub: sub}, nil
}

// WatchPositionManagerApproved is a free log subscription operation binding the contract event 0x540e692f36c2fa13e7583c4deeffd91ce6bc04f91e7d84f295d9d858372875fc.
//
// Solidity: event PositionManagerApproved(address indexed user, address indexed positionManager)
func (_Pool *PoolFilterer) WatchPositionManagerApproved(opts *bind.WatchOpts, sink chan<- *PoolPositionManagerApproved, user []common.Address, positionManager []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var positionManagerRule []interface{}
	for _, positionManagerItem := range positionManager {
		positionManagerRule = append(positionManagerRule, positionManagerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "PositionManagerApproved", userRule, positionManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolPositionManagerApproved)
				if err := _Pool.contract.UnpackLog(event, "PositionManagerApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePositionManagerApproved is a log parse operation binding the contract event 0x540e692f36c2fa13e7583c4deeffd91ce6bc04f91e7d84f295d9d858372875fc.
//
// Solidity: event PositionManagerApproved(address indexed user, address indexed positionManager)
func (_Pool *PoolFilterer) ParsePositionManagerApproved(log types.Log) (*PoolPositionManagerApproved, error) {
	event := new(PoolPositionManagerApproved)
	if err := _Pool.contract.UnpackLog(event, "PositionManagerApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolPositionManagerRevokedIterator is returned from FilterPositionManagerRevoked and is used to iterate over the raw logs and unpacked data for PositionManagerRevoked events raised by the Pool contract.
type PoolPositionManagerRevokedIterator struct {
	Event *PoolPositionManagerRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolPositionManagerRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolPositionManagerRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolPositionManagerRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolPositionManagerRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolPositionManagerRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolPositionManagerRevoked represents a PositionManagerRevoked event raised by the Pool contract.
type PoolPositionManagerRevoked struct {
	User            common.Address
	PositionManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPositionManagerRevoked is a free log retrieval operation binding the contract event 0x08c92c3870d10c79e9673fecea8f4ff261f8e6b661067d9ca63fd777882bff15.
//
// Solidity: event PositionManagerRevoked(address indexed user, address indexed positionManager)
func (_Pool *PoolFilterer) FilterPositionManagerRevoked(opts *bind.FilterOpts, user []common.Address, positionManager []common.Address) (*PoolPositionManagerRevokedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var positionManagerRule []interface{}
	for _, positionManagerItem := range positionManager {
		positionManagerRule = append(positionManagerRule, positionManagerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "PositionManagerRevoked", userRule, positionManagerRule)
	if err != nil {
		return nil, err
	}
	return &PoolPositionManagerRevokedIterator{contract: _Pool.contract, event: "PositionManagerRevoked", logs: logs, sub: sub}, nil
}

// WatchPositionManagerRevoked is a free log subscription operation binding the contract event 0x08c92c3870d10c79e9673fecea8f4ff261f8e6b661067d9ca63fd777882bff15.
//
// Solidity: event PositionManagerRevoked(address indexed user, address indexed positionManager)
func (_Pool *PoolFilterer) WatchPositionManagerRevoked(opts *bind.WatchOpts, sink chan<- *PoolPositionManagerRevoked, user []common.Address, positionManager []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var positionManagerRule []interface{}
	for _, positionManagerItem := range positionManager {
		positionManagerRule = append(positionManagerRule, positionManagerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "PositionManagerRevoked", userRule, positionManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolPositionManagerRevoked)
				if err := _Pool.contract.UnpackLog(event, "PositionManagerRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePositionManagerRevoked is a log parse operation binding the contract event 0x08c92c3870d10c79e9673fecea8f4ff261f8e6b661067d9ca63fd777882bff15.
//
// Solidity: event PositionManagerRevoked(address indexed user, address indexed positionManager)
func (_Pool *PoolFilterer) ParsePositionManagerRevoked(log types.Log) (*PoolPositionManagerRevoked, error) {
	event := new(PoolPositionManagerRevoked)
	if err := _Pool.contract.UnpackLog(event, "PositionManagerRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the Pool contract.
type PoolRepayIterator struct {
	Event *PoolRepay // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRepay)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolRepay)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRepay represents a Repay event raised by the Pool contract.
type PoolRepay struct {
	Reserve    common.Address
	User       common.Address
	Repayer    common.Address
	Amount     *big.Int
	UseATokens bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0xa534c8dbe71f871f9f3530e97a74601fea17b426cae02e1c5aee42c96c784051.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount, bool useATokens)
func (_Pool *PoolFilterer) FilterRepay(opts *bind.FilterOpts, reserve []common.Address, user []common.Address, repayer []common.Address) (*PoolRepayIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var repayerRule []interface{}
	for _, repayerItem := range repayer {
		repayerRule = append(repayerRule, repayerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Repay", reserveRule, userRule, repayerRule)
	if err != nil {
		return nil, err
	}
	return &PoolRepayIterator{contract: _Pool.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0xa534c8dbe71f871f9f3530e97a74601fea17b426cae02e1c5aee42c96c784051.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount, bool useATokens)
func (_Pool *PoolFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *PoolRepay, reserve []common.Address, user []common.Address, repayer []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var repayerRule []interface{}
	for _, repayerItem := range repayer {
		repayerRule = append(repayerRule, repayerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Repay", reserveRule, userRule, repayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRepay)
				if err := _Pool.contract.UnpackLog(event, "Repay", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRepay is a log parse operation binding the contract event 0xa534c8dbe71f871f9f3530e97a74601fea17b426cae02e1c5aee42c96c784051.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount, bool useATokens)
func (_Pool *PoolFilterer) ParseRepay(log types.Log) (*PoolRepay, error) {
	event := new(PoolRepay)
	if err := _Pool.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolReserveDataUpdatedIterator is returned from FilterReserveDataUpdated and is used to iterate over the raw logs and unpacked data for ReserveDataUpdated events raised by the Pool contract.
type PoolReserveDataUpdatedIterator struct {
	Event *PoolReserveDataUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolReserveDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolReserveDataUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolReserveDataUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolReserveDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolReserveDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolReserveDataUpdated represents a ReserveDataUpdated event raised by the Pool contract.
type PoolReserveDataUpdated struct {
	Reserve             common.Address
	LiquidityRate       *big.Int
	StableBorrowRate    *big.Int
	VariableBorrowRate  *big.Int
	LiquidityIndex      *big.Int
	VariableBorrowIndex *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReserveDataUpdated is a free log retrieval operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_Pool *PoolFilterer) FilterReserveDataUpdated(opts *bind.FilterOpts, reserve []common.Address) (*PoolReserveDataUpdatedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return &PoolReserveDataUpdatedIterator{contract: _Pool.contract, event: "ReserveDataUpdated", logs: logs, sub: sub}, nil
}

// WatchReserveDataUpdated is a free log subscription operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_Pool *PoolFilterer) WatchReserveDataUpdated(opts *bind.WatchOpts, sink chan<- *PoolReserveDataUpdated, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolReserveDataUpdated)
				if err := _Pool.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReserveDataUpdated is a log parse operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_Pool *PoolFilterer) ParseReserveDataUpdated(log types.Log) (*PoolReserveDataUpdated, error) {
	event := new(PoolReserveDataUpdated)
	if err := _Pool.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolReserveUsedAsCollateralDisabledIterator is returned from FilterReserveUsedAsCollateralDisabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralDisabled events raised by the Pool contract.
type PoolReserveUsedAsCollateralDisabledIterator struct {
	Event *PoolReserveUsedAsCollateralDisabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolReserveUsedAsCollateralDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolReserveUsedAsCollateralDisabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolReserveUsedAsCollateralDisabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolReserveUsedAsCollateralDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolReserveUsedAsCollateralDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolReserveUsedAsCollateralDisabled represents a ReserveUsedAsCollateralDisabled event raised by the Pool contract.
type PoolReserveUsedAsCollateralDisabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralDisabled is a free log retrieval operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_Pool *PoolFilterer) FilterReserveUsedAsCollateralDisabled(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*PoolReserveUsedAsCollateralDisabledIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ReserveUsedAsCollateralDisabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &PoolReserveUsedAsCollateralDisabledIterator{contract: _Pool.contract, event: "ReserveUsedAsCollateralDisabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralDisabled is a free log subscription operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_Pool *PoolFilterer) WatchReserveUsedAsCollateralDisabled(opts *bind.WatchOpts, sink chan<- *PoolReserveUsedAsCollateralDisabled, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ReserveUsedAsCollateralDisabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolReserveUsedAsCollateralDisabled)
				if err := _Pool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReserveUsedAsCollateralDisabled is a log parse operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_Pool *PoolFilterer) ParseReserveUsedAsCollateralDisabled(log types.Log) (*PoolReserveUsedAsCollateralDisabled, error) {
	event := new(PoolReserveUsedAsCollateralDisabled)
	if err := _Pool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolReserveUsedAsCollateralEnabledIterator is returned from FilterReserveUsedAsCollateralEnabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralEnabled events raised by the Pool contract.
type PoolReserveUsedAsCollateralEnabledIterator struct {
	Event *PoolReserveUsedAsCollateralEnabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolReserveUsedAsCollateralEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolReserveUsedAsCollateralEnabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolReserveUsedAsCollateralEnabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolReserveUsedAsCollateralEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolReserveUsedAsCollateralEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolReserveUsedAsCollateralEnabled represents a ReserveUsedAsCollateralEnabled event raised by the Pool contract.
type PoolReserveUsedAsCollateralEnabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralEnabled is a free log retrieval operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_Pool *PoolFilterer) FilterReserveUsedAsCollateralEnabled(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*PoolReserveUsedAsCollateralEnabledIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ReserveUsedAsCollateralEnabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &PoolReserveUsedAsCollateralEnabledIterator{contract: _Pool.contract, event: "ReserveUsedAsCollateralEnabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralEnabled is a free log subscription operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_Pool *PoolFilterer) WatchReserveUsedAsCollateralEnabled(opts *bind.WatchOpts, sink chan<- *PoolReserveUsedAsCollateralEnabled, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ReserveUsedAsCollateralEnabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolReserveUsedAsCollateralEnabled)
				if err := _Pool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReserveUsedAsCollateralEnabled is a log parse operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_Pool *PoolFilterer) ParseReserveUsedAsCollateralEnabled(log types.Log) (*PoolReserveUsedAsCollateralEnabled, error) {
	event := new(PoolReserveUsedAsCollateralEnabled)
	if err := _Pool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the Pool contract.
type PoolSupplyIterator struct {
	Event *PoolSupply // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSupply)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolSupply)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSupply represents a Supply event raised by the Pool contract.
type PoolSupply struct {
	Reserve      common.Address
	User         common.Address
	OnBehalfOf   common.Address
	Amount       *big.Int
	ReferralCode uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61.
//
// Solidity: event Supply(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_Pool *PoolFilterer) FilterSupply(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (*PoolSupplyIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Supply", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &PoolSupplyIterator{contract: _Pool.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61.
//
// Solidity: event Supply(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_Pool *PoolFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *PoolSupply, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Supply", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSupply)
				if err := _Pool.contract.UnpackLog(event, "Supply", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSupply is a log parse operation binding the contract event 0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61.
//
// Solidity: event Supply(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_Pool *PoolFilterer) ParseSupply(log types.Log) (*PoolSupply, error) {
	event := new(PoolSupply)
	if err := _Pool.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUserEModeSetIterator is returned from FilterUserEModeSet and is used to iterate over the raw logs and unpacked data for UserEModeSet events raised by the Pool contract.
type PoolUserEModeSetIterator struct {
	Event *PoolUserEModeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolUserEModeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUserEModeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolUserEModeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolUserEModeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUserEModeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUserEModeSet represents a UserEModeSet event raised by the Pool contract.
type PoolUserEModeSet struct {
	User       common.Address
	CategoryId uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUserEModeSet is a free log retrieval operation binding the contract event 0xd728da875fc88944cbf17638bcbe4af0eedaef63becd1d1c57cc097eb4608d84.
//
// Solidity: event UserEModeSet(address indexed user, uint8 categoryId)
func (_Pool *PoolFilterer) FilterUserEModeSet(opts *bind.FilterOpts, user []common.Address) (*PoolUserEModeSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "UserEModeSet", userRule)
	if err != nil {
		return nil, err
	}
	return &PoolUserEModeSetIterator{contract: _Pool.contract, event: "UserEModeSet", logs: logs, sub: sub}, nil
}

// WatchUserEModeSet is a free log subscription operation binding the contract event 0xd728da875fc88944cbf17638bcbe4af0eedaef63becd1d1c57cc097eb4608d84.
//
// Solidity: event UserEModeSet(address indexed user, uint8 categoryId)
func (_Pool *PoolFilterer) WatchUserEModeSet(opts *bind.WatchOpts, sink chan<- *PoolUserEModeSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "UserEModeSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUserEModeSet)
				if err := _Pool.contract.UnpackLog(event, "UserEModeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserEModeSet is a log parse operation binding the contract event 0xd728da875fc88944cbf17638bcbe4af0eedaef63becd1d1c57cc097eb4608d84.
//
// Solidity: event UserEModeSet(address indexed user, uint8 categoryId)
func (_Pool *PoolFilterer) ParseUserEModeSet(log types.Log) (*PoolUserEModeSet, error) {
	event := new(PoolUserEModeSet)
	if err := _Pool.contract.UnpackLog(event, "UserEModeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Pool contract.
type PoolWithdrawIterator struct {
	Event *PoolWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolWithdraw represents a Withdraw event raised by the Pool contract.
type PoolWithdraw struct {
	Reserve common.Address
	User    common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_Pool *PoolFilterer) FilterWithdraw(opts *bind.FilterOpts, reserve []common.Address, user []common.Address, to []common.Address) (*PoolWithdrawIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Withdraw", reserveRule, userRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolWithdrawIterator{contract: _Pool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_Pool *PoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PoolWithdraw, reserve []common.Address, user []common.Address, to []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Withdraw", reserveRule, userRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolWithdraw)
				if err := _Pool.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_Pool *PoolFilterer) ParseWithdraw(log types.Log) (*PoolWithdraw, error) {
	event := new(PoolWithdraw)
	if err := _Pool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
