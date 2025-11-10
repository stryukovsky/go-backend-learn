package binance

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math/big"
	"net/http"
	"net/url"
	"runtime/debug"
	"time"
)

var BinanceAddress string = "https://api.binance.com"
var QuoteEndpoint = "/api/v3/klines"

func GetQuoteId(tokenTicker string, baseTicker string) string {
	return tokenTicker + baseTicker
}

var (
	BinanceFetchFailed error = errors.New("Response received, but it was not 200 OK")
	MalformedPrice     error = errors.New("Malformed price string value")
)

func GetClosePrice(symbol string, instant *time.Time) (*big.Rat, error) {
	if symbol == "USDT" {
		return big.NewRat(1, 1), nil
	}
	params := url.Values{}
	params.Add("symbol", GetQuoteId(symbol, "USDT"))
	params.Add("interval", "1m")
	params.Add("startTime", fmt.Sprintf("%d", instant.UnixMilli()))
	params.Add("limit", "1")
	url, err := url.Parse(BinanceAddress + QuoteEndpoint + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	urlString := url.String()
	response, err := http.Get(urlString)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		slog.Warn(fmt.Sprintf("Cannot fetch quote from binance: %d %s", response.StatusCode, string(body)))
		slog.Warn(fmt.Sprintf("Request was GET %s", urlString))
		debug.PrintStack()
		return nil, BinanceFetchFailed
	}

	var quote [][]any
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, err
	}
	closePrice := big.NewRat(1, 1)
	closePrice, success := closePrice.SetString(quote[0][4].(string))
	if !success {
		return nil, MalformedPrice
	}
	return closePrice, nil
}
