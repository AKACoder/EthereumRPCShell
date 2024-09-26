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
		addr, blk, err := extractAddressAndBlock(params[0], params[1])
		if err != nil {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.Balance(*addr, *blk)
	},
}

var ETHGetTransactionCount = &EthRPCExecShell{
	name:        Method_eth_getTransactionCount,
	minParamLen: 2,
	maxParamLen: 2,
	defRet:      "0x0",
	execFn: func(params []any) (any, *RPCProviderError) {
		addr, blk, err := extractAddressAndBlock(params[0], params[1])
		if err != nil {
			return nil, ProviderInvalidParams
		}
		return rpcProvider.TransactionCount(*addr, *blk)
	},
}
