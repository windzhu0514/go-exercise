package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(12 & (12 - 1))
	rand.Seed(time.Now().UTC().UnixNano())

	var reservoir [10]int
	for i := 0; i < len(reservoir); i++ {
		reservoir[i] = i
	}
	fmt.Println(reservoir)

	for j := 10; j < 20; j++ {
		r := rand.Intn(j)
		if r < 10 {
			reservoir[r] = j
		}
	}

	fmt.Println(reservoir)
}
