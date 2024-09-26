package constants

const (
	EVMAddressLength    = 20
	EVMAddressHexLength = 42
	EVMHashLength       = 32
	EVMHashHexLength    = 66
	EVMIntLength        = 32
	EVMIntHexLength     = 32
	EVMUint64Length     = 8
	DefaultBlock        = "latest"
)

var ETHRPCNamesMap = map[string]bool{
	"web3_clientVersion":                      true,
	"web3_sha3":                               true,
	"net_version":                             true,
	"net_listening":                           true,
	"net_peerCount":                           true,
	"eth_protocolVersion":                     true,
	"eth_syncing":                             true,
	"eth_coinbase":                            true,
	"eth_chainId":                             true,
	"eth_mining":                              true,
	"eth_hashrate":                            true,
	"eth_gasPrice":                            true,
	"eth_accounts":                            true,
	"eth_blockNumber":                         true,
	"eth_getBalance":                          true,
	"eth_getStorageAt":                        true,
	"eth_getTransactionCount":                 true,
	"eth_getBlockTransactionCountByHash":      true,
	"eth_getBlockTransactionCountByNumber":    true,
	"eth_getUncleCountByBlockHash":            true,
	"eth_getUncleCountByBlockNumber":          true,
	"eth_getCode":                             true,
	"eth_sign":                                true,
	"eth_signTransaction":                     true,
	"eth_sendTransaction":                     true,
	"eth_sendRawTransaction":                  true,
	"eth_call":                                true,
	"eth_estimateGas":                         true,
	"eth_getBlockByHash":                      true,
	"eth_getBlockByNumber":                    true,
	"eth_getTransactionByHash":                true,
	"eth_getTransactionByBlockHashAndIndex":   true,
	"eth_getTransactionByBlockNumberAndIndex": true,
	"eth_getTransactionReceipt":               true,
	"eth_getUncleByBlockHashAndIndex":         true,
	"eth_getUncleByBlockNumberAndIndex":       true,
	"eth_newFilter":                           true,
	"eth_newBlockFilter":                      true,
	"eth_newPendingTransactionFilter":         true,
	"eth_uninstallFilter":                     true,
	"eth_getFilterChanges":                    true,
	"eth_getFilterLogs":                       true,
	"eth_getLogs":                             true,
}

const (
	Method_web3_clientVersion                      = "web3_clientVersion"
	Method_web3_sha3                               = "web3_sha3"
	Method_net_version                             = "net_version"
	Method_net_listening                           = "net_listening"
	Method_net_peerCount                           = "net_peerCount"
	Method_eth_protocolVersion                     = "eth_protocolVersion"
	Method_eth_syncing                             = "eth_syncing"
	Method_eth_coinbase                            = "eth_coinbase"
	Method_eth_chainId                             = "eth_chainId"
	Method_eth_mining                              = "eth_mining"
	Method_eth_hashrate                            = "eth_hashrate"
	Method_eth_gasPrice                            = "eth_gasPrice"
	Method_eth_accounts                            = "eth_accounts"
	Method_eth_blockNumber                         = "eth_blockNumber"
	Method_eth_getBalance                          = "eth_getBalance"
	Method_eth_getStorageAt                        = "eth_getStorageAt"
	Method_eth_getTransactionCount                 = "eth_getTransactionCount"
	Method_eth_getBlockTransactionCountByHash      = "eth_getBlockTransactionCountByHash"
	Method_eth_getBlockTransactionCountByNumber    = "eth_getBlockTransactionCountByNumber"
	Method_eth_getUncleCountByBlockHash            = "eth_getUncleCountByBlockHash"
	Method_eth_getUncleCountByBlockNumber          = "eth_getUncleCountByBlockNumber"
	Method_eth_getCode                             = "eth_getCode"
	Method_eth_sign                                = "eth_sign"
	Method_eth_signTransaction                     = "eth_signTransaction"
	Method_eth_sendTransaction                     = "eth_sendTransaction"
	Method_eth_sendRawTransaction                  = "eth_sendRawTransaction"
	Method_eth_call                                = "eth_call"
	Method_eth_estimateGas                         = "eth_estimateGas"
	Method_eth_getBlockByHash                      = "eth_getBlockByHash"
	Method_eth_getBlockByNumber                    = "eth_getBlockByNumber"
	Method_eth_getTransactionByHash                = "eth_getTransactionByHash"
	Method_eth_getTransactionByBlockHashAndIndex   = "eth_getTransactionByBlockHashAndIndex"
	Method_eth_getTransactionByBlockNumberAndIndex = "eth_getTransactionByBlockNumberAndIndex"
	Method_eth_getTransactionReceipt               = "eth_getTransactionReceipt"
	Method_eth_getUncleByBlockHashAndIndex         = "eth_getUncleByBlockHashAndIndex"
	Method_eth_getUncleByBlockNumberAndIndex       = "eth_getUncleByBlockNumberAndIndex"
	Method_eth_newFilter                           = "eth_newFilter"
	Method_eth_newBlockFilter                      = "eth_newBlockFilter"
	Method_eth_newPendingTransactionFilter         = "eth_newPendingTransactionFilter"
	Method_eth_uninstallFilter                     = "eth_uninstallFilter"
	Method_eth_getFilterChanges                    = "eth_getFilterChanges"
	Method_eth_getFilterLogs                       = "eth_getFilterLogs"
	Method_eth_getLogs                             = "eth_getLogs"
)
