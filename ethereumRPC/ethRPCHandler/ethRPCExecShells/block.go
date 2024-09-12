package ethRPCExecShells

import (
	. "github.com/AKACoder/EthereumRPCShell/common/constants"
	. "github.com/AKACoder/EthereumRPCShell/ethereumRPCProvider"
)

var ETHGetUncleCountByBlockHash = &EthRPCExecShell{
	name:        Method_eth_getUncleCountByBlockHash,
	minParamLen: 1,
	maxParamLen: 1,
	defRet:      "0x0",
	execFn: func(params []any) (any, *RPCProviderError) {
		hash, hashOk := params[0].(Hash256)

		if !hashOk {
			return nil, ProviderInvalidParams
		}

		if !hash.ValidHash() {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.UncleCountByBlockHash(hash)
	},
}

var ETHGetUncleCountByBlockNumber = &EthRPCExecShell{
	name:        Method_eth_getUncleCountByBlockNumber,
	minParamLen: 1,
	maxParamLen: 1,
	defRet:      "0x0",
	execFn: func(params []any) (any, *RPCProviderError) {
		var blk EthBlockNumString

		if !blk.FromAny(params[0]) {
			return nil, ProviderInvalidParams
		}

		if !blk.ValidBlock() {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.UncleCountByBlockNumber(blk)
	},
}

var ETHGetBlockByHash = &EthRPCExecShell{
	name:        Method_eth_getBlockByHash,
	minParamLen: 2,
	maxParamLen: 2,
	execFn: func(params []any) (any, *RPCProviderError) {
		var blkHash Hash256

		if !blkHash.FromAny(params[0]) {
			return nil, ProviderInvalidParams
		}

		if !blkHash.ValidHash() {
			return nil, ProviderInvalidParams
		}

		txFlag, flagOk := params[1].(bool)
		if !flagOk {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.BlockByHash(blkHash, txFlag)
	},
}

var ETHGetBlockByNumber = &EthRPCExecShell{
	name:        Method_eth_getBlockByNumber,
	minParamLen: 2,
	maxParamLen: 2,
	execFn: func(params []any) (any, *RPCProviderError) {
		var blk EthBlockNumString

		if !blk.FromAny(params[0]) {
			return nil, ProviderInvalidParams
		}

		if !blk.ValidBlock() {
			return nil, ProviderInvalidParams
		}

		txFlag, flagOk := params[1].(bool)
		if !flagOk {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.BlockByNumber(blk, txFlag)
	},
}

var ETHGetBlockTransactionCountByHash = &EthRPCExecShell{
	name:        Method_eth_getBlockTransactionCountByHash,
	minParamLen: 1,
	maxParamLen: 1,
	defRet:      "0x0",
	execFn: func(params []any) (any, *RPCProviderError) {
		var hash Hash256

		if !hash.FromAny(params[0]) {
			return nil, ProviderInvalidParams
		}

		if !hash.ValidHash() {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.BlockTransactionCountByHash(hash)
	},
}

var ETHGetBlockTransactionCountByNumber = &EthRPCExecShell{
	name:        Method_eth_getBlockTransactionCountByNumber,
	minParamLen: 1,
	maxParamLen: 1,
	defRet:      "0x0",
	execFn: func(params []any) (any, *RPCProviderError) {
		var blk EthBlockNumString

		if !blk.FromAny(params[0]) {
			return nil, ProviderInvalidParams
		}

		if !blk.ValidBlock() {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.BlockTransactionCountByNumber(blk)
	},
}

var ETHGetUncleByBlockHashAndIndex = &EthRPCExecShell{
	name:        Method_eth_getUncleByBlockHashAndIndex,
	minParamLen: 2,
	maxParamLen: 2,
	defRet:      "0x0",
	execFn: func(params []any) (any, *RPCProviderError) {
		var blkHash Hash256
		var idx HexInt

		if !blkHash.FromAny(params[0]) || !idx.FromAny(params[1]) {
			return nil, ProviderInvalidParams
		}

		if !blkHash.ValidHash() || !idx.ValidInt() {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.UncleByBlockHashAndIndex(blkHash, idx)
	},
}

var ETHGetUncleByBlockNumberAndIndex = &EthRPCExecShell{
	name:        Method_eth_getUncleByBlockNumberAndIndex,
	minParamLen: 2,
	maxParamLen: 2,
	defRet:      "0x0",
	execFn: func(params []any) (any, *RPCProviderError) {
		var blk EthBlockNumString
		var idx HexInt

		if !blk.FromAny(params[0]) || !idx.FromAny(params[1]) {
			return nil, ProviderInvalidParams
		}

		if !blk.ValidBlock() || !idx.ValidInt() {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.UncleByBlockNumberAndIndex(blk, idx)
	},
}
