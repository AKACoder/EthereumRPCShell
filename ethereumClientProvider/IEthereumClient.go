package ethereumClientProvider

type EthBasicTransaction struct {
	From                 HexAddress `json:"from"`
	To                   HexAddress `json:"to"`
	Input                HexData    `json:"input"`
	Gas                  HexInt     `json:"gas"`
	GasPrice             HexInt     `json:"gasPrice"`
	MaxPriorityFeePerGas *HexInt    `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         *HexInt    `json:"maxFeePerGas"`
	Value                HexInt     `json:"value"`
	Nonce                HexInt     `json:"nonce"`
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
	EthFullTransaction
	CumulativeGasUsed HexInt     `json:"cumulativeGasUsed"`
	EffectiveGasPrice HexInt     `json:"effectiveGasPrice"`
	GasUsed           HexInt     `json:"gasUsed"`
	ContractAddress   HexAddress `json:"contractAddress"`
	Logs              []EthLog   `json:"logs"`
	LogsBloom         HexData    `json:"logsBloom"`
	Type              HexInt     `json:"type"`
	Root              HexData    `json:"root"`
	Status            HexInt     `json:"status"`
}

type EthGetLogsParam struct {
	FromBlock *EthBlockNumString `json:"fromBlock"`
	ToBlock   *EthBlockNumString `json:"toBlock"`
	Address   any                `json:"address"`
	Topics    []Hash256          `json:"topics"`
	BlockHash *Hash256           `json:"blockhash"`
}

// IEthereumClient
// known unsupported rpc method:
// web3_sha3, net_listening, net_peerCount, eth_protocolVersion,
// eth_sign, eth_signTransaction,
// eth_newFilter, eth_newBlockFilter, eth_newPendingTransactionFilter, eth_uninstallFilter,
// eth_getFilterChanges, eth_getFilterLogs
type IEthereumClient interface {
	SupportCheck(string) bool
	Web3ClientVersion() (string, *EthClientError)
	NetVersion() (string, *EthClientError)
	ProtocolVersion() (string, *EthClientError)
	Syncing() (bool, any, *EthClientError)
	Coinbase() (HexAddress, *EthClientError)
	ChainId() (HexInt, *EthClientError)
	Mining() (bool, *EthClientError)
	HashRate() (HexInt, *EthClientError)
	GasPrice() (HexInt, *EthClientError)
	Accounts() ([]HexAddress, *EthClientError)
	BlockNumber() (HexInt, *EthClientError)
	Balance(HexAddress, EthBlockNumString) (HexInt, *EthClientError)
	StorageAt(HexAddress, HexInt, EthBlockNumString) (HexData, *EthClientError)
	TransactionCount(HexAddress, EthBlockNumString) (HexInt, *EthClientError)
	BlockTransactionCountByHash(Hash256) (HexInt, *EthClientError)
	BlockTransactionCountByNumber(EthBlockNumString) (HexInt, *EthClientError)
	UncleCountByBlockHash(Hash256) (HexInt, *EthClientError)
	UncleCountByBlockNumber(EthBlockNumString) (HexInt, *EthClientError)
	Code(HexAddress, EthBlockNumString) (HexData, *EthClientError)
	SendTransaction(EthBasicTransaction) (HexData, *EthClientError)
	SendRawTransaction(HexData) (HexData, *EthClientError)
	Call(EthBasicTransaction, EthBlockNumString) (HexData, *EthClientError)
	EstimateGas(EthBasicTransaction, EthBlockNumString) (HexInt, *EthClientError)
	BlockByHash(Hash256, bool) (*EthBlock, *EthClientError)
	BlockByNumber(EthBlockNumString, bool) (*EthBlock, *EthClientError)
	TransactionByHash(Hash256) (*EthFullTransaction, *EthClientError)
	TransactionByBlockHashAndIndex(Hash256, HexInt) (*EthFullTransaction, *EthClientError)
	TransactionByBlockNumberAndIndex(EthBlockNumString, HexInt) (*EthFullTransaction, *EthClientError)
	TransactionReceipt(Hash256) (*EthTransactionReceipt, *EthClientError)
	UncleByBlockHashAndIndex(Hash256, HexInt) (*EthBlock, *EthClientError)
	UncleByBlockNumberAndIndex(EthBlockNumString, HexInt) (*EthBlock, *EthClientError)
	Logs(EthGetLogsParam) ([]EthLog, *EthClientError)
}
