package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	work()
}

func work() {
	for {
		fmt.Println("123")
		time.Sleep(time.Second)
	}
}
