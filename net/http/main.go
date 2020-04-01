package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(5 * time.Second)
over:
	for {
		fmt.Println("for")
		select {
		case <-t.C:
			fmt.Println("over1")
			break over
		}
	}

	fmt.Println("over")
}
