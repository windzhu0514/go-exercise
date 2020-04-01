package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	// https://book.cebupacificair.com/Manage
	URL := "https://203.69.105.153/Manage"
	method := "GET"

	client := &http.Client{}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	http.DefaultTransport.(*http.Transport).Proxy = func(request *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8888")
	}
	req, err := http.NewRequest(method, URL, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Host = "book.cebupacificair.com"
	//req.Header.Add("Host", "www.flypgs.com")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
