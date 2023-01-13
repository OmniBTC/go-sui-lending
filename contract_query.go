package gosuilending

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"
	"strconv"
	"strings"

	"github.com/coming-chat/go-sui/types"
)

type (
	PoolInfo struct {
		Type            string
		PoolLiquidity   *big.Int
		PoolAddressType string
		DolaChainId     uint16
		DolaAddress     string // hex string. 0x123
	}

	DolaUserAddress struct {
		DolaChainId uint16
		DolaAddress string // hex string. 0x123
	}
)

func newDolaUserAddress(info interface{}) DolaUserAddress {
	mapInfo := info.(map[string]interface{})
	fields := mapInfo["fields"].(map[string]interface{})
	return DolaUserAddress{
		DolaChainId: uint16(fields["dola_chain_id"].(float64)),
		DolaAddress: newUserAddress(fields["dola_address"]),
	}
}

func newPoolInfo(info interface{}) PoolInfo {
	var poolInfo PoolInfo
	mapInfo := info.(map[string]interface{})
	poolInfo.Type = mapInfo["type"].(string)

	infoFields := mapInfo["fields"].(map[string]interface{})
	poolInfo.PoolLiquidity, _ = new(big.Int).SetString(infoFields["pool_liquidity"].(string), 10)

	poolAddress := infoFields["pool_address"].(map[string]interface{})
	poolInfo.PoolAddressType = poolAddress["type"].(string)

	poolAddressField := poolAddress["fields"].(map[string]interface{})
	poolInfo.DolaChainId = uint16(poolAddressField["dola_chain_id"].(float64))
	poolInfo.DolaAddress = newDolaAddress(poolInfo.DolaChainId, poolAddressField["dola_address"])
	return poolInfo
}

// newDolaAddress 把 []float64 转成 []byte 并转化成 hex string
func newDolaAddress(dolaChainId uint16, data interface{}) string {
	arrData := data.([]interface{})
	u8arr := make([]byte, len(arrData))
	for i := range arrData {
		u8arr[i] = byte(arrData[i].(float64))
	}
	switch dolaChainId {
	case 0, 1:
		return "0x" + strings.TrimPrefix(string(u8arr), "0x")
	default:
		return "0x" + hex.EncodeToString(u8arr)
	}
}

func newUserAddress(data interface{}) string {
	return newDolaAddress(10, data)
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

// GetPoolLiquidity return a pool liquidity on a chain
func (c *Contract) GetPoolLiquidity(ctx context.Context, signer types.Address, dolaChainId uint16, poolAddress string, callOptions CallOptions) (liquidity *big.Int, err error) {
	args := []any{
		*c.poolManagerInfo,
		dolaChainId,
		poolAddress,
	}
	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_pool_liquidity", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	err = parseLastEvent(effects, func(event types.Event) error {
		tokenLiquidity := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})["pool_liquidity"].(string)
		var b bool
		liquidity, b = big.NewInt(0).SetString(tokenLiquidity, 10)
		if !b {
			return errors.New("event parse failed: tokenLiquidity is not integer")
		}
		return nil
	})
	return
}

// GetAllPoolLiquidity return all chain liquidity of a dola pool
func (c *Contract) GetAllPoolLiquidity(ctx context.Context, signer types.Address, dolaPoolId uint16, callOptions CallOptions) (poolInfos []PoolInfo, err error) {
	args := []any{
		*c.poolManagerInfo,
		dolaPoolId,
	}
	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_all_pool_liquidity", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	err = parseLastEvent(effects, func(event types.Event) error {
		poolInfosData := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})["pool_infos"].([]interface{})
		poolInfos = make([]PoolInfo, len(poolInfosData))
		for i := range poolInfosData {
			poolInfos[i] = newPoolInfo(poolInfosData[i])
		}
		return nil
	})
	return
}

func (c *Contract) GetUserAllDebt(ctx context.Context, signer types.Address, dolaUserId string, callOptions CallOptions) (err error) {
	args := []any{
		*c.storage,
		dolaUserId,
	}
	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_user_all_debt", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	//todo parse user all debt
	err = parseLastEvent(effects, func(event types.Event) error {
		// fields := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})
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

func (c *Contract) GetUserAllCollateral(ctx context.Context, signer types.Address, dolaUserId string, callOptions CallOptions) (err error) {
	args := []any{
		*c.storage,
		dolaUserId,
	}
	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_user_all_collateral", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	//todo parse user all debt
	err = parseLastEvent(effects, func(event types.Event) error {
		// fields := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})
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

func (c *Contract) GetAllReserveInfo(ctx context.Context, signer types.Address, callOptions CallOptions) (err error) {
	args := []any{
		*c.poolManagerInfo,
		*c.storage,
	}
	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_all_reserve_info", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	// todo parse event
	err = parseLastEvent(effects, func(event types.Event) error {
		// fields := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})
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
		reserveInfo.BorrowApy, err = strconv.Atoi(fields["borrow_apy"].(string))
		if err != nil {
			return err
		}
		reserveInfo.Debt, b = new(big.Int).SetString(fields["debt"].(string), 10)
		if !b {
			return errors.New("parse reserve failed")
		}
		reserveInfo.Reserve, b = new(big.Int).SetString(fields["reserve"].(string), 10)
		if !b {
			return errors.New("parse reserve failed")
		}
		reserveInfo.SupplyApy, err = strconv.Atoi(fields["supply_apy"].(string))
		if err != nil {
			return err
		}
		reserveInfo.UtilizationRate, err = strconv.Atoi(fields["utilization_rate"].(string))
		if err != nil {
			return err
		}

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
		userLendingInfo.TotalCollateralValue, _ = new(big.Int).SetString(fields["total_collateral_value"].(string), 10)
		userLendingInfo.TotalDebtValue, _ = new(big.Int).SetString(fields["total_debt_value"].(string), 10)
		userLendingInfo.HealthFactor, _ = new(big.Int).SetString(fields["health_factor"].(string), 10)

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

func (c *Contract) GetOraclePrice(ctx context.Context, signer types.Address, dolaChainId uint16, callOptions CallOptions) (err error) {
	args := []any{
		*c.priceOracle,
		dolaChainId,
	}

	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_oracle_price", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	// todo parse event
	err = parseLastEvent(effects, func(event types.Event) error {
		// fields := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})
		// userId = fields["dola_user_id"].(string)
		return nil
	})
	return
}

func (c *Contract) GetAllOraclePrice(ctx context.Context, signer types.Address, callOptions CallOptions) (err error) {
	args := []any{
		*c.storage,
		*c.priceOracle,
	}

	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_all_oracle_price", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	// todo parse event
	err = parseLastEvent(effects, func(event types.Event) error {
		// fields := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})
		// userId = fields["dola_user_id"].(string)
		return nil
	})
	return
}

func (c *Contract) GetDolaUserId(ctx context.Context, signer types.Address, dolaChainId uint16, user string, callOptions CallOptions) (userId string, err error) {
	args := []any{
		*c.userManagerInfo,
		dolaChainId,
		user,
	}
	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_dola_user_id", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	err = parseLastEvent(effects, func(event types.Event) error {
		fields := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})
		userId = fields["dola_user_id"].(string)
		return nil
	})
	return
}

func (c *Contract) GetDolaUserAddresses(ctx context.Context, signer types.Address, dolaUserId string, callOptions CallOptions) (dolaUserAddresses []DolaUserAddress, err error) {
	args := []any{
		*c.userManagerInfo,
		dolaUserId,
	}
	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_dola_user_addresses", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	err = parseLastEvent(effects, func(event types.Event) error {
		fields := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})
		addresses := fields["dola_user_addresses"].([]interface{})
		dolaUserAddresses = make([]DolaUserAddress, len(addresses))
		for i, item := range addresses {
			dolaUserAddresses[i] = newDolaUserAddress(item)
		}
		return nil
	})
	return
}

func (c *Contract) GetUserHealthFactor(ctx context.Context, signer types.Address, dolaUserId string, callOptions CallOptions) (healthFactor string, err error) {
	args := []any{
		*c.storage,
		*c.priceOracle,
		dolaUserId,
	}
	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_user_health_factor", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	// todo parse user health factor
	err = parseLastEvent(effects, func(event types.Event) error {
		fields := event.(map[string]interface{})["moveEvent"].(map[string]interface{})["fields"].(map[string]interface{})
		healthFactor = fields["healthFactor"].(string)
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
