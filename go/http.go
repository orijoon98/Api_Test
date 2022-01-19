package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

func Post(method string) {
	PostMethod(Url(method + ".json"))
}

func PostMethod(json map[string]string) {
	url := json["url"]
	body := json["body"]

	client := resty.New()

	client.
		SetTimeout(5 * time.Second)

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)

	Check(url, body, res)
	
	result(res, err)
}

func result(res *resty.Response, err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Print("Status code : ")
		fmt.Println(res.StatusCode())
		fmt.Println("Response : ")
		fmt.Println(JsonPrettyPrint(res.String()))
		if(res.StatusCode() == 500) {
			os.Exit(1)
		}
	}

	// Wait()
}