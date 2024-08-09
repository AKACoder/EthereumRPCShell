package ethRPCHandler

import "github.com/gin-gonic/gin"

func RegisterAPI(router gin.IRouter) {
	rpcAPI.Register(router)
}
