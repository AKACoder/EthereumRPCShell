package ethRPCHandler

import "github.com/AKACoder/EthereumRPCShell/ethereumRPC/ethRPCHandler/ethRPCExecShells"

func LoadShells() {
	for _, shell := range ethRPCExecShells.AllShells {
		RegisterRPCShell(shell)
	}
}
