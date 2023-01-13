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
	t, err := types.NewHexData("0x66b49dd9b363e46038a1d31993362890890ad9af")
	AssertNil(err)
	return t
}

func getExternalInterfacePackageId() *types.HexData {
	externalInterfacePackageId, err := types.NewHexData("0xfc6568c500a90c4ec220a36eb969e4415a399f17")
	AssertNil(err)
	return externalInterfacePackageId
}

func getUSDTPoolId() uint16 {
	return 1
}

func getUserDolaId() string {
	return "6"
}

func getPoolManager() *types.HexData {
	poolManager, err := types.NewHexData("0x1cd53845462cac4fb0b8676c7858c1b5b1626c77")
	AssertNil(err)
	return poolManager
}

func getTestAddressAndGas() (*types.Address, *types.HexData) {
	address, err := types.NewAddressFromHex("0x4c62953a63373c9cbbbd04a971b9f72109cf9ef3")
	AssertNil(err)
	gasObj, err := types.NewHexData("0x1afaa2aa8a502439f9ffdd7b06a47726563f76cf")
	AssertNil(err)
	return address, gasObj
}

func getUSDTAddress() string {
	return "4a74f62ed7b44ee8dbfb0fc542172ab7ac1da096::coins::USDT"
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

func TestContract_GetPoolLiquidity(t *testing.T) {
	address, gasObj := getTestAddressAndGas()
	type fields struct {
		client                     *client.Client
		externalInterfacePackageId *types.HexData
		poolManagerInfo            *types.HexData
	}
	type args struct {
		ctx         context.Context
		signer      types.Address
		dolaChainId uint16
		poolAddress string
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
				poolManagerInfo:            getPoolManager(),
			},
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				dolaChainId: 0,
				poolAddress: getUSDTAddress(),
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
			_, err := c.GetPoolLiquidity(tt.args.ctx, tt.args.signer, tt.args.dolaChainId, tt.args.poolAddress, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetPoolLiquidity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestContract_GetAllPoolLiquidity(t *testing.T) {
	address, gasObj := getTestAddressAndGas()
	type fields struct {
		client                     *client.Client
		externalInterfacePackageId *types.HexData
		poolManagerInfo            *types.HexData
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
				externalInterfacePackageId: getExternalInterfacePackageId(),
				poolManagerInfo:            getPoolManager(),
			},
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				dolaPoolId:  getUSDTPoolId(),
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
			_, err := c.GetAllPoolLiquidity(tt.args.ctx, tt.args.signer, tt.args.dolaPoolId, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetAllPoolLiquidity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestContract_GetDolaUserAddresses(t *testing.T) {
	signer, gas := getTestAddressAndGas()
	type fields struct {
		client                     *client.Client
		externalInterfacePackageId *types.HexData
		userManagerInfo            *types.HexData
	}
	type args struct {
		ctx         context.Context
		signer      types.Address
		dolaUserId  string
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
				signer:      *signer,
				dolaUserId:  getUserDolaId(),
				callOptions: CallOptions{Gas: gas, GasBudget: 10000},
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
			_, err := c.GetDolaUserAddresses(tt.args.ctx, tt.args.signer, tt.args.dolaUserId, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetDolaUserAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
