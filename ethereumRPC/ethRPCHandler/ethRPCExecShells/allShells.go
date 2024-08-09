package ethRPCExecShells

import . "EthereumRPCShell/common/constants"

var AllShells = []*EthRPCExecShell{
	//unsupported method
	GenUnsupportedShell(Method_web3_sha3),
	GenUnsupportedShell(Method_web3_sha3),
	GenUnsupportedShell(Method_net_listening),
	GenUnsupportedShell(Method_net_peerCount),
	GenUnsupportedShell(Method_eth_protocolVersion),
	GenUnsupportedShell(Method_eth_sign),
	GenUnsupportedShell(Method_eth_signTransaction),
	GenUnsupportedShell(Method_eth_newFilter),
	GenUnsupportedShell(Method_eth_newBlockFilter),
	GenUnsupportedShell(Method_eth_newPendingTransactionFilter),
	GenUnsupportedShell(Method_eth_uninstallFilter),
	GenUnsupportedShell(Method_eth_getFilterChanges),
	GenUnsupportedShell(Method_eth_getFilterLogs),

	//misc
	Web3ClientVersion,
	NetVersion,
	ETHSyncing,
	ETHGetStorageAt,
	ETHGetCode,
	ETHGetLogs,

	//chain related
	ETHCoinbase,
	ETHChainId,
	ETHMining,
	ETHHashRate,
	ETHGasPrice,
	ETHBlockNumber,

	//account related
	ETHAccounts,
	ETHGetBalance,
	ETHGetTransactionCount,

	//block related
	ETHGetUncleCountByBlockHash,
	ETHGetUncleCountByBlockNumber,
	ETHGetBlockByHash,
	ETHGetBlockByNumber,
	ETHGetBlockTransactionCountByHash,
	ETHGetBlockTransactionCountByNumber,
	ETHGetUncleByBlockHashAndIndex,
	ETHGetUncleByBlockNumberAndIndex,

	//transaction related
	ETHSendTransaction,
	ETHSendRawTransaction,
	ETHCall,
	ETHEstimateGas,
	ETHGetTransactionByHash,
	ETHGetTransactionByBlockHashAndIndex,
	ETHGetTransactionByBlockNumberAndIndex,
	ETHGetTransactionReceipt,
}
