package ethRPCExecShells

import "EthereumRPCShell/ethereumClientProvider"

var rpcClient ethereumClientProvider.IEthereumClient = nil

func RegisterRPCClient(client ethereumClientProvider.IEthereumClient) {
	rpcClient = client
}
