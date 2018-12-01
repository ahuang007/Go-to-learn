package main

import "fmt"

// 变量声明以及初始化
var a = "GoToLearn"
var b string = "ahuang"
var c bool

func main() {
	// 这种定义只能出现在函数体中
	d := "xxxx" // 等价于c++中 auto d = "xxxx"

	fmt.Println(a, b, c, d)
}
