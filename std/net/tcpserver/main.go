//server.go
package main

import (
	"fmt"
	"net"
)

func handleConn(c net.Conn, count int) {
	defer c.Close()
	for {
		fmt.Println("connID:", count, "begin")
		var buf = make([]byte, 1024)
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println("connID:", count, "conn read error:", err)
			return
		}
		fmt.Println("connID:", count, "receive data:", string(buf[:n]))
	}
	fmt.Println("connID:", count, "conn close")
}

func main() {
	l, err := net.Listen("tcp", ":6666")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	count := 0
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		fmt.Println("accept a new connection")
		count++
		go handleConn(c, count)
	}
}
