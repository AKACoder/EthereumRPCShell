package types

import (
	"bytes"
	"github.com/AKACoder/EthereumRPCShell/common/constants"
	"github.com/AKACoder/EthereumRPCShell/common/shellErrors"
	"github.com/AKACoder/EthereumRPCShell/common/utils"
)

type Hash [constants.EVMHashLength]byte

func (h Hash) Bytes() []byte {
	return bytes.Clone(h[:])
}

func (h Hash) String() string {
	return utils.BytesToHex(h[:])
}

func (h Hash) MarshalText() ([]byte, error) {
	return []byte(h.String()), nil
}

func (h *Hash) SetBytes(b []byte) {
	if len(b) > constants.EVMHashLength {
		b = b[len(b)-constants.EVMHashLength:]
	}
	copy(h[constants.EVMHashLength-len(b):], b)
}

func (h *Hash) UnmarshalText(val []byte) error {
	var data []byte
	err := utils.HexToBytes(val, &data, false)
	if err == nil {
		h.SetBytes(data)
	}

	return err
}

func (h *Hash) FromParam(param any) error {
	hexStr, ok := param.(string)
	if !ok {
		return shellErrors.InvalidParameterType
	}

	if len(hexStr) != constants.EVMHashHexLength {
		return shellErrors.InvalidParameterType
	}

	err := h.UnmarshalText([]byte(hexStr))

	return err
}
