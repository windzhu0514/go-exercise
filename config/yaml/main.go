package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func main() {
	data, err := ioutil.ReadFile("./conf.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cfg)
}

type Config struct {
	Common struct {
		OrderCenterURL  string //订单中心URL
		ListenPort      int    //监听端口
		ReTryTimes      int    //重试次数
		LocalBaseURL    string
		RedisIPPort     string
		PinYinPath      string
		RegFailTimes    int
		DisableRegister bool // 禁用注册功能
	} `json:"common"`
	Modules  map[string]interface{} `json:"modules"`
	Services map[string]interface{} `json:"services"`
	MySQL    map[string]struct {
		MasterConStr string //主库
		SlaveConStr  string //从库
	}
	SMSServer struct {
		SMSServerURL string
	}

	EmailPop struct {
		Pop3Address string //pop3地址
	}
}
