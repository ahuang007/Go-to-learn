package main // 包名

import "fmt" // 引入其他模块

// init函数在main函数之前执行
func init() {
	fmt.Println("run before main")
}

// 一个package下面只能有一个main函数
// main函数是程序入口函数
func main() {
	fmt.Println("Hello, World!")

	// 同一个包 可以拆分成很多文件 不需要引入
	// 包内的函数以及变量都是包内可见的
	fmt.Println("part1 a1: ", a1)
	part1Say()
}
