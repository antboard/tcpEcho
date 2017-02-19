package main

import (
	"fmt"
	"net"
)

func newConnect(c net.Conn) {
	rip := c.RemoteAddr().String()
	c.Write([]byte("hello antriver.com"))
	for {
		req := make([]byte, 128)
		readLen, e := c.Read(req)
		if e != nil {
			fmt.Println(rip, "err Read e is ", e)
			c.Close()
			fmt.Println("exit link:", rip)
			return
		}

		if readLen == 0 {
			fmt.Println(rip, "disconnect!")
			c.Close()
			return // 链接已经断开
		}
		fmt.Println(string(req[:readLen]))
		c.Write(req[:readLen])
	}
}
func main() {
	fmt.Println("starting....")
	str := fmt.Sprintf(":%d", 3721)
	l, e := net.Listen("tcp", str)
	if e != nil {
		fmt.Println(e)
	}
	defer l.Close()

	for {
		conn, e := l.Accept()
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println("new connect!! ", conn.RemoteAddr().String())
		go newConnect(conn)
	}
}
