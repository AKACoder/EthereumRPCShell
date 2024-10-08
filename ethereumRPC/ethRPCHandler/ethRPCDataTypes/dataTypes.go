package ethRPCDataTypes

import . "github.com/AKACoder/EthereumRPCShell/ethereumRPCProvider"

type EthRPCCommon struct {
	ID      any    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
}

type EthRPCResult struct {
	EthRPCCommon
	Result any `json:"result"`
}

type EthRPCError struct {
	EthRPCCommon
	Error *RPCProviderError `json:"error"`
}

type EthRPCRequest struct {
	EthRPCCommon
	Method string `json:"method"`
	Params any    `json:"params"`
}
