package common

import (
	"os"

	"github.com/coming-chat/go-sui/account"
	"github.com/coming-chat/go-sui/client"
	gosuilending "github.com/omnibtc/go-sui-lending"
)

const DevnetRpcUrl = "https://fullnode.devnet.sui.io"

func GetDevClient() *client.Client {
	c, err := client.Dial(DevnetRpcUrl)
	PanicIfError(err)
	return c
}

func GetEnvAccount() *account.Account {
	account, err := account.NewAccountWithMnemonic(os.Getenv("m"))
	PanicIfError(err)
	return account
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetDefaultContract() *gosuilending.Contract {
	contract, err := gosuilending.NewContract(GetDevClient(), gosuilending.ContractConfig{
		LendingPortalPackageId:     "0x481619b177aabe0f4c6c06e0b141e3373644f90e",
		ExternalInterfacePackageId: "0xe558bd8e7a6a88a405ffd93cc71ecf1ade45686c",
		PoolManagerInfo:            "0x00e2cd853b00a1531b5a5579156a174891543e50",
		PoolState:                  "0x26220207229eece5c32f4200badb3333e30d1dd8",
		PriceOracle:                "0x6895216a4f584c747c196d6bb43a39ec59f94f11",
		Storage:                    "0xe9d6e200a86ef34a0f9388034cd66739e0c4782c",
		WormholeState:              "0x1d8a0ac5d10100111eae5509c637ee6841d9955e",
	})
	PanicIfError(err)
	return contract
}
