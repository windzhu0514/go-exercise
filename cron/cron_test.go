package cron_test

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

// 在任务开始就能计算出下一次执行的时间
// 不支持的情况：任务执行时间不固定，任务结束后等待固定时间后再执行任务

// DelayIfStillRunning只是把任务延迟了 任务的数量并没有减少
// 比如任务耗时10秒，每5秒执行一次，30秒后，执行了三次，还剩下3次没有执行，即使是删除了任务剩余的任务次数也会继续执行
func TestDelayIfStillRunning(t *testing.T) {
	cronLog := cron.PrintfLogger((log.New(os.Stdout, "cron: ", log.LstdFlags|log.Lshortfile)))
	c := cron.New(cron.WithLogger(cronLog), cron.WithChain(cron.DelayIfStillRunning(cron.DefaultLogger)))

	c.Start()

	entryID, err := c.AddFunc("@every 5s", func() {
		fmt.Println(time.Now(), "i am one")
		time.Sleep(time.Second * 10)
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		time.Sleep(time.Second * 30)
		fmt.Println("remove one")
		c.Remove(entryID)
		_, err := c.AddFunc("@every 5s", func() {
			fmt.Println(time.Now(), "i am two")
			time.Sleep(time.Second * 1)
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	select {}
}

func TestCustomSchedule(t *testing.T) {
	cronLog := cron.PrintfLogger((log.New(os.Stdout, "cron: ", log.LstdFlags|log.Lshortfile)))
	c := cron.New(cron.WithLogger(cronLog))
	c.Start()
	c.Schedule(Every(5*time.Second), cron.FuncJob(func() {
		fmt.Println(time.Now(), "i am one")
	}))

	select {}
}

type ConstantDelaySchedule struct {
	Delay time.Duration
}

func Every(duration time.Duration) ConstantDelaySchedule {
	if duration < time.Second {
		duration = time.Second
	}
	return ConstantDelaySchedule{
		Delay: duration - time.Duration(duration.Nanoseconds())%time.Second,
	}
}

func (schedule ConstantDelaySchedule) Next(t time.Time) time.Time {
	return t.Add(schedule.Delay - time.Duration(t.Nanosecond())*time.Nanosecond)
}

func TestGenCron(t *testing.T) {
	rates := []int{2, 50, 90, 130}
	for _, r := range rates {
		cronExpression := func() string {
			hour := r / 60
			minute := r % 60
			strMinute := strconv.Itoa(minute)
			strHour := "*/" + strconv.Itoa(hour)
			if hour == 0 {
				strMinute = "*/" + strconv.Itoa(minute)
				strHour = "*"
			}
			return fmt.Sprintf("%s %s * * *", strMinute, strHour)
		}()
		fmt.Println(cronExpression)
	}
}
