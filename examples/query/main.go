package main

import (
	"context"
	"fmt"

	"github.com/coming-chat/go-sui/types"
	gosuilending "github.com/omnibtc/go-sui-lending"
	"github.com/omnibtc/go-sui-lending/examples/common"
)

const (
	usdtAddress = "0x13e8531463853d9a3ff017d140be14a9357f6b1d::coins::USDT"
	btcAddress  = "0x13e8531463853d9a3ff017d140be14a9357f6b1d::coins::BTC"
)

func main() {
	acc := common.GetEnvAccount()
	client := common.GetDevClient()
	contract := common.GetDefaultContract()
	signer, err := types.NewHexData(acc.Address)
	common.PanicIfError(err)
	ctx := context.Background()
	coins, err := client.GetSuiCoinsOwnedByAddress(ctx, *signer)
	common.PanicIfError(err)
	gasCoin, err := coins.PickCoinNoLess(10000)
	common.PanicIfError(err)

	callOptions := gosuilending.CallOptions{
		Gas:       &gasCoin.Reference.ObjectId,
		GasBudget: 10000,
	}

	liquid, err := contract.GetDolaTokenLiquidity(ctx, *signer, usdtAddress, callOptions)
	common.PanicIfError(err)
	println("dola token liquidity:", liquid.String())

	appTokenLiquidity, err := contract.GetAppTokenLiquidity(ctx, *signer, 0, usdtAddress, callOptions)
	common.PanicIfError(err)
	println("app token liquidity:", appTokenLiquidity.String())

	debtAmount, debtValue, err := contract.GetUserTokenDebt(ctx, *signer, btcAddress, callOptions)
	common.PanicIfError(err)
	println("user btc token debt", debtAmount.String(), " ", debtValue.String())

	collateralAmount, collateralValue, err := contract.GetUserCollateral(ctx, *signer, usdtAddress, callOptions)
	common.PanicIfError(err)
	println("collateral: ", collateralAmount.String(), " ", collateralValue.String())

	userLendingInfo, err := contract.GetUserLendingInfo(ctx, *signer, callOptions)
	common.PanicIfError(err)
	fmt.Printf("lending info %v\n", userLendingInfo)

	reserveInfo, err := contract.GetReserveInfo(ctx, *signer, usdtAddress, callOptions)
	common.PanicIfError(err)
	fmt.Printf("%v\n", reserveInfo)

	// canBorrowAmount, err := contract.GetUserAllowedBorrow(ctx, *signer, usdtAddress, callOptions)
	// println("user can borrow usdt:", canBorrowAmount.String())
	// if err != nil {
	// 	println("reason:", err.Error())
	// }

	canBorrowAmount, err := contract.GetUserAllowedBorrow(ctx, *signer, btcAddress, callOptions)
	println("user can borrow btc:", canBorrowAmount.String())
	if err != nil {
		println("reason:", err.Error())
	}
}
