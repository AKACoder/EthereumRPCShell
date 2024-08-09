package middleware

import (
	"github.com/AKACoder/EthereumRPCShell/ethereumRPC/middleware/after"
	"github.com/AKACoder/EthereumRPCShell/ethereumRPC/middleware/before"
	"github.com/AKACoder/EthereumRPCShell/network/http"
)

func RegisterServiceMiddleWare() {
	http.RegisterMiddleware(before.Middlewares, after.Middlewares)
}
