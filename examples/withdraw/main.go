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
	usdtPool    = "0xf4bc9117ff693bd9086ebdb28aea09c1c7256d9a"
)

func main() {
	usdtPoolObject, err := types.NewHexData(usdtPool)
	common.PanicIfError(err)
	acc := common.GetEnvAccount()
	client := common.GetDevClient()
	contract := common.GetDefaultContract()
	common.PanicIfError(err)
	signer, err := types.NewHexData(acc.Address)
	common.PanicIfError(err)
	ctx := context.Background()
	coins, err := client.GetSuiCoinsOwnedByAddress(ctx, *signer)
	common.PanicIfError(err)
	gasCoin, err := coins.PickCoinNoLess(10000)
	common.PanicIfError(err)

	tx, err := contract.Withdraw(context.Background(), *signer, []string{
		usdtAddress,
	}, gosuilending.WithdrawArgs{
		WormholeMessageCoins:  []types.ObjectId{},
		WormholeMessageAmount: 0,
		Pool:                  *usdtPoolObject,
		DstChain:              1,
		Amount:                49999999,
	}, gosuilending.CallOptions{
		Gas:       &gasCoin.Reference.ObjectId,
		GasBudget: 10000,
	})
	common.PanicIfError(err)

	signedTx := tx.SignWith(acc.PrivateKey)
	resp, err := client.ExecuteTransaction(ctx, *signedTx, types.TxnRequestTypeWaitForLocalExecution)
	common.PanicIfError(err)
	fmt.Println(resp.EffectsCert.Certificate.TransactionDigest)
}
