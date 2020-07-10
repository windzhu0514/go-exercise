package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func main() {
	data, err := ioutil.ReadFile("./config/yaml/conf.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	var m yaml.MapSlice
	if err = yaml.Unmarshal(data, &m); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)
}
