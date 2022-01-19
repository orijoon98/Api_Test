package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/go-resty/resty/v2"
)

func Check(url string, body string, res *resty.Response) {
	fmt.Println("Request : " + url)
	
	url = HandleUrl(url)
	var params []interface{}
	var response Response

	params = unmarshalBody(body)
	response = unmarshalResponse(res, response)

	result := response.Result

	if result != nil {
		switch url {
		case "/web3_sha3":
			web3_sha3(result)
		case "/net_version":
			net_version(result)
		case "/eth_coinbase":
			eth_coinbase(result)
		case "/eth_gasPrice":
			eth_gasPrice(result)
		case "/eth_blockNumber":
			eth_blockNumber(result)
		case "/eth_getBalance":
			eth_getBalance(params, result)
		case "/eth_getTransactionCount":
			eth_getTransactionCount(params, result)
		case "/eth_sendRawTransaction":
			eth_sendRawTransaction(params, result)
		case "/eth_call":
			eth_call(params, result)
		case "/eth_estimateGas":
			eth_estimateGas(params, result)
		case "/eth_getBlockByHash":
			eth_getBlockByHash(params, result)
		case "/eth_getBlockByNumber":
			eth_getBlockByNumber(params, result)
		case "/eth_getTransactionByHash":
			eth_getTransactionByHash(params, result)
		case "/eth_getTransactionByBlockHashAndIndex":
			eth_getTransactionByBlockHashAndIndex(params, result)
		case "/eth_getTransactionByBlockNumberAndIndex":
			eth_getTransactionByBlockNumberAndIndex(params, result)
		case "/eth_getTransactionReceipt":
			eth_getTransactionReceipt(params, result)
		}
	}
}

func eth_getTransactionReceipt(params []interface{}, result interface{}) {
	if notHex(params[0].(string)) || len(params[0].(string)) != 66 {
		log.Fatal("params error : 올바른 transaction hash 값이 아닙니다.")
		os.Exit(1)
	}

	if result != nil {
		res := result.(map[string]interface{})

		tx := new(TransactionReceiptResponse)

		tx.TransactionHash = res["transactionHash"]
		tx.TransactionIndex = res["transactionIndex"]
		tx.BlockHash = res["blockHash"]
		tx.BlockNumber = res["blockNumber"]
		tx.From = res["from"]
		tx.To = res["to"]
		tx.CumulativeGasUsed = res["cumulativeGasUsed"]
		tx.GasUsed = res["gasUsed"]
		tx.ContractAddress = res["contractAddress"]
		// logs 관련 처리 필요
		// tx.Logs = res["logs"]
		tx.LogsBloom = res["logsBloom"]
		tx.Root = res["root"]
		tx.Status = res["status"]

		if notHex(tx.TransactionHash.(string)) || len(tx.TransactionHash.(string)) != 66 {
			log.Fatal("response error : 올바른 transactionHash 값이 아닙니다.")
			os.Exit(1)	
		}
		if notHex(tx.TransactionIndex.(string)) {
			log.Fatal("response error : 올바른 transactionIndex 값이 아닙니다.")
			os.Exit(1)	
		}
		if notHex(tx.BlockHash.(string)) || len(tx.BlockHash.(string)) != 66 {
			log.Fatal("response error : 올바른 blockHash 값이 아닙니다.")
			os.Exit(1)	
		}
		if notHex(tx.BlockNumber.(string)) {
			log.Fatal("response error : 올바른 blockNumber 값이 아닙니다.")
			os.Exit(1)	
		}		
		if notHex(tx.From.(string)) || len(tx.From.(string)) != 42 {
			log.Fatal("response error : 올바른 from 값이 아닙니다.")
			os.Exit(1)	
		}
		if notHex(tx.To.(string)) || len(tx.To.(string)) != 42 {
			log.Fatal("response error : 올바른 to 값이 아닙니다.")
			os.Exit(1)	
		}
		if notHex(tx.CumulativeGasUsed.(string)) {
			log.Fatal("response error : 올바른 cumulativeGasUsed 값이 아닙니다.")
			os.Exit(1)	
		}
		if notHex(tx.GasUsed.(string)) {
			log.Fatal("response error : 올바른 gasUsed 값이 아닙니다.")
			os.Exit(1)	
		}
		if tx.ContractAddress != nil {
			if notHex(tx.ContractAddress.(string)) || len(tx.ContractAddress.(string)) != 42 {
				log.Fatal("response error : 올바른 contractAddress 값이 아닙니다.")
				os.Exit(1)	
			}
		}
		if notHex(tx.LogsBloom.(string)) || len(tx.LogsBloom.(string)) != 514 {
			log.Fatal("response error : 올바른 logsBloom 값이 아닙니다.")
			os.Exit(1)	
		}
		if tx.Root != nil {
			if notHex(tx.Root.(string)) || len(tx.Root.(string)) != 66 {
				log.Fatal("response error : 올바른 root 값이 아닙니다.")
				os.Exit(1)	
			}
		}
		if tx.Status != nil {
			if tx.Status.(string) != "0x0" && tx.Status.(string) != "0x1" {
				log.Fatal("response error : 올바른 status 값이 아닙니다.")
				os.Exit(1)	
			}
		}
	}
}

func eth_getTransactionByBlockNumberAndIndex(params []interface{}, result interface{}) {
	var match [3]string = [3]string{"latest", "earliest", "pending"}
	for i := 0; i < len(match); i++ {
		if params[0].(string) == match[i] {
			break
		}
		if i == 2 {
			if notHex(params[0].(string)) {
				log.Fatal("params error : 올바른 block number 값이 아닙니다.")
				os.Exit(1)
			}
		}
	}
	if notHex(params[1].(string)) {
		log.Fatal("params error : 올바른 index position 값이 아닙니다.")
		os.Exit(1)
	}

	res := result.(map[string]interface{})

	tx := new(TransactionBlockResponse)
	
	tx.BlockHash = res["blockHash"]
	tx.BlockNumber = res["blockNumber"]
	tx.From = res["from"]
	tx.Gas = res["gas"]
	tx.GasPrice = res["gasPrice"]
	tx.Hash = res["hash"]
	tx.Input = res["input"]
	tx.Nonce = res["nonce"]
	tx.To = res["to"]
	tx.TransactionIndex = res["transactionIndex"]
	tx.Value = res["value"]
	tx.V = res["v"]
	tx.R = res["r"]
	tx.S = res["s"]

	if tx.BlockHash != nil {
		if notHex(tx.BlockHash.(string)) || len(tx.BlockHash.(string)) != 66 {
			log.Fatal("response error : 올바른 blockHash 값이 아닙니다.")
			os.Exit(1)	
		}
	}
	if tx.BlockNumber != nil {
		if notHex(tx.BlockNumber.(string)) {
			log.Fatal("response error : 올바른 blockNumber 값이 아닙니다.")
			os.Exit(1)	
		}
	}
	if notHex(tx.From.(string)) || len(tx.From.(string)) != 42 {
		log.Fatal("response error : 올바른 from 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.Gas.(string)) {
		log.Fatal("response error : 올바른 gas 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.GasPrice.(string)) {
		log.Fatal("response error : 올바른 gasPrice 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.Hash.(string)) || len(tx.Hash.(string)) != 66 {
		log.Fatal("response error : 올바른 hash 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.Input.(string)) {
		log.Fatal("response error : 올바른 input 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.Nonce.(string)) {
		log.Fatal("response error : 올바른 nonce 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.To.(string)) || len(tx.To.(string)) != 42 {
		log.Fatal("response error : 올바른 to 값이 아닙니다.")
		os.Exit(1)	
	}
	if tx.TransactionIndex != nil {
		if notHex(tx.TransactionIndex.(string)) {
			log.Fatal("response error : 올바른 transactionIndex 값이 아닙니다.")
			os.Exit(1)	
		}
	}
	if notHex(tx.Value.(string)) {
		log.Fatal("response error : 올바른 value 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.V.(string)) {
		log.Fatal("response error : 올바른 v 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.R.(string)) || len(tx.R.(string)) != 66 {
		log.Fatal("response error : 올바른 r 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.S.(string)) || len(tx.S.(string)) != 66 {
		log.Fatal("response error : 올바른 s 값이 아닙니다.")
		os.Exit(1)	
	}
}


func eth_getTransactionByBlockHashAndIndex(params []interface{}, result interface{}) {
	if notHex(params[0].(string)) || len(params[0].(string)) != 66 {
		log.Fatal("params error : 올바른 block hash 값이 아닙니다.")
		os.Exit(1)
	}
	if notHex(params[1].(string)) {
		log.Fatal("params error : 올바른 index position 값이 아닙니다.")
		os.Exit(1)
	}

	res := result.(map[string]interface{})

	tx := new(TransactionBlockResponse)
	
	tx.BlockHash = res["blockHash"]
	tx.BlockNumber = res["blockNumber"]
	tx.From = res["from"]
	tx.Gas = res["gas"]
	tx.GasPrice = res["gasPrice"]
	tx.Hash = res["hash"]
	tx.Input = res["input"]
	tx.Nonce = res["nonce"]
	tx.To = res["to"]
	tx.TransactionIndex = res["transactionIndex"]
	tx.Value = res["value"]
	tx.V = res["v"]
	tx.R = res["r"]
	tx.S = res["s"]

	if tx.BlockHash != nil {
		if notHex(tx.BlockHash.(string)) || len(tx.BlockHash.(string)) != 66 {
			log.Fatal("response error : 올바른 blockHash 값이 아닙니다.")
			os.Exit(1)	
		}
	}
	if tx.BlockNumber != nil {
		if notHex(tx.BlockNumber.(string)) {
			log.Fatal("response error : 올바른 blockNumber 값이 아닙니다.")
			os.Exit(1)	
		}
	}
	if notHex(tx.From.(string)) || len(tx.From.(string)) != 42 {
		log.Fatal("response error : 올바른 from 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.Gas.(string)) {
		log.Fatal("response error : 올바른 gas 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.GasPrice.(string)) {
		log.Fatal("response error : 올바른 gasPrice 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.Hash.(string)) || len(tx.Hash.(string)) != 66 {
		log.Fatal("response error : 올바른 hash 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.Input.(string)) {
		log.Fatal("response error : 올바른 input 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.Nonce.(string)) {
		log.Fatal("response error : 올바른 nonce 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.To.(string)) || len(tx.To.(string)) != 42 {
		log.Fatal("response error : 올바른 to 값이 아닙니다.")
		os.Exit(1)	
	}
	if tx.TransactionIndex != nil {
		if notHex(tx.TransactionIndex.(string)) {
			log.Fatal("response error : 올바른 transactionIndex 값이 아닙니다.")
			os.Exit(1)	
		}
	}
	if notHex(tx.Value.(string)) {
		log.Fatal("response error : 올바른 value 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.V.(string)) {
		log.Fatal("response error : 올바른 v 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.R.(string)) || len(tx.R.(string)) != 66 {
		log.Fatal("response error : 올바른 r 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.S.(string)) || len(tx.S.(string)) != 66 {
		log.Fatal("response error : 올바른 s 값이 아닙니다.")
		os.Exit(1)	
	}
}

func eth_getTransactionByHash(params []interface{}, result interface{}) {
	if notHex(params[0].(string)) || len(params[0].(string)) != 66 {
		log.Fatal("params error : 올바른 transaction hash 값이 아닙니다.")
		os.Exit(1)
	}

	res := result.(map[string]interface{})

	tx := new(TransactionBlockResponse)
	
	tx.BlockHash = res["blockHash"]
	tx.BlockNumber = res["blockNumber"]
	tx.From = res["from"]
	tx.Gas = res["gas"]
	tx.GasPrice = res["gasPrice"]
	tx.Hash = res["hash"]
	tx.Input = res["input"]
	tx.Nonce = res["nonce"]
	tx.To = res["to"]
	tx.TransactionIndex = res["transactionIndex"]
	tx.Value = res["value"]
	tx.V = res["v"]
	tx.R = res["r"]
	tx.S = res["s"]

	if tx.BlockHash != nil {
		if notHex(tx.BlockHash.(string)) || len(tx.BlockHash.(string)) != 66 {
			log.Fatal("response error : 올바른 blockHash 값이 아닙니다.")
			os.Exit(1)	
		}
	}
	if tx.BlockNumber != nil {
		if notHex(tx.BlockNumber.(string)) {
			log.Fatal("response error : 올바른 blockNumber 값이 아닙니다.")
			os.Exit(1)	
		}
	}
	if notHex(tx.From.(string)) || len(tx.From.(string)) != 42 {
		log.Fatal("response error : 올바른 from 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.Gas.(string)) {
		log.Fatal("response error : 올바른 gas 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.GasPrice.(string)) {
		log.Fatal("response error : 올바른 gasPrice 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.Hash.(string)) || len(tx.Hash.(string)) != 66 {
		log.Fatal("response error : 올바른 hash 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.Input.(string)) {
		log.Fatal("response error : 올바른 input 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.Nonce.(string)) {
		log.Fatal("response error : 올바른 nonce 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.To.(string)) || len(tx.To.(string)) != 42 {
		log.Fatal("response error : 올바른 to 값이 아닙니다.")
		os.Exit(1)	
	}
	if tx.TransactionIndex != nil {
		if notHex(tx.TransactionIndex.(string)) {
			log.Fatal("response error : 올바른 transactionIndex 값이 아닙니다.")
			os.Exit(1)	
		}
	}
	if notHex(tx.Value.(string)) {
		log.Fatal("response error : 올바른 value 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.V.(string)) {
		log.Fatal("response error : 올바른 v 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.R.(string)) || len(tx.R.(string)) != 66 {
		log.Fatal("response error : 올바른 r 값이 아닙니다.")
		os.Exit(1)	
	}
	if notHex(tx.S.(string)) || len(tx.S.(string)) != 66 {
		log.Fatal("response error : 올바른 s 값이 아닙니다.")
		os.Exit(1)	
	}
}


func eth_getBlockByNumber(params []interface{}, result interface{}) {
	var match [3]string = [3]string{"latest", "earliest", "pending"}
	for i := 0; i < len(match); i++ {
		if params[0].(string) == match[i] {
			break
		}
		if i == 2 {
			if notHex(params[0].(string)) {
				log.Fatal("params error : 올바른 block number 값이 아닙니다.")
				os.Exit(1)
			}
		}
	}
	boolean, ok := params[1].(bool)
	if !ok {
		log.Fatal("params error : bool 파라미터가 필요합니다.")
		os.Exit(1)
	}
	
	res := result.(map[string]interface{})

	block := new(BlockResponse)

	block.Number = res["number"]
	block.Hash = res["hash"]
	block.ParentHash = res["parentHash"]
	block.Nonce = res["nonce"]
	block.Sha3Uncles = res["sha3Uncles"]
	block.LogsBloom = res["logsBloom"]
	block.TransactionsRoot = res["transactionsRoot"]
	block.StateRoot = res["stateRoot"]
	block.ReceiptsRoot = res["receiptsRoot"]
	block.Miner = res["miner"]
	block.Difficulty = res["difficulty"]
	block.TotalDifficulty = res["totalDifficulty"]
	block.ExtraData = res["extraData"]
	block.Size = res["size"]
	block.GasLimit = res["gasLimit"]
	block.GasUsed = res["gasUsed"]
	block.Timestamp = res["timestamp"]
	// uncles 관련 처리 필요
	// uncles := res["uncles"].([]interface{})

	if boolean {
		res := result.(map[string]interface{})
		list := res["transactions"].([]interface{})
		for i := 0; i < len(list); i++ {
			// transactions object 관련 처리 필요

			// m := list[i].(map[string]string)
			// hash := m["hash"]
			// nonce := m["nonce"]
			// blockHash := m["blockHash"]
			// blockNumber := m["blockNumber"]
			// transactionIndex := m["transactionIndex"]
			// from := m["from"]
			// to := m["to"]
			// value := m["value"]
			// gas := m["gas"]
			// gasPrice := m["gasPrice"]
			// input := m["input"]
			// v := m["v"]
			// r := m["r"]
			// s := m["s"]
		}
	} else {
		res := result.(map[string]interface{})
		list := res["transactions"].([]interface{})
		for i := 0; i < len(list); i++ {
			if len(list[i].(string)) != 66 || notHex(list[i].(string)) {
				log.Fatal("response error : 올바른 transactions 값이 아닙니다.")
				os.Exit(1)	
			}
		}
	}

	if block.Number != nil {
		if notHex(block.Number.(string)) {
			log.Fatal("response error : 올바른 number 값이 아닙니다.")
			os.Exit(1)
		}
	}
	if block.Hash != nil {
		if notHex(block.Hash.(string)) || len(block.Hash.(string)) != 66 {
			log.Fatal("response error : 올바른 hash 값이 아닙니다.")
			os.Exit(1)
		}
 	}
	if notHex(block.ParentHash.(string)) || len(block.ParentHash.(string)) != 66 {
			log.Fatal("response error : 올바른 parentHash 값이 아닙니다.")
			os.Exit(1)
	}
	if block.Nonce != nil {
		if notHex(block.Nonce.(string)) || len(block.Nonce.(string)) != 18 {
			log.Fatal("response error : 올바른 nonce 값이 아닙니다.")
			os.Exit(1)
		}
 	}
	if notHex(block.Sha3Uncles.(string)) || len(block.Sha3Uncles.(string)) != 66 {
			log.Fatal("response error : 올바른 sha3Uncles 값이 아닙니다.")
			os.Exit(1)
	}
	if block.LogsBloom != nil {
		if notHex(block.LogsBloom.(string)) || len(block.LogsBloom.(string)) != 514 {
			log.Fatal("response error : 올바른 logsBloom 값이 아닙니다.")
			os.Exit(1)
		}
 	}
	if notHex(block.TransactionsRoot.(string)) || len(block.TransactionsRoot.(string)) != 66 {
			log.Fatal("response error : 올바른 transactionsRoot 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.StateRoot.(string)) || len(block.StateRoot.(string)) != 66 {
			log.Fatal("response error : 올바른 stateRoot 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.ReceiptsRoot.(string)) || len(block.ReceiptsRoot.(string)) != 66 {
			log.Fatal("response error : 올바른 receiptsRoot 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.Miner.(string)) || len(block.Miner.(string)) != 42 {
			log.Fatal("response error : 올바른 miner 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.Difficulty.(string)) {
			log.Fatal("response error : 올바른 difficulty 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.TotalDifficulty.(string)) {
			log.Fatal("response error : 올바른 totalDifficulty 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.ExtraData.(string)) {
			log.Fatal("response error : 올바른 extraData 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.Size.(string)) {
			log.Fatal("response error : 올바른 size 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.GasLimit.(string)) {
			log.Fatal("response error : 올바른 gasLimit 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.GasUsed.(string)) {
			log.Fatal("response error : 올바른 gasUsed 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.Timestamp.(string)) {
			log.Fatal("response error : 올바른 timestamp 값이 아닙니다.")
			os.Exit(1)
	}
}

func eth_getBlockByHash(params []interface{}, result interface{}) {
	if len(params[0].(string)) != 66 {
		log.Fatal("params error : 32 Bytes block hash 가 아닙니다.")
		os.Exit(1)
	}
	if notHex(params[0].(string)) {
		log.Fatal("params error : 올바른 block hash 가 아닙니다.")
		os.Exit(1)
	}
	boolean, ok := params[1].(bool)
	if !ok {
		log.Fatal("params error : bool 파라미터가 필요합니다.")
		os.Exit(1)
	}
	
	res := result.(map[string]interface{})

	block := new(BlockResponse)

	block.Number = res["number"]
	block.Hash = res["hash"]
	block.ParentHash = res["parentHash"]
	block.Nonce = res["nonce"]
	block.Sha3Uncles = res["sha3Uncles"]
	block.LogsBloom = res["logsBloom"]
	block.TransactionsRoot = res["transactionsRoot"]
	block.StateRoot = res["stateRoot"]
	block.ReceiptsRoot = res["receiptsRoot"]
	block.Miner = res["miner"]
	block.Difficulty = res["difficulty"]
	block.TotalDifficulty = res["totalDifficulty"]
	block.ExtraData = res["extraData"]
	block.Size = res["size"]
	block.GasLimit = res["gasLimit"]
	block.GasUsed = res["gasUsed"]
	block.Timestamp = res["timestamp"]
	// uncles 관련 처리 필요
	// uncles := res["uncles"].([]interface{})

	if boolean {
		res := result.(map[string]interface{})
		list := res["transactions"].([]interface{})
		for i := 0; i < len(list); i++ {
			// transactions object 관련 처리 필요

			// m := list[i].(map[string]string)
			// hash := m["hash"]
			// nonce := m["nonce"]
			// blockHash := m["blockHash"]
			// blockNumber := m["blockNumber"]
			// transactionIndex := m["transactionIndex"]
			// from := m["from"]
			// to := m["to"]
			// value := m["value"]
			// gas := m["gas"]
			// gasPrice := m["gasPrice"]
			// input := m["input"]
			// v := m["v"]
			// r := m["r"]
			// s := m["s"]
		}
	} else {
		res := result.(map[string]interface{})
		list := res["transactions"].([]interface{})
		for i := 0; i < len(list); i++ {
			if len(list[i].(string)) != 66 || notHex(list[i].(string)) {
				log.Fatal("response error : 올바른 transactions 값이 아닙니다.")
				os.Exit(1)	
			}
		}
	}

	if block.Number != nil {
		if notHex(block.Number.(string)) {
			log.Fatal("response error : 올바른 number 값이 아닙니다.")
			os.Exit(1)
		}
	}
	if block.Hash != nil {
		if notHex(block.Hash.(string)) || len(block.Hash.(string)) != 66 {
			log.Fatal("response error : 올바른 hash 값이 아닙니다.")
			os.Exit(1)
		}
 	}
	if notHex(block.ParentHash.(string)) || len(block.ParentHash.(string)) != 66 {
			log.Fatal("response error : 올바른 parentHash 값이 아닙니다.")
			os.Exit(1)
	}
	if block.Nonce != nil {
		if notHex(block.Nonce.(string)) || len(block.Nonce.(string)) != 18 {
			log.Fatal("response error : 올바른 nonce 값이 아닙니다.")
			os.Exit(1)
		}
 	}
	if notHex(block.Sha3Uncles.(string)) || len(block.Sha3Uncles.(string)) != 66 {
			log.Fatal("response error : 올바른 sha3Uncles 값이 아닙니다.")
			os.Exit(1)
	}
	if block.LogsBloom != nil {
		if notHex(block.LogsBloom.(string)) || len(block.LogsBloom.(string)) != 514 {
			log.Fatal("response error : 올바른 logsBloom 값이 아닙니다.")
			os.Exit(1)
		}
 	}
	if notHex(block.TransactionsRoot.(string)) || len(block.TransactionsRoot.(string)) != 66 {
			log.Fatal("response error : 올바른 transactionsRoot 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.StateRoot.(string)) || len(block.StateRoot.(string)) != 66 {
			log.Fatal("response error : 올바른 stateRoot 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.ReceiptsRoot.(string)) || len(block.ReceiptsRoot.(string)) != 66 {
			log.Fatal("response error : 올바른 receiptsRoot 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.Miner.(string)) || len(block.Miner.(string)) != 42 {
			log.Fatal("response error : 올바른 miner 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.Difficulty.(string)) {
			log.Fatal("response error : 올바른 difficulty 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.TotalDifficulty.(string)) {
			log.Fatal("response error : 올바른 totalDifficulty 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.ExtraData.(string)) {
			log.Fatal("response error : 올바른 extraData 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.Size.(string)) {
			log.Fatal("response error : 올바른 size 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.GasLimit.(string)) {
			log.Fatal("response error : 올바른 gasLimit 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.GasUsed.(string)) {
			log.Fatal("response error : 올바른 gasUsed 값이 아닙니다.")
			os.Exit(1)
	}
	if notHex(block.Timestamp.(string)) {
			log.Fatal("response error : 올바른 timestamp 값이 아닙니다.")
			os.Exit(1)
	}
}

func eth_estimateGas(params []interface{}, result interface{}) {
	object := params[0].(map[string]interface{})
	from, ok := object["from"].(string)
	if ok {
		if len(from) != 42 {
			log.Fatal("params error : 올바른 address 값이 아닙니다.")
			os.Exit(1)
		}
	}
	to, ok := object["to"].(string)
	if ok {
		if len(to) != 42 {
			log.Fatal("params error : 올바른 address 값이 아닙니다.")
			os.Exit(1)
		}
	}
	gas, ok := object["gas"].(string)
	if ok {
		if notHex(gas) {
			log.Fatal("params error : 올바른 gas 값이 아닙니다.")
			os.Exit(1)
		}
	}
	gasPrice, ok := object["gasPrice"].(string)
	if ok {
		if notHex(gasPrice) {
			log.Fatal("params error : 올바른 gasPrice 값이 아닙니다.")
			os.Exit(1)
		}
	}
	value, ok := object["value"].(string)
	if ok {
		if notHex(value) {
			log.Fatal("params error : 올바른 value 값이 아닙니다.")
			os.Exit(1)
		}
	}
	data, ok := object["data"].(string)
	if ok {
		if notHex(data) {
			log.Fatal("params error : 올바른 data 값이 아닙니다.")
			os.Exit(1)
		}
	}

	quantityOrTag, ok := params[1].(string)

	if ok {
		var match [3]string = [3]string{"latest", "earliest", "pending"}
		for i := 0; i < len(match); i++ {
			if quantityOrTag == match[i] {
				break
			}
			if i == 2 {
				if notHex(quantityOrTag) {
					log.Fatal("params error : 올바른 block number 값이 아닙니다.")
					os.Exit(1)
				}
			}
		}
	}

	if notHex(result.(string)) {
		log.Fatal("response error : 올바른 current balance 값이 아닙니다.")
		os.Exit(1)
	}
}

func eth_call(params2 []interface{}, result interface{}) {
	object := params2[0].(map[string]interface{})
	from, ok := object["from"].(string)
	if ok {
		if len(from) != 42 {
			log.Fatal("params error : 올바른 address 값이 아닙니다.")
			os.Exit(1)
		}
	}
	to := object["to"].(string)
	if len(to) != 42 {
		log.Fatal("params error : 올바른 address 값이 아닙니다.")
		os.Exit(1)
	}
	gas, ok := object["gas"].(string)
	if ok {
		if notHex(gas) {
			log.Fatal("params error : 올바른 gas 값이 아닙니다.")
			os.Exit(1)
		}
	}
	gasPrice, ok := object["gasPrice"].(string)
	if ok {
		if notHex(gasPrice) {
			log.Fatal("params error : 올바른 gasPrice 값이 아닙니다.")
			os.Exit(1)
		}
	}
	value, ok := object["value"].(string)
	if ok {
		if notHex(value) {
			log.Fatal("params error : 올바른 value 값이 아닙니다.")
			os.Exit(1)
		}
	}
	data, ok := object["data"].(string)
	if ok {
		if notHex(data) {
			log.Fatal("params error : 올바른 data 값이 아닙니다.")
			os.Exit(1)
		}
	}

	quantityOrTag := params2[1].(string)

	var match [3]string = [3]string{"latest", "earliest", "pending"}
	for i := 0; i < len(match); i++ {
		if quantityOrTag == match[i] {
			break
		}
		if i == 2 {
			if notHex(quantityOrTag) {
				log.Fatal("params error : 올바른 block number 값이 아닙니다.")
				os.Exit(1)
			}
		}
	}
	if notHex(result.(string)) {
		log.Fatal("response error : 올바른 current balance 값이 아닙니다.")
		os.Exit(1)
	}
}

func eth_sendRawTransaction(params []interface{}, result interface{}) {
	if len(params) != 1 {
		log.Fatal("params error : signed transaction data 1개만 필요합니다.")
		os.Exit(1)
	}
	if notHex(params[0].(string)) {
		log.Fatal("params error : 올바른 signed transaction data 값이 아닙니다.")
		os.Exit(1)
	}
	if notHex(result.(string)) || (len(result.(string)) != 66 && result.(string) != "0x0") {
		log.Fatal("response error : 올바른 transaction hash 값이 아닙니다.")
		os.Exit(1)
	}
}

func eth_getTransactionCount(params []interface{}, result interface{}) {
	if len(params[0].(string)) != 42 {
		log.Fatal("params error : 20 Bytes address 값이 아닙니다.")
		os.Exit(1)
	}
	var match [3]string = [3]string{"latest", "earliest", "pending"}
	for i := 0; i < len(match); i++ {
		if params[1] == match[i] {
			break
		}
		if i == 2 {
			if(notHex(params[1].(string))) {
				log.Fatal("params error : 올바른 block number 값이 아닙니다.")
				os.Exit(1)
			}
		}
	}
	if notHex(result.(string)) {
		log.Fatal("response error : 올바른 current balance 값이 아닙니다.")
		os.Exit(1)
	}
}

func eth_getBalance(params []interface{}, result interface{}) {
	if len(params[0].(string)) != 42 {
		log.Fatal("params error : 올바른 address 값이 아닙니다.")
		os.Exit(1)
	}
	var match [3]string = [3]string{"latest", "earliest", "pending"}
	for i := 0; i < len(match); i++ {
		if params[1] == match[i] {
			break
		}
		if i == 2 {
			if notHex(params[1].(string)) {
				log.Fatal("params error : 올바른 block number 값이 아닙니다.")
				os.Exit(1)
			}
		}
	}
	if notHex(result.(string)) {
		log.Fatal("response error : 올바른 current balance 값이 아닙니다.")
		os.Exit(1)
	}
}

func eth_blockNumber(result interface{}) {
	if notHex(result.(string)) {
		log.Fatal("response error : 올바른 block number 값이 아닙니다.")
		os.Exit(1)
	}
}

func eth_gasPrice(result interface{}) {
	if notHex(result.(string)) {
		log.Fatal("response error : 올바른 gas price 값이 아닙니다.")
		os.Exit(1)
	}
}

func eth_coinbase(result interface{}) {
	if len(result.(string)) != 42 {
		log.Fatal("response error : 올바른 coinbase address 값이 아닙니다.")
		os.Exit(1)
	}
}

func net_version(result interface{}) {
	var err bool = true
	var match [6]string = [6]string{"1", "2", "3", "4", "42", "5777"}
	for i := 0; i < len(match); i++ {
		if result.(string) == match[i] {
			err = false
		}
	}
	if err {
		log.Fatal("response error : 올바른 network id 값이 아닙니다.")
		os.Exit(1)
	}
}

func web3_sha3(result interface{}) {
	if notHex(result.(string)) || len(result.(string)) != 66 {
		log.Fatal("response error : 올바른 sha3 해시값이 아닙니다.")
		os.Exit(1)
	}
}

func unmarshalBody(str string) []interface{} {
	var body Body
	err := json.Unmarshal([]byte(str), &body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return body.Params
}

func unmarshalResponse(res *resty.Response, response Response) Response {
	err := json.Unmarshal([]byte(res.String()), &response)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return response
}

func notHex(number string) bool {
	matched, _ := regexp.MatchString("(?:0[xX])?[0-9a-fA-F]+", number)
	
	return !matched
}

func HandleUrl(url string) string {
	result := ""

	cnt := 0
	for i := 0; i < len(url); i++ {
		if string(url[i]) == "/" {
			cnt++
		}
		if cnt >= 3 {
			result += string(url[i])
		}
	}

	return result
}