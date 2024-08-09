package middleware

import (
	"EthereumRPCShell/ethereumRPC/middleware/after"
	"EthereumRPCShell/ethereumRPC/middleware/before"
	"EthereumRPCShell/network/http"
)

func RegisterServiceMiddleWare() {
	http.RegisterMiddleware(before.Middlewares, after.Middlewares)
}
