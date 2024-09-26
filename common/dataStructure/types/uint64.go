package types

import (
	"encoding/binary"
	"fmt"
	"github.com/AKACoder/EthereumRPCShell/common/constants"
	"github.com/AKACoder/EthereumRPCShell/common/shellErrors"
	"github.com/AKACoder/EthereumRPCShell/common/utils"
)

type Uint64 uint64

func (u Uint64) Bytes() []byte {
	return binary.BigEndian.AppendUint64(nil, uint64(u))
}

func (u Uint64) String() string {
	return fmt.Sprintf("0x%x", uint64(u))
}

func (u Uint64) MarshalText() ([]byte, error) {
	return []byte(u.String()), nil
}

func (u *Uint64) SetBytes(b []byte) {
	var iBytes [constants.EVMUint64Length]byte
	utils.BytesCopy(iBytes[:], b)
	*u = Uint64(binary.BigEndian.Uint64(iBytes[:]))
}

func (u *Uint64) UnmarshalText(val []byte) error {
	var hexStr = string(val)
	i, err := utils.HexParamToUint64(hexStr)
	if err == nil {
		*u = Uint64(i)
	}

	return err
}

func (u *Uint64) FromParam(param any) error {
	hexStr, ok := param.(string)
	if !ok {
		return shellErrors.InvalidParameterType
	}

	err := u.UnmarshalText([]byte(hexStr))

	return err
}
