package main

import "fmt"

type Book struct {
	name    string // 书名
	author  string // 作者
	subject string // 学科
	id      int    // 编号
}

func main() {
	// 创建一个结构体并按key顺序初始化(不用写key)
	book1 := Book{"go从入门到精通", "ahuang", "计算机", 1001}
	// 用key=>value格式初始化结构体
	book2 := Book{
		name:    "c++从入门到放弃",
		author:  "ahuang",
		subject: "计算机",
		id:      1002}
	//忽略的字段有默认值(字符串为"" 数字为0)
	book3 := Book{name: "mysql从入门到删库跑路", id: 1003}

	fmt.Println(book1) // fmt.Println可以直接打印结构体
	printBook(book2)   // 结构体做函数参数 访问结构体变量用"."
	printBook1(&book3) // 结构体指针做函数参数 访问结构体变量也用"."
}

func printBook(book Book) {
	fmt.Printf("Book name: %s\n", book.name)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject: %s\n", book.subject)
	fmt.Printf("Book id: %d\n", book.id)
}

func printBook1(book *Book) {
	fmt.Printf("Book name: %s\n", book.author)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject: %s\n", book.subject)
	fmt.Printf("Book id: %d\n", book.id)
}
