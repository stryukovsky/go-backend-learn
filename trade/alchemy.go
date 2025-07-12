package trade

import (
	"context"
	"net"

	"github.com/sourcegraph/jsonrpc2"
)

func GetTransfersForAccount(worker Worker, wallet TrackedWallet) ([]ERC20Transfer, error) {

	conn, err := net.Dial("tcp", worker.AlchemyApiUrl)
	if err != nil {
		return []ERC20Transfer{}, err
	}

	client := jsonrpc2.NewConn(
		context.Background(), jsonrpc2.NewBufferedStream(conn, jsonrpc2.VSCodeObjectCodec{}), jsonrpc2.Handler(func(context.Context, *jsonrpc2.Conn, *jsonrpc2.Request) {

		}))

}
