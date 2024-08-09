package ethRPCDataTypes

import . "github.com/AKACoder/EthereumRPCShell/ethereumClientProvider"

const RPCDefaultErrorId = "error"

func RPCError(id any, err *EthClientError) *EthRPCError {
	return &EthRPCError{
		EthRPCCommon: EthRPCCommon{
			ID:      id,
			JSONRPC: "2.0",
		},
		Error: err,
	}
}
