package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"math/rand"
)

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	// data, err := ioutil.ReadFile("./test3/file.go")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "./test3/file.go", nil, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, decl := range f.Decls {
		fmt.Println(decl)
	}

}

type param struct {
}

type resp struct {
}

func Retry(f func() bool) {
	for i := 0; i < 3; i++ {
		if !f() {
			return
		}
	}
}

func req() (string, error) {
	var resp string
	var err error
	Retry(func() bool {
		n := rand.Intn(2)
		if n == 1 {
			resp = "可以了"
			return false
		}

		fmt.Println("retry retry")
		return true
	})

	return resp, err
}
