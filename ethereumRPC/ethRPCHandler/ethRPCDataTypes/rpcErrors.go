package ethRPCDataTypes

import . "github.com/AKACoder/EthereumRPCShell/ethereumRPCProvider"

const RPCDefaultErrorId = "error"

func NewRPCError(id any, err *RPCProviderError) *EthRPCError {
	return &EthRPCError{
		EthRPCCommon: EthRPCCommon{
			ID:      id,
			JSONRPC: "2.0",
		},
		Error: err,
	}
}
