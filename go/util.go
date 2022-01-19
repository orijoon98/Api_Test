package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

func Url(fileName string) map[string]string {
	baseDir := "./api/"
	fileName = baseDir + fileName

	// host := "https://bsc-dataseed.binance.org"
	host := "http://127.0.0.1:7545"

	file, err := ioutil.ReadFile(fileName)

	url := host

	if err == nil {
		str := string(file)
		left := 0
		for i := 0; i < len(str); i++ {
			if string(str[i]) == "\"" {
				left++
			}
			if left == 3 {
				if string(str[i]) != "\"" {
					url += string(str[i])
				}
			}
		}
	}

	body := body(fileName)

	result := map[string]string {
		"url": url,
		"body": body,
	}

	return result
}

func body(fileName string) string {

	file, err := ioutil.ReadFile(fileName)

	body := ""

	if err == nil {
		str := string(file)
		left := 0
		for i := 0; i < len(str); i++ {
			if string(str[i]) == "{" {
				left++
			}
			if left >= 2 {
				body += string(str[i])
			}
			if str[i] == '}' {
				left--;
			}
		}
	}

	return body
}

func JsonPrettyPrint(str string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(str), "", "\t")
	if err != nil {
		return str
	}
	return out.String()
}

func Wait() {
	time.Sleep(time.Second * 2)
}

func RandInt(n int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomInt := r1.Intn(n)

	return randomInt
}