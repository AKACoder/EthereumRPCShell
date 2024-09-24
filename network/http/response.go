/*
 * Copyright 2020 The SealABC Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package http

import (
	"github.com/AKACoder/EthereumRPCShell/shellErrors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ServiceOKCode = 0
)

type ServiceResult struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Response struct {
	ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		ctx: ctx,
	}
}

//service response will always set http status to 200.
func (c *Response) afterMiddlewareFailed() {
	err := shellErrors.ServiceErrors.Middleware
	c.ctx.JSON(http.StatusOK, &ServiceResult{
		Code:    err.Code(),
		Message: err.Error(),
		Data:    nil,
	})
}

func (c *Response) Success(data interface{}) {
	c.ServiceResult(0, "", data)
}

func (c *Response) JSON(data interface{}) {
	c.ctx.JSON(http.StatusOK, data)
}

func (c *Response) ServiceResult(code int64, msg string, data interface{}) {
	sRet := ServiceResult{
		Code:    code,
		Message: msg,
		Data:    data,
	}

	if !callAfterMiddleware(&sRet, c) {
		c.afterMiddlewareFailed()
	} else {
		c.ctx.JSON(http.StatusOK, sRet)
	}
}
