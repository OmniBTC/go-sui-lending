package gosuilending

import (
	"context"

	"github.com/coming-chat/go-sui/v2/sui_types"
	"github.com/coming-chat/go-sui/v2/types"
)

type (
	BindingArgs struct {
		WormholeMessageCoins  []sui_types.ObjectID // vector<Coin<SUI>>
		WormholeMessageAmount string
		DolaChainId           uint16
		BindAddress           string
	}

	UnbindingArgs struct {
		WormholeMessageCoins  []sui_types.ObjectID // vector<Coin<SUI>>
		WormholeMessageAmount string
		DolaChainId           uint16
		UnbindAddress         string
	}
)

func (c *Contract) SendBinding(ctx context.Context, signer sui_types.SuiAddress, typeArgs []string, bindingArgs BindingArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		*c.poolState,
		*c.wormholeState,
		bindingArgs.WormholeMessageCoins,
		bindingArgs.WormholeMessageAmount,
		bindingArgs.DolaChainId,
		bindingArgs.BindAddress,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.bridgePoolPackageId, "bridge_pool", "send_binding", typeArgs, args, callOptions.Gas, types.NewSafeSuiBigInt(callOptions.GasBudget))
	return resp, err
}

func (c *Contract) SendingUnbinding(ctx context.Context, signer sui_types.SuiAddress, typeArgs []string, unbindingArgs UnbindingArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		*c.poolState,
		*c.wormholeState,
		unbindingArgs.WormholeMessageCoins,
		unbindingArgs.WormholeMessageAmount,
		unbindingArgs.DolaChainId,
		unbindingArgs.UnbindAddress,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.bridgePoolPackageId, "bridge_pool", "send_unbinding", typeArgs, args, callOptions.Gas, types.NewSafeSuiBigInt(callOptions.GasBudget))
	return resp, err
}
