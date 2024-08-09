package shellErrors

import "EthereumRPCShell/common/dataStructure/enum"

var CommonErrors = struct {
	InvalidHexPrefix enum.ErrorElement `msg:"Invalid hex prefix"`
	InvalidHexString enum.ErrorElement `msg:"Invalid hex string"`
	InvalidHexLength enum.ErrorElement `msg:"expected length does not match"`
}{}

func loadCommonErrors(startCode int64) {
	enum.BuildErrorEnum(&CommonErrors, startCode)
}
