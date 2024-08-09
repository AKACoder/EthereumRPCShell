package utils

import (
	"EthereumRPCShell/common/constants"
	"strconv"
)

func Str2Uint(s string) (uint, error) {
	ret, err := strconv.ParseUint(s, constants.DecimalBase, constants.BitLen32)
	return uint(ret), err
}

func Str2Uint32(s string) (uint32, error) {
	ret, err := strconv.ParseUint(s, constants.DecimalBase, constants.BitLen32)
	return uint32(ret), err
}

func Str2Uint64(s string) (uint64, error) {
	return strconv.ParseUint(s, constants.DecimalBase, constants.BitLen64)
}

func Str2Int32(s string) (int32, error) {
	ret, err := strconv.ParseInt(s, constants.DecimalBase, constants.BitLen32)
	return int32(ret), err
}

func Str2Int64(s string) (int64, error) {
	return strconv.ParseInt(s, constants.DecimalBase, constants.BitLen64)
}

func Str2Int(s string) (int, error) {
	ret, err := strconv.ParseInt(s, constants.DecimalBase, constants.BitLen32)
	return int(ret), err
}
