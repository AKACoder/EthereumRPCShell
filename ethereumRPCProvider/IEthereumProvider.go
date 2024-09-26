package ethereumRPCProvider

import (
	"github.com/AKACoder/EthereumRPCShell/common/dataStructure/types"
)

type EthBasicTransaction struct {
	Type                 types.Uint64   `json:"type"`
	Nonce                types.Uint64   `json:"nonce"`
	From                 types.Address  `json:"from"`
	To                   *types.Address `json:"to"`
	AccessList           []any          `json:"accessList"`
	Input                types.Data     `json:"input"`
	Data                 types.Data     `json:"data"`
	Gas                  types.Uint64   `json:"gas"`
	GasPrice             *types.BigInt  `json:"gasPrice"`
	MaxPriorityFeePerGas *types.BigInt  `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         *types.BigInt  `json:"maxFeePerGas"`
	Value                *types.BigInt  `json:"value"`
}

type EthBlock struct {
	Number                types.Uint64   `json:"number"`
	Hash                  types.Hash     `json:"hash"`
	ParentHash            types.Hash     `json:"parentHash"`
	Nonce                 types.Data     `json:"nonce"`
	Sha3Uncles            *types.Hash    `json:"sha3Uncles"`
	LogsBloom             *types.Data    `json:"logsBloom"`
	TransactionsRoot      types.Hash     `json:"transactionsRoot"`
	StateRoot             types.Hash     `json:"stateRoot"`
	ReceiptsRoot          types.Hash     `json:"receiptsRoot"`
	Miner                 *types.Address `json:"miner"`
	Difficulty            types.Data     `json:"difficulty"`
	TotalDifficulty       types.Data     `json:"totalDifficulty"`
	ExtraData             types.Data     `json:"extraData"`
	MixHash               types.Data     `json:"mixHash"`
	Size                  types.Uint64   `json:"size"`
	GasLimit              types.Uint64   `json:"gasLimit"`
	GasUsed               types.Uint64   `json:"gasUsed"`
	BaseFeePerGas         *types.BigInt  `json:"baseFeePerGas"`
	Withdrawals           []any          `json:"withdrawals"`
	WithdrawalsRoot       *types.Hash    `json:"withdrawalsRoot"`
	BlobGasUsed           types.Uint64   `json:"blobGasUsed"`
	ExcessBlobGas         types.Uint64   `json:"excessBlobGas"`
	ParentBeaconBlockRoot *types.Hash    `json:"parentBeaconBlockRoot"`
	Timestamp             types.Uint64   `json:"timestamp"`
	Transactions          []any          `json:"transactions"`
	Uncles                []types.Hash   `json:"uncles"`
}

type EthFullTransaction struct {
	BlockHash   types.Hash    `json:"blockHash"`
	BlockNumber types.Uint64  `json:"blockNumber"`
	ChainId     *types.BigInt `json:"chainId"`

	EthBasicTransaction
	Hash             types.Hash   `json:"hash"`
	TransactionIndex types.Uint64 `json:"transactionIndex"`
	YParity          types.Uint64 `json:"yParity"`
	V                types.Uint64 `json:"v"`
	R                types.Data   `json:"r"`
	S                types.Data   `json:"s"`
}

type EthLog struct {
	Removed          bool          `json:"removed"`
	LogIndex         types.Uint64  `json:"logIndex"`
	TransactionIndex types.Uint64  `json:"transactionIndex"`
	TransactionHash  types.Hash    `json:"transactionHash"`
	BlockHash        types.Hash    `json:"blockHash"`
	BlockNumber      types.Uint64  `json:"blockNumber"`
	Address          types.Address `json:"address"`
	Data             types.Data    `json:"data"`
	Topics           []types.Data  `json:"topics"`
}

type EthTransactionReceipt struct {
	Root              *types.Hash    `json:"root,omitempty"`
	BlockHash         types.Hash     `json:"blockHash"`
	BlockNumber       types.Uint64   `json:"blockNumber"`
	ChainId           *types.BigInt  `json:"chainId"`
	Type              types.Uint64   `json:"type"`
	TxHash            types.Hash     `json:"transactionHash"`
	TxIndex           types.Uint64   `json:"transactionIndex"`
	From              types.Address  `json:"from"`
	To                *types.Address `json:"to"`
	ContractAddress   *types.Address `json:"contractAddress"`
	Logs              []EthLog       `json:"logs"`
	LogsBloom         types.Data     `json:"logsBloom"`
	GasUsed           types.Uint64   `json:"gasUsed"`
	CumulativeGasUsed types.Uint64   `json:"cumulativeGasUsed"`
	EffectiveGasPrice *types.BigInt  `json:"effectiveGasPrice"`
	BlobGasPrice      *types.BigInt  `json:"blobGasPrice"`
	Status            types.Uint64   `json:"status"`
}

type EthGetLogsParam struct {
	FromBlock *EthBlockNumString `json:"fromBlock"`
	ToBlock   *EthBlockNumString `json:"toBlock"`
	Address   any                `json:"address"`
	Topics    []types.Hash       `json:"topics"`
	BlockHash *types.Hash        `json:"blockhash"`
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
	BlockNumber() (types.Uint64, *RPCProviderError)
	Balance(addr types.Address, blk EthBlockNumString) (*types.BigInt, *RPCProviderError)
	StorageAt(addr types.Address, key types.Key, blk EthBlockNumString) (types.Data, *RPCProviderError)
	TransactionCount(addr types.Address, blk EthBlockNumString) (types.Uint64, *RPCProviderError)
	BlockTransactionCountByHash(hash types.Hash) (types.Uint64, *RPCProviderError)
	BlockTransactionCountByNumber(blk EthBlockNumString) (types.Uint64, *RPCProviderError)
	UncleCountByBlockHash(blkHash types.Hash) (types.Uint64, *RPCProviderError)
	UncleCountByBlockNumber(blk EthBlockNumString) (types.Uint64, *RPCProviderError)
	Code(addr types.Address, blk EthBlockNumString) (types.Data, *RPCProviderError)
	SendTransaction(tx EthBasicTransaction) (types.Hash, *RPCProviderError)
	SendRawTransaction(data types.Data) (types.Hash, *RPCProviderError)
	Call(tx EthBasicTransaction, blk EthBlockNumString) (types.Data, *RPCProviderError)
	EstimateGas(tx EthBasicTransaction, blk EthBlockNumString) (types.Uint64, *RPCProviderError)
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
