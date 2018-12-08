package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

/*
客户端网络编程的核心函数
tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
conn, err := net.DialTCP("tcp", nil, tcpAddr)
conn.Write([]byte(words))
*/

func sender(conn net.Conn, uid int, i int) {
	var words string
	if i == 1 {
		words = "hello world! I'm user" + strconv.Itoa(uid)
	} else {
		words = "This is user" + strconv.Itoa(uid) + " ping msg"
	}

	conn.Write([]byte(words))
	fmt.Printf("send words:%s\n", words)
}

// 获取时间戳
func getTs() int {
	t := time.Now()
	ts := strconv.FormatInt(t.UnixNano(), 10)
	ts = ts[:10]
	i, err := strconv.Atoi(ts)
	if err == nil {
		return i
	} else {
		err.Error()
		return i
	}
}

func main() {
	rand.Seed(int64(getTs()))

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

	go Handle(conn)

	var username string = "ahuang"
	var msg string
	for {
		msg = ""
		fmt.Scan(&msg)
		conn.Write([]byte(username + " say: " + msg))
		if msg == "quit" {
			break // 程序结束运行
		}
	}
}

func Handle(conn net.Conn) {
	for {
		data := make([]byte, 1024) // 创建一个字节流
		msg_read, err := conn.Read(data)
		if msg_read == 0 || err != nil {
			break
		}
		fmt.Println("Receiv client msg:", string(data[:msg_read]))
	}
}
