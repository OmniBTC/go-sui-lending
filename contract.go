package gosuilending

import (
	"context"
	"errors"
	"math/big"

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

type ReserveInfo struct {
	BorrowApy       int      // 200 -> 200/10000=2.0%
	Debt            *big.Int // 100000000 -> 100000000/1e8 = 1
	Reserve         *big.Int // 100000000 -> 100000000/1e8 = 1
	SupplyApy       int      // 100 -> 100/10000=1.0%
	UtilizationRate int      // 100 -> 100/10000=1.0%
	DolaPoolId      uint16
}

type UserLendingInfo struct {
	TotalCollateralValue *big.Int
	TotalDebtValue       *big.Int
	CollateralInfos      []CollateralItem
	DebtInfos            []DebtItem
}

type CollateralItem struct {
	Type             string
	CollateralAmount *big.Int
	CollateralValue  *big.Int
	DolaPoolId       uint16
}

type DebtItem struct {
	Type       string
	DebtAmount *big.Int
	DebtValue  *big.Int
	DolaPoolId uint16
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

type Contract struct {
	client *client.Client

	lendingPortalPackageId     *types.HexData
	externalInterfacePackageId *types.HexData
	poolManagerInfo            *types.HexData
	poolState                  *types.HexData
	priceOracle                *types.HexData
	storage                    *types.HexData
	wormholeState              *types.HexData
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

func (c *Contract) GetDolaTokenLiquidity(ctx context.Context, signer types.Address, dolaPoolId uint16, callOptions CallOptions) (liquidity *big.Int, err error) {
	args := []any{
		*c.poolManagerInfo,
		dolaPoolId,
	}

	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_dola_token_liquidity", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	err = parseLastEvent(effects, func(event types.Event) error {
		tokenLiquidity := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})["token_liquidity"].(string)
		liquidity, _ = big.NewInt(0).SetString(tokenLiquidity, 10)
		return nil
	})
	return
}

func (c *Contract) GetAppTokenLiquidity(ctx context.Context, signer types.Address, appId uint16, dolaPoolId uint16, callOptions CallOptions) (liquidity *big.Int, err error) {
	args := []any{
		*c.poolManagerInfo,
		appId,
		dolaPoolId,
	}
	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_app_token_liquidity", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	err = parseLastEvent(effects, func(event types.Event) error {
		tokenLiquidity := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})["token_liquidity"].(string)
		var b bool
		liquidity, b = big.NewInt(0).SetString(tokenLiquidity, 10)
		if !b {
			return errors.New("event parse failed: tokenLiquidity is not integer")
		}
		return nil
	})
	return
}

func (c *Contract) GetUserTokenDebt(ctx context.Context, signer types.Address, dolaUserId string, dolaPoolId uint16, callOptions CallOptions) (debtAmount *big.Int, debtValue *big.Int, err error) {
	args := []any{
		*c.storage,
		*c.priceOracle,
		dolaUserId,
		dolaPoolId,
	}
	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_user_token_debt", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	err = parseLastEvent(effects, func(event types.Event) error {
		fields := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})
		debtAmountStr := fields["debt_amount"].(string)
		debtValueStr := fields["debt_value"].(string)
		var ok bool
		debtAmount, ok = big.NewInt(0).SetString(debtAmountStr, 10)
		if !ok {
			return errors.New("debtAmount parse error")
		}
		debtValue, ok = big.NewInt(0).SetString(debtValueStr, 10)
		if !ok {
			return errors.New("debtValue parse error")
		}
		return nil
	})
	return
}

func (c *Contract) GetUserCollateral(ctx context.Context, signer types.Address, dolaUserId string, dolaPoolId uint16, callOptions CallOptions) (collateralAmount *big.Int, collateralValue *big.Int, err error) {
	args := []any{
		*c.storage,
		*c.priceOracle,
		dolaUserId,
		dolaPoolId,
	}
	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_user_collateral", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	err = parseLastEvent(effects, func(event types.Event) error {
		fields := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})
		collateralAmountStr := fields["collateral_amount"].(string)
		collateralValueStr := fields["collateral_value"].(string)
		var ok bool
		collateralAmount, ok = big.NewInt(0).SetString(collateralAmountStr, 10)
		if !ok {
			return errors.New("collateralAmount parse error")
		}
		collateralValue, ok = big.NewInt(0).SetString(collateralValueStr, 10)
		if !ok {
			return errors.New("collateralValue parse error")
		}
		return nil
	})
	return
}

func (c *Contract) GetReserveInfo(ctx context.Context, signer types.Address, dolaPoolId uint16, callOptions CallOptions) (reserveInfo *ReserveInfo, err error) {
	args := []any{
		*c.poolManagerInfo,
		*c.storage,
		dolaPoolId,
	}
	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_reserve_info", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	err = parseLastEvent(effects, func(event types.Event) error {
		reserveInfo = &ReserveInfo{}
		var b bool
		fields := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})
		reserveInfo.BorrowApy = int(fields["borrow_apy"].(float64))
		reserveInfo.Debt, b = new(big.Int).SetString(fields["debt"].(string), 10)
		if !b {
			return errors.New("parse reserve failed")
		}
		reserveInfo.Reserve, b = new(big.Int).SetString(fields["reserve"].(string), 10)
		if !b {
			return errors.New("parse reserve failed")
		}
		reserveInfo.SupplyApy = int(fields["supply_apy"].(float64))
		reserveInfo.UtilizationRate = int(fields["utilization_rate"].(float64))
		reserveInfo.DolaPoolId = uint16(fields["dola_pool_id"].(float64))
		return nil
	})
	return
}

func (c *Contract) GetUserAllowedBorrow(ctx context.Context, signer types.Address, dolaUserId string, borrowPoolId uint16, callOptions CallOptions) (amount *big.Int, err error) {
	args := []any{
		*c.poolManagerInfo,
		*c.storage,
		*c.priceOracle,
		dolaUserId,
		borrowPoolId,
	}
	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_user_allowed_borrow", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	err = parseLastEvent(effects, func(event types.Event) error {
		fields := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})
		amount, _ = big.NewInt(0).SetString(fields["borrow_amount"].(string), 10)
		if amount.Cmp(big.NewInt(0)) == 0 {
			if fields["reason"] != "" && fields["reason"] != nil {
				return errors.New(fields["reason"].(string))
			}
		}
		return nil
	})
	return
}

func (c *Contract) GetUserLendingInfo(ctx context.Context, signer types.Address, dolaUserId string, callOptions CallOptions) (userLendingInfo *UserLendingInfo, err error) {
	args := []any{
		*c.storage,
		*c.priceOracle,
		dolaUserId,
	}
	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_user_lending_info", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	err = parseLastEvent(effects, func(event types.Event) error {
		fields := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})
		userLendingInfo = &UserLendingInfo{}
		userLendingInfo.TotalCollateralValue = new(big.Int).SetUint64(uint64(fields["total_collateral_value"].(float64)))
		userLendingInfo.TotalDebtValue = new(big.Int).SetUint64(uint64(fields["total_debt_value"].(float64)))

		if fields["collateral_infos"] != "" {
			infos := fields["collateral_infos"].([]interface{})
			userLendingInfo.CollateralInfos = make([]CollateralItem, 0, len(infos))
			for _, info := range infos {
				infoMap := info.(map[string]interface{})
				innerFields := infoMap["fields"].(map[string]interface{})
				dolaPoolId := uint16(innerFields["dola_pool_id"].(float64))
				amount, _ := new(big.Int).SetString(innerFields["collateral_amount"].(string), 10)
				value, _ := new(big.Int).SetString(innerFields["collateral_value"].(string), 10)

				userLendingInfo.CollateralInfos = append(userLendingInfo.CollateralInfos, CollateralItem{
					Type:             infoMap["type"].(string),
					CollateralAmount: amount,
					CollateralValue:  value,
					DolaPoolId:       dolaPoolId,
				})
			}
		}

		if fields["debt_infos"] != "" {
			infos := fields["debt_infos"].([]interface{})
			userLendingInfo.DebtInfos = make([]DebtItem, 0, len(infos))
			for _, info := range infos {
				infoMap := info.(map[string]interface{})
				innerFields := infoMap["fields"].(map[string]interface{})
				dolaPoolId := uint16(innerFields["dola_pool_id"].(float64))
				amount, _ := new(big.Int).SetString(innerFields["debt_amount"].(string), 10)
				value, _ := new(big.Int).SetString(innerFields["debt_value"].(string), 10)

				userLendingInfo.DebtInfos = append(userLendingInfo.DebtInfos, DebtItem{
					Type:       infoMap["type"].(string),
					DebtAmount: amount,
					DebtValue:  value,
					DolaPoolId: dolaPoolId,
				})
			}
		}

		return nil
	})
	return
}

func parseLastEvent(effects *types.TransactionEffects, f func(event types.Event) error) (err error) {
	if effects.Status.Status != "success" {
		return errors.New(effects.Status.Error)
	}

	if len(effects.Events) == 0 {
		return errors.New("invalid events")
	}

	defer func() {
		if merr := recover(); merr != nil {
			err = errors.New("event parse failed")
		}
	}()

	return f(effects.Events[len(effects.Events)-1])
}
