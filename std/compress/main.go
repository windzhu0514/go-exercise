package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
)

func main() {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)

	w.Write([]byte("1000000"))

	w.Flush()
	w.Close()
	fmt.Println("gzip size:", len(b.Bytes()))
	fmt.Println(string(b.Bytes()))

	r, err := gzip.NewReader(&b)
	if err != nil {
		return
	}

	s, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	fmt.Println(string(s))
}
