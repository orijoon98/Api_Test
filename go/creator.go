package main

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890"
const HexBytes = "abcdefABCDEF0123456789"

func CreateNull() interface{} {
	var result interface{} = nil
	return result
}

func CreateStrings(n int) interface{} {
    str := make([]byte, n)
    for i := range str {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
        str[i] = letterBytes[r1.Intn((len(letterBytes)))]
    }
    return string(str)
}

func CreateBooleans() interface{} {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	i := r1.Intn(2)
	return i == 1
}

func CreateHex(n int) interface{} {
	prefix := "0x"
	str := make([]byte, n)
    for i := range str {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
        str[i] = HexBytes[r1.Intn((len(HexBytes)))]
    }
    return prefix + string(str)
}

func CreateStringArray(n int) interface{} {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomInt := r1.Intn(100)
	arr := make([]string, randomInt)
	for i := range arr {
		str := make([]byte, n)
    	for j := range str {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			str[j] = letterBytes[r1.Intn((len(letterBytes)))]
		}
		arr[i] = string(str)
	}
	return arr
}

func CreateHexArray(n int) interface{} {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomInt := r1.Intn(100)
	arr := make([]string, randomInt)
	for i := range arr {
		prefix := "0x"
		str := make([]byte, n)
		for i := range str {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			str[i] = HexBytes[r1.Intn((len(HexBytes)))]
		}
		arr[i] = prefix + string(str) 
	}
	return arr
}