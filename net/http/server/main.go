package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"time"
)

var count int

var address = flag.String("addr", ":9290", "address to bind to.")

func main() {
	//upload.DefaultHttpServerHandles()
	//fmt.Println(http.ListenAndServe(":9290", nil))
	flag.Parse()
	server := http.Server{Addr: *address}
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("hello"))
	})
	http.HandleFunc("/count", func(w http.ResponseWriter, request *http.Request) {
		count++
		w.Write([]byte(strconv.Itoa(count)))
	})
	http.HandleFunc("/add", func(w http.ResponseWriter, request *http.Request) {
		count++
		fmt.Println("count:", count)
	})

	http.HandleFunc("/longtime", func(w http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Hour)
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("需要认证"))
			return
		}

		if username != "test" || password != "123456" {
			w.Write([]byte("认证账号或密码错误"))
			return
		}

		w.Write([]byte("哈哈哈"))
	})

	http.HandleFunc()
	fmt.Println(server.ListenAndServe())
}
