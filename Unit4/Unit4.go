package main

import "fmt"
import "unsafe"

//常量做枚举值
//数字0、1、2 分别代表未知性别、女性、男性
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

//常量可以用len(),cap(),unsafe.Sizeof()函数计算表达式的值。
//常量表达式中，函数必须是内置函数，否则编译不过
const (
	a1 = "abc"
	a2 = len(a1)
	a3 = unsafe.Sizeof(a2)
)

// iota 特殊常量 在const关键字出现时将被重置为0（const内部的第一行之前）
// const 每新增一行常量声明将使iota计数一次(iota可以理解为const语句块中的行索引)
const (
	b1 = iota
	b2
	b3
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" // 多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)
	println(a1, a2, a3)

	arr := []int{1, 2, 3} //var arr = [3]int{1, 2, 3}
	println(cap(arr))     // cap()计算数组长度

	println(b1, b2, b3)
}
