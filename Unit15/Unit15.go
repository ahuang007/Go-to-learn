package main

import "fmt"

// range 范围 类型lua的 ipairs+pairs

func main() {
	//range用在数组或者切片上 相当于ipairs
	nums := []int{2, 3, 4}
	sum := 0
	//
	for _, num := range nums { // 如果不想用索引变量(index) 则用"_"省略
		sum += num
	}
	fmt.Println("sum = ", sum)
	for i, num := range nums { // 第一个参数是数组索引 第二个是数组对应下标的值
		if num == 3 {
			fmt.Println("find 3 index: ", i)
		}
	}

	// range也可以用在map的键值对上[第一个参数是map的key，第二个是key对应的value。]
	mps := map[string]string{"a": "iphone", "b": "android"} //map定义
	for k, v := range mps {
		fmt.Printf("key:%s => value:%s\n", k, v)
	}

	// range用来遍历unicode字符串[第一个参数是字符的索引，第二个是字符（Unicode的值）本身。]
	for i, c := range "ahuang" {
		fmt.Printf("i:%d, c:%c, c:%d\n", i, c, c) //c: %c 字符 %d ascii码
	}
}
