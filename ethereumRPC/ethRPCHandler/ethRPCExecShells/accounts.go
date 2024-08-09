package ethRPCExecShells

import (
	. "EthereumRPCShell/common/constants"
	. "EthereumRPCShell/ethereumClientProvider"
)

var ETHAccounts = &EthRPCExecShell{
	Name:   Method_eth_accounts,
	defRet: []HexAddress{},
	execFn: func([]any) (any, *EthClientError) {
		return rpcClient.Accounts()
	},
}

var ETHGetBalance = &EthRPCExecShell{
	Name:        Method_eth_getBalance,
	minParamLen: 2,
	maxParamLen: 2,
	defRet:      "0x0",
	execFn: func(params []any) (any, *EthClientError) {
		var addr HexAddress
		var blk EthBlockNumString

		if !addr.FromAny(params[0]) || !blk.FromAny(params[1]) {
			return nil, ClientInvalidParams
		}

		if !addr.ValidAddr() || !blk.ValidBlock() {
			return nil, ClientInvalidParams
		}

		return rpcClient.Balance(addr, blk)
	},
}

var ETHGetTransactionCount = &EthRPCExecShell{
	Name:        Method_eth_getTransactionCount,
	minParamLen: 2,
	maxParamLen: 2,
	defRet:      "0x0",
	execFn: func(params []any) (any, *EthClientError) {
		var addr HexAddress
		var blk EthBlockNumString

		if !addr.FromAny(params[0]) || !blk.FromAny(params[1]) {
			return nil, ClientInvalidParams
		}

		if !addr.ValidAddr() || !blk.ValidBlock() {
			return nil, ClientInvalidParams
		}

		return rpcClient.TransactionCount(addr, blk)
	},
}
