package gosuilending

import (
	"errors"
	"math/big"
	"strconv"

	"github.com/coming-chat/go-sui/v2/types"
)

const (
	CallTypeSupply    = 0
	CallTypeWithdraw  = 1
	CallTypeBorrow    = 2
	CallTypeRepay     = 3
	CallTypeLiquidite = 4
	CallTypeBinding   = 5
	CallTypeUnbinding = 6
)

type (
	EventHeader struct {
		Timestamp uint64
		TxDigest  string
		Id        types.EventId
	}

	MoveEventHeader struct {
		EventHeader
		PackageId         string
		TransactionModule string
		Sender            string
		Type              string
		Bcs               string
	}

	// Sui local chain lending event
	LocalLendingEvent struct {
		MoveEventHeader MoveEventHeader
		Nonce           uint64
		Sender          string
		DolaPoolAddress []byte
		Amount          uint64
		CallType        int
	}

	// Sui chain call remote borrow/withdraw
	LendingPortalEvent struct {
		MoveEventHeader MoveEventHeader
		Nonce           uint64
		Sender          string
		DolaPoolAddress []byte
		SourceChainId   uint16
		DstChainId      uint16
		Receiver        []byte
		Amount          uint64
		CallType        int
	}

	// bridge send event, from other chain transaction
	LendingCoreEvent struct {
		MoveEventHeader MoveEventHeader
		Nonce           uint64
		SenderUserId    uint64
		SourceChainId   uint16
		DstChainId      uint16
		DolaPoolId      uint16
		Receiver        []byte
		Amount          uint64
		LiquidateUserId uint64
		CallType        int
	}

	// core event
	LendingCoreExecuteEvent struct {
		MoveEventHeader MoveEventHeader
		UserId          uint64
		Amount          *big.Int
		PoolId          uint16
		ViolatorId      uint64
		CallType        int
	}
)

func ParseLendingCoreExecuteEvent(event types.SuiEvent) (result *LendingCoreExecuteEvent, err error) {
	defer func() {
		if tmpErr := recover(); tmpErr != nil {
			err = tmpErr.(error)
		}
	}()
	result = &LendingCoreExecuteEvent{}
	fields := event.ParsedJson.(map[string]interface{})
	if result.MoveEventHeader, err = parseMoveEventHeader(event); err != nil {
		return
	}
	if result.UserId, err = strconv.ParseUint(fields["user_id"].(string), 10, 64); err != nil {
		return
	}
	if result.ViolatorId, err = strconv.ParseUint(fields["violator_id"].(string), 10, 64); err != nil {
		return
	}
	result.PoolId = uint16(fields["pool_id"].(float64))
	result.CallType = int(fields["call_type"].(float64))

	// contract use u256 for compute, all amount is in u64
	var b bool
	result.Amount, b = big.NewInt(0).SetString(fields["amount"].(string), 0)
	if !b {
		err = errors.New("parse amount failed")
	}
	return
}

func ParseLendingCoreEvent(event types.SuiEvent) (result *LendingCoreEvent, err error) {
	defer func() {
		if tmpErr := recover(); tmpErr != nil {
			err = tmpErr.(error)
		}
	}()
	result = &LendingCoreEvent{}
	fields := event.ParsedJson.(map[string]interface{})
	if result.MoveEventHeader, err = parseMoveEventHeader(event); err != nil {
		return
	}
	result.SourceChainId = uint16(fields["source_chain_id"].(float64))
	result.DstChainId = uint16(fields["dst_chain_id"].(float64))
	result.Receiver = parseByteSlice(fields["receiver"].([]interface{}))
	if result.Nonce, err = strconv.ParseUint(fields["nonce"].(string), 10, 64); err != nil {
		return
	}
	// contract use u256 for compute, all amount is in u64
	if result.Amount, err = strconv.ParseUint(fields["amount"].(string), 10, 64); err != nil {
		return
	}
	result.CallType = int(fields["call_type"].(float64))
	if result.SenderUserId, err = strconv.ParseUint(fields["sender_user_id"].(string), 10, 64); err != nil {
		return
	}
	result.DolaPoolId = uint16(fields["dola_pool_id"].(float64))
	if result.LiquidateUserId, err = strconv.ParseUint(fields["liquidate_user_id"].(string), 10, 64); err != nil {
		return
	}
	return
}

func ParseLendingPortalEvent(event types.SuiEvent) (result *LendingPortalEvent, err error) {
	defer func() {
		if tmpErr := recover(); tmpErr != nil {
			err = tmpErr.(error)
		}
	}()
	result = &LendingPortalEvent{}
	fields := event.ParsedJson.(map[string]interface{})
	if result.MoveEventHeader, err = parseMoveEventHeader(event); err != nil {
		return
	}
	result.SourceChainId = uint16(fields["source_chain_id"].(float64))
	result.DstChainId = uint16(fields["dst_chain_id"].(float64))
	result.Receiver = parseByteSlice(fields["receiver"].([]interface{}))
	if result.Nonce, err = strconv.ParseUint(fields["nonce"].(string), 10, 64); err != nil {
		return
	}
	if result.Amount, err = strconv.ParseUint(fields["amount"].(string), 10, 64); err != nil {
		return
	}
	result.Sender = fields["sender"].(string)
	result.CallType = int(fields["call_type"].(float64))
	result.DolaPoolAddress = parseByteSlice(fields["dola_pool_address"].([]interface{}))
	return
}

func ParseLocalLendingEvent(event types.SuiEvent) (result *LocalLendingEvent, err error) {
	defer func() {
		if tmpErr := recover(); tmpErr != nil {
			err = tmpErr.(error)
		}
	}()
	result = &LocalLendingEvent{}
	fields := event.ParsedJson.(map[string]interface{})
	if result.MoveEventHeader, err = parseMoveEventHeader(event); err != nil {
		return
	}
	if result.Nonce, err = strconv.ParseUint(fields["nonce"].(string), 10, 64); err != nil {
		return
	}
	if result.Amount, err = strconv.ParseUint(fields["amount"].(string), 10, 64); err != nil {
		return
	}
	result.Sender = fields["sender"].(string)
	result.CallType = int(fields["call_type"].(float64))
	result.DolaPoolAddress = parseByteSlice(fields["dola_pool_address"].([]interface{}))
	return
}

func parseByteSlice(s []interface{}) []byte {
	result := make([]byte, len(s))
	for i, v := range s {
		result[i] = byte(v.(float64))
	}
	return result
}

func parseMoveEventHeader(event types.SuiEvent) (result MoveEventHeader, err error) {
	if result.EventHeader, err = parseEventHeader(event); err != nil {
		return
	}

	result.PackageId = event.PackageId.String()
	result.TransactionModule = event.TransactionModule
	result.Sender = event.Sender.String()
	result.Type = event.Type
	result.Bcs = event.Bcs
	return
}

func parseEventHeader(event types.SuiEvent) (result EventHeader, err error) {
	result.Timestamp = event.TimestampMs.Uint64()
	result.TxDigest = event.Id.TxDigest.String()
	result.Id = event.Id
	return
}
