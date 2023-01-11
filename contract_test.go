package gosuilending

import (
	"context"
	"fmt"
	"testing"

	"github.com/coming-chat/go-sui/client"
	"github.com/coming-chat/go-sui/types"
)

const devnetRpcUrl = "https://fullnode.devnet.sui.io"

func getDevClient() *client.Client {
	c, err := client.Dial(devnetRpcUrl)
	AssertNil(err)
	return c
}

func AssertNil(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func TestContract_GetDolaTokenLiquidity(t *testing.T) {
	externalInterfacePackageId, err := types.NewHexData("0x20c2b9cb6d88de7dcf2b6ba98900058e1d58781c")
	AssertNil(err)
	poolManager, err := types.NewHexData("0x6f68637e8f8f98ac62d7a08efaeacebbcd620ce9")
	AssertNil(err)
	address, err := types.NewAddressFromHex("0x4c62953a63373c9cbbbd04a971b9f72109cf9ef3")
	AssertNil(err)
	gasObj, err := types.NewHexData("0x085c11efc7d4d405f75c91f8a1990de89182acd5")
	AssertNil(err)
	type fields struct {
		client                     *client.Client
		poolManagerInfo            *types.HexData
		externalInterfacePackageId *types.HexData
	}
	type args struct {
		ctx         context.Context
		signer      types.Address
		dolaPoolId  uint16
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "case",
			fields: fields{
				client:                     getDevClient(),
				poolManagerInfo:            poolManager,
				externalInterfacePackageId: externalInterfacePackageId,
			},
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				dolaPoolId:  1,
				callOptions: CallOptions{Gas: gasObj, GasBudget: 10000},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Contract{
				client:                     tt.fields.client,
				poolManagerInfo:            tt.fields.poolManagerInfo,
				externalInterfacePackageId: tt.fields.externalInterfacePackageId,
			}
			_, err := c.GetDolaTokenLiquidity(tt.args.ctx, tt.args.signer, tt.args.dolaPoolId, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetDolaTokenLiquidity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestContract_GetAppTokenLiquidity(t *testing.T) {
	externalInterfacePackageId, err := types.NewHexData("0x20c2b9cb6d88de7dcf2b6ba98900058e1d58781c")
	AssertNil(err)
	poolManager, err := types.NewHexData("0x6f68637e8f8f98ac62d7a08efaeacebbcd620ce9")
	AssertNil(err)
	address, err := types.NewAddressFromHex("0x4c62953a63373c9cbbbd04a971b9f72109cf9ef3")
	AssertNil(err)
	gasObj, err := types.NewHexData("0x085c11efc7d4d405f75c91f8a1990de89182acd5")
	AssertNil(err)
	type fields struct {
		client                     *client.Client
		externalInterfacePackageId *types.HexData
		poolManagerInfo            *types.HexData
	}
	type args struct {
		ctx         context.Context
		signer      types.Address
		appId       uint16
		dolaPoolId  uint16
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "case",
			fields: fields{
				client:                     getDevClient(),
				poolManagerInfo:            poolManager,
				externalInterfacePackageId: externalInterfacePackageId,
			},
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				appId:       0,
				dolaPoolId:  1,
				callOptions: CallOptions{Gas: gasObj, GasBudget: 10000},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Contract{
				client:                     tt.fields.client,
				externalInterfacePackageId: tt.fields.externalInterfacePackageId,
				poolManagerInfo:            tt.fields.poolManagerInfo,
			}
			_, err := c.GetAppTokenLiquidity(tt.args.ctx, tt.args.signer, tt.args.appId, tt.args.dolaPoolId, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetAppTokenLiquidity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
