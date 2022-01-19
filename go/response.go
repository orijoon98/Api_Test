package main

type Response struct {
	Id int `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Result interface{} `json:"result"`
}

type TransactionReceiptResponse struct {
	TransactionHash interface{}
	TransactionIndex interface{}
	BlockHash interface{}
	BlockNumber interface{}
	From interface{}
	To interface{}
	CumulativeGasUsed interface{}
	GasUsed interface{}
	ContractAddress interface{}
	Logs []interface{}
	LogsBloom interface{}
	Root interface{}
	Status interface{}
}

type TransactionBlockResponse struct {
	BlockHash interface{}
	BlockNumber interface{}
	From interface{}
	Gas interface{}
	GasPrice interface{}
	Hash interface{}
	Input interface{}
	Nonce interface{}
	To interface{}
	TransactionIndex interface{}
	Value interface{}
	V interface{}
	R interface{}
	S interface{}
}

type BlockResponse struct {
	Number interface{}
	Hash interface{}
	ParentHash interface{}
	Nonce interface{}
	Sha3Uncles interface{}
	LogsBloom interface{}
	TransactionsRoot interface{}
	StateRoot interface{}
	ReceiptsRoot interface{}
	Miner interface{}
	Difficulty interface{}
	TotalDifficulty interface{}
	ExtraData interface{}
	Size interface{}
	GasLimit interface{}
	GasUsed interface{}
	Timestamp interface{}
	Transactions []interface{}
	Uncles []interface{}
}