package main

import "fmt"

var arr = [5]int{1, 2, 3, 4, 5}

// 等价于 var arr = []{1,2,3,4,5}
// 初始化数组中{}的元素个数不能大于[]中的数字
// 如果忽略[]中的数字 不设置数组大小，go语言会根据元素的个数来设置数组的大小

func main() {
	arr[0] = 99      // go数组索引从0开始
	fmt.Println(arr) // Println可以直接打印数组

	var iarr [10]int //数组声明
	var i, j int

	for i = 0; i < cap(iarr); i++ { // cap(arr) 数组长度
		iarr[i] = i + 100 // 数组赋值
	}

	for j = 0; j < cap(iarr); j++ {
		fmt.Printf("Element[%d] = %d\n", j, iarr[j]) //数组用下标访问
	}

	// 二维数组
	var a = [5][2]int{
		{0, 0},
		{1, 2},
		{2, 4},
		{3, 6},
		{4, 8},
	}
	var m, n int
	for m = 0; m < cap(a); m++ {
		for n = 0; n < cap(a[0]); n++ {
			fmt.Printf("a[%d][%d] = %d\n", m, n, a[m][n])
		}
	}

	var scores = []int{100, 95, 80, 73, 69, 52, 88}
	var avg float32
	avg = getAverage(scores)
	fmt.Printf("平均分为：%f\n", avg)
}

//数组作为函数参数
func getAverage(arr []int) float32 {
	var i, sum int
	var avg float32

	arr_len := cap(arr)
	for i = 0; i < arr_len; i++ {
		sum += arr[i]
	}

	avg = float32(sum / arr_len) // 注意必须强制类型转换
	return avg
}
