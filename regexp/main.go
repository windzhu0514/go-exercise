package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "ewp_proxy_err_msg=提取记录失败，请至机场办理乘机手续。"
	reg := regexp.MustCompile("ewp_proxy_err_msg=(.*)。")
	fmt.Println(reg.FindStringSubmatch(str))

}
