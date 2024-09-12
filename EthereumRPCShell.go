package EthereumRPCShell

import (
	"github.com/AKACoder/EthereumRPCShell/common/wLog"
	. "github.com/AKACoder/EthereumRPCShell/ethRPCShellConfig"
	. "github.com/AKACoder/EthereumRPCShell/ethereumRPC"
	"github.com/AKACoder/EthereumRPCShell/ethereumRPC/ethRPCHandler"
	"github.com/AKACoder/EthereumRPCShell/ethereumRPC/ethRPCHandler/ethRPCExecShells"
	"github.com/AKACoder/EthereumRPCShell/ethereumRPC/middleware"
	. "github.com/AKACoder/EthereumRPCShell/ethereumRPCProvider"
	"github.com/AKACoder/EthereumRPCShell/network/http"
	"github.com/AKACoder/EthereumRPCShell/shellErrors"
	"sync"
)

func StartRPCShell(provider IEthereumRPCProvider, cfg Config) (*sync.WaitGroup, error) {
	wLog.SetUpLogger(cfg.LOG)

	shellErrors.Load()
	ethRPCExecShells.RegisterProvider(provider)
	ethRPCHandler.LoadShells()

	middleware.RegisterServiceMiddleWare()

	server := http.Server{Config: &cfg.HTTP}
	server.Config.RouteRegisters = []http.RouteRegister{RoutesRegister}
	return server.Start(nil)
}
