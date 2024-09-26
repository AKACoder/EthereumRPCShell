package ethereumRPCProvider

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
	Coinbase() (HexAddress, *RPCProviderError)
	ChainId() (HexInt, *RPCProviderError)
	Mining() (bool, *RPCProviderError)
	HashRate() (HexInt, *RPCProviderError)
	GasPrice() (HexInt, *RPCProviderError)
	Accounts() ([]HexAddress, *RPCProviderError)
	BlockNumber() (HexInt, *RPCProviderError)
	Balance(HexAddress, EthBlockNumString) (HexInt, *RPCProviderError)
	StorageAt(HexAddress, HexInt, EthBlockNumString) (HexData, *RPCProviderError)
	TransactionCount(HexAddress, EthBlockNumString) (HexInt, *RPCProviderError)
	BlockTransactionCountByHash(Hash256) (HexInt, *RPCProviderError)
	BlockTransactionCountByNumber(EthBlockNumString) (HexInt, *RPCProviderError)
	UncleCountByBlockHash(Hash256) (HexInt, *RPCProviderError)
	UncleCountByBlockNumber(EthBlockNumString) (HexInt, *RPCProviderError)
	Code(HexAddress, EthBlockNumString) (HexData, *RPCProviderError)
	SendTransaction(EthBasicTransaction) (HexData, *RPCProviderError)
	SendRawTransaction(HexData) (HexData, *RPCProviderError)
	Call(EthBasicTransaction, EthBlockNumString) (HexData, *RPCProviderError)
	EstimateGas(EthBasicTransaction, EthBlockNumString) (HexInt, *RPCProviderError)
	BlockByHash(Hash256, bool) (*EthBlock, *RPCProviderError)
	BlockByNumber(EthBlockNumString, bool) (*EthBlock, *RPCProviderError)
	TransactionByHash(Hash256) (*EthFullTransaction, *RPCProviderError)
	TransactionByBlockHashAndIndex(Hash256, HexInt) (*EthFullTransaction, *RPCProviderError)
	TransactionByBlockNumberAndIndex(EthBlockNumString, HexInt) (*EthFullTransaction, *RPCProviderError)
	TransactionReceipt(Hash256) (*EthTransactionReceipt, *RPCProviderError)
	UncleByBlockHashAndIndex(Hash256, HexInt) (*EthBlock, *RPCProviderError)
	UncleByBlockNumberAndIndex(EthBlockNumString, HexInt) (*EthBlock, *RPCProviderError)
	Logs(EthGetLogsParam) ([]EthLog, *RPCProviderError)
}
