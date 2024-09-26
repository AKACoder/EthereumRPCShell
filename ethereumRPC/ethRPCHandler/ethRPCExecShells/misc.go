package ethRPCExecShells

import (
	"encoding/json"
	. "github.com/AKACoder/EthereumRPCShell/common/constants"
	"github.com/AKACoder/EthereumRPCShell/common/dataStructure/types"
	. "github.com/AKACoder/EthereumRPCShell/ethereumRPCProvider"
)

var Web3ClientVersion = &EthRPCExecShell{
	name:   Method_web3_clientVersion,
	defRet: "unknown",
	execFn: func([]any) (any, *RPCProviderError) {
		return rpcProvider.Web3ClientVersion()
	},
}

var NetVersion = &EthRPCExecShell{
	name:   Method_net_version,
	defRet: "unknown",
	execFn: func([]any) (any, *RPCProviderError) {
		return rpcProvider.NetVersion()
	},
}

var ETHSyncing = &EthRPCExecShell{
	name:   Method_eth_syncing,
	defRet: false,
	execFn: func([]any) (any, *RPCProviderError) {
		status, data, err := rpcProvider.Syncing()
		if err != nil {
			return nil, err
		}

		if data != nil {
			return data, nil
		}

		return status, err
	},
}

var ETHGetStorageAt = &EthRPCExecShell{
	name:        Method_eth_getStorageAt,
	minParamLen: 3,
	maxParamLen: 3,
	defRet:      "0x0",
	execFn: func(params []any) (any, *RPCProviderError) {
		var addr types.Address
		err := addr.FromParam(params[0])
		if err != nil {
			return nil, ProviderInvalidParams
		}

		var pos types.Key
		err = pos.FromParam(params[1])
		if err != nil {
			return nil, ProviderInvalidParams
		}

		var blk EthBlockNumString
		if !blk.FromAny(params[2]) {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.StorageAt(addr, pos, blk)
	},
}

var ETHGetCode = &EthRPCExecShell{
	name:        Method_eth_getCode,
	minParamLen: 2,
	maxParamLen: 2,
	defRet:      "0x",
	execFn: func(params []any) (any, *RPCProviderError) {
		addr, blk, err := extractAddressAndBlock(params[0], params[1])
		if err != nil {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.Code(*addr, *blk)
	},
}

var ETHGetLogs = &EthRPCExecShell{
	name:        Method_eth_getLogs,
	minParamLen: 1,
	maxParamLen: 1,
	defRet:      []EthLog{},
	execFn: func(params []any) (any, *RPCProviderError) {
		logParam := &EthGetLogsParam{}
		paramJSON, _ := json.Marshal(params[0])

		err := json.Unmarshal(paramJSON, logParam)
		if err != nil {
			return nil, ProviderInvalidParams
		}

		return rpcProvider.Logs(*logParam)
	},
}
