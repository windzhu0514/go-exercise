package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	go httpsServer()

	// 客户端加载证书
	pool := x509.NewCertPool()
	caCertPath := "cert.pem"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://10.101.72.62:6500/test")
	// resp, err := http.Get("https://10.101.72.62:6500/test")
	// 没有加载证书会提示 Get error: Get https://10.101.72.62:6500/test: x509: certificate signed by unknown authority
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func httpsServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("response from https server"))
		fmt.Println("i'm https server")
	})

	fmt.Println(http.ListenAndServeTLS(":6500", "cert.pem", "key.pem", mux))
}
