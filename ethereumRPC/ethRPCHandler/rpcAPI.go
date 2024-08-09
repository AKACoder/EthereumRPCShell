package ethRPCHandler

import (
	. "github.com/AKACoder/EthereumRPCShell/ethereumClientProvider"
	. "github.com/AKACoder/EthereumRPCShell/ethereumRPC/ethRPCHandler/ethRPCDataTypes"
	"github.com/AKACoder/EthereumRPCShell/ethereumRPC/ethRPCHandler/ethRPCUtils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ethereumRPCHandler struct {
	ctx *gin.Context
}

var rpcAPI ethereumRPCHandler

func (e *ethereumRPCHandler) Response(data any) {
	e.ctx.JSON(http.StatusOK, data)
}

func (e *ethereumRPCHandler) Handle(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			e.Response(RPCError(RPCDefaultErrorId, ClientInternalErr))
		}
	}()

	e.ctx = ctx
	req := &EthRPCRequest{}
	ok := ethRPCUtils.GetReqData(ctx, req)
	if !ok {
		e.Response(RPCError(req.ID, ClientParseErr))
		return
	}

	shell := rpcShells[req.Method]
	if shell == nil {
		e.Response(RPCError(req.ID, ClientMethodNotFound))
		return
	}

	successData, errorData := shell.Execute(req)
	if errorData != nil {
		e.Response(errorData)
	} else {
		e.Response(successData)
	}
}

func (e *ethereumRPCHandler) Register(router gin.IRouter) {
	router.POST("/eth/rpc", e.Handle)
}
