package ethRPCExecShells

import (
	"encoding/json"
	. "github.com/AKACoder/EthereumRPCShell/common/constants"
	. "github.com/AKACoder/EthereumRPCShell/ethereumRPCProvider"
)

var ETHSendTransaction = &EthRPCExecShell{
	name:        Method_eth_sendTransaction,
	minParamLen: 1,
	maxParamLen: 1,
	execFn: func(params []any) (any, *RPCProviderError) {
		txJSON, _ := json.Marshal(params[0])
		tx := &EthBasicTransaction{}
		err := json.Unmarshal(txJSON, tx)

		if err != nil {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.SendTransaction(*tx)
	},
}

var ETHSendRawTransaction = &EthRPCExecShell{
	name:        Method_eth_sendRawTransaction,
	minParamLen: 1,
	maxParamLen: 1,
	execFn: func(params []any) (any, *RPCProviderError) {
		var txData HexData
		if !txData.FromAny(params[0]) {
			return nil, ProviderInvalidParams
		}

		if !txData.ValidData() {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.SendRawTransaction(txData)
	},
}

func getETHCallParam(params []any) (*EthBasicTransaction, *EthBlockNumString, *RPCProviderError) {
	var blk EthBlockNumString
	if !blk.FromAny(params[1]) {
		return nil, nil, ProviderInvalidParams
	}

	if !blk.ValidBlock() {
		return nil, nil, ProviderInvalidParams
	}

	txJSON, _ := json.Marshal(params[0])
	tx := &EthBasicTransaction{}
	err := json.Unmarshal(txJSON, tx)

	if err != nil {
		return nil, nil, ProviderInvalidParams
	}

	return tx, &blk, nil
}

var ETHCall = &EthRPCExecShell{
	name:        Method_eth_call,
	minParamLen: 2,
	maxParamLen: 2,
	execFn: func(params []any) (any, *RPCProviderError) {
		tx, blk, err := getETHCallParam(params)
		if err != nil {
			return nil, err
		}
		return rpcProvider.Call(*tx, *blk)
	},
}

var ETHEstimateGas = &EthRPCExecShell{
	name:        Method_eth_estimateGas,
	minParamLen: 2,
	maxParamLen: 2,
	execFn: func(params []any) (any, *RPCProviderError) {
		tx, blk, err := getETHCallParam(params)
		if err != nil {
			return nil, err
		}
		return rpcProvider.EstimateGas(*tx, *blk)
	},
}

var ETHGetTransactionByHash = &EthRPCExecShell{
	name:        Method_eth_getTransactionByHash,
	minParamLen: 1,
	maxParamLen: 1,
	execFn: func(params []any) (any, *RPCProviderError) {
		var txHash Hash256
		if !txHash.FromAny(params[0]) {
			return nil, ProviderInvalidParams
		}

		if !txHash.ValidHash() {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.TransactionByHash(txHash)
	},
}

var ETHGetTransactionByBlockHashAndIndex = &EthRPCExecShell{
	name:        Method_eth_getTransactionByBlockHashAndIndex,
	minParamLen: 2,
	maxParamLen: 2,
	execFn: func(params []any) (any, *RPCProviderError) {
		var blkHash Hash256
		var idx HexInt

		if !blkHash.FromAny(params[0]) || idx.FromAny(params[1]) {
			return nil, ProviderInvalidParams
		}

		if !blkHash.ValidHash() {
			return nil, ProviderInvalidParams
		}

		if !idx.ValidInt() {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.TransactionByBlockHashAndIndex(blkHash, idx)
	},
}

var ETHGetTransactionByBlockNumberAndIndex = &EthRPCExecShell{
	name:        Method_eth_getTransactionByBlockNumberAndIndex,
	minParamLen: 2,
	maxParamLen: 2,
	execFn: func(params []any) (any, *RPCProviderError) {
		var blk EthBlockNumString
		var idx HexInt

		if !blk.FromAny(params[0]) || idx.FromAny(params[1]) {
			return nil, ProviderInvalidParams
		}

		if !blk.ValidBlock() {
			return nil, ProviderInvalidParams
		}

		if !idx.ValidInt() {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.TransactionByBlockNumberAndIndex(blk, idx)
	},
}

var ETHGetTransactionReceipt = &EthRPCExecShell{
	name:        Method_eth_getTransactionReceipt,
	minParamLen: 1,
	maxParamLen: 1,
	execFn: func(params []any) (any, *RPCProviderError) {
		var txHash Hash256

		if !txHash.FromAny(params[0]) {
			return nil, ProviderInvalidParams
		}

		if !txHash.ValidHash() {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.TransactionReceipt(txHash)
	},
}
