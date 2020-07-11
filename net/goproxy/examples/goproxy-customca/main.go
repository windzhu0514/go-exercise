package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/elazarl/goproxy"
)

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	logfile := flag.String("l", "./log.log", "log file")
	addr := flag.String("addr", ":8080", "proxy listen address")
	flag.Parse()

	f, err := os.OpenFile(*logfile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.SetOutput(f)

	setCA(caCert, caKey)
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (request *http.Request, response *http.Response) {
		log.Println(req.URL.String(), "req Heads:")
		for k, v := range req.Header {
			log.Println(k, v)
		}
		log.Println(req.URL.String(), "req Body:")
		body, _ := httputil.DumpRequest(req, true)
		log.Println(string(body))
		return req, nil
	})
	proxy.OnResponse().DoFunc(func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		log.Println(resp.Request.URL.String(), "resp Heads:")
		for k, v := range resp.Header {
			log.Println(k, v)
		}
		log.Println(resp.Request.URL.String(), "resp Body:")
		body, _ := httputil.DumpResponse(resp, true)
		log.Println(string(body))

		return resp
	})
	proxy.Verbose = *verbose
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
