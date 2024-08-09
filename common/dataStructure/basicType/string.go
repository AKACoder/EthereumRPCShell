package basicType

import (
	"encoding/hex"
	"github.com/AKACoder/EthereumRPCShell/common/wLog"
	"strings"
)

type SearchKey string

func (s SearchKey) Trim() string {
	trimmed := strings.TrimSpace(string(s))
	if len(trimmed) == 0 {
		return ""
	}
	base := []rune(trimmed)
	if len(base) > 20 {
		return string(base[:20])
	}

	return string(base)
}

func (s SearchKey) Valid() bool {
	return true
}

type EthAddress string

func (e EthAddress) Valid() bool {
	return ValidHex(string(e), 20, false)
}

func (e EthAddress) Str() string {
	return string(e)
}

func ValidHex(hexStr string, byteLen int, allowBlank bool) bool {
	hexLen := len(hexStr)

	if hexLen == 0 {
		return allowBlank
	}

	if hexLen%2 != 0 {
		return false
	}

	bytes, err := hex.DecodeString(hexStr[2:])
	if err != nil {
		wLog.Log.Warn("decode hex failed:", err)
		return false
	}

	if len(bytes) != byteLen {
		return false
	}

	return true
}
