package main

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
)

func main() {
	log()
}

func log() {
	_, file, line, ok := runtime.Caller(1)
	fmt.Println(file, ok)
	file = path.Base(file) + ":" + strconv.Itoa(line)
	fmt.Println(file)

	pc, _, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	fmt.Println(f.Name())
}
