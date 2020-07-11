package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	cronLog := cron.VerbosePrintfLogger((log.New(os.Stdout, "cron: ", log.LstdFlags)))
	c := cron.New(cron.WithLogger(cronLog))

	c.Start()

	c.AddFunc("0 1 */3 * *", func() {
		fmt.Println(time.Now(), "i am one")
	})

	time.Sleep(time.Second * 5)
}
