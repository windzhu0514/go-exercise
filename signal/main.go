package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	f, err := os.OpenFile("./log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(f)

	// kill -1 syscall.SIGHUP					output:hangup
	// ctrl+c syscall.SIGINT					output:interrupt
	// syscall.SIGKILL 该信号不能被捕获处理
	// kill syscall.SIGTERM  kill 默认发送TERM	 output:terminated
	// ctrl+\ syscall.SIGQUIT					output:quit
	// kill -10 syscall.SIGUSR1					output:user defined signal 1
	// kill -12 syscall.SIGUSR2					output:user defined signal 2
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR2, syscall.SIGUSR1)
	for {
		sig := <-ch
		log.Println(sig)
	}
}
