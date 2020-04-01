package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"go-exercise/protobuff/ctripproto"

	"github.com/gogo/protobuf/proto"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	reqData := `登录成功！�
M2790179236`

	var body ctripproto.Body
	if err := json.Unmarshal([]byte(reqData), &body); err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%+v", body)

	pdata, err := proto.Marshal(&body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("protobuf编码后数据:", string(pdata))

	// data := encode([]byte(reqData))
	var buff bytes.Buffer
	//fmt.Fprintf(&buff, "%2c", 9)

	gz := gzip.NewWriter(&buff)
	gz.Write([]byte("9 3201M2506839260         32001031810063385093ff3ce3bd57a7c3c9    707.002 4008            950030011526288048520       5E3C32296B9EC08B761E5EC279604C05DE7F2E0A8F1CF4B20C2D8321D9177A2C"))
	gz.Write([]byte(reqData))
	//gz.Write(pdata)
	gz.Flush()
	gz.Close()

	fmt.Println("加密前数据:", string(buff.Bytes()))

	//data := encode(buff.Bytes())

	log.Println("加密后的数据:", string(data))

	//  Arrays.asList("114.80.10.33","101.226.248.27", "140.206.211.33", "140.207.228.72","221.130.198.227" , "117.184.207.146")
	conn, err := net.Dial("tcp", "117.184.207.146:443")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		log.Println(err)
		return
	}

	buf := make([]byte, 2048)
	bufLen, err := conn.Read(buf)

	if err != nil {
		log.Println(err)
		return
	}

	//var buff bytes.Buffer
	// buff.Write(data)
	// r, err := gzip.NewReader(&buff)
	// if err != nil {
	// 	log.Println(err)
	// }

	// data, err = ioutil.ReadAll(r)
	// if err != nil {
	// 	log.Println(err)
	// }

	fmt.Println(string(buf[:bufLen]))
}
