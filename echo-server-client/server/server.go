package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

/*
服务端网络编程的核心函数
netListen, err := net.Listen("tcp", "localhost:1024")
conn, err := netListen.Accept()
n, err := conn.Read(buffer)
*/

// 连接池
var ConnMap map[int]net.Conn = make(map[int]net.Conn)

//todo: 增加{uid => conn}
// quit 删除连接
// 超时不发消息的断开连接

func main() {
	// 建立socket 监听端口
	netListen, err := net.Listen("tcp", "127.0.0.1:1024")
	CheckError(err)
	defer netListen.Close() // return语句执行完之前 会执行defer指定的函数

	Log("Waiting for Clients:")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		Log(conn.RemoteAddr().String(), " tcp connect success ")
		go handleConnection(conn) // goroutine 只增加了一个go 就支持了多个客户端连接
	}
}

// 处理连接
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error:", err)
			conn.Close()
			return
		}

		var msg string = string(buffer[:n])
		Log(conn.RemoteAddr().String(), " receive data string:\n ", msg)
		sender(conn, msg)
	}
}

func sender(conn net.Conn, msg string) {
	conn.Write([]byte(msg))
	fmt.Printf("server receiver msgs:%s\n", msg)
}

func Log(v ...interface{}) {
	log.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
