package gosuilending

import (
	"context"
	"fmt"
	"testing"

	"github.com/coming-chat/go-sui/client"
	"github.com/coming-chat/go-sui/types"
)

const (
	devnetRpcUrl              = "https://fullnode.testnet.sui.io"
	devLendingPortalPackageId = "0x5b81c31943358fcf8f20d2c9b92adf6b47062aa4b01afb2e5c901920807c3e09"
	devExternalInterfaces     = "0x25bf584ec396b75ee3ab0367a5b6ebdb82f8fd8bf42f5bd00c281464d604a994"
	devWormholeBridge         = "0xe198cbf3b61678ba33be2a53965c4d68a2b55d00aea67af9038f54c4dba1ec61" //

	devUSDTAddress     = "54fc06a12aeed0752c6db5d949fcf4554bd320ca69676ee9d3085ba946b91af0::coins::USDT"
	devUSDTPoolId      = 1
	devPoolManager     = "0x7b84fae163835c88ea5a8b05a257ccd211bb3c330c63d1f2f19c725d8cf15d11"
	devPoolState       = "0x846bca7df86db919d0bc44d3b328664c4b7bd85b82ea5a20208ccb31d2535d27"
	devPriceOracle     = "0xf16e8e7741c31361ae59547f2ec21c5402d719bccf53e04d53b2e9c369116ae6"
	devStorage         = "0x219a09c981bf165d9bbc40341593777d9391af4c8b8463d1dbbb974b8c34900b"
	devUserManagerInfo = "0x270434bfc0de627d8236e02f10e45beeb341462f6a6457b7659089be781f8468"
	devWormholeState   = "0xb35a426ed4b8b310645ebd978f29944de17bff73397271b5d59695b753d39ace"

	devTestUserId      = "8"
	devTestUserAddress = "0x79e54dcebd85b45b6f447358d529a6c08687e3f98c6e9cd790238299fdedeabc"
	devTestGasObj      = "0x04219e6b31495353970ffbb911de16b45aaa868ac99ce71179bd300881a7eb49"
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
	return address, CallOptions{Gas: gasObj, GasBudget: 30_000_000}
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
