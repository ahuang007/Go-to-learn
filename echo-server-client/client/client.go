package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func sender(conn net.Conn) {
	words := "hello world!"
	conn.Write([]byte(words))
	fmt.Printf("send words:%s\n", words)
}

func main() {
	server := "127.0.0.1:1024"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, " Fatal error:%s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, " Fatal error:%s", err.Error())
		os.Exit(1)
	}

	fmt.Println(" connect success")
	for {
		sender(conn)
		time.Sleep(time.Second * 5)
	}
}
