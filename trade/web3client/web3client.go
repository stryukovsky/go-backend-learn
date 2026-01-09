package web3client

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stryukovsky/go-backend-learn/trade"
)

type ClientWithURL struct {
	Client *ethclient.Client
	Url    string
}

func (c *ClientWithURL) URL() string { return c.Url }

type MultiURLClient struct {
	clients []*ClientWithURL
}

func NewMultiURLClient(urls []string) (*MultiURLClient, error) {
	clients := make([]*ClientWithURL, 0, len(urls))
	for _, url := range urls {
		client, err := ethclient.Dial(url)
		if err != nil {
			slog.Warn(fmt.Sprintf("Failed to connect to JSON RPC %s: %v", url, err))
		} else {
			clients = append(clients, &ClientWithURL{Client: client, Url: url})
		}
	}
	if len(clients) == 0 {
		return nil, fmt.Errorf("No Ethereum JSONRPC URL is reachable")
	}
	return &MultiURLClient{clients}, nil
}

func (c *MultiURLClient) ChainID() (*big.Int, error) {
	return trade.RetryEthCall(
		func() []*ClientWithURL { return c.clients },
		func(client *ClientWithURL) (*big.Int, error) { return client.Client.ChainID(context.Background()) })
}

func (c *MultiURLClient) BlockNumber() (uint64, error) {
	return trade.RetryEthCall(
		func() []*ClientWithURL { return c.clients },
		func(client *ClientWithURL) (uint64, error) { return client.Client.BlockNumber(context.Background()) })
}

func (c *MultiURLClient) RandomClient() *ClientWithURL {
	return trade.RandomChoice(c.clients)
}

func (c *MultiURLClient) Length() int {
	return len(c.clients)
}

func (c *MultiURLClient) Iter() []*ClientWithURL {
	return c.clients
}
