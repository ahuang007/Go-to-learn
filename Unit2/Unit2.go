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

	for i := 1; i < 10; i++ {
		fmt.Printf("%d ", Fibonacci(uint32(i))) // 类型转换 type_name(expression)
	}
	fmt.Println()
}

// 递归调用
func Fibonacci(n uint32) uint32 {
	if n <= 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}
