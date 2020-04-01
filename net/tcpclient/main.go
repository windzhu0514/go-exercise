package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	dial := net.Dialer{Timeout: time.Second * 5, KeepAlive: time.Second}
	for i := 0; i < 10; i++ {
		go func(index int) {
			conn, err := dial.Dial("tcp", ":6666")
			if err != nil {
				fmt.Println("dial error:", err)
				return
			}
			defer conn.Close()

			conn.Write([]byte(strings.Repeat(strconv.Itoa(index), 6)))
			log.Println("client:", index, "sent")
		}(i)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-sig)
}
