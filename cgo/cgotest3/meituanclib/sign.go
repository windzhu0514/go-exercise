package meituanclib

// #cgo LDFLAGS: -L. -lpolarssl -lz -lstdc++
// #cgo CPPFLAGS: -I.
// #include "sign.h"
// #include "stdlib.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

func Sign(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("input is nil or empty")
	}

	// void sign(const char* in, int in_size, char* out, int* out_size);
	cstrData := C.CString(input)
	defer C.free(unsafe.Pointer(cstrData))

	var encoded [30]C.char
	var outlen = C.size_t(30)
	C.sign(cstrData, C.size_t(len(input)), &encoded[0], &outlen)

	if outlen > 0 && outlen < 30 {
		return C.GoStringN(&encoded[0], C.int(outlen)), nil
	} else {
		return "", fmt.Errorf("outlen:%d is not valid", outlen)
	}
}

func Siua(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("input is nil or empty")
	}

	// void siua(const char* in, int in_size,char* out,  int * out_size);
	cstrData := C.CString(input)
	defer C.free(unsafe.Pointer(cstrData))

	inputLen := len(input)

	var encoded = make([]C.char, inputLen*2)
	var outlen C.size_t
	C.siua(cstrData, C.size_t(len(input)), &encoded[0], &outlen)

	if outlen > 0 && outlen < C.size_t(inputLen)*2 {
		return C.GoStringN(&encoded[0], C.int(outlen)), nil
	} else {
		return "", fmt.Errorf("outlen:%d is not valid", outlen)
	}
}

func EncyptDataAES2(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("input is nil or empty")
	}

	cstrInput := C.CString(input)
	defer C.free(unsafe.Pointer(cstrInput))

	outputLen := len(input) * 2
	var encoded = make([]C.char, outputLen)
	var outlen C.size_t

	C.encyptDataAES2(cstrInput, C.size_t(len(input)), &encoded[0], &outlen)

	if outlen > 0 && outlen < C.size_t(outputLen) {
		return C.GoStringN(&encoded[0], C.int(outlen)), nil
	} else {
		return "", fmt.Errorf("outlen:%d is not valid", outlen)
	}
}

func PayEncryptRequestWithRandom(uuid, timestamp string) (string, error) {
	if len(uuid) == 0 || len(timestamp) == 0 {
		return "", errors.New("input is nil or empty")
	}

	cstrUUID := C.CString(uuid)
	defer C.free(unsafe.Pointer(cstrUUID))

	cstrTimestamp := C.CString(timestamp)
	defer C.free(unsafe.Pointer(cstrTimestamp))

	var encoded = make([]C.char, 400)
	var outlen C.int

	C.pay_encrypt_request_with_random(cstrUUID, cstrTimestamp, &encoded[0], &outlen)

	if outlen > 0 && outlen < 400 {
		return C.GoStringN(&encoded[0], outlen), nil
	} else {
		return "", fmt.Errorf("outlen:%d is not valid", outlen)
	}
}
