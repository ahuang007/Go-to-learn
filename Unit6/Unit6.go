package main

import "fmt"

/*
switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。。

switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加break
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。
类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3。

Type Switch
switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
*/

func main() {
	var grade string = "B"
	var marks int = 90
	switch marks {
	case 90:
		grade = "A" // 不需要加break
	case 80:
		grade = "B"
	case 50, 60, 70: // 同时测试多个可能符合条件的值，用逗号分隔
		grade = "C"
	default:
		grade = "D"
	}

	switch grade {
	case "A":
		fmt.Printf("你的等级是:%s\t优秀\n", grade)
	case "B", "C":
		fmt.Printf("你的等级是:%s\t良好\n", grade)
	case "D":
		fmt.Printf("你的等级是:%s\t及格\n", grade)
	case "F":
		fmt.Printf("你的等级是:%s\t不及格\n", grade)
	default:
		fmt.Printf("你的等级是:%s\t差\n", grade)
	}

	var a int = -1
	switch { // swtich的便利不是必须的 只要保证条件是相同类型即可
	case a > 3:
		fmt.Println("a>3")
	case a == 3:
		fmt.Println("a=3")
	case a > 0 && a < 3:
		fmt.Println("0<a<3")
	default:
		fmt.Println("a<0")
	}

	//Go语言提供了一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
	var x interface{} = "abc"
	switch i := x.(type) { // 使用type进行类型查询时，只能在switch中使用，且使用类型查询的变量类型必须是interface{}
	case nil:
		fmt.Printf("x的类型：%T\n", i)
	case int:
		fmt.Printf("x的类型：int\n")
	case float64:
		fmt.Printf("x的类型：float64\n")
	case bool, string:
		fmt.Printf("x的类型：%T\n", i)
	case func(int) float64:
		fmt.Printf("x的类型：func(int)\n")
	default:
		fmt.Printf("未知类型\n")
	}
}
