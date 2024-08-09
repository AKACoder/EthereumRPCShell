package ethereumClientProvider

import (
	. "EthereumRPCShell/common/constants"
)

const (
	unknownStr   = "unknown"
	zeroStr      = "0x00"
	emptyDataStr = "0x"
)

func mustImplementErr(method string) *EthClientError {
	return &EthClientError{
		Code:    -932004,
		Message: "method " + method + " must be implemented by the client, but it is not.",
	}
}

type IMustImplementMethods interface {
	SupportCheck(name string) bool
	ChainId() (HexInt, *EthClientError)
	BlockNumber() (HexInt, *EthClientError)
	Balance(HexAddress, EthBlockNumString) (HexInt, *EthClientError)
	SendRawTransaction(HexData) (HexData, *EthClientError)
	Call(EthBasicTransaction, EthBlockNumString) (HexData, *EthClientError)
	BlockByHash(Hash256, bool) (*EthBlock, *EthClientError)
	BlockByNumber(EthBlockNumString, bool) (*EthBlock, *EthClientError)
	TransactionByHash(Hash256) (*EthFullTransaction, *EthClientError)
	TransactionReceipt(Hash256) (*EthTransactionReceipt, *EthClientError)
	Logs(EthGetLogsParam) ([]EthLog, *EthClientError)
}

type ETHClientPresetMethods struct{}

func (e *ETHClientPresetMethods) Web3ClientVersion() (string, *EthClientError) {
	return unknownStr, nil
}

func (e *ETHClientPresetMethods) NetVersion() (string, *EthClientError) {
	return unknownStr, nil
}

func (e *ETHClientPresetMethods) ProtocolVersion() (string, *EthClientError) {
	return unknownStr, nil
}

func (e *ETHClientPresetMethods) Syncing() (bool, any, *EthClientError) {
	return false, nil, nil
}

func (e *ETHClientPresetMethods) Coinbase() (HexAddress, *EthClientError) {
	return ETHAddressZero, nil
}

func (e *ETHClientPresetMethods) Mining() (bool, *EthClientError) {
	return false, nil
}

func (e *ETHClientPresetMethods) HashRate() (HexInt, *EthClientError) {
	return zeroStr, nil
}

func (e *ETHClientPresetMethods) GasPrice() (HexInt, *EthClientError) {
	return zeroStr, nil
}

func (e *ETHClientPresetMethods) Accounts() ([]HexAddress, *EthClientError) {
	return []HexAddress{}, nil
}

func (e *ETHClientPresetMethods) StorageAt(HexAddress, HexInt, EthBlockNumString) (HexData, *EthClientError) {
	return emptyDataStr, ClientMethodNotSupport
}

func (e *ETHClientPresetMethods) TransactionCount(HexAddress, EthBlockNumString) (HexInt, *EthClientError) {
	return zeroStr, nil
}

func (e *ETHClientPresetMethods) BlockTransactionCountByHash(Hash256) (HexInt, *EthClientError) {
	return zeroStr, nil
}

func (e *ETHClientPresetMethods) BlockTransactionCountByNumber(EthBlockNumString) (HexInt, *EthClientError) {
	return zeroStr, nil
}

func (e *ETHClientPresetMethods) UncleCountByBlockHash(Hash256) (HexInt, *EthClientError) {
	return zeroStr, nil
}

func (e *ETHClientPresetMethods) UncleCountByBlockNumber(EthBlockNumString) (HexInt, *EthClientError) {
	return zeroStr, nil
}

func (e *ETHClientPresetMethods) Code(HexAddress, EthBlockNumString) (HexData, *EthClientError) {
	return emptyDataStr, ClientMethodNotSupport
}

func (e *ETHClientPresetMethods) SendTransaction(EthBasicTransaction) (HexData, *EthClientError) {
	return emptyDataStr, ClientMethodNotSupport
}

func (e *ETHClientPresetMethods) EstimateGas(EthBasicTransaction, EthBlockNumString) (HexInt, *EthClientError) {
	return zeroStr, nil
}

func (e *ETHClientPresetMethods) TransactionByBlockHashAndIndex(Hash256, HexInt) (*EthFullTransaction, *EthClientError) {
	return nil, ClientMethodNotSupport
}

func (e *ETHClientPresetMethods) TransactionByBlockNumberAndIndex(EthBlockNumString, HexInt) (*EthFullTransaction, *EthClientError) {
	return nil, ClientMethodNotSupport
}

func (e *ETHClientPresetMethods) UncleByBlockHashAndIndex(Hash256, HexInt) (*EthBlock, *EthClientError) {
	return nil, nil
}

func (e *ETHClientPresetMethods) UncleByBlockNumberAndIndex(EthBlockNumString, HexInt) (*EthBlock, *EthClientError) {
	return nil, nil
}
