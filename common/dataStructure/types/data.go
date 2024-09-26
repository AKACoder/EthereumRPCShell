package types

import (
	"bytes"
	"github.com/AKACoder/EthereumRPCShell/common/shellErrors"
	"github.com/AKACoder/EthereumRPCShell/common/utils"
)

type Data []byte

func (d Data) Bytes() []byte {
	return bytes.Clone(d)
}

func (d Data) String() string {
	return utils.BytesToHex(d)
}

func (d Data) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Data) SetBytes(b []byte) {
	*d = bytes.Clone(b)
}

func (d *Data) UnmarshalText(val []byte) error {
	var data []byte
	err := utils.HexToBytes(val, &data, false)
	if err == nil {
		d.SetBytes(data)
	}

	return err
}

func (d *Data) FromParam(param any) error {
	hexStr, ok := param.(string)
	if !ok {
		return shellErrors.InvalidParameterType
	}

	err := d.UnmarshalText([]byte(hexStr))

	return err
}
