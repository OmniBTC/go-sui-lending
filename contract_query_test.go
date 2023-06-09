package gosuilending

import (
	"context"
	"fmt"
	"testing"

	"github.com/coming-chat/go-sui/v2/client"
	"github.com/coming-chat/go-sui/v2/sui_types"
)

const (
	devnetRpcUrl              = "https://fullnode.mainnet.sui.io"
	devLendingPortalPackageId = "0xc5b2a5049cd71586362d0c6a38e34cfaae7ea9ce6d5401a350506a15f817bf72"
	devExternalInterfaces     = "0x93b49ef245f169342cb07e70b6a4835d4071594451a9df738acbb5ecdcac2e88"
	devWormholeBridge         = "0x5306f64e312b581766351c07af79c72fcb1cd25147157fdc2f8ad76de9a3fb6a" //

	devUSDTAddress     = "c060006111016b8a020ad5b33834984a437aaa7d3c74c18e09a95d48aceab08c::coin::COIN"
	devUSDTPoolId      = 1
	devPoolManager     = "0x1be839a23e544e8d4ba7fab09eab50626c5cfed80f6a22faf7ff71b814689cfb"
	devPoolState       = "0x5c9d9db2dd5f34154ee59686334f3504026809fa67afe5332837191ee6220586"
	devPriceOracle     = "0x42afbffd3479b06f40c5576799b02ea300df36cf967adcd1ae15445270f572e2"
	devStorage         = "0xe5a189b1858b207f2cf8c05a09d75bae4271c7a9a8f84a8c199c6896dc7c37e6"
	devUserManagerInfo = "0xee633dc3fd1218d3bd9703fb9b98e6c8d7fdd8c8bf1ca2645ee40d65fb533a3e"
	devWormholeState   = "0xaeab97f96cf9877fee2883315d459552b2b921edc16d7ceac6eab944dd88919c"

	devTestUserId      = "72"
	devTestUserAddress = "0x79e54dcebd85b45b6f447358d529a6c08687e3f98c6e9cd790238299fdedeabc"
	devTestGasObj      = "0x09db26ce25076c41d7cd9008ae6aa521e73940686d91a49800198ec3710cc8a3"
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

func toHex(str string) *sui_types.ObjectID {
	hexData, err := sui_types.NewObjectIdFromHex(str)
	AssertNil(err)
	return hexData
}

func getTestAddressAndCallOptions() (*sui_types.SuiAddress, CallOptions) {
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
		signer      sui_types.SuiAddress
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
		signer      sui_types.SuiAddress
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
		signer      sui_types.SuiAddress
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
		signer      sui_types.SuiAddress
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
		signer      sui_types.SuiAddress
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
		signer      sui_types.SuiAddress
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
		signer      sui_types.SuiAddress
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
		signer      sui_types.SuiAddress
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
		signer      sui_types.SuiAddress
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
		signer      sui_types.SuiAddress
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
		signer      sui_types.SuiAddress
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
		signer      sui_types.SuiAddress
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
		signer      sui_types.SuiAddress
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
