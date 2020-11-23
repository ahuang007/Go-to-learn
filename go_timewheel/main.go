package main

import (
	"fmt"
	"time"
)

// 定义心跳包，设置心跳超时时间，处理函数
var wheelHeartBeat = New(time.Second*1, 30, func(data interface{}) {
	fmt.Println("timeout close")
})

func main() {
	wheelHeartBeat.Start()
}
