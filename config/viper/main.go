package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type SystemConfig struct {
	Common struct {
		OrderCenterURL     string //订单中心URL
		ListenPort         int    //监听端口
		ReTryTimes         int    //重试次数
		RedisIPPort        string
		SiteIDs            string
		DisableInvalidLine bool // 是否禁用无效线路
		IsNoDB             bool // 数据不写入数据库
	}
	QueryOptions struct {
		QueryPreDays            int //查询提前天数
		FilterGNum              int // 线路筛选goroutine数量 默认值
		DepartureUseFirstLetter bool
		ArrivalUseFirstLetter   bool
		Periodic                bool   // 定时查询
		CronExpression          string // 每次开启时间 linux cron语法
	}
	MYSQL struct {
		MasterConStr string //主库
		SlaveConStr  string //从库
	}
	SiteOptions map[string]map[string]struct {
		QueryPreDays            string //查询提前天数
		FilterGNum              int    // 线路筛选goroutine数量 默认值
		DepartureUseFirstLetter bool
		ArrivalUseFirstLetter   bool
		Periodic                bool   // 定时查询
		CronExpression          string // 每次开启时间 linux cron语法
	}
}

func main() {
	//viper.SetConfigName("conf")
	viper.SetConfigFile("./conf/conf2.toml")
	viper.SetConfigFile("./conf/conf.toml")

	//viper.AddConfigPath("./conf/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("ConfigFileNotFoundError")
		} else {
			fmt.Println(err)
		}
		return
	}

	fmt.Println(viper.GetString("Common.OrdercenterURL"))
}
