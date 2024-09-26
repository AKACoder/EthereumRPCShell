package ethRPCExecShells

import (
	"encoding/json"
	. "github.com/AKACoder/EthereumRPCShell/common/constants"
	"github.com/AKACoder/EthereumRPCShell/common/dataStructure/types"
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
		var txData types.Data
		err := txData.FromParam(params[0])
		if err != nil {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.SendRawTransaction(txData)
	},
}

func getETHCallParam(params []any) (*EthBasicTransaction, *EthBlockNumString, *RPCProviderError) {
	txJSON, _ := json.Marshal(params[0])
	tx := &EthBasicTransaction{}
	err := json.Unmarshal(txJSON, tx)

	if err != nil {
		return nil, nil, ProviderInvalidParams
	}

	var blk EthBlockNumString = DefaultBlock
	if len(params) == 1 {
		return tx, &blk, nil
	}

	if !blk.FromAny(params[1]) {
		return nil, nil, ProviderInvalidParams
	}

	return tx, &blk, nil
}

var ETHCall = &EthRPCExecShell{
	name:        Method_eth_call,
	minParamLen: 1,
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
	minParamLen: 1,
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
		txHash, err := extractHash(params[0])
		if err != nil {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.TransactionByHash(*txHash)
	},
}

var ETHGetTransactionByBlockHashAndIndex = &EthRPCExecShell{
	name:        Method_eth_getTransactionByBlockHashAndIndex,
	minParamLen: 2,
	maxParamLen: 2,
	execFn: func(params []any) (any, *RPCProviderError) {
		blkHash, idx, err := extractHashAndInt(params[0], params[1])
		if err != nil {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.TransactionByBlockHashAndIndex(*blkHash, idx)
	},
}

var ETHGetTransactionByBlockNumberAndIndex = &EthRPCExecShell{
	name:        Method_eth_getTransactionByBlockNumberAndIndex,
	minParamLen: 2,
	maxParamLen: 2,
	execFn: func(params []any) (any, *RPCProviderError) {
		blk, idx, err := extractBlkNumAndInt(params[0], params[1])

		if err != nil {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.TransactionByBlockNumberAndIndex(*blk, idx)
	},
}

var ETHGetTransactionReceipt = &EthRPCExecShell{
	name:        Method_eth_getTransactionReceipt,
	minParamLen: 1,
	maxParamLen: 1,
	execFn: func(params []any) (any, *RPCProviderError) {
		txHash, err := extractHash(params[0])

		if err != nil {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.TransactionReceipt(*txHash)
	},
}
