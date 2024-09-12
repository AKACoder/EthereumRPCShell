package ethRPCExecShells

import (
	. "github.com/AKACoder/EthereumRPCShell/ethereumRPC/ethRPCHandler/ethRPCDataTypes"
	. "github.com/AKACoder/EthereumRPCShell/ethereumRPCProvider"
)

type RPCExecuteFunc func(param []any) (any, *RPCProviderError)

type EthRPCExecShell struct {
	name        string
	unsupported bool
	minParamLen int
	maxParamLen int
	defRet      any
	execFn      RPCExecuteFunc
}

func NewEthRPCExecShell(
	name string,
	minParamLen int,
	maxParamLen int,
	defRet any,
	execFun RPCExecuteFunc,
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

func (e EthRPCExecShell) paramLenCheck(paramLen int) *RPCProviderError {
	if paramLen < e.minParamLen || paramLen > e.maxParamLen {
		return ProviderInvalidParams
	}

	return nil
}

func (e EthRPCExecShell) Execute(req *EthRPCRequest) (*EthRPCResult, *EthRPCError) {
	if e.unsupported {
		return nil, NewRPCError(req.ID, ProviderMethodNotSupport)
	}

	params, paramOk := req.Params.([]any)
	if !paramOk {
		return nil, NewRPCError(req.ID, ProviderInvalidParams)
	}

	if rpcErr := e.paramLenCheck(len(params)); rpcErr != nil {
		return nil, NewRPCError(req.ID, ProviderInvalidReq)
	}

	if rpcProvider == nil {
		if e.defRet != nil {
			return &EthRPCResult{
				EthRPCCommon: req.EthRPCCommon,
				Result:       e.defRet,
			}, nil
		} else {
			return nil, NewRPCError(req.ID, ProviderMethodNotFound)
		}
	}

	if !rpcProvider.SupportCheck(e.name) {
		return nil, NewRPCError(req.ID, ProviderMethodNotSupport)
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
