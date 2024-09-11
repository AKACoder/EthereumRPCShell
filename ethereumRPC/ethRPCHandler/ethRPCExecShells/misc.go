package ethRPCExecShells

import (
	"encoding/json"
	. "github.com/AKACoder/EthereumRPCShell/common/constants"
	. "github.com/AKACoder/EthereumRPCShell/ethereumClientProvider"
)

var Web3ClientVersion = &EthRPCExecShell{
	name:   Method_web3_clientVersion,
	defRet: "unknown",
	execFn: func([]any) (any, *EthClientError) {
		return rpcClient.Web3ClientVersion()
	},
}

var NetVersion = &EthRPCExecShell{
	name:   Method_net_version,
	defRet: "unknown",
	execFn: func([]any) (any, *EthClientError) {
		return rpcClient.NetVersion()
	},
}

var ETHSyncing = &EthRPCExecShell{
	name:   Method_eth_syncing,
	defRet: false,
	execFn: func([]any) (any, *EthClientError) {
		status, data, err := rpcClient.Syncing()
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
	execFn: func(params []any) (any, *EthClientError) {
		var addr HexAddress
		var pos HexInt
		var blk EthBlockNumString

		if !addr.FromAny(params[0]) || !pos.FromAny(params[1]) || !blk.FromAny(params[2]) {
			return nil, ClientInvalidParams
		}

		if !addr.ValidAddr() || !pos.ValidInt() || !blk.ValidBlock() {
			return nil, ClientInvalidParams
		}

		return rpcClient.StorageAt(addr, pos, blk)
	},
}

var ETHGetCode = &EthRPCExecShell{
	name:        Method_eth_getCode,
	minParamLen: 2,
	maxParamLen: 2,
	defRet:      "0x",
	execFn: func(params []any) (any, *EthClientError) {
		var addr HexAddress
		var blk EthBlockNumString

		if !addr.FromAny(params[0]) || !blk.FromAny(params[1]) {
			return nil, ClientInvalidParams
		}

		if !addr.ValidAddr() || !blk.ValidBlock() {
			return nil, ClientInvalidParams
		}

		return rpcClient.Code(addr, blk)
	},
}

var ETHGetLogs = &EthRPCExecShell{
	name:        Method_eth_getLogs,
	minParamLen: 1,
	maxParamLen: 1,
	defRet:      []EthLog{},
	execFn: func(params []any) (any, *EthClientError) {
		logParam := &EthGetLogsParam{}
		paramJSON, _ := json.Marshal(params[0])

		err := json.Unmarshal(paramJSON, logParam)
		if err != nil {
			return nil, ClientInvalidParams
		}

		return rpcClient.Logs(*logParam)
	},
}
