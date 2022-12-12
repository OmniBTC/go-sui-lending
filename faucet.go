package gosuilending

import (
	"context"

	"github.com/coming-chat/go-sui/client"
	"github.com/coming-chat/go-sui/types"
)

type Faucet interface {
	Claim(ctx context.Context, signer types.Address, typeArgs []string, callOptions CallOptions) (*types.TransactionBytes, error)
}

type innerFaucetContract struct {
	client    *client.Client
	packageId *types.HexData
	faucetId  *types.ObjectId
}

func NewFaucet(client *client.Client, packageId, faucetId string) (Faucet, error) {
	c := &innerFaucetContract{client: client}
	var err error
	if c.packageId, err = types.NewHexData(packageId); err != nil {
		return nil, err
	}
	if c.faucetId, err = types.NewHexData(faucetId); err != nil {
		return nil, err
	}
	return c, nil
}

func (i *innerFaucetContract) Claim(ctx context.Context, signer types.Address, typeArgs []string, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		*i.faucetId,
	}
	resp, err := i.client.MoveCall(ctx, signer, *i.packageId, "faucet", "claim", typeArgs, args, callOptions.Gas, callOptions.GasBudget)
	return resp, err
}
