package ethereumRPCProvider

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/AKACoder/EthereumRPCShell/common/constants"
	"github.com/AKACoder/EthereumRPCShell/common/utils"
	"github.com/AKACoder/EthereumRPCShell/shellErrors"
)

type HexString string

func (h HexString) Valid(expLen int) error {
	if len(h) < 2 {
		return shellErrors.CommonErrors.InvalidHexPrefix
	}

	if h[:2] != "0x" {
		return shellErrors.CommonErrors.InvalidHexPrefix
	}

	if h == "0x" && expLen == 0 {
		return nil
	}

	ret, err := hex.DecodeString(string(h[2:]))
	if err != nil {
		return shellErrors.CommonErrors.InvalidHexString
	}

	if expLen != 0 {
		if len(ret) != expLen {
			return shellErrors.CommonErrors.InvalidHexLength
		}
	}

	return nil
}

func (h *HexString) FromAny(v any) bool {
	vStr, ok := v.(string)
	if !ok {
		return false
	}

	*h = HexString(vStr)
	return true
}

type HexAddress = HexString

func (h HexAddress) ValidAddr() bool {
	return h.Valid(constants.EVMAddressLength) == nil
}

type Hash256 = HexString

func (h Hash256) ValidHash() bool {
	return h.Valid(constants.EVMHashLength) == nil
}

type HexData = HexString

func (h HexData) ValidData() bool {
	return h.Valid(0) == nil
}

func (h HexData) ToBytes() ([]byte, error) {
	if len(h) == 0 {
		return nil, nil
	}
	return hex.DecodeString(string(h[2:]))
}

type HexInt = HexString

func (h HexInt) ValidInt() bool {
	hexLen := len(h)
	if hexLen <= 2 {
		return false
	}

	hCopy := h
	if hexLen%2 != 0 {
		hCopy = h[:2] + "0" + h[2:]
	}

	err := hCopy.Valid(0)
	return err == nil
}

type EthBlockNumString string

func (e *EthBlockNumString) FromAny(v any) bool {
	vStr, ok := v.(string)
	if !ok {
		return false
	}

	blkStr := EthBlockNumString(vStr)
	if !blkStr.ValidBlock() {
		return false
	}

	*e = blkStr
	return true
}

func (e EthBlockNumString) Uint64(defBlk uint64) uint64 {
	blk, err := utils.HexParamToUint64(string(e))
	if err == nil {
		return blk
	}

	if string(e) == "earliest" {
		return 0
	}

	return defBlk
}

func (e EthBlockNumString) ValidBlock() bool {
	var validBlockString = map[EthBlockNumString]bool{
		"earliest":  true,
		"latest":    true,
		"safe":      true,
		"finalized": true,
		"pending":   true,
	}

	if HexInt(e).ValidInt() {
		return true
	}

	return validBlockString[e]
}

type RPCProviderError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *RPCProviderError) Error() error {
	return errors.New(fmt.Sprintf("code: %d, message: %s", e.Code, e.Message))
}

var ProviderParseErr = &RPCProviderError{Code: -32700, Message: "Parse error"}
var ProviderInvalidReq = &RPCProviderError{Code: -32600, Message: "Invalid request"}
var ProviderMethodNotFound = &RPCProviderError{Code: -32601, Message: "Method not found"}
var ProviderInvalidParams = &RPCProviderError{Code: -32602, Message: "Invalid params"}
var ProviderInternalErr = &RPCProviderError{Code: -32603, Message: "Internal error"}
var ProviderMethodNotSupport = &RPCProviderError{Code: -32004, Message: "Method not supported"}
