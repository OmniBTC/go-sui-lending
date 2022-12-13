package main

import (
	"context"
	"fmt"

	"github.com/coming-chat/go-sui/types"
	gosuilending "github.com/omnibtc/go-sui-lending"
	"github.com/omnibtc/go-sui-lending/examples/common"
)

const (
	btcAddress = "0x13e8531463853d9a3ff017d140be14a9357f6b1d::coins::BTC"
	btcPool    = "0x2240c0e485c4c86a68edba2f8797ca3bcab5366a"
)

func main() {
	btcPoolObject, err := types.NewHexData(btcPool)
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

	// get btc coins
	btcCoins, err := client.GetCoinsOwnedByAddress(ctx, *signer, btcAddress)
	common.PanicIfError(err)
	btcCoinObjectIds := []types.ObjectId{}
	for _, coin := range btcCoins {
		btcCoinObjectIds = append(btcCoinObjectIds, coin.Reference.ObjectId)
	}

	tx, err := contract.Repay(context.Background(), *signer, []string{
		btcAddress,
	}, gosuilending.RepayArgs{
		WormholeMessageCoins:  []types.ObjectId{},
		WormholeMessageAmount: 0,
		Pool:                  *btcPoolObject,
		RepayCoins:            btcCoinObjectIds,
		RepayAmount:           200,
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
