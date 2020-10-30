package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "9000", "server listen port")
	flag.Parse()
	log.Fatal(http.ListenAndServe(":"+*port, http.FileServer(http.Dir("./files"))))
}
