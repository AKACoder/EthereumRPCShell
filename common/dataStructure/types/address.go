package types

import (
	"bytes"
	"github.com/AKACoder/EthereumRPCShell/common/constants"
	"github.com/AKACoder/EthereumRPCShell/common/shellErrors"
	"github.com/AKACoder/EthereumRPCShell/common/utils"
)

type Address [constants.EVMAddressLength]byte

func (a Address) Bytes() []byte {
	return bytes.Clone(a[:])
}

func (a Address) String() string {
	return utils.BytesToHex(a[:])
}

func (a Address) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}

func (a *Address) SetBytes(b []byte) {
	if len(b) > constants.EVMAddressLength {
		b = b[len(b)-constants.EVMAddressLength:]
	}
	copy(a[constants.EVMAddressLength-len(b):], b)
}

func (a *Address) UnmarshalText(val []byte) error {
	var data []byte
	err := utils.HexToBytes(val, &data, false)
	if err == nil {
		a.SetBytes(data)
	}

	return err
}

func (a *Address) FromParam(param any) error {
	hexStr, ok := param.(string)
	if !ok {
		return shellErrors.InvalidParameterType
	}

	if len(hexStr) != constants.EVMAddressHexLength {
		return shellErrors.InvalidParameterType
	}

	err := a.UnmarshalText([]byte(hexStr))

	return err
}
