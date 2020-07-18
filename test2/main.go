package main

import (
	"fmt"
	"go-exercise/httpclient"
	"time"
)

func main() {
	for i := 0; i < 10000; i++ {
		req := httpclient.Get("https://www.baidu.com/")
		fmt.Println(req.GetStatusCode())
		time.Sleep(time.Second)
	}
	fmt.Println("over")
}
