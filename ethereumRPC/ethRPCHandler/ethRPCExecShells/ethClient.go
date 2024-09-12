package ethRPCExecShells

import "github.com/AKACoder/EthereumRPCShell/ethereumRPCProvider"

var rpcProvider ethereumRPCProvider.IEthereumRPCProvider = nil

func RegisterProvider(provider ethereumRPCProvider.IEthereumRPCProvider) {
	rpcProvider = provider
}
