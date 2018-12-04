package main

import "fmt"

//slice 类型 c++ vector
//内置类型切片("动态数组"),长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
func main() {
	var slice0 = make([]int, 3, 5) // 切片初始化 make(type, len, cap)
	printSlice(slice0)

	var slice1 []int   // 空切片
	printSlice(slice1) // len:0 cap:0
	if slice1 == nil {
		fmt.Printf("切片是空的\n")
	}

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8} // 创建切片
	printSlice(numbers)

	fmt.Println("numbers == ", numbers)
	fmt.Println("numbers[1:4] == ", numbers[1:4]) // [1,4)
	fmt.Println("numbers[:3] == ", numbers[:3])   // [0,3)
	fmt.Println("numbers[4:]", numbers[4:])       // [4, len())

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)
	numbers2 := numbers[:2] // 截取切片numbers初始化新切片numbers2
	printSlice(numbers2)
	numbers3 := numbers[2:5] // 如果截取下标超过cap() 则编译报错：数组越界
	printSlice(numbers3)

	// copy 拷贝切片
	// append  向切片追加新元素
	var slice2 []int
	printSlice(slice2)

	slice2 = append(slice2, 0) //  允许追加空切片
	printSlice(slice2)

	slice2 = append(slice2, 1) // 向slice 添加一个元素
	printSlice(slice2)

	slice2 = append(slice2, 2, 3, 4) // 同时向slice添加多个元素
	printSlice(slice2)

	// slice扩容类似c++ vectoro扩容
	// 创建切片slice3  是之前切片的两倍容量
	slice3 := make([]int, len(slice2), cap(slice2)*2)
	copy(slice3, slice2) // 拷贝slice2的内容到slice3
	printSlice(slice3)
}

func printSlice(x []int) {
	// len() 切片长度(当前元素的个数)
	// cap() 切片容量(切片可以容纳的元素最大值)
	// %v 相应值的默认格式(%T 相应值的类型的 Go 语法表示)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(x), cap(x), x)
}
