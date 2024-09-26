package types

import (
	"bytes"
	"github.com/AKACoder/EthereumRPCShell/common/constants"
	"github.com/AKACoder/EthereumRPCShell/common/shellErrors"
	"github.com/AKACoder/EthereumRPCShell/common/utils"
)

type Key [constants.EVMIntLength]byte

func (k Key) Bytes() []byte {
	return bytes.Clone(k[:])
}

func (k Key) String() string {
	return utils.BytesToHex(k[:])
}

func (k Key) MarshalText() ([]byte, error) {
	return []byte(k.String()), nil
}

func (k *Key) SetBytes(b []byte) {
	if len(b) > constants.EVMIntLength {
		b = b[len(b)-constants.EVMIntLength:]
	}
	copy(k[constants.EVMIntLength-len(b):], b)
}

func (k *Key) UnmarshalText(val []byte) error {
	var data []byte
	err := utils.HexToBytes(val, &data, false)
	if err == nil {
		k.SetBytes(data)
	}

	return err
}

func (k *Key) FromParam(param any) error {
	hexStr, ok := param.(string)
	if !ok {
		return shellErrors.InvalidParameterType
	}

	if len(hexStr) != constants.EVMIntHexLength {
		return shellErrors.InvalidParameterType
	}

	err := k.UnmarshalText([]byte(hexStr))

	return err
}
