package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/elazarl/goproxy"
)

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	logfile := flag.String("l", "", "log file")
	addr := flag.String("a", ":8080", "proxy listen address")
	printBody := flag.Bool("b", false, "print request body and response body")
	certFile := flag.String("certfile", "", "the TLS cert file")
	keyFile := flag.String("keyfile", "", "the TLS key file")
	// "json,html,xml,image"
	bodyFilter := flag.String("f", "", "filter body by contenttype")

	flag.Parse()

	if *certFile != "" && *keyFile == "" || *certFile == "" && *keyFile != "" {
		log.Fatal("Both the certificate file and the key file must be specified")
	}

	var fliterTypes []string
	if *bodyFilter != "" {
		if !*printBody {
			log.Fatal("-b flag must be specified when use body filter")
		}

		fliterTypes = strings.Split(*bodyFilter, ",")
	}

	isPrintBody := func(contentType string) bool {
		if *printBody && len(fliterTypes) == 0 {
			return true
		}

		for _, v := range fliterTypes {
			if strings.Contains(contentType, v) {
				return true
			}
		}

		return false
	}

	if *logfile != "" {
		f, err := os.OpenFile(*logfile, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
			return
		}

		log.SetOutput(f)
	}

	setCA(*certFile, *keyFile)

	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (request *http.Request, response *http.Response) {
		var reqInfo string
		if req != nil {
			r, _ := httputil.DumpRequest(req, *printBody)
			reqInfo = strings.TrimSpace(string(r))
		}
		if reqInfo != "" {
			reqInfo = ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n" + reqInfo
		}
		log.Println(reqInfo)
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		return req, nil
	})
	proxy.OnResponse().DoFunc(func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		var respInfo string
		if resp != nil {
			r, _ := httputil.DumpResponse(resp, isPrintBody(resp.Header.Get("Content-Type")))
			respInfo = strings.TrimSpace(string(r))
		}
		if respInfo != "" {
			respInfo = "<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n" + respInfo
		}
		log.Println(respInfo)
		log.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
		return resp
	})
	proxy.Verbose = *verbose

	log.Println("listening on", *addr)
	log.Fatal(http.ListenAndServe(*addr, proxy))
}

func setCA(certFile, keyFile string) error {
	goproxyCa, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}

	if goproxyCa.Leaf, err = x509.ParseCertificate(goproxyCa.Certificate[0]); err != nil {
		return err
	}
	goproxy.GoproxyCa = goproxyCa
	goproxy.OkConnect = &goproxy.ConnectAction{Action: goproxy.ConnectAccept, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.MitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.HTTPMitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectHTTPMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.RejectConnect = &goproxy.ConnectAction{Action: goproxy.ConnectReject, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	return nil
}
