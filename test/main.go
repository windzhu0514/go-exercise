package main

import (
	"fmt"
	"go-exercise/test/modules"
	_ "go-exercise/test/modules/stringbank"
)

func main() {
	modInfo, err := modules.GetModule("string.bank")
	if err != nil {
		fmt.Println(err)
		return
	}

	sb, err := modInfo.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sb.Size())
	fmt.Println(sb.Get(1))
}
