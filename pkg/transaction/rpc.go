package transaction

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type RPCBackend interface {
	BlockHashAt(ctx context.Context, blockNumber *big.Int) (common.Hash, error)
}

type rpcBackend struct {
	client *rpc.Client
}

func NewRPCBackend(client *rpc.Client) RPCBackend {
	return &rpcBackend{
		client: client,
	}
}

type blockHeaderResponse struct {
	BlockHash common.Hash `json:"blockHash"`
}

func (c rpcBackend) BlockHashAt(ctx context.Context, blockNumber *big.Int) (common.Hash, error) {
	var resp blockHeaderResponse
	err := c.client.CallContext(ctx, &resp, "eth_getBlockByNumber", blockNumber, false)
	if err != nil {
		return common.Hash{}, err
	}
	return resp.BlockHash, nil
}
