package ethRPCExecShells

import (
	. "github.com/AKACoder/EthereumRPCShell/common/constants"
	. "github.com/AKACoder/EthereumRPCShell/ethereumClientProvider"
)

var ETHCoinbase = &EthRPCExecShell{
	name:   Method_eth_coinbase,
	defRet: ETHAddressZero,
	execFn: func([]any) (any, *EthClientError) {
		return rpcClient.Coinbase()
	},
}

var ETHChainId = &EthRPCExecShell{
	name: Method_eth_chainId,
	execFn: func([]any) (any, *EthClientError) {
		return rpcClient.ChainId()
	},
}

var ETHMining = &EthRPCExecShell{
	name:   Method_eth_mining,
	defRet: false,
	execFn: func([]any) (any, *EthClientError) {
		return rpcClient.Mining()
	},
}

var ETHHashRate = &EthRPCExecShell{
	name:   Method_eth_hashrate,
	defRet: "0",
	execFn: func([]any) (any, *EthClientError) {
		return rpcClient.HashRate()
	},
}

var ETHGasPrice = &EthRPCExecShell{
	name:   Method_eth_gasPrice,
	defRet: "0",
	execFn: func([]any) (any, *EthClientError) {
		return rpcClient.GasPrice()
	},
}

var ETHBlockNumber = &EthRPCExecShell{
	name:   Method_eth_blockNumber,
	defRet: "0",
	execFn: func([]any) (any, *EthClientError) {
		return rpcClient.BlockNumber()
	},
}
