package ethRPCHandler

import (
	"github.com/AKACoder/EthereumRPCShell/ethereumRPC/ethRPCHandler/ethRPCExecShells"
)

var rpcShells = make(map[string]*ethRPCExecShells.EthRPCExecShell)

func RegisterRPCShell(shell *ethRPCExecShells.EthRPCExecShell) {
	rpcShells[shell.Name()] = shell
}
