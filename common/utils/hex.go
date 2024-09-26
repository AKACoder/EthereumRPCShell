package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/AKACoder/EthereumRPCShell/common/shellErrors"
	"strconv"
)

func BytesToHex(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}

func HexToBytes(val []byte, receiver *[]byte, allowOdd bool) error {
	valLen := len(val)
	if valLen < 2 {
		return hex.ErrLength
	}

	if valLen == 2 {
		if string(val) == "0x" || string(val) == "0X" {
			*receiver = nil
			return nil
		}

		return hex.ErrLength
	}

	isOdd := len(val)%2 != 0

	if isOdd {
		if allowOdd {
			val = val[1:]
			val[0] = '0'
		} else {
			return hex.ErrLength
		}
	} else {
		val = val[2:]
	}

	data, err := hex.DecodeString(string(val))
	if err != nil {
		return err
	}

	*receiver = data
	return nil
}

func HexParamToUint64(val any) (uint64, error) {
	hexStr, ok := val.(string)
	if !ok {
		fmt.Println("go error here?")
		return 0, shellErrors.InvalidParameterType
	}

	strLen := len(hexStr)
	if strLen < 2 {
		return 0, shellErrors.InvalidParameterType
	}

	if strLen == 2 {
		if hexStr == "0x" || hexStr == "0X" {
			return 0, nil
		}

		return 0, shellErrors.InvalidParameterType
	}

	uintValue, err := strconv.ParseUint(hexStr[2:], 16, 64)
	if err == nil {
		return uintValue, nil
	}

	return 0, err
}

func BytesCopy(dst []byte, data []byte) {
	dstLen := len(dst)
	dataLen := len(data)
	if dataLen > dstLen {
		data = data[dataLen-dstLen:]
	}

	copy(dst[dstLen-len(data):], data)
}
