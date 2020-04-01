package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/2 * * * * *", func() {
		fmt.Println(time.Now(), "i am one")
	})
	go func() {
		c.Start()
	}()

	time.Sleep(time.Second * 5)
	entryID, err := c.AddFunc("*/2 * * * * *", func() {
		fmt.Println(time.Now(), "i am two")
	})
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 15)

	c.Remove(entryID)

	select {}
}
