package ethereumRPCProvider

import (
	"github.com/AKACoder/EthereumRPCShell/common/dataStructure/types"
)

const (
	unknownStr = "unknown"
)

type IMustImplementMethods interface {
	SupportCheck(name string) bool
	ChainId() (*types.BigInt, *RPCProviderError)
	BlockNumber() (types.Uint64, *RPCProviderError)
	Balance(types.Address, EthBlockNumString) (*types.BigInt, *RPCProviderError)
	SendRawTransaction(types.Data) (types.Hash, *RPCProviderError)
	Call(EthBasicTransaction, EthBlockNumString) (types.Hash, *RPCProviderError)
	BlockByHash(types.Hash, bool) (*EthBlock, *RPCProviderError)
	BlockByNumber(EthBlockNumString, bool) (*EthBlock, *RPCProviderError)
	TransactionByHash(types.Hash) (*EthFullTransaction, *RPCProviderError)
	TransactionReceipt(types.Hash) (*EthTransactionReceipt, *RPCProviderError)
	Logs(EthGetLogsParam) ([]EthLog, *RPCProviderError)
}

type ETHRPCPresetMethods struct{}

func (e *ETHRPCPresetMethods) Web3ClientVersion() (string, *RPCProviderError) {
	return unknownStr, nil
}

func (e *ETHRPCPresetMethods) NetVersion() (string, *RPCProviderError) {
	return unknownStr, nil
}

func (e *ETHRPCPresetMethods) ProtocolVersion() (string, *RPCProviderError) {
	return unknownStr, nil
}

func (e *ETHRPCPresetMethods) Syncing() (bool, any, *RPCProviderError) {
	return false, nil, nil
}

func (e *ETHRPCPresetMethods) Coinbase() (types.Address, *RPCProviderError) {
	return types.Address{}, nil
}

func (e *ETHRPCPresetMethods) Mining() (bool, *RPCProviderError) {
	return false, nil
}

func (e *ETHRPCPresetMethods) HashRate() (*types.BigInt, *RPCProviderError) {
	return types.NewBigInt(), nil
}

func (e *ETHRPCPresetMethods) GasPrice() (*types.BigInt, *RPCProviderError) {
	return types.NewBigInt(), nil
}

func (e *ETHRPCPresetMethods) Accounts() ([]types.Address, *RPCProviderError) {
	return []types.Address{}, nil
}

func (e *ETHRPCPresetMethods) StorageAt(types.Address, types.Key, EthBlockNumString) (types.Data, *RPCProviderError) {
	return types.Data{}, ProviderMethodNotSupport
}

func (e *ETHRPCPresetMethods) TransactionCount(types.Address, EthBlockNumString) (types.Uint64, *RPCProviderError) {
	return 0, nil
}

func (e *ETHRPCPresetMethods) BlockTransactionCountByHash(types.Hash) (types.Uint64, *RPCProviderError) {
	return 0, nil
}

func (e *ETHRPCPresetMethods) BlockTransactionCountByNumber(EthBlockNumString) (types.Uint64, *RPCProviderError) {
	return 0, nil
}

func (e *ETHRPCPresetMethods) UncleCountByBlockHash(types.Hash) (types.Uint64, *RPCProviderError) {
	return 0, nil
}

func (e *ETHRPCPresetMethods) UncleCountByBlockNumber(EthBlockNumString) (types.Uint64, *RPCProviderError) {
	return 0, nil
}

func (e *ETHRPCPresetMethods) Code(types.Address, EthBlockNumString) (types.Data, *RPCProviderError) {
	return types.Data{}, ProviderMethodNotSupport
}

func (e *ETHRPCPresetMethods) SendTransaction(EthBasicTransaction) (types.Hash, *RPCProviderError) {
	return types.Hash{}, ProviderMethodNotSupport
}

func (e *ETHRPCPresetMethods) EstimateGas(EthBasicTransaction, EthBlockNumString) (types.Uint64, *RPCProviderError) {
	return 0, nil
}

func (e *ETHRPCPresetMethods) TransactionByBlockHashAndIndex(types.Hash, uint64) (*EthFullTransaction, *RPCProviderError) {
	return nil, ProviderMethodNotSupport
}

func (e *ETHRPCPresetMethods) TransactionByBlockNumberAndIndex(EthBlockNumString, uint64) (*EthFullTransaction, *RPCProviderError) {
	return nil, ProviderMethodNotSupport
}

func (e *ETHRPCPresetMethods) UncleByBlockHashAndIndex(types.Hash, uint64) (*EthBlock, *RPCProviderError) {
	return nil, nil
}

func (e *ETHRPCPresetMethods) UncleByBlockNumberAndIndex(EthBlockNumString, uint64) (*EthBlock, *RPCProviderError) {
	return nil, nil
}
