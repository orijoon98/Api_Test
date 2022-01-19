package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func Test(method string) {
	TestMethod(Url(method + ".json"))
}

func TestMethod(json map[string]string) {
	url := json["url"]
	body := json["body"]

	var test map[string]interface{}
	random := unmarshalBodyTest(body)
	test = random.(map[string]interface{})

	url = HandleUrl(url)

	switch url {
	case "/eth_getBalance":
		dataQuantityTag(url, test)
	case "/eth_getTransactionCount":
		dataQuantityTag(url, test)
	case "/eth_call":
		objectQuantityTag(url, test)
	case "/eth_estimateGas":
		objectQuantityTag(url, test)
	case "/eth_getBlockByHash":
		dataBoolean(url, test)
	case "/eth_getBlockByNumber":
		quantityTagBoolean(url, test)
	case "/eth_getTransactionByHash":
		data(url, test)
	case "/eth_getTransactionByBlockHashAndIndex":
		dataQunatity(url, test)
	case "/eth_getTransactionByBlockNumberAndIndex":
		quantityTagQuantity(url, test)
	case "/eth_getTransactionReceipt":
		data(url, test)
	}
}

func quantityTagQuantity(url string, test map[string]interface{}) {
	random := test["random"].(bool)
	quantityType := (int)(test["quantity-type"].(float64))
	testCount := (int)(test["test-count"].(float64))

	var quantityTag, quantity string

	if random {
		for i := 0; i < testCount; i++ {
			quantityTag = randomQuantityOrTag()
			switch quantityType {
			case 1:
				quantity = CreateHex(1).(string)
			case 2:
				quantity = CreateStrings(1).(string)
			}
			var params []string
			params = append(params, quantityTag)
			params = append(params, quantity)
			SetParamsStrings("eth_getTransactionByBlockNumberAndIndex", params)
			Post("eth_getTransactionByBlockNumberAndIndex")
		}
	} else {
		fmt.Print("1. quantity|tag 파라미터 값 입력: ")
		fmt.Scanln(&quantityTag)
		fmt.Print("2. quantity 파라미터 값 입력: ")
		fmt.Scanln(&quantity)
		var params []string
		params = append(params, quantityTag)
		params = append(params, quantity)
		SetParamsStrings("eth_getTransactionByBlockNumberAndIndex", params)
		Post("eth_getTransactionByBlockNumberAndIndex")
	}
}

func dataQunatity(url string, test map[string]interface{}) {
	random := test["random"].(bool)
	dataType := (int)(test["data-type"].(float64))
	dataByte := (int)(test["data-byte"].(float64))
	quantityType := (int)(test["quantity-type"].(float64))
	testCount := (int)(test["test-count"].(float64))

	var data, quantity string

	if random {
		for i := 0; i < testCount; i++ {
			switch dataType {
			case 1:
				data = CreateHex(dataByte * 2).(string)
			case 2:
				data = CreateStrings(dataByte * 2).(string)
			}
			switch quantityType {
			case 1:
				quantity = CreateHex(1).(string)
			case 2:
				quantity = CreateStrings(1).(string)
			}
			var params []string
			params = append(params, data)
			params = append(params, quantity)
			SetParamsStrings("eth_getTransactionByBlockHashAndIndex", params)
			Post("eth_getTransactionByBlockHashAndIndex")
		}
	} else {
		fmt.Print("1. data 파라미터 값 입력: ")
		fmt.Scanln(&data)
		fmt.Print("2. quantity 파라미터 값 입력: ")
		fmt.Scanln(&quantity)
		var params []string
		params = append(params, data)
		params = append(params, quantity)
		SetParamsStrings("eth_getTransactionByBlockHashAndIndex", params)
		Post("eth_getTransactionByBlockHashAndIndex")
	}
}

func data(url string, test map[string]interface{}) {
	random := test["random"].(bool)
	dataType := (int)(test["data-type"].(float64))
	dataByte := (int)(test["data-byte"].(float64))
	testCount := (int)(test["test-count"].(float64))

	var data string

	if random {
		for i := 0; i < testCount; i++ {
			switch dataType {
			case 1:
				data = CreateHex(dataByte * 2).(string)
			case 2:
				data = CreateStrings(dataByte * 2).(string)
			}
			var params []string
			params = append(params, data)
			switch url {
			case "/eth_getTransactionByHash":
				SetParamsStrings("eth_getTransactionByHash", params)
				Post("eth_getTransactionByHash")
			case "/eth_getTransactionReceipt":
				SetParamsStrings("eth_getTransactionReceipt", params)
				Post("eth_getTransactionReceipt")
			}
			
		}
	} else {
		fmt.Print("1. data 파라미터 값 입력: ")
		fmt.Scanln(&data)
		var params []string
		params = append(params, data)
		switch url {
		case "/eth_getTransactionByHash":
			SetParamsStrings("eth_getTransactionByHash", params)
			Post("eth_getTransactionByHash")
		case "/eth_getTransactionReceipt":
			SetParamsStrings("eth_getTransactionReceipt", params)
			Post("eth_getTransactionReceipt")
		}
	}
}

func quantityTagBoolean(url string, test map[string]interface{}) {
	random := test["random"].(bool)
	testCount := (int)(test["test-count"].(float64))

	var quantityTag string
	var boolean bool

	if random {
		for i := 0; i < testCount; i++ {
			SetParamsStringAndBoolean("eth_getBlockByNumber", randomQuantityOrTag(), randomBoolean())
			Post("eth_getBlockByNumber")
		}
	} else {
		fmt.Print("1. quantity|tag 파라미터 값 입력: ")
		fmt.Scanln(&quantityTag)
		fmt.Print("2. boolean 파라미터 값 입력: ")
		fmt.Scanln(&boolean)
		SetParamsStringAndBoolean("eth_getBlockByNumber", quantityTag, boolean)
		Post("eth_getBlockByNumber")
	}
}

func dataBoolean(url string, test map[string]interface{}) {
	random := test["random"].(bool)
	dataType := (int)(test["data-type"].(float64))
	dataByte := (int)(test["data-byte"].(float64))
	testCount := (int)(test["test-count"].(float64))

	var data string
	var boolean bool

	if random {
		for i := 0; i < testCount; i++ {
			
			switch dataType {
			case 1:
				data = CreateHex(dataByte * 2).(string)
			case 2:
				data = CreateStrings(dataByte * 2).(string)
			}
			boolean = randomBoolean()
			SetParamsStringAndBoolean("eth_getBlockByHash", data, boolean)
			Post("eth_getBlockByHash")
		}
	} else {
		fmt.Print("1. data 파라미터 값 입력: ")
		fmt.Scanln(&data)
		fmt.Print("2. boolean 파라미터 값 입력: ")
		fmt.Scanln(&boolean)
		SetParamsStringAndBoolean("eth_getBlockByHash", data, boolean)
		Post("eth_getBlockByHash")
	}
}

func objectQuantityTag(url string, test map[string]interface{}) {
	random := test["random"].(bool)
	fromType := (int)(test["from-type"].(float64))
	fromByte := (int)(test["from-byte"].(float64))
	toType := (int)(test["to-type"].(float64))
	toByte := (int)(test["to-byte"].(float64))
	gasType := (int)(test["gas-type"].(float64))
	gasPriceType := (int)(test["gasPrice-type"].(float64))
	valueType := (int)(test["value-type"].(float64))
	dataType := (int)(test["data-type"].(float64))
	testCount := (int)(test["test-count"].(float64))

	keys := []string{"from", "to", "gas", "gasPrice", "value", "data"}
	var values []interface{}
	var quantityTag string

	if random {
		var from, to, gas, gasPrice, value, data interface{}
		for i := 0; i < testCount; i++ {
			switch fromType {
			case 1:
				from = CreateHex(fromByte * 2)
			case 2:
				from = CreateStrings(fromByte * 2)
			case 3:
				from = nil
			}
			values = append(values, from)

			switch toType {
			case 1:
				to = CreateHex(toByte * 2)
			case 2:
				to = CreateStrings(toByte * 2)
			}
			values = append(values, to)

			switch gasType {
			case 1:
				gas = CreateHex(4)
			case 2:
				gas = CreateStrings(4)
			case 3:
				gas = nil
			}
			values = append(values, gas)
			
			switch gasPriceType {
			case 1:
				gasPrice = CreateHex(11)
			case 2:
				gasPrice = CreateStrings(11)
			case 3:
				gasPrice = nil
			}
			values = append(values, gasPrice)

			switch valueType {
			case 1:
				value = CreateHex(8)
			case 2:
				value = CreateStrings(8)
			case 3:
				value = nil
			}
			values = append(values, value)

			switch dataType {
			case 1:
				data = CreateHex(82)
			case 2:
				data = CreateStrings(82)
			case 3:
				data = nil
			}
			values = append(values, data)
			quantityTag = randomQuantityOrTag()
			switch url {
			case "/eth_call":
				SetParamsMapAndString("eth_call", keys, values, quantityTag)
				Post("eth_call")
			case "/eth_estimateGas":
				SetParamsMapAndString("eth_estimateGas", keys, values, quantityTag)
				Post("eth_estimateGas")
			}
		}
	} else {
		var from, to, gas, gasPrice, value, data string
		fmt.Print("1. from 파라미터 값 입력(null 가능): ")
		fmt.Scanln(&from)
		values = append(values, from)
		fmt.Print("2. to 파라미터 값 입력: ")
		fmt.Scanln(&to)
		values = append(values, to)
		fmt.Print("3. gas 파라미터 값 입력(null 가능): ")
		fmt.Scanln(&gas)
		values = append(values, gas)
		fmt.Print("4. gasPrice 파라미터 값 입력(null 가능): ")
		fmt.Scanln(&gasPrice)
		values = append(values, gasPrice)
		fmt.Print("5. value 파라미터 값 입력(null 가능): ")
		fmt.Scanln(&value)
		values = append(values, value)
		fmt.Print("6. data 파라미터 값 입력(null 가능): ")
		fmt.Scanln(&data)
		values = append(values, data)
		fmt.Print("7. QUANTITY|TAG 파라미터 값 입력: ")
		fmt.Scanln(&quantityTag)
		for i := 0; i < len(keys); i++ {
			if values[i].(string) == "null" {
				values[i] = nil
			}
		}
		switch url {
		case "/eth_call":
			SetParamsMapAndString("eth_call", keys, values, quantityTag)
			Post("eth_call")
		}
	}
}

func dataQuantityTag(url string, test map[string]interface{}) {
	random := test["random"].(bool)
	dataType := (int)(test["data-type"].(float64))
	dataByte := (int)(test["data-byte"].(float64))
	testCount := (int)(test["test-count"].(float64))


	var param1, param2 string
	var params = []string{param1, param2}

	if random {
		switch dataType {
		case 1:
			for i := 0; i < testCount; i++ {
				param1 = CreateHex(dataByte * 2).(string)
				param2 = randomQuantityOrTag()
				params = []string{param1, param2}
				switch url {
				case "/eth_getBalance":
					SetParamsStrings("eth_getBalance", params)
					Post("eth_getBalance")
				case "/eth_getTransactionCount":
					SetParamsStrings("eth_getTransactionCount", params)
					Post("eth_getTransactionCount")
				}
			}
		case 2:
			for i := 0; i < testCount; i++ {
				param1 = CreateStrings(dataByte * 2).(string)
				param2 = randomQuantityOrTag()
				params = []string{param1, param2}
				switch url {
				case "/eth_getBalance":
					SetParamsStrings("eth_getBalance", params)
					Post("eth_getBalance")
				case "/eth_getTransactionCount":
					SetParamsStrings("eth_getTransactionCount", params)
					Post("eth_getTransactionCount")
				}
			}
		}
	} else {
		fmt.Print("1. DATA 파라미터 값 입력: ")
		fmt.Scanln(&param1)
		fmt.Print("2. QUANTITY|TAG 파라미터 값 입력: ")
		fmt.Scanln(&param2)
		params = []string{param1, param2}
		switch url {
		case "/eth_getBalance":
			SetParamsStrings("eth_getBalance", params)
			Post("eth_getBalance")
		case "/eth_getTransactionCount":
			SetParamsStrings("eth_getTransactionCount", params)
			Post("eth_getTransactionCount")
		}
	}
}

func randomQuantityOrTag() string {
	random := RandInt(4)
	var tags = []string{"latest", "earliest", "pending"}
	if random != 3 {
		return tags[random]
	} else {
		return CreateHex(RandInt(100)).(string)
	}
}

func randomBoolean() bool {
	random := RandInt(2)
	if random == 1 {
		return true
	} else {
		return false
	}
}

func unmarshalBodyTest(str string) interface{} {
	var body Body
	err := json.Unmarshal([]byte(str), &body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return body.Test
}

