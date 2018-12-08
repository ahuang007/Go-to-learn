package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

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
		handleConnection(conn)
	}
}

// 处理连接
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error:", err)
			return
		}

		Log(conn.RemoteAddr().String(), " receive data string:\n ", string(buffer[:n]))
	}
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
