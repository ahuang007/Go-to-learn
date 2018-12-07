package main

import "fmt"

// go错误处理
/*
Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
error类型是一个接口类型，定义如下：
type error interface {
    Error() string
}

在编码中通过实现 error 接口类型来生成错误信息。
函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息
*/

type DiviceError struct {
	dividee int //被除数
	divider int //除数
}

// 实现'error'接口
func (de DiviceError) Error() string {
	strFormat := `
	Cannot proceed, the divider iszero,
	dividee:%d
	divider:%d
	`
	return fmt.Sprintf(strFormat, de.dividee, de.divider)
}

func Divide(dividee int, divider int) (result int, errorMsg string) {
	if divider == 0 {
		de := DiviceError{
			dividee: dividee,
			divider: divider,
		}
		errorMsg = de.Error()
		return // 这里没有这样写 "return 0, errorMsg" 下面却得到了正确的结果 why?
	} else {
		return dividee / divider, ""
	}
}

func main() {
	// 下面语句等价于
	/*
			result, errorMsg := Divide(100, 9)
		 	if errorMsg == "" {
				// ...
			}
	*/

	var dividee int = 100
	var divider int = 9
	if result, errorMsg := Divide(dividee, divider); errorMsg == "" {
		fmt.Printf("%d/%d = %d\n", dividee, divider, result)
	}

	divider = 0
	if _, errorMsg := Divide(dividee, divider); errorMsg != "" {
		fmt.Printf("errorMsg is: %s\n", errorMsg)
	}
}
