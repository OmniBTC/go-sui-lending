package gosuilending

import (
	"errors"
	"strconv"

	"github.com/coming-chat/go-sui/types"
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
		TxDigest  types.Base64Data
		Id        types.EventID
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
		SourceChainId   uint16
		DstChainId      uint16
		PoolAddress     []byte
		Receiver        []byte
		Amount          uint64
		CallType        int
	}
)

func ParseLendingCoreEvent(event interface{}) (result *LendingCoreEvent, err error) {
	defer func() {
		if tmpErr := recover(); tmpErr != nil {
			err = tmpErr.(error)
		}
	}()
	result = &LendingCoreEvent{}
	eventMap := event.(map[string]interface{})
	var fields map[string]interface{}
	if result.MoveEventHeader, fields, err = parseMoveEventHeader(eventMap); err != nil {
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
	result.CallType = int(fields["call_type"].(float64))
	result.PoolAddress = parseByteSlice(fields["pool_address"].([]interface{}))
	return
}

func ParseLendingPortalEvent(event interface{}) (result *LendingPortalEvent, err error) {
	defer func() {
		if tmpErr := recover(); tmpErr != nil {
			err = tmpErr.(error)
		}
	}()
	result = &LendingPortalEvent{}
	eventMap := event.(map[string]interface{})
	var fields map[string]interface{}
	if result.MoveEventHeader, fields, err = parseMoveEventHeader(eventMap); err != nil {
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

func ParseLocalLendingEvent(event interface{}) (result *LocalLendingEvent, err error) {
	defer func() {
		if tmpErr := recover(); tmpErr != nil {
			err = tmpErr.(error)
		}
	}()
	result = &LocalLendingEvent{}
	eventMap := event.(map[string]interface{})
	var fields map[string]interface{}
	if result.MoveEventHeader, fields, err = parseMoveEventHeader(eventMap); err != nil {
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

func parseMoveEventHeader(event map[string]interface{}) (result MoveEventHeader, fields map[string]interface{}, err error) {
	if result.EventHeader, err = parseEventHeader(event); err != nil {
		return
	}

	moveEvent := event["event"].(map[string]interface{})["moveEvent"].(map[string]interface{})
	fields = moveEvent["fields"].(map[string]interface{})

	result.PackageId = moveEvent["packageId"].(string)
	result.TransactionModule = moveEvent["transactionModule"].(string)
	result.Sender = moveEvent["sender"].(string)
	result.Type = moveEvent["type"].(string)
	result.Bcs = moveEvent["bcs"].(string)
	return
}

func parseEventHeader(event map[string]interface{}) (result EventHeader, err error) {
	result.Timestamp = uint64(event["timestamp"].(float64))
	if d, err := types.NewBase64Data(event["txDigest"].(string)); err != nil {
		return result, err
	} else {
		result.TxDigest = *d
	}
	if id, ok := event["id"]; ok {
		idMap := id.(map[string]interface{})
		txd, err := types.NewBase64Data(idMap["txDigest"].(string))
		if err != nil {
			return result, err
		}
		result.Id.TxDigest = *txd
		result.Id.EventSeq = int64(idMap["eventSeq"].(float64))
	} else {
		err = errors.New("event no id field")
	}
	return
}
