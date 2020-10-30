package main

import (
	"fmt"
	"time"

	"go-exercise/httpclient/httpclient"
)

func main() {
	httpclient.OpenConnlog = 1
	for i := 0; i < 150; i++ {
		req := httpclient.Get("https://www.baidu.com")
		statusCode, err := req.GetStatusCode()
		fmt.Println(statusCode, err)
		time.Sleep(2 * time.Second)
	}
}
