package protocols

type DeFiProtocolHandler[BlockchainInteraction any, FinanceInteraction any] interface {
	FetchBlockchainInteractions(
		chainId string,
		participants []string,
		fromBlock uint64,
		toBlock uint64,
	) ([]BlockchainInteraction, error)
	PopulateWithFinanceInfo(interactions []BlockchainInteraction) ([]FinanceInteraction, error)
	Name() string
}
