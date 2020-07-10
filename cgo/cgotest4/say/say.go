package say

// #include"say.h"
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func Say(s string) {
	cs := C.CString(s)
	C.saySomething(cs)
	fmt.Println(C.GoString(cs))
	C.free(unsafe.Pointer(cs))
}
