package ethRPCExecShells

import (
	. "github.com/AKACoder/EthereumRPCShell/common/constants"
	. "github.com/AKACoder/EthereumRPCShell/ethereumRPCProvider"
)

var ETHAccounts = &EthRPCExecShell{
	name:   Method_eth_accounts,
	defRet: []HexAddress{},
	execFn: func([]any) (any, *RPCProviderError) {
		return rpcProvider.Accounts()
	},
}

var ETHGetBalance = &EthRPCExecShell{
	name:        Method_eth_getBalance,
	minParamLen: 2,
	maxParamLen: 2,
	defRet:      "0x0",
	execFn: func(params []any) (any, *RPCProviderError) {
		var addr HexAddress
		var blk EthBlockNumString

		if !addr.FromAny(params[0]) || !blk.FromAny(params[1]) {
			return nil, ProviderInvalidParams
		}

		if !addr.ValidAddr() || !blk.ValidBlock() {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.Balance(addr, blk)
	},
}

var ETHGetTransactionCount = &EthRPCExecShell{
	name:        Method_eth_getTransactionCount,
	minParamLen: 2,
	maxParamLen: 2,
	defRet:      "0x0",
	execFn: func(params []any) (any, *RPCProviderError) {
		var addr HexAddress
		var blk EthBlockNumString

		if !addr.FromAny(params[0]) || !blk.FromAny(params[1]) {
			return nil, ProviderInvalidParams
		}

		if !addr.ValidAddr() || !blk.ValidBlock() {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.TransactionCount(addr, blk)
	},
}
