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
		PoolLiquidity      *big.Int
		DolaChainId        uint16
		DolaAddress        string   // hex string. 0x123
		PoolEquilibriumFee *big.Int // u256
		PoolWeight         *big.Int // u256
	}

	DolaUserAddress struct {
		DolaChainId uint16
		DolaAddress string // hex string. 0x123
	}

	DolaTokenPrice struct {
		Decimal    int // the decimal of price,
		DolaPoolId uint16
		Price      *big.Int
	}

	ReserveInfo struct {
		BorrowApy             int // 200 -> 200/10000=2.0%
		BorrowCoefficient     *big.Int
		CollateralCoefficient *big.Int
		Debt                  *big.Int // 100000000 -> 100000000/1e8 = 1
		Reserve               *big.Int // 100000000 -> 100000000/1e8 = 1
		SupplyApy             int      // 100 -> 100/10000=1.0%
		UtilizationRate       int      // 100 -> 100/10000=1.0%
		DolaPoolId            uint16
		Pools                 []PoolInfo
	}

	CollateralItem struct {
		CollateralAmount *big.Int
		CollateralValue  *big.Int
		DolaPoolId       uint16
		BorrowApy        int
		SupplyApy        int
	}

	UserLendingInfo struct {
		TotalCollateralValue *big.Int
		TotalDebtValue       *big.Int
		HealthFactor         *big.Int
		NetApy               int // profit_state decide positive or negtive
		TotalBorrowApy       int
		TotalSupplyApy       int

		CollateralInfos []CollateralItem
		DebtInfos       []DebtItem
	}

	DebtItem struct {
		DebtAmount *big.Int
		DebtValue  *big.Int
		DolaPoolId uint16
		BorrowApy  int
		SupplyApy  int
	}
)

func newDebtItem(info interface{}) (debtItem DebtItem, err error) {
	var b bool
	fields := info.(map[string]interface{})
	debtItem.BorrowApy, err = strconv.Atoi(fields["borrow_apy"].(string))
	if err != nil {
		return
	}
	debtItem.DebtAmount, b = new(big.Int).SetString(fields["debt_amount"].(string), 10)
	if !b {
		err = errors.New("parse debt_amount fail")
		return
	}
	debtItem.DebtValue, b = new(big.Int).SetString(fields["debt_value"].(string), 10)
	if !b {
		err = errors.New("parse debt_value fail")
		return
	}
	debtItem.DolaPoolId = uint16(fields["dola_pool_id"].(float64))
	debtItem.SupplyApy, err = strconv.Atoi(fields["supply_apy"].(string))
	if err != nil {
		return
	}
	return
}

func newCollateralItem(parsedJson interface{}) (collateral CollateralItem, err error) {
	var b bool
	fields := parsedJson.(map[string]interface{})
	collateral.BorrowApy, err = strconv.Atoi(fields["borrow_apy"].(string))
	if err != nil {
		return
	}
	collateral.CollateralAmount, b = new(big.Int).SetString(fields["collateral_amount"].(string), 10)
	if !b {
		err = errors.New("parse collateral_amount fail")
		return
	}
	collateral.CollateralValue, b = new(big.Int).SetString(fields["collateral_value"].(string), 10)
	if !b {
		err = errors.New("parse collateral_value fail")
		return
	}
	collateral.DolaPoolId = uint16(fields["dola_pool_id"].(float64))
	collateral.SupplyApy, err = strconv.Atoi(fields["supply_apy"].(string))
	if err != nil {
		return
	}
	return
}

func newReserveInfo(parsedJson interface{}) (reserveInfo ReserveInfo, err error) {
	var b bool
	fields := parsedJson.(map[string]interface{})
	reserveInfo.BorrowApy, err = strconv.Atoi(fields["borrow_apy"].(string))
	if err != nil {
		return
	}
	reserveInfo.BorrowCoefficient, b = new(big.Int).SetString(fields["borrow_coefficient"].(string), 10)
	if !b {
		return reserveInfo, errors.New("fail to parse borrow_coefficient")
	}
	reserveInfo.CollateralCoefficient, b = new(big.Int).SetString(fields["collateral_coefficient"].(string), 10)
	if !b {
		return reserveInfo, errors.New("fail to parse collateral_coefficient")
	}
	reserveInfo.Debt, b = new(big.Int).SetString(fields["debt"].(string), 10)
	if !b {
		return reserveInfo, errors.New("parse reserve failed")
	}
	reserveInfo.Reserve, b = new(big.Int).SetString(fields["reserve"].(string), 10)
	if !b {
		return reserveInfo, errors.New("parse reserve failed")
	}
	reserveInfo.SupplyApy, err = strconv.Atoi(fields["supply_apy"].(string))
	if err != nil {
		return
	}
	reserveInfo.UtilizationRate, err = strconv.Atoi(fields["utilization_rate"].(string))
	if err != nil {
		return
	}
	reserveInfo.DolaPoolId = uint16(fields["dola_pool_id"].(float64))
	poolsInfo := fields["pools"].([]interface{})
	pools := make([]PoolInfo, len(poolsInfo))
	for i := range poolsInfo {
		pools[i] = newPoolInfo(poolsInfo[i])
	}
	reserveInfo.Pools = pools

	return
}

func newDolaTokenPrice(priceInfo interface{}) DolaTokenPrice {
	fields := priceInfo.(map[string]interface{})
	price, _ := new(big.Int).SetString(fields["price"].(string), 10)
	return DolaTokenPrice{
		Decimal:    int(fields["decimal"].(float64)),
		DolaPoolId: uint16(fields["dola_pool_id"].(float64)),
		Price:      price,
	}
}

func newDolaUserAddress(info interface{}) DolaUserAddress {
	fields := info.(map[string]interface{})
	return DolaUserAddress{
		DolaChainId: uint16(fields["dola_chain_id"].(float64)),
		DolaAddress: newUserAddress(fields["dola_address"]),
	}
}

func newPoolInfo(info interface{}) PoolInfo {
	var poolInfo PoolInfo
	infoFields := info.(map[string]interface{})

	poolInfo.PoolLiquidity, _ = new(big.Int).SetString(infoFields["pool_liquidity"].(string), 10)

	poolAddress := infoFields["pool_address"].(map[string]interface{})
	poolInfo.DolaChainId = uint16(poolAddress["dola_chain_id"].(float64))
	poolInfo.DolaAddress = newDolaAddress(poolInfo.DolaChainId, poolAddress["dola_address"])

	poolInfo.PoolWeight, _ = new(big.Int).SetString(infoFields["pool_equilibrium_fee"].(string), 10)

	poolInfo.PoolEquilibriumFee, _ = new(big.Int).SetString(infoFields["pool_weight"].(string), 10)

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

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		tokenLiquidity := event.ParsedJson.(map[string]interface{})["token_liquidity"].(string)
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

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		tokenLiquidity := event.ParsedJson.(map[string]interface{})["token_liquidity"].(string)
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

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		tokenLiquidity := event.ParsedJson.(map[string]interface{})["pool_liquidity"].(string)
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

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		poolInfosData := event.ParsedJson.(map[string]interface{})["pool_infos"].([]interface{})
		poolInfos = make([]PoolInfo, len(poolInfosData))
		for i := range poolInfosData {
			poolInfos[i] = newPoolInfo(poolInfosData[i])
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

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		fields := event.ParsedJson.(map[string]interface{})
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

func (c *Contract) GetUserCollateral(ctx context.Context, signer types.Address, dolaUserId string, dolaPoolId uint16, callOptions CallOptions) (collateral CollateralItem, err error) {
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

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		collateral, err = newCollateralItem(event.ParsedJson)
		return nil
	})
	return
}

func (c *Contract) GetAllReserveInfo(ctx context.Context, signer types.Address, callOptions CallOptions) (reserveInfos []ReserveInfo, err error) {
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

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		fields := event.ParsedJson.(map[string]interface{})
		responseReserveInfos := fields["reserve_infos"].([]interface{})
		reserveInfos = make([]ReserveInfo, len(responseReserveInfos))
		for i := range responseReserveInfos {
			reserveInfos[i], err = newReserveInfo(responseReserveInfos[i])
			if err != nil {
				return err
			}
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

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		res, err := newReserveInfo(event.ParsedJson)
		if err != nil {
			return err
		}
		reserveInfo = &res
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

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		fields := event.ParsedJson.(map[string]interface{})
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

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		fields := event.ParsedJson.(map[string]interface{})
		userLendingInfo = &UserLendingInfo{}
		userLendingInfo.TotalCollateralValue, _ = new(big.Int).SetString(fields["total_collateral_value"].(string), 10)
		userLendingInfo.TotalDebtValue, _ = new(big.Int).SetString(fields["total_debt_value"].(string), 10)
		userLendingInfo.HealthFactor, _ = new(big.Int).SetString(fields["health_factor"].(string), 10)
		profitState := fields["profit_state"].(bool)
		userLendingInfo.NetApy, err = strconv.Atoi(fields["net_apy"].(string))
		if err != nil {
			return err
		}
		if !profitState {
			userLendingInfo.NetApy = -userLendingInfo.NetApy
		}
		userLendingInfo.TotalBorrowApy, err = strconv.Atoi(fields["total_borrow_apy"].(string))
		if err != nil {
			return err
		}
		userLendingInfo.TotalSupplyApy, err = strconv.Atoi(fields["total_supply_apy"].(string))
		if err != nil {
			return err
		}

		if fields["collateral_infos"] != "" {
			infos := fields["collateral_infos"].([]interface{})
			userLendingInfo.CollateralInfos = make([]CollateralItem, 0, len(infos))
			for _, info := range infos {
				collateralInfo, err := newCollateralItem(info)
				if err != nil {
					return err
				}

				userLendingInfo.CollateralInfos = append(userLendingInfo.CollateralInfos, collateralInfo)
			}
		}

		if fields["debt_infos"] != "" {
			infos := fields["debt_infos"].([]interface{})
			userLendingInfo.DebtInfos = make([]DebtItem, 0, len(infos))
			for _, info := range infos {
				debtItem, err := newDebtItem(info)
				if err != nil {
					return err
				}

				userLendingInfo.DebtInfos = append(userLendingInfo.DebtInfos, debtItem)
			}
		}

		return nil
	})
	return
}

func (c *Contract) GetOraclePrice(ctx context.Context, signer types.Address, dolaPoolId uint16, callOptions CallOptions) (dolaTokenPrice DolaTokenPrice, err error) {
	args := []any{
		*c.priceOracle,
		dolaPoolId,
	}

	tx, err := c.client.MoveCall(ctx, signer, *c.externalInterfacePackageId, "interfaces", "get_oracle_price", []string{}, args, callOptions.Gas, callOptions.GasBudget)
	if err != nil {
		return
	}

	effects, err := c.client.DryRunTransaction(ctx, tx)
	if err != nil {
		return
	}

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		dolaTokenPrice = newDolaTokenPrice(event.ParsedJson)
		return nil
	})
	return
}

func (c *Contract) GetAllOraclePrice(ctx context.Context, signer types.Address, callOptions CallOptions) (prices []DolaTokenPrice, err error) {
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

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		fields := event.ParsedJson.(map[string]interface{})
		tokenPrices := fields["token_prices"].([]interface{})
		prices = make([]DolaTokenPrice, len(tokenPrices))
		for i, item := range tokenPrices {
			prices[i] = newDolaTokenPrice(item)
		}
		return nil
	})
	return
}

// GetDolaUserId return dola_user_id for (dola_chain_id, address) pair
// if not exist, an error return
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

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		fields := event.ParsedJson.(map[string]interface{})
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

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		fields := event.ParsedJson.(map[string]interface{})
		addresses := fields["dola_user_addresses"].([]interface{})
		dolaUserAddresses = make([]DolaUserAddress, len(addresses))
		for i, item := range addresses {
			dolaUserAddresses[i] = newDolaUserAddress(item)
		}
		return nil
	})
	return
}

func (c *Contract) GetUserHealthFactor(ctx context.Context, signer types.Address, dolaUserId string, callOptions CallOptions) (healthFactor *big.Int, err error) {
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

	err = parseLastEvent(effects, func(event types.SuiEvent) error {
		fields := event.ParsedJson.(map[string]interface{})
		healthFactor, _ = new(big.Int).SetString(fields["health_factor"].(string), 10)
		return nil
	})
	return
}

func parseLastEvent(dryRunResponse *types.DryRunTransactionBlockResponse, f func(event types.SuiEvent) error) (err error) {
	if dryRunResponse.Effects.Status.Status != "success" {
		return errors.New(dryRunResponse.Effects.Status.Error)
	}

	if len(dryRunResponse.Events) == 0 {
		return errors.New("invalid events")
	}

	defer func() {
		if merr := recover(); merr != nil {
			err = errors.New("event parse failed")
		}
	}()

	return f(dryRunResponse.Events[len(dryRunResponse.Events)-1])
}
