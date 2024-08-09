package ethRPCExecShells

import (
	"encoding/json"
	. "github.com/AKACoder/EthereumRPCShell/common/constants"
	. "github.com/AKACoder/EthereumRPCShell/ethereumClientProvider"
)

var ETHSendTransaction = &EthRPCExecShell{
	Name:        Method_eth_sendTransaction,
	minParamLen: 1,
	maxParamLen: 1,
	execFn: func(params []any) (any, *EthClientError) {
		txJSON, _ := json.Marshal(params[0])
		tx := &EthBasicTransaction{}
		err := json.Unmarshal(txJSON, tx)

		if err != nil {
			return nil, ClientInvalidParams
		}

		return rpcClient.SendTransaction(*tx)
	},
}

var ETHSendRawTransaction = &EthRPCExecShell{
	Name:        Method_eth_sendRawTransaction,
	minParamLen: 1,
	maxParamLen: 1,
	execFn: func(params []any) (any, *EthClientError) {
		var txData HexData
		if !txData.FromAny(params[0]) {
			return nil, ClientInvalidParams
		}

		if !txData.ValidData() {
			return nil, ClientInvalidParams
		}

		return rpcClient.SendRawTransaction(txData)
	},
}

func getETHCallParam(params []any) (*EthBasicTransaction, *EthBlockNumString, *EthClientError) {
	var blk EthBlockNumString
	if !blk.FromAny(params[1]) {
		return nil, nil, ClientInvalidParams
	}

	if !blk.ValidBlock() {
		return nil, nil, ClientInvalidParams
	}

	txJSON, _ := json.Marshal(params[0])
	tx := &EthBasicTransaction{}
	err := json.Unmarshal(txJSON, tx)

	if err != nil {
		return nil, nil, ClientInvalidParams
	}

	return tx, &blk, nil
}

var ETHCall = &EthRPCExecShell{
	Name:        Method_eth_call,
	minParamLen: 2,
	maxParamLen: 2,
	execFn: func(params []any) (any, *EthClientError) {
		tx, blk, err := getETHCallParam(params)
		if err != nil {
			return nil, err
		}
		return rpcClient.Call(*tx, *blk)
	},
}

var ETHEstimateGas = &EthRPCExecShell{
	Name:        Method_eth_estimateGas,
	minParamLen: 2,
	maxParamLen: 2,
	execFn: func(params []any) (any, *EthClientError) {
		tx, blk, err := getETHCallParam(params)
		if err != nil {
			return nil, err
		}
		return rpcClient.EstimateGas(*tx, *blk)
	},
}

var ETHGetTransactionByHash = &EthRPCExecShell{
	Name:        Method_eth_getTransactionByHash,
	minParamLen: 1,
	maxParamLen: 1,
	execFn: func(params []any) (any, *EthClientError) {
		var txHash Hash256
		if !txHash.FromAny(params[0]) {
			return nil, ClientInvalidParams
		}

		if !txHash.ValidHash() {
			return nil, ClientInvalidParams
		}

		return rpcClient.TransactionByHash(txHash)
	},
}

var ETHGetTransactionByBlockHashAndIndex = &EthRPCExecShell{
	Name:        Method_eth_getTransactionByBlockHashAndIndex,
	minParamLen: 2,
	maxParamLen: 2,
	execFn: func(params []any) (any, *EthClientError) {
		var blkHash Hash256
		var idx HexInt

		if !blkHash.FromAny(params[0]) || idx.FromAny(params[1]) {
			return nil, ClientInvalidParams
		}

		if !blkHash.ValidHash() {
			return nil, ClientInvalidParams
		}

		if !idx.ValidInt() {
			return nil, ClientInvalidParams
		}

		return rpcClient.TransactionByBlockHashAndIndex(blkHash, idx)
	},
}

var ETHGetTransactionByBlockNumberAndIndex = &EthRPCExecShell{
	Name:        Method_eth_getTransactionByBlockNumberAndIndex,
	minParamLen: 2,
	maxParamLen: 2,
	execFn: func(params []any) (any, *EthClientError) {
		var blk EthBlockNumString
		var idx HexInt

		if !blk.FromAny(params[0]) || idx.FromAny(params[1]) {
			return nil, ClientInvalidParams
		}

		if !blk.ValidBlock() {
			return nil, ClientInvalidParams
		}

		if !idx.ValidInt() {
			return nil, ClientInvalidParams
		}

		return rpcClient.TransactionByBlockNumberAndIndex(blk, idx)
	},
}

var ETHGetTransactionReceipt = &EthRPCExecShell{
	Name:        Method_eth_getTransactionReceipt,
	minParamLen: 1,
	maxParamLen: 1,
	execFn: func(params []any) (any, *EthClientError) {
		var txHash Hash256

		if !txHash.FromAny(params[0]) {
			return nil, ClientInvalidParams
		}

		if !txHash.ValidHash() {
			return nil, ClientInvalidParams
		}

		return rpcClient.TransactionReceipt(txHash)
	},
}
