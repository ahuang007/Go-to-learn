package main

import "fmt"

func main() {
	var a int = 20
	var iptr *int                        //声明指针变量
	fmt.Printf("指针iptr指向的地址：%x\n", iptr) // 0 空指针
	if iptr == nil {                     // nil为空指针 此处不能写成 iptr == 0 因为类型不一致
		fmt.Printf("iptr为nil\n")
	}

	iptr = &a                      // 给指针赋值
	fmt.Printf("a变量的地址是：%x\n", &a) // %x 16进制
	fmt.Printf("指针iptr指向的地址:%x\n", iptr)
	fmt.Printf("指针iptr指向的地址的值:%d\n", *iptr)

	const MAX int = 3
	arr := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int // 指针数组
	for i = 0; i < MAX; i++ {
		ptr[i] = &arr[i]
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("arr[%d] = %d\n", i, *ptr[i])
	}

	var b int
	var ptr1 *int   // 指向整型变量的指针
	var pptr1 **int // 指向指针变量的指针(其中这个指针变量指向整型变量)
	b = 3000
	ptr1 = &b
	pptr1 = &ptr1
	fmt.Printf("变量b = %d\n", b)
	fmt.Printf("指针变量*ptr1 = %d\n", *ptr1)
	fmt.Printf("指向指针的指针变量**ptr1 = %d\n", **pptr1)

	var m int = 100
	var n int = 200
	fmt.Printf("before swap m:%d, n:%d\n", m, n)
	swap(&m, &n)
	fmt.Printf("after swap m:%d, n:%d\n", m, n)
}

// 指针作为函数参数(用引用也可以达到相同的目的 值传递不行)
func swap(x *int, y *int) {
	*x, *y = *y, *x // 简洁的写法 等价于下面注释代码

	/*
		var tmp int
		tmp = *x
		*x = *y
		*y = tmp
	*/
}
