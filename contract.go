package gosuilending

import (
	"context"

	"github.com/coming-chat/go-sui/client"
	"github.com/coming-chat/go-sui/types"
)

type Contract interface {
}

type CallOptions struct {
	gas       *types.ObjectId
	gasBudget uint64
}

type SupplyArgs struct {
	WormholeMessageCoins  []types.ObjectId // vector<Coin<SUI>>
	WormholeMessageAmount uint64
	Pool                  types.HexData
	DepositCoins          []types.ObjectId // vector<Coin<CoinType>>
	DepositAmount         uint64
}

type WithdrawArgs struct {
	Pool                  types.HexData
	DstChain              uint64
	WormholeMessageCoins  []types.ObjectId // vector<Coin<SUI>>
	WormholeMessageAmount uint64
	Amount                uint64
}

type BorrowArgs struct {
	Pool                  types.HexData
	DstChain              uint64
	WormholeMessageCoins  []types.ObjectId // vector<Coin<SUI>>
	WormholeMessageAmount uint64
	Amount                uint64
}

type RepayArgs struct {
	Pool                  types.HexData
	WormholeMessageCoins  []types.ObjectId // vector<Coin<SUI>>
	WormholeMessageAmount uint64
	RepayCoins            []types.ObjectId // vector<Coin<CoinType>>
	RepayAmount           uint64
}

type ContractConfig struct {
	LendingPortalPackageId     string
	ExternalInterfacePackageId string
	PoolManagerInfo            string
	PoolState                  string
	PriceOracle                string
	Storage                    string
	WormholeState              string
}

type innerContract struct {
	client *client.Client

	lendingPortalPackageId     *types.HexData
	externalInterfacePackageId *types.HexData
	poolManagerInfo            *types.HexData
	poolState                  *types.HexData
	priceOracle                *types.HexData
	storage                    *types.HexData
	wormholeState              *types.HexData
}

func NewContract(client *client.Client, config ContractConfig) (Contract, error) {
	contract := &innerContract{client: client}
	var err error
	if contract.lendingPortalPackageId, err = types.NewHexData(config.LendingPortalPackageId); err != nil {
		return nil, err
	}
	if contract.externalInterfacePackageId, err = types.NewHexData(config.ExternalInterfacePackageId); err != nil {
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
	return contract, nil
}

func (c *innerContract) Supply(ctx context.Context, signer types.Address, typeArgs []string, supplyArgs SupplyArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		*c.poolState,
		*c.wormholeState,
		supplyArgs.WormholeMessageCoins,
		supplyArgs.WormholeMessageAmount,
		supplyArgs.Pool,
		supplyArgs.DepositCoins,
		supplyArgs.DepositAmount,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.lendingPortalPackageId, "lending", "supply", typeArgs, args, callOptions.gas, callOptions.gasBudget)
	return resp, err
}

func (c *innerContract) Withdraw(ctx context.Context, signer types.Address, typeArgs []string, withdrawArgs WithdrawArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		withdrawArgs.Pool,
		*c.poolState,
		*c.wormholeState,
		withdrawArgs.DstChain,
		withdrawArgs.WormholeMessageCoins,
		withdrawArgs.WormholeMessageAmount,
		withdrawArgs.Amount,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.lendingPortalPackageId, "lending", "withdraw", typeArgs, args, callOptions.gas, callOptions.gasBudget)
	return resp, err
}

func (c *innerContract) Borrow(ctx context.Context, signer types.Address, typeArgs []string, borrowArgs BorrowArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		borrowArgs.Pool,
		*c.poolState,
		*c.wormholeState,
		borrowArgs.DstChain,
		borrowArgs.WormholeMessageCoins,
		borrowArgs.WormholeMessageAmount,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.lendingPortalPackageId, "lending", "borrow", typeArgs, args, callOptions.gas, callOptions.gasBudget)
	return resp, err
}

func (c *innerContract) Repay(ctx context.Context, signer types.Address, typeArgs []string, repayArgs RepayArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		repayArgs.Pool,
		*c.poolState,
		*c.wormholeState,
		repayArgs.WormholeMessageCoins,
		repayArgs.WormholeMessageAmount,
		repayArgs.RepayCoins,
		repayArgs.RepayAmount,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.lendingPortalPackageId, "lending", "repay", typeArgs, args, callOptions.gas, callOptions.gasBudget)
	return resp, err
}
