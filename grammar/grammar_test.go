package grammar

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ss := []int{2, 3, 5, 8}
loop:
	for i, v := range s {
		for j, vv := range ss {
			if v == vv {
				fmt.Println(i, v, j, vv)
				continue loop
			}
		}
	}
}
