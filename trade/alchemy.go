package trade
//
// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log/slog"
// 	"net/http"
// 	"strings"
//
// 	"github.com/chenzhijie/go-web3"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/redis/go-redis/v9"
// )
//
// type AssetTransferParams struct {
// 	FromBlock   string `json:"fromBlock"`
// 	ToBlock     string `json:"toBlock"`
// 	FromAddress string `json:"fromAddress"`
// 	ToAddress   string `json:"toAddress"`
// }
//
// type RawContract struct {
// 	Value   string `json:"value"`
// 	Address string `json:"address"`
// 	Decimal string `json:"decimal"`
// }
//
// type AssetTransfer struct {
// 	Category    string       `json:"category"`
// 	Hash        string       `json:"hash"`
// 	BlockNum    string       `json:"blockNum"`
// 	From        string       `json:"from"`
// 	To          string       `json:"to"`
// 	Value       float64      `json:"value"`
// 	RawContract *RawContract `json:"rawContract,omitempty"`
// }
//
// type Response struct {
// 	PageKey   string          `json:"pageKey"`
// 	Transfers []AssetTransfer `json:"transfers"`
// }
//
// type Request struct {
// 	JsonRPC string `json:"jsonrpc"`
// 	Method  string `json:"method"`
// 	Params  []any  `json:"params"`
// 	Id      int    `json:"id"`
// }
//
// var EmptyArr []ERC20Transfer = make([]ERC20Transfer, 0)
//
// func AlchemyGetTransfersForAccount(w3 *web3.Web3, cache *redis.Client, worker Worker, wallet TrackedWallet) ([]ERC20Transfer, error) {
// 	if strings.TrimSpace(worker.AlchemyApiUrl) == "" {
// 		return EmptyArr, fmt.Errorf("Alchemy API URL is not set ")
//
// 	}
// 	slog.Info(fmt.Sprintf("Out-dated wallet %s", wallet.Address))
//
// 	requestBody := Request{
// 		JsonRPC: "2.0",
// 		Id:      1,
// 		Method:  "alchemy_getAssetTransfers",
// 		Params: []any{
// 			"0x0",
// 			"latest",
// 			wallet.Address,
// 			"",
// 			true,
// 			[]string{"erc20"},
// 		},
// 	}
//
// 	requestBodyJSON, err := json.Marshal(requestBody)
// 	if err != nil {
// 		return EmptyArr, err
// 	}
// 	req, err := http.NewRequest(http.MethodPost, worker.AlchemyApiUrl, bytes.NewBuffer(requestBodyJSON))
// 	if err != nil {
// 		return EmptyArr, err
// 	}
// 	res, err := http.Post(req.URL.String(), "application/json", req.Body)
// 	if err != nil {
// 		return EmptyArr, err
// 	}
//
// 	responseRaw, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return EmptyArr, err
// 	}
// 	response := Response{}
// 	err = json.Unmarshal(responseRaw, &response)
// 	if err != nil {
// 		return EmptyArr, err
// 	}
// 	slog.Info(fmt.Sprintf("Found %d transfers via alchemy", len(response.Transfers)))
//
// 	result := make([]ERC20Transfer, len(response.Transfers))
// 	for i, transfer := range response.Transfers {
// 		if transfer.RawContract == nil {
// 			continue
// 		}
// 		blockNumber := common.HexToHash(transfer.BlockNum).Big()
// 		timestamp, err := GetCachedBlockTimestamp(w3, cache, blockNumber.Uint64())
// 		if err != nil {
// 			continue
// 		}
// 		result[i] = ERC20Transfer{
// 			TokenAddress: transfer.RawContract.Address,
// 			Name:         "unknown",
// 			Symbol:       "unknown",
// 			Decimals:     DBInt{common.HexToHash(transfer.RawContract.Decimal).Big()},
//
// 			Sender:    transfer.From,
// 			Recipient: transfer.To,
//
// 			Amount:    DBInt{common.HexToHash(transfer.RawContract.Value).Big()},
// 			Block:     DBInt{common.HexToHash(transfer.BlockNum).Big()},
// 			Timestamp: *timestamp,
// 			TxId:      transfer.Hash,
// 		}
//
// 	}
// 	return result, nil
//
// }
