package main

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var ErrHashMismatch = errors.New("new file hash mismatch after patch")
var BinURL = "http://61.155.197.245/airticket/3_0.31314.bin.gz"
var Md5Sign = "57c4bf2fb79e6b244b7fe921ccc65881"

type People struct {
	names []string
	age   int
	home  map[string]string
}

func main() {
	var s = []People{{names: []string{"1111", "2222"}, age: 111, home: make(map[string]string)}}
	for _, v := range s {
		v.names = append(v.names, "xxxx")
		v.age = 10
		v.home["11"] = "222"
	}

	fmt.Println(s)
}

func fetchAndVerifyFullBin() ([]byte, error) {
	bin, err := fetchBin()
	if err != nil {
		return nil, err
	}
	md5, _ := hex.DecodeString(Md5Sign)
	verified := verifyMd5(bin, md5)
	if !verified {
		return nil, ErrHashMismatch
	}
	return bin, nil
}

func fetchBin() ([]byte, error) {
	r, err := fetch(BinURL)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer r.Close()
	buf := new(bytes.Buffer)
	gz, err := gzip.NewReader(r)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if _, err = io.Copy(buf, gz); err != nil {
		log.Println(err)
		return nil, err
	}

	return buf.Bytes(), nil
}

func fetch(url string) (io.ReadCloser, error) {
	log.Println("fetch url: " + url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("bad http status from %s: %v", url, resp.Status)
	}
	return resp.Body, nil
}

func verifyMd5(bin []byte, sha []byte) bool {
	h := md5.New()
	h.Write(bin)
	// fmt.string(h.Sum(nil)))
	// fmt.Println(hex.EncodeToString(h.Sum(nil)))
	log.Println("h.Sum(nil) = " + hex.EncodeToString(h.Sum(nil)))
	log.Println("dest sum = " + hex.EncodeToString(sha))
	return bytes.Equal(h.Sum(nil), sha)
}
