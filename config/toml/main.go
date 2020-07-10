package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	Common struct {
		OrderCenterURL     string // 订单中心URL
		ListenPort         int    // 监听端口
		ReTryTimes         int    // 重试次数
		RedisIPPort        string
		SiteIDs            string
		DisableInvalidLine bool             // 是否禁用无效线路
		IsNoDB             bool             // 数据不写入数据库
		SiteType           map[string][]int // 站点类型和该类型的站点
	}
}

func main() {
	var c config
	meta, err := toml.DecodeFile("111config.toml", &c)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("xxxxxxxxxx")
			return
		}
		fmt.Println(err)
		return
	}

	fmt.Println(meta)
	fmt.Printf("%+v\n", c)

}
