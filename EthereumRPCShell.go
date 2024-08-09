package EthereumRPCShell

import (
	"EthereumRPCShell/common/wLog"
	. "EthereumRPCShell/ethRPCShellConfig"
	. "EthereumRPCShell/ethereumClientProvider"
	. "EthereumRPCShell/ethereumRPC"
	"EthereumRPCShell/ethereumRPC/ethRPCHandler"
	"EthereumRPCShell/ethereumRPC/ethRPCHandler/ethRPCExecShells"
	"EthereumRPCShell/ethereumRPC/middleware"
	"EthereumRPCShell/network/http"
	"sync"
)

func StartRPCShell(client IEthereumClient, cfg Config) (*sync.WaitGroup, error) {
	wLog.SetUpLogger(cfg.LOG)

	ethRPCExecShells.RegisterRPCClient(client)
	ethRPCHandler.LoadShells()

	middleware.RegisterServiceMiddleWare()

	server := http.Server{Config: &cfg.HTTP}
	server.Config.RouteRegisters = []http.RouteRegister{RoutesRegister}
	return server.Start(nil)
}
