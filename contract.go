package gosuilending

import (
	"context"

	"github.com/coming-chat/go-sui/v2/client"
	"github.com/coming-chat/go-sui/v2/sui_types"
	"github.com/coming-chat/go-sui/v2/types"
)

type CallOptions struct {
	Gas       *sui_types.ObjectID
	GasBudget uint64
}

type SupplyArgs struct {
	Pool          sui_types.ObjectID
	DepositCoins  []*sui_types.ObjectID // vector<Coin<CoinType>>
	DepositAmount string
}

type WithdrawArgs struct {
	Pool           sui_types.ObjectID
	Receiver       string
	DstChain       string
	Amount         string
	RelayFeeCoins  []*sui_types.ObjectID
	RelayFeeAmount string
}

type BorrowArgs struct {
	Pool     sui_types.ObjectID
	Receiver string
	DstChain string
	Amount   string
}

type RepayArgs struct {
	Pool        sui_types.ObjectID
	RepayCoins  []*sui_types.ObjectID // vector<Coin<CoinType>>
	RepayAmount string
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
	CoreState                  string
	LendingPortal              string
	LendingCore                string
	Clock                      string
	PoolApproval               string
}

type Contract struct {
	client *client.Client

	lendingPortalPackageId     *sui_types.ObjectID
	externalInterfacePackageId *sui_types.ObjectID
	bridgePoolPackageId        *sui_types.ObjectID
	poolManagerInfo            *sui_types.ObjectID
	poolState                  *sui_types.ObjectID
	priceOracle                *sui_types.ObjectID
	storage                    *sui_types.ObjectID
	wormholeState              *sui_types.ObjectID
	userManagerInfo            *sui_types.ObjectID
	coreState                  *sui_types.ObjectID
	lendingPortal              *sui_types.ObjectID
	clock                      *sui_types.ObjectID
	poolApproval               *sui_types.ObjectID
}

func NewContract(client *client.Client, config ContractConfig) (*Contract, error) {
	contract := &Contract{client: client}
	var err error
	if contract.lendingPortalPackageId, err = sui_types.NewObjectIdFromHex(config.LendingPortalPackageId); err != nil {
		return nil, err
	}
	if contract.externalInterfacePackageId, err = sui_types.NewObjectIdFromHex(config.ExternalInterfacePackageId); err != nil {
		return nil, err
	}
	if contract.bridgePoolPackageId, err = sui_types.NewObjectIdFromHex(config.BridgePoolPackageId); err != nil {
		return nil, err
	}
	if contract.poolManagerInfo, err = sui_types.NewObjectIdFromHex(config.PoolManagerInfo); err != nil {
		return nil, err
	}
	if contract.poolState, err = sui_types.NewObjectIdFromHex(config.PoolState); err != nil {
		return nil, err
	}
	if contract.priceOracle, err = sui_types.NewObjectIdFromHex(config.PriceOracle); err != nil {
		return nil, err
	}
	if contract.storage, err = sui_types.NewObjectIdFromHex(config.Storage); err != nil {
		return nil, err
	}
	if contract.wormholeState, err = sui_types.NewObjectIdFromHex(config.WormholeState); err != nil {
		return nil, err
	}
	if contract.userManagerInfo, err = sui_types.NewObjectIdFromHex(config.UserManagerInfo); err != nil {
		return nil, err
	}
	if contract.coreState, err = sui_types.NewObjectIdFromHex(config.CoreState); err != nil {
		return nil, err
	}
	if contract.lendingPortal, err = sui_types.NewObjectIdFromHex(config.LendingPortal); err != nil {
		return nil, err
	}
	if contract.clock, err = sui_types.NewObjectIdFromHex(config.Clock); err != nil {
		return nil, err
	}
	if contract.poolApproval, err = sui_types.NewObjectIdFromHex(config.PoolApproval); err != nil {
		return nil, err
	}
	return contract, nil
}

func (c *Contract) Supply(ctx context.Context, signer sui_types.SuiAddress, typeArgs []string, supplyArgs SupplyArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		*c.storage,
		*c.priceOracle,
		*c.clock,
		*c.lendingPortal,
		*c.userManagerInfo,
		*c.poolManagerInfo,
		supplyArgs.Pool,
		supplyArgs.DepositCoins,
		supplyArgs.DepositAmount,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.lendingPortalPackageId, "lending", "supply", typeArgs, args, callOptions.Gas, types.NewSafeSuiBigInt(callOptions.GasBudget))
	return resp, err
}

func (c *Contract) WithdrawLocal(ctx context.Context, signer sui_types.SuiAddress, typeArgs []string, withdrawArgs WithdrawArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		*c.storage,
		*c.priceOracle,
		*c.clock,
		*c.lendingPortal,
		*c.poolManagerInfo,
		*c.userManagerInfo,
		withdrawArgs.Pool,
		withdrawArgs.Amount,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.lendingPortalPackageId, "lending", "withdraw_local", typeArgs, args, callOptions.Gas, types.NewSafeSuiBigInt(callOptions.GasBudget))
	return resp, err
}

func (c *Contract) WithdrawRemote(ctx context.Context, signer sui_types.SuiAddress, typeArgs []string, withdrawArgs WithdrawArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		*c.storage,
		*c.priceOracle,
		*c.clock,
		*c.coreState,
		*c.lendingPortal,
		*c.wormholeState,
		*c.poolManagerInfo,
		*c.userManagerInfo,
		withdrawArgs.Pool,
		withdrawArgs.Receiver,
		withdrawArgs.DstChain,
		withdrawArgs.Amount,
		withdrawArgs.RelayFeeCoins,
		withdrawArgs.RelayFeeAmount,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.lendingPortalPackageId, "lending", "withdraw_local", typeArgs, args, callOptions.Gas, types.NewSafeSuiBigInt(callOptions.GasBudget))
	return resp, err
}

func (c *Contract) BorrowLocal(ctx context.Context, signer sui_types.SuiAddress, typeArgs []string, borrowArgs BorrowArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		*c.poolApproval,
		*c.storage,
		*c.priceOracle,
		*c.clock,
		*c.lendingPortal,
		*c.poolManagerInfo,
		*c.userManagerInfo,
		borrowArgs.Pool,
		borrowArgs.Amount,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.lendingPortalPackageId, "lending", "borrow_local", typeArgs, args, callOptions.Gas, types.NewSafeSuiBigInt(callOptions.GasBudget))
	return resp, err
}

func (c *Contract) Repay(ctx context.Context, signer sui_types.SuiAddress, typeArgs []string, repayArgs RepayArgs, callOptions CallOptions) (*types.TransactionBytes, error) {
	args := []any{
		*c.storage,
		*c.priceOracle,
		*c.lendingPortal,
		*c.userManagerInfo,
		*c.poolManagerInfo,
		repayArgs.Pool,
		repayArgs.RepayCoins,
		repayArgs.RepayAmount,
	}
	resp, err := c.client.MoveCall(ctx, signer, *c.lendingPortalPackageId, "lending", "repay", typeArgs, args, callOptions.Gas, types.NewSafeSuiBigInt(callOptions.GasBudget))
	return resp, err
}
