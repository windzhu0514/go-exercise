package main

import "github.com/skoo87/log4go"

func main() {
	log4go.SetLevel(log4go.DEBUG)
	log4go.Register(log4go.NewConsoleWriter())
	log4go.Error("123123")
	log4go.Error("123123")
	log4go.Error("123123")
	log4go.Close()
}
