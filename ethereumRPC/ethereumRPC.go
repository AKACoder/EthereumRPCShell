package ethereumRPC

import (
	"EthereumRPCShell/ethereumRPC/ethRPCHandler"
	"github.com/gin-gonic/gin"
)

const (
	apiBaseUri = "/shell"
)

func RoutesRegister(router gin.IRouter) {
	gr := router.Group(apiBaseUri)
	ethRPCHandler.RegisterAPI(gr)
}
