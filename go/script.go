package main

import (
	"bufio"
	"fmt"
	"os"
)

func RandomScript() {
	WriteScript("eth_getBalance")
	WriteScript("eth_getTransactionCount")
	WriteScript("eth_call")
	WriteScript("eth_estimateGas")
	WriteScript("eth_getBlockByHash")
	WriteScript("eth_getBlockByNumber")
	WriteScript("eth_getTransactionByHash")
	WriteScript("eth_getTransactionByBlockHashAndIndex")
	WriteScript("eth_getTransactionByBlockNumberAndIndex")
	WriteScript("eth_getTransactionReceipt")
}

func WriteScript(url string) {
	fileName := "./api/" + url + ".json"

	file, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)
    w := bufio.NewWriter(file)

    rw := bufio.NewReadWriter(r, w)

	testParams := ""

	switch url {
	case "eth_getBalance":
		testParams = 
		"      \"random\": true,\n" +
		"      \"data-type\": 1,\n" +
		"      \"data-byte\": 20,\n" +
		"      \"test-count\": 1\n"		
	case "eth_getTransactionCount":
		testParams = 
		"      \"random\": true,\n" +
		"      \"data-type\": 1,\n" +
		"      \"data-byte\": 20,\n" +
		"      \"test-count\": 1\n"
	case "eth_call":
		testParams = 
		"      \"random\": true,\n" +
		"      \"from-type\": 1,\n" +
		"      \"from-byte\": 20,\n" +
		"      \"to-type\": 1,\n" +
		"      \"to-byte\": 20,\n" +
		"      \"gas-type\": 1,\n" +
		"      \"gasPrice-type\": 1,\n" +
		"      \"value-type\": 1,\n" +
		"      \"data-type\": 1,\n" +
		"      \"test-count\": 1\n"
	case "eth_estimateGas":
		testParams = 
		"      \"random\": true,\n" +
		"      \"from-type\": 1,\n" +
		"      \"from-byte\": 20,\n" +
		"      \"to-type\": 1,\n" +
		"      \"to-byte\": 20,\n" +
		"      \"gas-type\": 1,\n" +
		"      \"gasPrice-type\": 1,\n" +
		"      \"value-type\": 1,\n" +
		"      \"data-type\": 1,\n" +
		"      \"test-count\": 1\n"
	case "eth_getBlockByHash":
		testParams = 
		"      \"random\": true,\n" +
		"      \"data-type\": 1,\n" +
		"      \"data-byte\": 32,\n" +
		"      \"test-count\": 1\n"
	case "eth_getBlockByNumber":
		testParams = 
		"      \"random\": true,\n" +
		"      \"test-count\": 1\n"
	case "eth_getTransactionByHash":
		testParams = 
		"      \"random\": true,\n" +
		"      \"data-type\": 1,\n" +
		"      \"data-byte\": 32,\n" +
		"      \"test-count\": 1\n"
	case "eth_getTransactionByBlockHashAndIndex":
		testParams = 
		"      \"random\": true,\n" +
		"      \"data-type\": 1,\n" +
		"      \"data-byte\": 32,\n" +
		"      \"quantity-type\": 1,\n" +
		"      \"test-count\": 1\n"
	case "eth_getTransactionByBlockNumberAndIndex":
		testParams = 
		"      \"random\": true,\n" +
		"      \"quantity-type\": 1,\n" +
		"      \"test-count\": 1\n"
	case "eth_getTransactionReceipt":
		testParams = 
		"      \"random\": true,\n" +
		"      \"data-type\": 1,\n" +
		"      \"data-byte\": 32,\n" +
		"      \"test-count\": 1\n"
	}

	rw.WriteString(
		"{\n  \"url\": \"/"+ url + "\",\n" +
		"  \"body\": {\n" +
		"    \"jsonrpc\": \"2.0\",\n" +
		"    \"method\": \"" + url + "\",\n" +
		"    \"params\": [],\n" +
		"    \"id\": 1,\n" +
		"    \"test\": {\n" +
		testParams +
		"    }\n" +
		"  }\n}")

	rw.Flush()
}

func SetParamsStrings(url string, params []string) {
	fileName := "./api/" + url + ".json"

	file, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_RDWR,
		os.FileMode(0644))
	
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	str := ""
	for s.Scan() {
		str += s.Text()
	}

	before := ""

	for i := 0; i < len(str); i++ {
		if str[i] == '[' {
			break
		}
		before += string(str[i])
	}

	after := ""
	cnt := 0
	flag := false

	for i := 0; i < len(str); i++ {
		if str[i] == '[' {
			cnt++
			flag = true
		}
		if str[i] == ']' {
			cnt--
		}
		if (cnt == 0 && flag && str[i] != ']') {
			after += string(str[i])
		}
	}

	middle := "["

	for i := 0; i < len(params); i++ {
		middle += "\""
		middle += params[i]
		if i != len(params) - 1 {
			middle += "\", "
		} else {
			middle += "\""
		}
	}

	middle += "]"

	file2, err2 := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	
	if err2 != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()

	r := bufio.NewReader(file2)
    w := bufio.NewWriter(file2)

    rw := bufio.NewReadWriter(r, w)

	rw.WriteString(before + middle + after)
	rw.Flush()
}

func SetParamsMapAndString(url string, keys []string, values []interface{}, strParam string) {
	fileName := "./api/" + url + ".json"

	file, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_RDWR,
		os.FileMode(0644))
	
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	str := ""
	for s.Scan() {
		str += s.Text()
	}

	before := ""

	for i := 0; i < len(str); i++ {
		if str[i] == '[' {
			break
		}
		before += string(str[i])
	}

	after := ""
	cnt := 0
	flag := false

	for i := 0; i < len(str); i++ {
		if str[i] == '[' {
			cnt++
			flag = true
		}
		if str[i] == ']' {
			cnt--
		}
		if (cnt == 0 && flag && str[i] != ']') {
			after += string(str[i])
		}
	}

	middle := "[\n  {\n"

	for i := 0; i < len(keys); i++ {
		middle += "\"" + keys[i] + "\": "
		if i != len(keys) - 1 {
			if values[i] == nil {
			middle += "null,"
			} else {
			middle += "\"" + values[i].(string) + "\", "
			}
		} else {
			if values[i] == nil {
			middle += "null"
			} else {
			middle += "\"" + values[i].(string) + "\"\n"
			}
		}
	}
	
	middle += "  },\n"

	middle += "\"" + strParam + "\"\n"
	
	middle += "]"

	file2, err2 := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	
	if err2 != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()

	r := bufio.NewReader(file2)
    w := bufio.NewWriter(file2)

    rw := bufio.NewReadWriter(r, w)

	rw.WriteString(before + middle + after)
	rw.Flush()
}

func SetParamsStringAndBoolean(url string, stringParam string, boolParam bool) {
	fileName := "./api/" + url + ".json"

	file, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_RDWR,
		os.FileMode(0644))
	
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	str := ""
	for s.Scan() {
		str += s.Text()
	}

	before := ""

	for i := 0; i < len(str); i++ {
		if str[i] == '[' {
			break
		}
		before += string(str[i])
	}

	after := ""
	cnt := 0
	flag := false

	for i := 0; i < len(str); i++ {
		if str[i] == '[' {
			cnt++
			flag = true
		}
		if str[i] == ']' {
			cnt--
		}
		if (cnt == 0 && flag && str[i] != ']') {
			after += string(str[i])
		}
	}

	middle := "["
	middle += "\"" + stringParam + "\", "
	if boolParam {
		middle += "true"
	} else {
		middle += "false"
	}

	middle += "]"

	file2, err2 := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	
	if err2 != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()

	r := bufio.NewReader(file2)
    w := bufio.NewWriter(file2)

    rw := bufio.NewReadWriter(r, w)

	rw.WriteString(before + middle + after)
	rw.Flush()
}