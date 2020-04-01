package main

import "C"
import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/robertkrimen/otto"
)

func main() {
	data, err := ioutil.ReadFile("D:\\testfiles\\test.html")
	if err != nil {
		log.Println(err)
		return
	}

	content := string(data)
	beginIndex := strings.LastIndex(content, "var payInfo = ")
	if beginIndex < 0 {
		return
	}

	endIndex := strings.Index(content[beginIndex:], "};")
	if endIndex < 0 {
		return

	}

	str := content[beginIndex : beginIndex+endIndex+2]

	vm := otto.New()
	vm.Run(str + `ppp=JSON.stringify(payInfo);`)
	v, err := vm.Get("ppp")
	if err != nil {
		return
	}
	fmt.Println(v.String())

	// if err := json.Unmarshal([]byte(v.String()), &p); err != nil {
	// 	return p, err
	// }
	//
	// v, err = vm.RunString(`_fmOpt.getinfo()`)
	// log.Println(v, err)
	//
	// str := string(data)
	// v, err = vm.RunString(str)
	// log.Println(v, err)
	// // v, err = vm.Call("eval", nil)
	//
	// v, err = vm.RunString(`_fmOpt.getinfo()`)
	// log.Println(v, err)
}
