package main

import (
	"fmt"
	"net/http"
)

const defaultAddr = "127.0.0.1:0"

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RemoteAddr)
	}))
	http.ListenAndServe(defaultAddr, nil)
}
