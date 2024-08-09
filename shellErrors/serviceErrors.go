package shellErrors

import "EthereumRPCShell/common/dataStructure/enum"

var ServiceErrors = struct {
	NoData           enum.ErrorElement `msg:"no data"`
	Internal         enum.ErrorElement `msg:"internal error"`
	Middleware       enum.ErrorElement `msg:"middleware error"`
	Unknown          enum.ErrorElement `msg:"unknown error"`
	InvalidParam     enum.ErrorElement `msg:"invalid parameter"`
	InvalidSignature enum.ErrorElement `msg:"invalid signature"`
}{}

func loadServiceErrors(startCode int64) {
	enum.BuildErrorEnum(&ServiceErrors, startCode)
}
