package ethRPCHandler

import (
	. "github.com/AKACoder/EthereumRPCShell/ethereumRPC/ethRPCHandler/ethRPCDataTypes"
	"github.com/AKACoder/EthereumRPCShell/ethereumRPC/ethRPCHandler/ethRPCUtils"
	. "github.com/AKACoder/EthereumRPCShell/ethereumRPCProvider"
	"github.com/AKACoder/EthereumRPCShell/network/http"
	"github.com/gin-gonic/gin"
)

type ethereumRPCHandler struct{}

var rpcAPI ethereumRPCHandler

func (e *ethereumRPCHandler) Handle(ctx *gin.Context) {
	res := http.NewResponse(ctx)
	defer func() {
		if err := recover(); err != nil {
			res.JSON(NewRPCError(RPCDefaultErrorId, ProviderInternalErr))
		}
	}()

	req := &EthRPCRequest{}
	ok := ethRPCUtils.GetReqData(ctx, req)
	if !ok {
		res.JSON(NewRPCError(req.ID, ProviderParseErr))
		return
	}

	shell := rpcShells[req.Method]
	if shell == nil {
		res.JSON(NewRPCError(req.ID, ProviderMethodNotFound))
		return
	}

	successData, errorData := shell.Execute(req)
	if errorData != nil {
		res.JSON(errorData)
	} else {
		res.JSON(successData)
	}
}

func (e *ethereumRPCHandler) Register(router gin.IRouter) {
	router.POST("/eth/rpc", e.Handle)
}
