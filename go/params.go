package main

type Body struct {
	JsonRpc string `json:"jsonrpc"`
	Method string `json:"method"`
	Params  []interface{} `json:"params"`
	Id int `json:"id"`
	Test interface{} `json:"test"`
}