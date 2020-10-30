package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestReplace(t *testing.T) {
	reg := regexp.MustCompile(`{query\.([\w-]*)}`)
	fmt.Println(reg.ReplaceAllString("query.(x-)", "{http.request.uri.query.$1}"))
}

func TestSubMatch(t *testing.T) {
	str := "javascript:__doPostBack('ctl00$c$MC1$paxContainer$ctl03$ctl00$MAB_SEAT~P0011','')"
	reg := regexp.MustCompile(`\('(\S*)',`)
	fmt.Println(reg.FindStringSubmatch(str))
}
