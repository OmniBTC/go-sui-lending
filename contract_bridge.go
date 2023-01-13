package gosuilending

import (
	"context"

	"github.com/coming-chat/go-sui/types"
)

type (
	BindingArgs struct {
		WormholeMessageCoins  []types.ObjectId // vector<Coin<SUI>>
		WormholeMessageAmount string
		DolaChainId           uint16
		BindAddress           string
	}

	UnbindingArgs struct {
		WormholeMessageCoins  []types.ObjectId // vector<Coin<SUI>>
		WormholeMessageAmount string
		DolaChainId           uint16
		UnbindAddress         string
	}
)

func (c *Contract) SendBinding(ctx context.Context, signer types.Address, typeArgs []string, bindingArgs BindingArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		*c.poolState,
		*c.wormholeState,
		bindingArgs.WormholeMessageCoins,
		bindingArgs.WormholeMessageAmount,
		bindingArgs.DolaChainId,
		bindingArgs.BindAddress,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.bridgePoolPackageId, "bridge_pool", "send_binding", typeArgs, args, callOptions.Gas, callOptions.GasBudget)
	return resp, err
}

func (c *Contract) SendingUnbinding(ctx context.Context, signer types.Address, typeArgs []string, unbindingArgs UnbindingArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		*c.poolState,
		*c.wormholeState,
		unbindingArgs.WormholeMessageCoins,
		unbindingArgs.WormholeMessageAmount,
		unbindingArgs.DolaChainId,
		unbindingArgs.UnbindAddress,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.bridgePoolPackageId, "bridge_pool", "send_unbinding", typeArgs, args, callOptions.Gas, callOptions.GasBudget)
	return resp, err
}
