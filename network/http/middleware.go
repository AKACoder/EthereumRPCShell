package http

import (
	"github.com/AKACoder/EthereumRPCShell/common/rpcLog"
	"github.com/gin-gonic/gin"
)

type AfterHandler func(result *ServiceResult, res *Response)
type BeforeHandler gin.HandlerFunc

type middlewareList struct {
	before []gin.HandlerFunc
	after  []AfterHandler
}

var middlewares middlewareList

func RegisterMiddleware(before []gin.HandlerFunc, after []AfterHandler) {
	middlewares.before = before
	middlewares.after = after
}

func callAfterMiddleware(finalResult *ServiceResult, res *Response) (ret bool) {
	defer func() {
		if e := recover(); e != nil {
			rpcLog.Log.Errorln("got an error @ after middleware: ", e)
			ret = false
		}
	}()

	for _, f := range middlewares.after {
		f(finalResult, res)
	}

	return true
}
