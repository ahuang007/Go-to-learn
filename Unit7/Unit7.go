package main

import (
	"fmt"
	"os"
	"time"
)

/*
Channel是Go中的一个核心类型，你可以把它看成一个管道，通过它并发核心单元就可以发送或者接收数据进行通讯(communication)。
它的操作符是箭头 <- 。
ch := make(chan int) // 初始化channel
ch <- v    // 发送值v到Channel ch中
v := <-ch  // 从Channel ch中接收数据，并将数据赋值给v
close(ch) // 关闭channel

select是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。
select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。

每个case都必须是一个通信
所有channel表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通信可以进行，它就执行；其他被忽略。
如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
否则：
如果有default子句，则执行该语句。
如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。

*/

func launch() {
	fmt.Println("nuclear launch detected")
}

func startCD(canLunch chan int) {
	c := time.Tick(time.Second * 1)
	for cd := 10; cd > 0; cd-- {
		<-c
		fmt.Println(cd)
	}
	canLunch <- -1
}

func isAbort(abort chan int) {
	os.Stdin.Read(make([]byte, 1))
	abort <- -1
}

func main() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1]n")
	case c2 <- i2:
		fmt.Printf("send ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

	fmt.Println("start cd")
	abort := make(chan int)
	canLunch := make(chan int)
	go startCD(canLunch)
	//go isAbort(abort)
	select {
	case <-canLunch:
		fmt.Println("cd end")
		break
	case <-abort:
		fmt.Println("Launch aborted!")
		break
	}
	launch()
}
