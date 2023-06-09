package gosuilending

import (
	"context"

	"github.com/coming-chat/go-sui/v2/client"
	"github.com/coming-chat/go-sui/v2/sui_types"
	"github.com/coming-chat/go-sui/v2/types"
)

type Faucet interface {
	Claim(ctx context.Context, signer sui_types.SuiAddress, typeArgs []string, callOptions CallOptions) (*types.TransactionBytes, error)
}

type innerFaucetContract struct {
	client    *client.Client
	packageId *sui_types.ObjectID
	faucetId  *sui_types.ObjectID
}

func NewFaucet(client *client.Client, packageId, faucetId string) (Faucet, error) {
	c := &innerFaucetContract{client: client}
	var err error
	if c.packageId, err = sui_types.NewObjectIdFromHex(packageId); err != nil {
		return nil, err
	}
	if c.faucetId, err = sui_types.NewObjectIdFromHex(faucetId); err != nil {
		return nil, err
	}
	return c, nil
}

func (i *innerFaucetContract) Claim(ctx context.Context, signer sui_types.SuiAddress, typeArgs []string, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		*i.faucetId,
	}
	resp, err := i.client.MoveCall(ctx, signer, *i.packageId, "faucet", "claim", typeArgs, args, callOptions.Gas, types.NewSafeSuiBigInt(callOptions.GasBudget))
	return resp, err
}
