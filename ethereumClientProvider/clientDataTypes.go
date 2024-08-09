package ethereumClientProvider

import (
	"encoding/hex"
	"github.com/AKACoder/EthereumRPCShell/common/constants"
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
	return h.Valid(constants.EVMHash256Length) == nil
}

type HexData = HexString

func (h HexData) ValidData() bool {
	return h.Valid(0) == nil
}

func (h HexData) ToBytes() ([]byte, error) {
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

	*e = EthBlockNumString(vStr)
	return true
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

type EthClientError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var ClientParseErr = &EthClientError{Code: -32700, Message: "Parse error"}
var ClientInvalidReq = &EthClientError{Code: -32600, Message: "Invalid request"}
var ClientMethodNotFound = &EthClientError{Code: -32601, Message: "Method not found"}
var ClientInvalidParams = &EthClientError{Code: -32602, Message: "Invalid params"}
var ClientInternalErr = &EthClientError{Code: -32603, Message: "Internal error"}
var ClientMethodNotSupport = &EthClientError{Code: -32004, Message: "Method not supported"}
