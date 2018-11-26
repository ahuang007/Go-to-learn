package main

import "fmt"

// 算术运算符 +,-,*,/,%,++,--
// 关系运算符 ==,!=,>,<,>=,<=
// 逻辑运算符 &&,||,!
// 位运算符  &,|,^,<<,>>
// 赋值运算符 =,+=,-=,*=,/=,%=,<<=,>>=,&=,^=,|=
// 其他运算符 &【返回变量存储地址】,*【指针变量】
// 运算符优先级【当不确定优先级时加括号】

func main() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第1行 %d\n", c)
	c = a - b
	fmt.Printf("第2行 %d\n", c)
	c = a * b
	fmt.Printf("第3行 %d\n", c)
	c = a / b
	fmt.Printf("第4行 %d\n", c)
	c = a % b
	fmt.Printf("第5行 %d\n", c)
	a++
	fmt.Printf("第6行 %d\n", a)
	a = 21
	a--
	fmt.Printf("第7行 %d\n", a)

	if a == b {
		fmt.Printf("第8行 a等于b\n")
	} else {
		fmt.Printf("第8行 a不等于b\n")
	}

	if a < b {
		fmt.Printf("第9行 a小于b\n")
	} else {
		fmt.Printf("第9行 a不小于b\n")
	}

	if a < b {
		fmt.Printf("第10行 a小于b\n")
	} else {
		fmt.Printf("第10行 a不小于b\n")
	}

	a = 5
	b = 20
	if a <= b {
		fmt.Printf("第11行 a小于b\n")
	}

	if b >= a {
		fmt.Printf("第12行 a不小于b\n")
	}

	var x bool = true
	var y bool = false
	if x && y {
		fmt.Printf("第13行 条件为true\n")
	}
	if x || y {
		fmt.Printf("第14行 条件为true\n")
	}

	x = false
	y = true
	if x && y {
		fmt.Printf("第15行 条件为true\n")
	} else {
		fmt.Printf("第15行 条件为false\n")
	}

	if !(x && y) {
		fmt.Printf("第16行 条件为true\n")
	}

	var m uint = 60                         // 60 = 0011 1100
	var n uint = 13                         // 13 = 0000 1101
	fmt.Printf("m&n:%d, %b\n", m&n, m&n)    // 12 = 0000 1100
	fmt.Printf("m|n:%d, %b\n", m|n, m|n)    // 61 = 0011 1101
	fmt.Printf("m^n:%d, %b\n", m^n, m^n)    // 49 = 0011 0001
	fmt.Printf("m<<2:%d, %b\n", m<<2, m<<2) // 240= 1111 0000
	fmt.Printf("m>>2:%d, %b\n", m>>2, m>>2) // 15 = 0000 1111

	var d int
	d = a
	fmt.Printf("第17行 d:%d\n", d)
	d += a
	fmt.Printf("第17行 d:%d\n", d)
	d -= a
	fmt.Printf("第18行 d:%d\n", d)
	d *= a
	fmt.Printf("第19行 d:%d\n", d)
	d /= a
	fmt.Printf("第20行 d:%d\n", d)
	d = 200
	d <<= 2
	fmt.Printf("第21行 d:%d\n", d)
	d >>= 2
	fmt.Printf("第22行 d:%d\n", d)
	d ^= 2
	fmt.Printf("第23行 d:%d\n", d)
	d |= 2
	fmt.Printf("第24行 d:%d\n", d)
	d &= 2
	fmt.Printf("第25行 d:%d\n", d)

	var a1 int = 4
	var a2 int32
	var a3 float32
	var ptr *int
	fmt.Printf("第26行 a1变量类型为：%T\n", a1)
	fmt.Printf("第27行 a2变量类型为：%T\n", a2)
	fmt.Printf("第28行 a3变量类型为：%T\n", a3)

	ptr = &a1
	fmt.Printf("a1的值：%d, %s\n", a1, a1)
	fmt.Printf("ptr的地址：%p\n", ptr)
	fmt.Printf("ptr的值：%d\n", *ptr)
}

// fmt.Printf格式化输出
/*
一般占位符
符号	说明
%v	相应值的默认格式
%+v	在打印结构体时，默认格式，会添加字段名
%#v	相应值的 Go 语法表示
%T	相应值的类型的 Go 语法表示
%%	字面上的百分号，并非值的占位符

布尔占位符
符号	说明
%t	单词 true 或 false

整数占位符
符号	说明
%b	二进制表示
%c	相应 Unicode 码点所表示的字符
%d	十进制表示
%o	八进制表示
%q	单引号围绕的字符字面值，由 Go 语法安全地转义
%x	十六进制表示，字母形式为小写 a-f
%X	十六进制表示，字母形式为大写 A-F
%U	Unicode 格式：U+1234，等同于 "U+%04X"

浮点数及其复合构成占位符
符号	说明
%b	无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat 的 'b' 转换格式一致。例如 -123456p-78
%e	科学计数法，例如 -1234.456e+78
%E	科学计数法，例如 -1234.456E+78
%f	有小数点而无指数，例如 123.456
%g	根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的 0）输出
%G	根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的 0）输出

字符串与字节切片占位符
符号	说明
%s	字符串或切片的无解译字节
%q	双引号围绕的字符串，由 Go 语法安全地转义
%x	十六进制，小写字母，每字节两个字符
%X	十六进制，大写字母，每字节两个字符

指针
符号	说明
%p	十六进制表示，前缀 0x
*/
