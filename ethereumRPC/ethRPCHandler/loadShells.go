package ethRPCHandler

import (
	. "EthereumRPCShell/ethereumRPC/ethRPCHandler/ethRPCExecShells"
)

func LoadShells() {
	for _, shell := range AllShells {
		RegisterRPCShell(shell)
	}
}
