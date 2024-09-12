package ethereumRPCProvider

import . "github.com/AKACoder/EthereumRPCShell/common/constants"

const (
	unknownStr   = "unknown"
	zeroStr      = "0x00"
	emptyDataStr = "0x"
)

type IMustImplementMethods interface {
	SupportCheck(name string) bool
	ChainId() (HexInt, *RPCProviderError)
	BlockNumber() (HexInt, *RPCProviderError)
	Balance(HexAddress, EthBlockNumString) (HexInt, *RPCProviderError)
	SendRawTransaction(HexData) (HexData, *RPCProviderError)
	Call(EthBasicTransaction, EthBlockNumString) (HexData, *RPCProviderError)
	BlockByHash(Hash256, bool) (*EthBlock, *RPCProviderError)
	BlockByNumber(EthBlockNumString, bool) (*EthBlock, *RPCProviderError)
	TransactionByHash(Hash256) (*EthFullTransaction, *RPCProviderError)
	TransactionReceipt(Hash256) (*EthTransactionReceipt, *RPCProviderError)
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

func (e *ETHRPCPresetMethods) Coinbase() (HexAddress, *RPCProviderError) {
	return ETHAddressZero, nil
}

func (e *ETHRPCPresetMethods) Mining() (bool, *RPCProviderError) {
	return false, nil
}

func (e *ETHRPCPresetMethods) HashRate() (HexInt, *RPCProviderError) {
	return zeroStr, nil
}

func (e *ETHRPCPresetMethods) GasPrice() (HexInt, *RPCProviderError) {
	return zeroStr, nil
}

func (e *ETHRPCPresetMethods) Accounts() ([]HexAddress, *RPCProviderError) {
	return []HexAddress{}, nil
}

func (e *ETHRPCPresetMethods) StorageAt(HexAddress, HexInt, EthBlockNumString) (HexData, *RPCProviderError) {
	return emptyDataStr, ProviderMethodNotSupport
}

func (e *ETHRPCPresetMethods) TransactionCount(HexAddress, EthBlockNumString) (HexInt, *RPCProviderError) {
	return zeroStr, nil
}

func (e *ETHRPCPresetMethods) BlockTransactionCountByHash(Hash256) (HexInt, *RPCProviderError) {
	return zeroStr, nil
}

func (e *ETHRPCPresetMethods) BlockTransactionCountByNumber(EthBlockNumString) (HexInt, *RPCProviderError) {
	return zeroStr, nil
}

func (e *ETHRPCPresetMethods) UncleCountByBlockHash(Hash256) (HexInt, *RPCProviderError) {
	return zeroStr, nil
}

func (e *ETHRPCPresetMethods) UncleCountByBlockNumber(EthBlockNumString) (HexInt, *RPCProviderError) {
	return zeroStr, nil
}

func (e *ETHRPCPresetMethods) Code(HexAddress, EthBlockNumString) (HexData, *RPCProviderError) {
	return emptyDataStr, ProviderMethodNotSupport
}

func (e *ETHRPCPresetMethods) SendTransaction(EthBasicTransaction) (HexData, *RPCProviderError) {
	return emptyDataStr, ProviderMethodNotSupport
}

func (e *ETHRPCPresetMethods) EstimateGas(EthBasicTransaction, EthBlockNumString) (HexInt, *RPCProviderError) {
	return zeroStr, nil
}

func (e *ETHRPCPresetMethods) TransactionByBlockHashAndIndex(Hash256, HexInt) (*EthFullTransaction, *RPCProviderError) {
	return nil, ProviderMethodNotSupport
}

func (e *ETHRPCPresetMethods) TransactionByBlockNumberAndIndex(EthBlockNumString, HexInt) (*EthFullTransaction, *RPCProviderError) {
	return nil, ProviderMethodNotSupport
}

func (e *ETHRPCPresetMethods) UncleByBlockHashAndIndex(Hash256, HexInt) (*EthBlock, *RPCProviderError) {
	return nil, nil
}

func (e *ETHRPCPresetMethods) UncleByBlockNumberAndIndex(EthBlockNumString, HexInt) (*EthBlock, *RPCProviderError) {
	return nil, nil
}
