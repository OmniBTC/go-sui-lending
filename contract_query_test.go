package gosuilending

import (
	"context"
	"fmt"
	"testing"

	"github.com/coming-chat/go-sui/client"
	"github.com/coming-chat/go-sui/types"
)

const (
	devnetRpcUrl          = "https://fullnode.testnet.sui.io"
	devLendingPortal      = "0x725996f982c461ddb1060cb64a7b47246e7332be"
	devExternalInterfaces = "0xaa65494974bfa11425bfbff836db69cf7950f3ef"
	devWormholeBridge     = "0x10ec199c006b64d40511ca7f2f0527051577d23f"

	devUSDTAddress     = "72b846eca3c7f91961ec3cae20441be96a21e1fe::coins::USDT"
	devUSDTPoolId      = 1
	devPoolManager     = "0xa6c1415b41a0a768fb49dcdcc1d2587f8956e739"
	devPoolState       = "0x3037ba7392653eb9a4850b7c471a839959b09de0"
	devPriceOracle     = "0x39b21e8cf71ca3c6d6c6bd03a01753d9526a5502"
	devStorage         = "0xf709a28d31c38a1be1e61fee1e1e77217d9ba554"
	devUserManagerInfo = "0x789ef32e90f5d97f0475f127f501ce31257da033"
	devWormholeState   = "0x69d54fb067de394c88d18f6217c950a780bc148c"

	devTestUserId      = "230"
	devTestUserAddress = "0x4c62953a63373c9cbbbd04a971b9f72109cf9ef3"
	devTestGasObj      = "0x0e304469df7958ab1beeac819dd92877eb3bd165"
)

func getDevContract() *Contract {
	return &Contract{
		client:                     getDevClient(),
		lendingPortalPackageId:     toHex(devLendingPortal),
		externalInterfacePackageId: toHex(devExternalInterfaces),
		bridgePoolPackageId:        toHex(devWormholeBridge),
		poolManagerInfo:            toHex(devPoolManager),
		poolState:                  toHex(devPoolState),
		priceOracle:                toHex(devPriceOracle),
		storage:                    toHex(devStorage),
		wormholeState:              toHex(devWormholeState),
		userManagerInfo:            toHex(devUserManagerInfo),
	}
}

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

func getUSDTPoolId() uint16 {
	return devUSDTPoolId
}

func getUserDolaId() string {
	return devTestUserId
}

func toHex(str string) *types.HexData {
	hexData, err := types.NewHexData(str)
	AssertNil(err)
	return hexData
}

func getTestAddressAndCallOptions() (*types.Address, CallOptions) {
	address := toHex(devTestUserAddress)
	gasObj := toHex(devTestGasObj)
	return address, CallOptions{Gas: gasObj, GasBudget: 10000}
}

func getUSDTAddress() string {
	return devUSDTAddress
}

func TestContract_GetDolaTokenLiquidity(t *testing.T) {
	address, callOptions := getTestAddressAndCallOptions()
	type args struct {
		ctx         context.Context
		signer      types.Address
		dolaPoolId  uint16
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case",
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				dolaPoolId:  1,
				callOptions: callOptions,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := getDevContract()
			_, err := c.GetDolaTokenLiquidity(tt.args.ctx, tt.args.signer, tt.args.dolaPoolId, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetDolaTokenLiquidity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestContract_GetAppTokenLiquidity(t *testing.T) {
	address, callOptions := getTestAddressAndCallOptions()
	type args struct {
		ctx         context.Context
		signer      types.Address
		appId       uint16
		dolaPoolId  uint16
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case",
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				appId:       0,
				dolaPoolId:  1,
				callOptions: callOptions,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := getDevContract()
			_, err := c.GetAppTokenLiquidity(tt.args.ctx, tt.args.signer, tt.args.appId, tt.args.dolaPoolId, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetAppTokenLiquidity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestContract_GetDolaUserId(t *testing.T) {
	address, callOptions := getTestAddressAndCallOptions()
	type args struct {
		ctx         context.Context
		signer      types.Address
		dolaChainId uint16
		user        string
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case",
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				dolaChainId: 0,
				user:        address.String(),
				callOptions: callOptions,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := getDevContract()
			_, err := c.GetDolaUserId(tt.args.ctx, tt.args.signer, tt.args.dolaChainId, tt.args.user, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetDolaUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestContract_GetPoolLiquidity(t *testing.T) {
	address, callOptions := getTestAddressAndCallOptions()
	type args struct {
		ctx         context.Context
		signer      types.Address
		dolaChainId uint16
		poolAddress string
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case",
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				dolaChainId: 0,
				poolAddress: getUSDTAddress(),
				callOptions: callOptions,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := getDevContract()
			_, err := c.GetPoolLiquidity(tt.args.ctx, tt.args.signer, tt.args.dolaChainId, tt.args.poolAddress, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetPoolLiquidity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestContract_GetAllPoolLiquidity(t *testing.T) {
	address, callOptions := getTestAddressAndCallOptions()
	type args struct {
		ctx         context.Context
		signer      types.Address
		dolaPoolId  uint16
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case",
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				dolaPoolId:  getUSDTPoolId(),
				callOptions: callOptions,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := getDevContract()
			_, err := c.GetAllPoolLiquidity(tt.args.ctx, tt.args.signer, tt.args.dolaPoolId, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetAllPoolLiquidity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestContract_GetDolaUserAddresses(t *testing.T) {
	signer, callOptions := getTestAddressAndCallOptions()
	type args struct {
		ctx         context.Context
		signer      types.Address
		dolaUserId  string
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case",
			args: args{
				ctx:         context.Background(),
				signer:      *signer,
				dolaUserId:  getUserDolaId(),
				callOptions: callOptions,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := getDevContract()
			_, err := c.GetDolaUserAddresses(tt.args.ctx, tt.args.signer, tt.args.dolaUserId, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetDolaUserAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestContract_GetUserHealthFactor(t *testing.T) {
	address, callOptions := getTestAddressAndCallOptions()
	type args struct {
		ctx         context.Context
		signer      types.Address
		dolaUserId  string
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case",
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				dolaUserId:  getUserDolaId(),
				callOptions: callOptions,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := getDevContract()
			_, err := c.GetUserHealthFactor(tt.args.ctx, tt.args.signer, tt.args.dolaUserId, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetUserHealthFactor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestContract_GetAllOraclePrice(t *testing.T) {
	address, callOptions := getTestAddressAndCallOptions()
	type args struct {
		ctx         context.Context
		signer      types.Address
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case",
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				callOptions: callOptions,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := getDevContract()
			if _, err := c.GetAllOraclePrice(tt.args.ctx, tt.args.signer, tt.args.callOptions); (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetAllOraclePrice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContract_GetOraclePrice(t *testing.T) {
	address, callOptions := getTestAddressAndCallOptions()
	type args struct {
		ctx         context.Context
		signer      types.Address
		dolaPoolId  uint16
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case",
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				dolaPoolId:  getUSDTPoolId(),
				callOptions: callOptions,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := getDevContract()
			if _, err := c.GetOraclePrice(tt.args.ctx, tt.args.signer, tt.args.dolaPoolId, tt.args.callOptions); (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetOraclePrice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContract_GetAllReserveInfo(t *testing.T) {
	address, callOptions := getTestAddressAndCallOptions()
	type args struct {
		ctx         context.Context
		signer      types.Address
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case",
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				callOptions: callOptions,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := getDevContract()
			if _, err := c.GetAllReserveInfo(tt.args.ctx, tt.args.signer, tt.args.callOptions); (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetAllReserveInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContract_GetReserveInfo(t *testing.T) {
	address, callOptions := getTestAddressAndCallOptions()
	type args struct {
		ctx         context.Context
		signer      types.Address
		dolaPoolId  uint16
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case",
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				dolaPoolId:  getUSDTPoolId(),
				callOptions: callOptions,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := getDevContract()
			_, err := c.GetReserveInfo(tt.args.ctx, tt.args.signer, tt.args.dolaPoolId, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetReserveInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestContract_GetUserCollateral(t *testing.T) {
	address, callOptions := getTestAddressAndCallOptions()
	type args struct {
		ctx         context.Context
		signer      types.Address
		dolaUserId  string
		dolaPoolId  uint16
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case",
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				dolaUserId:  getUserDolaId(),
				dolaPoolId:  getUSDTPoolId(),
				callOptions: callOptions,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := getDevContract()
			_, err := c.GetUserCollateral(tt.args.ctx, tt.args.signer, tt.args.dolaUserId, tt.args.dolaPoolId, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetUserCollateral() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestContract_GetUserLendingInfo(t *testing.T) {
	address, callOptions := getTestAddressAndCallOptions()
	type args struct {
		ctx         context.Context
		signer      types.Address
		dolaUserId  string
		callOptions CallOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case",
			args: args{
				ctx:         context.Background(),
				signer:      *address,
				dolaUserId:  getUserDolaId(),
				callOptions: callOptions,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := getDevContract()
			_, err := c.GetUserLendingInfo(tt.args.ctx, tt.args.signer, tt.args.dolaUserId, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetUserLendingInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
