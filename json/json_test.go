package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSameName(t *testing.T) {
	// 编码后如果出现了相同的json key，这个key不会编码到json字符串中。
	// 结果 {}
	type Price struct {
		TicketPrice  float64 `json:"ticketPrice"`
		TicketPrice2 float32 `json:"ticketPrice"`
	}

	var p = Price{10.2, 25.6}
	d, err := json.Marshal(p)
	fmt.Println(string(d), err)

	// Output:
	// {} <nil>
}

func TestEmptyArraySlive(t *testing.T) {
	arr := [3]int{}
	s := []int{}

	d, err := json.Marshal(arr)
	fmt.Println(string(d), err)
	d, err = json.Marshal(s)
	fmt.Println(string(d), err)

	// Output:
	// [0,0,0] <nil>
	// [] <nil>
}

func TestAnonymous(t *testing.T) {
	type Cat struct {
	}

	type Animal struct {
		Cat `json:"cat"`
	}

	type Animal2 struct {
		Cat interface{}
	}

	var a Animal
	fmt.Println(a)
}
