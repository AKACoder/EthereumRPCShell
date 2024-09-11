package ethRPCExecShells

import (
	. "github.com/AKACoder/EthereumRPCShell/ethereumClientProvider"
	. "github.com/AKACoder/EthereumRPCShell/ethereumRPC/ethRPCHandler/ethRPCDataTypes"
)

type RPCClientExecuteFunc func(param []any) (any, *EthClientError)

type EthRPCExecShell struct {
	name        string
	unsupported bool
	minParamLen int
	maxParamLen int
	defRet      any
	execFn      RPCClientExecuteFunc
}

func NewEthRPCExecShell(
	name string,
	minParamLen int,
	maxParamLen int,
	defRet any,
	execFun RPCClientExecuteFunc,
) *EthRPCExecShell {

	return &EthRPCExecShell{
		name:        name,
		unsupported: false,
		minParamLen: minParamLen,
		maxParamLen: maxParamLen,
		defRet:      defRet,
		execFn:      execFun,
	}
}

func (e EthRPCExecShell) Name() string {
	return e.name
}

func (e EthRPCExecShell) paramLenCheck(paramLen int) *EthClientError {
	if paramLen < e.minParamLen || paramLen > e.maxParamLen {
		return ClientInvalidParams
	}

	return nil
}

func (e EthRPCExecShell) Execute(req *EthRPCRequest) (*EthRPCResult, *EthRPCError) {
	if e.unsupported {
		return nil, RPCError(req.ID, ClientMethodNotSupport)
	}

	params, paramOk := req.Params.([]any)
	if !paramOk {
		return nil, RPCError(req.ID, ClientInvalidParams)
	}

	if rpcErr := e.paramLenCheck(len(params)); rpcErr != nil {
		return nil, RPCError(req.ID, ClientInvalidReq)
	}

	if rpcClient == nil {
		if e.defRet != nil {
			return &EthRPCResult{
				EthRPCCommon: req.EthRPCCommon,
				Result:       e.defRet,
			}, nil
		} else {
			return nil, RPCError(req.ID, ClientMethodNotFound)
		}
	}

	if !rpcClient.SupportCheck(e.name) {
		return nil, RPCError(req.ID, ClientMethodNotSupport)
	}

	ret, err := e.execFn(params)
	if err != nil {
		return nil, &EthRPCError{
			EthRPCCommon: req.EthRPCCommon,
			Error:        err,
		}
	} else {
		return &EthRPCResult{
			EthRPCCommon: req.EthRPCCommon,
			Result:       ret,
		}, nil
	}
}
