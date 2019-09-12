package main

import (
	"fmt"

	"github.com/go-ini/ini"
)

type SiteOption struct {
	CacheID  string
	CacheMin int
	CacheMax int

	RegGNum           int // 注册线程数
	MobileNoChannelID int // 手机号供应商
	SMSChannelID      int // 短信供应商
	EngineType        int
	BeginHour         int
	EndHour           int
	MaxCountDay       int // 每天最大注册数量
	MaxCount          int // 最大注册数量
	MaxFailCount      int // 最大注册返回失败次数
}

type JobNumbers struct {
	ErrorReport string
}

type SystemConfig struct {
	// 手机号供应商
	Suppliers map[string]struct {
		URL   string
		Token string
	}
}

func main() {
	var conf SystemConfig
	if err := ini.MapTo(&conf, "conf.ini"); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", conf)
}
