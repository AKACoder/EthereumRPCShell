package ethereumRPCProvider

import (
	"github.com/AKACoder/EthereumRPCShell/common/dataStructure/types"
)

type EthBasicTransaction struct {
	Type                 HexInt      `json:"type"`
	Nonce                HexInt      `json:"nonce"`
	From                 HexAddress  `json:"from"`
	To                   *HexAddress `json:"to"`
	AccessList           []any       `json:"accessList"`
	Input                HexData     `json:"input"`
	Data                 HexData     `json:"data"`
	Gas                  HexInt      `json:"gas"`
	GasPrice             HexInt      `json:"gasPrice"`
	MaxPriorityFeePerGas *HexInt     `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         *HexInt     `json:"maxFeePerGas"`
	Value                HexInt      `json:"value"`
}

type EthBlock struct {
	Number                *HexInt     `json:"number"`
	Hash                  *Hash256    `json:"hash"`
	ParentHash            *Hash256    `json:"parentHash"`
	Nonce                 *HexInt     `json:"nonce"`
	Sha3Uncles            *Hash256    `json:"sha3Uncles"`
	LogsBloom             *HexData    `json:"logsBloom"`
	TransactionsRoot      *Hash256    `json:"transactionsRoot"`
	StateRoot             *Hash256    `json:"stateRoot"`
	ReceiptsRoot          *Hash256    `json:"receiptsRoot"`
	Miner                 *HexAddress `json:"miner"`
	Difficulty            *HexInt     `json:"difficulty"`
	TotalDifficulty       *HexInt     `json:"totalDifficulty"`
	ExtraData             *HexData    `json:"extraData"`
	MixHash               *HexData    `json:"mixHash"`
	Size                  *HexInt     `json:"size"`
	GasLimit              *HexInt     `json:"gasLimit"`
	GasUsed               *HexInt     `json:"gasUsed"`
	BaseFeePerGas         *HexInt     `json:"baseFeePerGas"`
	Withdrawals           []any       `json:"withdrawals"`
	WithdrawalsRoot       *HexInt     `json:"withdrawalsRoot"`
	BlobGasUsed           *HexInt     `json:"blobGasUsed"`
	ExcessBlobGas         *HexInt     `json:"excessBlobGas"`
	ParentBeaconBlockRoot *HexInt     `json:"parentBeaconBlockRoot"`
	Timestamp             *HexInt     `json:"timestamp"`
	Transactions          []any       `json:"transactions"`
	Uncles                []Hash256   `json:"uncles"`
}

type EthFullTransaction struct {
	BlockHash   *Hash256 `json:"blockHash"`
	BlockNumber *HexInt  `json:"blockNumber"`
	ChainId     *HexInt  `json:"chainId"`

	EthBasicTransaction
	Hash             Hash256 `json:"hash"`
	TransactionIndex HexInt  `json:"transactionIndex"`
	YParity          HexInt  `json:"yParity"`
	V                HexData `json:"v"`
	R                HexData `json:"r"`
	S                HexData `json:"s"`
}

type EthLog struct {
	Removed          bool        `json:"removed"`
	LogIndex         *HexInt     `json:"logIndex"`
	TransactionIndex *HexInt     `json:"transactionIndex"`
	TransactionHash  *Hash256    `json:"transactionHash"`
	BlockHash        *Hash256    `json:"blockHash"`
	BlockNumber      *HexInt     `json:"blockNumber"`
	Address          *HexAddress `json:"address"`
	Data             HexData     `json:"data"`
	Topics           []HexData   `json:"topics"`
}

type EthTransactionReceipt struct {
	BlockHash         *Hash256    `json:"blockHash"`
	BlockNumber       *HexInt     `json:"blockNumber"`
	ChainId           *HexInt     `json:"chainId"`
	Type              HexInt      `json:"type"`
	TxHash            Hash256     `json:"transactionHash"`
	TxIndex           Hash256     `json:"transactionIndex"`
	From              HexAddress  `json:"from"`
	To                *HexAddress `json:"to"`
	ContractAddress   HexAddress  `json:"contractAddress"`
	Logs              []EthLog    `json:"logs"`
	LogsBloom         HexData     `json:"logsBloom"`
	Root              HexData     `json:"root"`
	GasUsed           HexInt      `json:"gasUsed"`
	CumulativeGasUsed HexInt      `json:"cumulativeGasUsed"`
	EffectiveGasPrice HexInt      `json:"effectiveGasPrice"`
	BlobGasPrice      HexInt      `json:"blobGasPrice"`
	Status            HexInt      `json:"status"`
}

type EthGetLogsParam struct {
	FromBlock *EthBlockNumString `json:"fromBlock"`
	ToBlock   *EthBlockNumString `json:"toBlock"`
	Address   any                `json:"address"`
	Topics    []Hash256          `json:"topics"`
	BlockHash *Hash256           `json:"blockhash"`
}

// IEthereumRPCProvider
// known unsupported rpc method:
// web3_sha3, net_listening, net_peerCount, eth_protocolVersion,
// eth_sign, eth_signTransaction,
// eth_newFilter, eth_newBlockFilter, eth_newPendingTransactionFilter, eth_uninstallFilter,
// eth_getFilterChanges, eth_getFilterLogs
type IEthereumRPCProvider interface {
	SupportCheck(string) bool
	Web3ClientVersion() (string, *RPCProviderError)
	NetVersion() (string, *RPCProviderError)
	ProtocolVersion() (string, *RPCProviderError)
	Syncing() (bool, any, *RPCProviderError)
	Coinbase() (types.Address, *RPCProviderError)
	ChainId() (*types.BigInt, *RPCProviderError)
	Mining() (bool, *RPCProviderError)
	HashRate() (*types.BigInt, *RPCProviderError)
	GasPrice() (*types.BigInt, *RPCProviderError)
	Accounts() ([]types.Address, *RPCProviderError)
	BlockNumber() (uint64, *RPCProviderError)
	Balance(addr types.Address, blk EthBlockNumString) (*types.BigInt, *RPCProviderError)
	StorageAt(addr types.Address, key types.Key, blk EthBlockNumString) (types.Data, *RPCProviderError)
	TransactionCount(addr types.Address, blk EthBlockNumString) (uint64, *RPCProviderError)
	BlockTransactionCountByHash(hash types.Hash) (uint64, *RPCProviderError)
	BlockTransactionCountByNumber(blk EthBlockNumString) (uint64, *RPCProviderError)
	UncleCountByBlockHash(blkHash types.Hash) (uint64, *RPCProviderError)
	UncleCountByBlockNumber(blk EthBlockNumString) (uint64, *RPCProviderError)
	Code(addr types.Address, blk EthBlockNumString) (types.Data, *RPCProviderError)
	SendTransaction(tx EthBasicTransaction) (types.Hash, *RPCProviderError)
	SendRawTransaction(data types.Data) (types.Hash, *RPCProviderError)
	Call(tx EthBasicTransaction, blk EthBlockNumString) (types.Data, *RPCProviderError)
	EstimateGas(tx EthBasicTransaction, blk EthBlockNumString) (uint64, *RPCProviderError)
	BlockByHash(blkHash types.Hash, fullTx bool) (*EthBlock, *RPCProviderError)
	BlockByNumber(blk EthBlockNumString, fullTx bool) (*EthBlock, *RPCProviderError)
	TransactionByHash(txHash types.Hash) (*EthFullTransaction, *RPCProviderError)
	TransactionByBlockHashAndIndex(hash types.Hash, idx uint64) (*EthFullTransaction, *RPCProviderError)
	TransactionByBlockNumberAndIndex(blk EthBlockNumString, idx uint64) (*EthFullTransaction, *RPCProviderError)
	TransactionReceipt(txHash types.Hash) (*EthTransactionReceipt, *RPCProviderError)
	UncleByBlockHashAndIndex(blkHash types.Hash, idx uint64) (*EthBlock, *RPCProviderError)
	UncleByBlockNumberAndIndex(blk EthBlockNumString, idx uint64) (*EthBlock, *RPCProviderError)
	Logs(EthGetLogsParam) ([]EthLog, *RPCProviderError)
}
