package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := `"{\"success\":1,\"gt\":\"25ba81caec944f8d74c98befd841a667\",\"challenge\":\"2234ca37dcaef92166a615e42c032ed6\",\"new_captcha\":true}"`
	fmt.Println(str)
	fmt.Println(strconv.Unquote(str))
}
