package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "5月29日01:00之前（不影响酒店留房）"
	reg := regexp.MustCompile("(\\d{1,2}月\\d{1,2}日\\d{1,2}:\\d{1,2})之前")
	fmt.Println(reg.FindStringSubmatch(str))
}
