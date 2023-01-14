package gosuilending

import (
	"context"
	"fmt"
	"testing"

	"github.com/coming-chat/go-sui/client"
	"github.com/coming-chat/go-sui/types"
)

const devnetRpcUrl = "https://fullnode.devnet.sui.io"

func getDevContract() *Contract {
	return &Contract{
		client: getDevClient(),
		// lendingPortalPackageId: getContract().lendingPortalPackageId,
		externalInterfacePackageId: getExternalInterfacePackageId(),
		// bridgePoolPackageId: getContract().bridgePoolPackageId,
		poolManagerInfo: getPoolManager(),
		// poolState: getContract().poolState,
		priceOracle: getPriceOracle(),
		storage:     getStorage(),
		// wormholeState: getContract().wormholeState,
		userManagerInfo: getUserManagerInfo(),
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

func getPriceOracle() *types.HexData {
	priceOracle, err := types.NewHexData("0x44f0e3fcd7fc3d297bfead7d6ea3ff339b353aff")
	AssertNil(err)
	return priceOracle
}

func getSuiDolaChainId() uint16 {
	return 0
}

func getStorage() *types.HexData {
	storage, err := types.NewHexData("0x22e55281cb7974950c5a6849406fed7eb64f1ac5")
	AssertNil(err)
	return storage
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

func getTestAddressAndCallOptions() (*types.Address, CallOptions) {
	address, err := types.NewAddressFromHex("0x4c62953a63373c9cbbbd04a971b9f72109cf9ef3")
	AssertNil(err)
	gasObj, err := types.NewHexData("0x1afaa2aa8a502439f9ffdd7b06a47726563f76cf")
	AssertNil(err)
	return address, CallOptions{Gas: gasObj, GasBudget: 10000}
}

func getUSDTAddress() string {
	return "4a74f62ed7b44ee8dbfb0fc542172ab7ac1da096::coins::USDT"
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
