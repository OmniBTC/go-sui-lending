package main

import (
	"context"
	"fmt"

	"github.com/coming-chat/go-sui/types"
	gosuilending "github.com/omnibtc/go-sui-lending"
	"github.com/omnibtc/go-sui-lending/examples/common"
)

const (
	faucetPackageId = "0x13e8531463853d9a3ff017d140be14a9357f6b1d"
	faucetObjectId  = "0x581a7ba6df5a9f2fe2d53637bfa3ce62240a4c3c"
	usdtAddress     = "0x13e8531463853d9a3ff017d140be14a9357f6b1d::coins::USDT"
)

func main() {
	acc := common.GetEnvAccount()
	client := common.GetDevClient()
	contract, err := gosuilending.NewFaucet(client, faucetPackageId, faucetObjectId)
	common.PanicIfError(err)
	signer, err := types.NewHexData(acc.Address)
	common.PanicIfError(err)
	ctx := context.Background()
	coins, err := client.GetSuiCoinsOwnedByAddress(ctx, *signer)
	common.PanicIfError(err)
	gasCoin, err := coins.PickCoinNoLess(1000)
	common.PanicIfError(err)
	tx, err := contract.Claim(context.Background(), *signer, []string{
		usdtAddress,
	}, gosuilending.CallOptions{
		Gas:       &gasCoin.Reference.ObjectId,
		GasBudget: 1000,
	})
	common.PanicIfError(err)

	signedTx := tx.SignWith(acc.PrivateKey)
	resp, err := client.ExecuteTransaction(ctx, *signedTx, types.TxnRequestTypeWaitForLocalExecution)
	common.PanicIfError(err)
	fmt.Println(resp.EffectsCert.Certificate.TransactionDigest)
}
