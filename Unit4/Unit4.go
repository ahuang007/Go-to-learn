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
	a2 = len(a1)           // 3 字符串长度
	a3 = unsafe.Sizeof(a1) // 16
	//字符串类型在go里是个结构，包含指向底层数组的指针和长度
	//这两部分每部分都是8字节，所以字符串长度为16字节
)

// iota 特殊常量 在const关键字出现时将被重置为0（const内部的第一行之前）
// const 每新增一行常量声明将使iota计数一次(iota可以理解为const语句块中的行索引)
const (
	b1 = iota // 0
	b2        // 1
	b3        // 2
)

const (
	d1 = 1 << iota //1
	d2 = 3 << iota //6
	d3             //12
	d4             //24
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
	println(a1, a2, a3, a4)

	arr := []int{1, 2, 3} //var arr = [3]int{1, 2, 3}
	println(cap(arr))     // cap()计算数组长度

	println(b1, b2, b3)

	const (
		c1 = iota // 0
		c2        // 1
		c3        // 2
		c4 = "ha" // "ha" iota += 1
		c5        // "ha" iota += 1
		c6 = 100  // 100  iota += 1
		c7        // 100  iota += 1
		c8 = iota // 7, 恢复计数
		c9        // 8
	)

	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9)
	fmt.Println(d1, d2, d3, d4)
}
