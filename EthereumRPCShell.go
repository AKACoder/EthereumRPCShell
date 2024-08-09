package EthereumRPCShell

import (
	"github.com/AKACoder/EthereumRPCShell/common/wLog"
	. "github.com/AKACoder/EthereumRPCShell/ethRPCShellConfig"
	. "github.com/AKACoder/EthereumRPCShell/ethereumClientProvider"
	. "github.com/AKACoder/EthereumRPCShell/ethereumRPC"
	"github.com/AKACoder/EthereumRPCShell/ethereumRPC/ethRPCHandler"
	"github.com/AKACoder/EthereumRPCShell/ethereumRPC/ethRPCHandler/ethRPCExecShells"
	"github.com/AKACoder/EthereumRPCShell/ethereumRPC/middleware"
	"github.com/AKACoder/EthereumRPCShell/network/http"
	"github.com/AKACoder/EthereumRPCShell/shellErrors"
	"sync"
)

func StartRPCShell(client IEthereumClient, cfg Config) (*sync.WaitGroup, error) {
	wLog.SetUpLogger(cfg.LOG)

	shellErrors.Load()
	ethRPCExecShells.RegisterRPCClient(client)
	ethRPCHandler.LoadShells()

	middleware.RegisterServiceMiddleWare()

	server := http.Server{Config: &cfg.HTTP}
	server.Config.RouteRegisters = []http.RouteRegister{RoutesRegister}
	return server.Start(nil)
}
