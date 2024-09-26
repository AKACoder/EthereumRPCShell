package types

import (
	"github.com/AKACoder/EthereumRPCShell/common/constants"
	"github.com/AKACoder/EthereumRPCShell/common/utils"
	"math/big"
)

type BigInt struct {
	i *big.Int
}

func NewBigInt() *BigInt {
	return &BigInt{
		big.NewInt(0),
	}
}

func BigIntOf(i *big.Int) *BigInt {
	ret := NewBigInt()
	if i != nil {
		ret.i.Set(i)
	}

	return ret
}

func (b BigInt) String() string {
	return b.i.Text(constants.DecimalBase)
}

func (b BigInt) Int() *big.Int {
	ret := big.NewInt(0).Set(b.i)
	return ret
}

func (b BigInt) Uint64() uint64 {
	return b.i.Uint64()
}

func (b *BigInt) MarshalText() ([]byte, error) {
	return []byte("0x" + b.i.Text(constants.HexBase)), nil
}

func (b *BigInt) SetUint64(x uint64) *BigInt {
	b.i.SetUint64(x)
	return b
}

func (b *BigInt) SetBytes(d []byte) *BigInt {
	if len(d) > constants.EVMIntLength {
		d = d[len(d)-constants.EVMIntLength:]
	}

	b.i.SetBytes(d)

	return b
}

func (b *BigInt) UnmarshalText(val []byte) error {
	var data []byte
	err := utils.HexToBytes(val, &data, true)
	if err != nil {
		return err
	}

	b.i.SetBytes(data)

	return nil
}
