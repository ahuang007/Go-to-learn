package main

import "fmt"

func main() {
	var b int = 15
	var a int
	numbers := [6]int{1, 2, 3, 5}

	/* for循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a的值为：%d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a的值为：%d\n", a)
	}

	var c int
	for {
		if c < 10 {
			c++
			if c%2 == 0 { // c为偶数不打印c的值
				continue
			}
			fmt.Printf("c的值为：%d\n", c)
		} else {
			break
		}
	}

	for i, v := range numbers {
		fmt.Printf("index:%d = value:%d\n", i, v)
	}

	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不为素数
			}
		}
		if j > (i / j) { // 只有当前面循环不跳出 才是素数 前面循环执行完 j = (i/j) + 1
			fmt.Printf("%d 是素数\n", i)
		}
	}

	var x int = 10
lable1:
	for x < 20 {
		if x == 15 {
			x++
			goto lable1 //continue用goto实现
		}
		fmt.Printf("x的值：%d\n", x)
		x++
	}
}
