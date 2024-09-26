package ethereumRPCProvider

import (
	"errors"
	"fmt"
	"github.com/AKACoder/EthereumRPCShell/common/utils"
)

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

	var blkStr = string(e)
	_, err := utils.HexParamToUint64(blkStr)
	if err == nil {
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
