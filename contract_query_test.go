package gosuilending

import (
	"context"
	"fmt"
	"testing"

	"github.com/coming-chat/go-sui/client"
	"github.com/coming-chat/go-sui/types"
)

const (
	devnetRpcUrl              = "https://fullnode.devnet.sui.io"
	devLendingPortalPackageId = "0x0df92f7b748c47d3caf0450acd666bc080b8d923"
	devExternalInterfaces     = "0x6238067382b88d81f4ee7410b179e5278705d6ef"
	devWormholeBridge         = "0x73d584432c695975829daaf05bf5f2e2d35e9057" //

	devUSDTAddress     = "b77dc99976a25a4162f23fb19de535bf21d15766::coins::USDT"
	devUSDTPoolId      = 1
	devPoolManager     = "0x5fbb6b21ff9242bdf69322c1bef804c8d2beceab"
	devPoolState       = "0x03f88e0cb2e35537e2cf6167ce165d238ee70b6d"
	devPriceOracle     = "0x31d6132d6181e0c0a61db5cdc19c261671b5f243"
	devStorage         = "0xae50809178927a8c5418d742f691f1c9edecd4bd"
	devUserManagerInfo = "0x8d3aab85b96f6f202994416ae540d649ca2d18cb"
	devWormholeState   = "0xe558313313037b879950b07e91358641263c62be"

	devTestUserId      = "13"
	devTestUserAddress = "0x4c62953a63373c9cbbbd04a971b9f72109cf9ef3"
	devTestGasObj      = "0x071e88fb503e74b6cb77a57b177ea501cbce9aee"
)

func getDevContract() *Contract {
	return &Contract{
		client:                     getDevClient(),
		lendingPortalPackageId:     toHex(devLendingPortalPackageId),
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
			v, err := c.GetDolaUserAddresses(tt.args.ctx, tt.args.signer, tt.args.dolaUserId, tt.args.callOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contract.GetDolaUserAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			println(v)
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
