package main

/*
#cgo CPPFLAGS: -I./base64
#include <sign.h>
#include<stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(Getskcy("GET", "http://apihotel.meituan.com/hbsearch/HotelSearch", "__reqTraceID=6346b894-ff05-4bcd-94d3-742bb211ec76"))

	// array := []C.int{1, 2, 3, 4, 5}
	// C.printarray(&array[0], C.int(len(array)))
}

func Getskcy(method, url, rawQuery string) string {
	// GET http://apihotel.meituan.com/hbsearch/HotelSearch __reqTraceID=6346b894-ff05-4bcd-94d3-742bb211ec76&
	data := method + " " + url + " " + rawQuery
	cData := C.CString(data)
	defer C.free(unsafe.Pointer(cData))

	// var skcy string
	// _skcy := C.CString(skcy)
	// defer C.free(unsafe.Pointer(_skcy))
	//outLen := C.int(0)

	var _skcy [30]C.char
	var outLen C.int
	C.sign(cData, C.int(len(data)), &_skcy[0], &outLen)

	skcy := C.GoStringN(&_skcy[0], outLen-1)

	return skcy
}

// 5m6zKbGVLhL/9166rIQPXN5rLrs=
