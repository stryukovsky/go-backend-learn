package trade

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strconv"

	"github.com/joomcode/errorx"
)

var BinanceAddress string = "https://api.binance.com"
var QuoteEndpoint = "/api/v3/klines"

func GetQuoteId(tokenTicker string, baseTicker string) string {
	return tokenTicker + baseTicker

}

func CreateDeal(transfer ERC20Transfer) (*Deal, error) {
	params := url.Values{}
	params.Add("symbol", GetQuoteId(transfer.Symbol, "USDT"))
	params.Add("interval", "1m")
	params.Add("startTime", strconv.FormatInt(transfer.Timestamp.Unix(), 10))
	params.Add("limit", "1")
	baseUrl, err := url.Parse(BinanceAddress + QuoteEndpoint + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	response, err := http.Get(baseUrl.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errorx.IllegalState.New("Cannot fetch data from binance ")
	}

	var quote [][]any
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, err
	}
	closePrice := big.NewRat(1, 1)
	closePrice, success := closePrice.SetString(quote[0][4].(string))
	if !success {
		return nil, errorx.IllegalState.New("Bad close price for quote ")
	}

	volumeToken := big.NewRat(1, 1)
	volumeToken = volumeToken.SetFrac(transfer.Amount.Int, new(big.Int).Exp(big.NewInt(10), transfer.Decimals.Int, nil))

	volumeUSD := new(big.Rat).Mul(volumeToken, closePrice)
	return &Deal{
		Price:              DBNumeric{closePrice},
		VolumeUSD:          DBNumeric{volumeUSD},
		VolumeTokens:       DBNumeric{volumeToken},
		BlockchainTransfer: transfer,
	}, nil
}
