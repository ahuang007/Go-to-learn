package main

import (
	"fmt"
)

// go 接口 任何其他类型 只要实现了这些方法就是实现了这个接口

type Phone interface {
	call()
}

type XiaoMi struct {
}

func (xiaomi XiaoMi) call() {
	fmt.Println("I'm xiaomi, you can use me to call others")
}

type HuaWei struct {
}

func (huawei HuaWei) call() {
	fmt.Println("I'm huawei, I can call you")
}

type iPad struct {
}

// 方法：包含了接受者的函数
func (ipad iPad) play() {
	fmt.Println("I'm iPad, you can use me to play")
}

func main() {
	var phone Phone

	phone = new(XiaoMi)
	phone.call()

	phone = new(HuaWei)
	phone.call()

	//phone = new(iPad) // iPad没有实现Phone的call方法 所以不能用Phone这个接口
	/*
		cannot use new(iPad) (type *iPad) as type Phone in assignment:
		iPad does not implement Phone (missing call method)
	*/

	ipad := new(iPad)
	ipad.play()
}
