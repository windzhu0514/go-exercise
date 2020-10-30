package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	pprofHandler()
	mux.HandleFunc("/debug/pprof/", index())
	mux.HandleFunc("/debug/pprof/cmdline", cmdline())
	mux.HandleFunc("/debug/pprof/profile", profile())
	mux.HandleFunc("/debug/pprof/symbol", symbol())
	mux.HandleFunc("/debug/pprof/trace", trace())
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("一边去"))
	})

	fmt.Println(http.ListenAndServe(":9001", mux))
}
