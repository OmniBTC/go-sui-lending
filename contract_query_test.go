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

func getUserManagerInfo() *types.HexData {
	t, err := types.NewHexData("0x489f3e671c6d493d6b28aa1c682f71ce80c8e263")
	AssertNil(err)
	return t
}

func getExternalInterfacePackageId() *types.HexData {
	externalInterfacePackageId, err := types.NewHexData("0x20c2b9cb6d88de7dcf2b6ba98900058e1d58781c")
	AssertNil(err)
	return externalInterfacePackageId
}

func getPoolManager() *types.HexData {
	poolManager, err := types.NewHexData("0x6f68637e8f8f98ac62d7a08efaeacebbcd620ce9")
	AssertNil(err)
	return poolManager
}

func getTestAddressAndGas() (*types.Address, *types.HexData) {
	address, err := types.NewAddressFromHex("0x4c62953a63373c9cbbbd04a971b9f72109cf9ef3")
	AssertNil(err)
	gasObj, err := types.NewHexData("0x085c11efc7d4d405f75c91f8a1990de89182acd5")
	AssertNil(err)
	return address, gasObj
}

func TestContract_GetDolaTokenLiquidity(t *testing.T) {
	address, gasObj := getTestAddressAndGas()
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
				poolManagerInfo:            getPoolManager(),
				externalInterfacePackageId: getExternalInterfacePackageId(),
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
	address, gasObj := getTestAddressAndGas()
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
				poolManagerInfo:            getPoolManager(),
				externalInterfacePackageId: getExternalInterfacePackageId(),
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

func TestContract_GetDolaUserId(t *testing.T) {
	address, gasObj := getTestAddressAndGas()
	type fields struct {
		client                     *client.Client
		externalInterfacePackageId *types.HexData
		userManagerInfo            *types.HexData
	}
	type args struct {
		ctx         context.Context
		signer      types.Address
		dolaChainId uint16
		user        string
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
				externalInterfacePackageId: getExternalInterfacePackageId(),
				userManagerInfo:            getUserManagerInfo(),
			},
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				dolaChainId: 0,
				user:        address.String(),
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
				userManagerInfo:            tt.fields.userManagerInfo,
			}
			_, err := c.GetDolaUserId(tt.args.ctx, tt.args.signer, tt.args.dolaChainId, tt.args.user, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetDolaUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
