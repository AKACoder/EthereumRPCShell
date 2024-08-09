package ethRPCExecShells

import (
	. "github.com/AKACoder/EthereumRPCShell/common/constants"
	. "github.com/AKACoder/EthereumRPCShell/ethereumClientProvider"
)

var ETHGetUncleCountByBlockHash = &EthRPCExecShell{
	Name:        Method_eth_getUncleCountByBlockHash,
	minParamLen: 1,
	maxParamLen: 1,
	defRet:      "0x0",
	execFn: func(params []any) (any, *EthClientError) {
		hash, hashOk := params[0].(Hash256)

		if !hashOk {
			return nil, ClientInvalidParams
		}

		if !hash.ValidHash() {
			return nil, ClientInvalidParams
		}

		return rpcClient.UncleCountByBlockHash(hash)
	},
}

var ETHGetUncleCountByBlockNumber = &EthRPCExecShell{
	Name:        Method_eth_getUncleCountByBlockNumber,
	minParamLen: 1,
	maxParamLen: 1,
	defRet:      "0x0",
	execFn: func(params []any) (any, *EthClientError) {
		var blk EthBlockNumString

		if !blk.FromAny(params[0]) {
			return nil, ClientInvalidParams
		}

		if !blk.ValidBlock() {
			return nil, ClientInvalidParams
		}

		return rpcClient.UncleCountByBlockNumber(blk)
	},
}

var ETHGetBlockByHash = &EthRPCExecShell{
	Name:        Method_eth_getBlockByHash,
	minParamLen: 2,
	maxParamLen: 2,
	execFn: func(params []any) (any, *EthClientError) {
		var blkHash Hash256

		if !blkHash.FromAny(params[0]) {
			return nil, ClientInvalidParams
		}

		if !blkHash.ValidHash() {
			return nil, ClientInvalidParams
		}

		txFlag, flagOk := params[1].(bool)
		if !flagOk {
			return nil, ClientInvalidParams
		}

		return rpcClient.BlockByHash(blkHash, txFlag)
	},
}

var ETHGetBlockByNumber = &EthRPCExecShell{
	Name:        Method_eth_getBlockByNumber,
	minParamLen: 2,
	maxParamLen: 2,
	execFn: func(params []any) (any, *EthClientError) {
		var blk EthBlockNumString

		if !blk.FromAny(params[0]) {
			return nil, ClientInvalidParams
		}

		if !blk.ValidBlock() {
			return nil, ClientInvalidParams
		}

		txFlag, flagOk := params[1].(bool)
		if !flagOk {
			return nil, ClientInvalidParams
		}

		return rpcClient.BlockByNumber(blk, txFlag)
	},
}

var ETHGetBlockTransactionCountByHash = &EthRPCExecShell{
	Name:        Method_eth_getBlockTransactionCountByHash,
	minParamLen: 1,
	maxParamLen: 1,
	defRet:      "0x0",
	execFn: func(params []any) (any, *EthClientError) {
		var hash Hash256

		if !hash.FromAny(params[0]) {
			return nil, ClientInvalidParams
		}

		if !hash.ValidHash() {
			return nil, ClientInvalidParams
		}

		return rpcClient.BlockTransactionCountByHash(hash)
	},
}

var ETHGetBlockTransactionCountByNumber = &EthRPCExecShell{
	Name:        Method_eth_getBlockTransactionCountByNumber,
	minParamLen: 1,
	maxParamLen: 1,
	defRet:      "0x0",
	execFn: func(params []any) (any, *EthClientError) {
		var blk EthBlockNumString

		if !blk.FromAny(params[0]) {
			return nil, ClientInvalidParams
		}

		if !blk.ValidBlock() {
			return nil, ClientInvalidParams
		}

		return rpcClient.BlockTransactionCountByNumber(blk)
	},
}

var ETHGetUncleByBlockHashAndIndex = &EthRPCExecShell{
	Name:        Method_eth_getUncleByBlockHashAndIndex,
	minParamLen: 2,
	maxParamLen: 2,
	defRet:      "0x0",
	execFn: func(params []any) (any, *EthClientError) {
		var blkHash Hash256
		var idx HexInt

		if !blkHash.FromAny(params[0]) || !idx.FromAny(params[1]) {
			return nil, ClientInvalidParams
		}

		if !blkHash.ValidHash() || !idx.ValidInt() {
			return nil, ClientInvalidParams
		}

		return rpcClient.UncleByBlockHashAndIndex(blkHash, idx)
	},
}

var ETHGetUncleByBlockNumberAndIndex = &EthRPCExecShell{
	Name:        Method_eth_getUncleByBlockNumberAndIndex,
	minParamLen: 2,
	maxParamLen: 2,
	defRet:      "0x0",
	execFn: func(params []any) (any, *EthClientError) {
		var blk EthBlockNumString
		var idx HexInt

		if !blk.FromAny(params[0]) || !idx.FromAny(params[1]) {
			return nil, ClientInvalidParams
		}

		if !blk.ValidBlock() || !idx.ValidInt() {
			return nil, ClientInvalidParams
		}

		return rpcClient.UncleByBlockNumberAndIndex(blk, idx)
	},
}
