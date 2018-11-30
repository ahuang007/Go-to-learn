package main

import "fmt"

func main() {
	var a int = 10
	var b int = 15
	var c int
	c = max(a, b)
	fmt.Println("c:%d", c)

	var s1 string = "abc"
	var s2 string = "def"
	fmt.Println("before swap:", s1, s2)
	fmt.Println(swap(s1, s2))          // 函数作为值
	fmt.Println("after swap:", s1, s2) // 值传递 不会影响实际参数

	fmt.Println("before swap:", s1, s2)
	fmt.Println(swap1(&s1, &s2))       // 函数作为值
	fmt.Println("after swap:", s1, s2) // 引用传递 会影响实际参数
}

func max(a int, b int) int { // 等价写法 func max(a, b int) int 类型参数相同则只需要写一个类型
	if a > b {
		return a
	} else {
		return b
	}
}

func swap(x, y string) (string, string) { // 函数支持多个返回值
	var tmp string
	tmp = x
	x = y
	y = tmp

	return x, y
}

func swap1(x *string, y *string) (string, string) { // 注意：swap1不能与swap同名 也就是不能通过参数不同来重载
	var tmp string
	tmp = *x
	*x = *y
	*y = tmp

	return *x, *y
}
