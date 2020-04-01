package main

// #cgo CFLAGS:-g -O1
// #cgo CXXFLAGS:-g -O1
// #cgo LDFLAGS:-g -O1
// #include "librsasign.h"
// #include "stdlib.h"
import "C"

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	fmt.Println("hahah", HuazhuDecrypt("EOqXguubro6Q8LS3P7SYvw=="))
}

func HuazhuDecrypt(src string) string {
	fmt.Println("HuazhuDecrypt src:", src)
	cstrCrptSrc := C.CString(src)
	defer C.free(unsafe.Pointer(cstrCrptSrc))

	var outLen int
	var outbuf *int
	C.getSign(cstrCrptSrc, (**C.int)(unsafe.Pointer(&outbuf)), (*C.int)(unsafe.Pointer(&outLen)))

	var out string
	for i := 0; i < outLen; i++ {
		d := *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(i)))
		fmt.Printf("%d ", d)
		out += hex.EncodeToString([]byte{byte(d)})
	}
	fmt.Println()
	fmt.Println("HuazhuDecrypt ret:", strings.ToUpper(out))

	data := make([]byte, outLen*4)

	fmt.Println("len:", outLen)

	for i := 0; i < outLen; i++ {
		d := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(i)*unsafe.Sizeof(outbuf)))
		j := i * 4
		data[j] = byte(d >> 24)
		data[j+1] = byte(d >> 16)
		data[j+2] = byte(d >> 8)
		data[j+3] = byte(d)
	}
	C.free(unsafe.Pointer(outbuf))

	return base64.StdEncoding.EncodeToString(data)
}
