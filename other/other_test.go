package other

import (
	"fmt"
	"testing"
)

type Queryer struct {
}

func (q Queryer) Do() error {
	return nil
}

func TestAnonymous(t *testing.T) {
	err := Queryer{}.Do()
	if err != nil {
		fmt.Println(err)
		return
	}

	// wrong way
	// if err := Queryer{}.Do();err != nil {

	// }
}
