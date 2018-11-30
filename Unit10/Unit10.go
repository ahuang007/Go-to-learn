package main

import "fmt"

// 闭包
func getSeq() func() int {
	i := 0 // i为闭包函数的upvalue
	return func() int {
		i++
		return i
	}
}

type Circle struct {
	radius float64
}

const pi float64 = 3.141592653

// 该method属于Circle类型对象中的方法
// Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，
// 接受者可以是命名类型或者结构体类型的一个值或者是一个指针
func (c Circle) getArea() float64 {
	// 如果pi定义为 float32 则编译不过 go不会像c++一样 低精度向高精度转换
	return pi * c.radius * c.radius
}

func main() {
	getNextNo := getSeq()    //getNextNo 作为一个函数 i = 0
	fmt.Println(getNextNo()) // i = 1
	fmt.Println(getNextNo()) // i = 2
	fmt.Println(getNextNo()) // i = 3

	getNextNo1 := getSeq()    // 创建新的函数getNextNo1  i = 0
	fmt.Println(getNextNo1()) // i = 1
	fmt.Println(getNextNo1()) // i = 2
	fmt.Println(getNextNo())  // i = 4 i的可见域为getSeq函数以及闭包内 i并不是一个全局变量

	add_func := add(1, 2)
	fmt.Println(add_func(1, 1))
	fmt.Println(add_func(0, 0))
	fmt.Println(add_func(2, 2))

	var c Circle
	c.radius = 10
	fmt.Printf("圆的面积:%f\n", c.getArea())
}

func add(x1, x2 int) func(x3, x4 int) (int, int, int) {
	i := 0 // 注意 :和= 不能分开! 如果分开 i : = 0 那么i为lable
	return func(x3, x4 int) (int, int, int) {
		i++
		return i, x1 + x2, x3 + x4
	}
}
