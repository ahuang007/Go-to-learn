package main

import (
	"fmt"
)

func main() {
	mapScore := make(map[string]int)
	mapScore["zhangsan"] = 98
	mapScore["lisi"] = 86
	mapScore["wangwu"] = 100

	mapPlayer := map[int64]string{
		10001: "zhangsan",
		10002: "lisi",
		10003: "wangwu",
	}

	for k, v := range mapPlayer {
		fmt.Println(k, v, mapScore[v])
	}

	//golang 的slice是一个指向底层的数组的指针结构体。
	//这个结构体有三个属性，1.指向数组指针，2.len： slice中元素的数量 3.cap：slice占用内存数量。
	//当切片长度超过容量时，底层数组会重新分配内存空间(扩容)

	my_slice := make([]int, 10, 100)
	fmt.Println(len(my_slice))
	fmt.Println(cap(my_slice))

	for i := 1; i < 10; i++ {
		my_slice[i-1] = i
	}

	my_slice[2] = 66

	fmt.Println(my_slice)
}
