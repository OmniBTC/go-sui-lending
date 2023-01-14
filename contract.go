package gosuilending

import (
	"context"

	"github.com/coming-chat/go-sui/client"
	"github.com/coming-chat/go-sui/types"
)

type CallOptions struct {
	Gas       *types.ObjectId
	GasBudget uint64
}

type SupplyArgs struct {
	WormholeMessageCoins  []types.ObjectId // vector<Coin<SUI>>
	WormholeMessageAmount string
	Pool                  types.HexData
	DepositCoins          []types.ObjectId // vector<Coin<CoinType>>
	DepositAmount         string
}

type WithdrawArgs struct {
	Pool                  types.HexData
	Receiver              string
	DstChain              string
	WormholeMessageCoins  []types.ObjectId // vector<Coin<SUI>>
	WormholeMessageAmount string
	Amount                string
}

type BorrowArgs struct {
	Pool                  types.HexData
	Receiver              string
	DstChain              string
	WormholeMessageCoins  []types.ObjectId // vector<Coin<SUI>>
	WormholeMessageAmount string
	Amount                string
}

type RepayArgs struct {
	Pool                  types.HexData
	WormholeMessageCoins  []types.ObjectId // vector<Coin<SUI>>
	WormholeMessageAmount string
	RepayCoins            []types.ObjectId // vector<Coin<CoinType>>
	RepayAmount           string
}

type ContractConfig struct {
	LendingPortalPackageId     string
	ExternalInterfacePackageId string
	BridgePoolPackageId        string
	PoolManagerInfo            string
	PoolState                  string
	PriceOracle                string
	Storage                    string
	WormholeState              string
	UserManagerInfo            string
}

type Contract struct {
	client *client.Client

	lendingPortalPackageId     *types.HexData
	externalInterfacePackageId *types.HexData
	bridgePoolPackageId        *types.HexData
	poolManagerInfo            *types.HexData
	poolState                  *types.HexData
	priceOracle                *types.HexData
	storage                    *types.HexData
	wormholeState              *types.HexData
	userManagerInfo            *types.HexData
}

func NewContract(client *client.Client, config ContractConfig) (*Contract, error) {
	contract := &Contract{client: client}
	var err error
	if contract.lendingPortalPackageId, err = types.NewHexData(config.LendingPortalPackageId); err != nil {
		return nil, err
	}
	if contract.externalInterfacePackageId, err = types.NewHexData(config.ExternalInterfacePackageId); err != nil {
		return nil, err
	}
	if contract.bridgePoolPackageId, err = types.NewHexData(config.BridgePoolPackageId); err != nil {
		return nil, err
	}
	if contract.poolManagerInfo, err = types.NewHexData(config.PoolManagerInfo); err != nil {
		return nil, err
	}
	if contract.poolState, err = types.NewHexData(config.PoolState); err != nil {
		return nil, err
	}
	if contract.priceOracle, err = types.NewHexData(config.PriceOracle); err != nil {
		return nil, err
	}
	if contract.storage, err = types.NewHexData(config.Storage); err != nil {
		return nil, err
	}
	if contract.wormholeState, err = types.NewHexData(config.WormholeState); err != nil {
		return nil, err
	}
	if contract.userManagerInfo, err = types.NewHexData(config.UserManagerInfo); err != nil {
		return nil, err
	}
	return contract, nil
}

func (c *Contract) Supply(ctx context.Context, signer types.Address, typeArgs []string, supplyArgs SupplyArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		*c.poolState,
		*c.wormholeState,
		supplyArgs.WormholeMessageCoins,
		supplyArgs.WormholeMessageAmount,
		supplyArgs.Pool,
		supplyArgs.DepositCoins,
		supplyArgs.DepositAmount,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.lendingPortalPackageId, "lending", "supply", typeArgs, args, callOptions.Gas, callOptions.GasBudget)
	return resp, err
}

func (c *Contract) Withdraw(ctx context.Context, signer types.Address, typeArgs []string, withdrawArgs WithdrawArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		withdrawArgs.Pool,
		*c.poolState,
		*c.wormholeState,
		withdrawArgs.Receiver,
		withdrawArgs.DstChain,
		withdrawArgs.WormholeMessageCoins,
		withdrawArgs.WormholeMessageAmount,
		withdrawArgs.Amount,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.lendingPortalPackageId, "lending", "withdraw", typeArgs, args, callOptions.Gas, callOptions.GasBudget)
	return resp, err
}

func (c *Contract) Borrow(ctx context.Context, signer types.Address, typeArgs []string, borrowArgs BorrowArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		borrowArgs.Pool,
		*c.poolState,
		*c.wormholeState,
		borrowArgs.Receiver,
		borrowArgs.DstChain,
		borrowArgs.WormholeMessageCoins,
		borrowArgs.WormholeMessageAmount,
		borrowArgs.Amount,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.lendingPortalPackageId, "lending", "borrow", typeArgs, args, callOptions.Gas, callOptions.GasBudget)
	return resp, err
}

func (c *Contract) Repay(ctx context.Context, signer types.Address, typeArgs []string, repayArgs RepayArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		repayArgs.Pool,
		*c.poolState,
		*c.wormholeState,
		repayArgs.WormholeMessageCoins,
		repayArgs.WormholeMessageAmount,
		repayArgs.RepayCoins,
		repayArgs.RepayAmount,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.lendingPortalPackageId, "lending", "repay", typeArgs, args, callOptions.Gas, callOptions.GasBudget)
	return resp, err
}
