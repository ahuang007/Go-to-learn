package main

import "fmt"

// 同一个包下 不允许有多个main函数
/*
func main() {
	fmt.Println("part1 main")
}
*/

var a1 int = 10

func part1Say() {
	fmt.Println("i'm part1")
}
