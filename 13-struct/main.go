package main

import "fmt"

// Person 结构体，里面有属性name和age
type Person struct {
	// 定义属性和方法
	name string
	age  int
}

func main() {
	// 零值
	var person Person
	fmt.Println(person) //{ 0}

	// 初始化
	//第一种初始化方式
	var p = Person{age: 1, name: "golang"}
	fmt.Println(p) // {golang 1}
	// 第二种使用new,new返回的是是指针。
	var pointer = new(Person)
	//下方赋值使用(*pointer).age = 1是对的，go为了简介允许不使用*。
	pointer.age = 1
	pointer.name = "go"
	fmt.Println(*pointer) // {go 1}
	// 第三种使用&符号
	var p2 = &Person{age: 1, name: "golang"}
	fmt.Println(p2)

	// 匿名结构体
	anonymousStruct := struct{ name string }{name: "golang"}
	fmt.Println(anonymousStruct) //{golang}

	// 匿名字段访问
	t1 := T1{age: 1, string: "golang"}
	fmt.Println(t1) // {golang 1}

	// 调用拥有函数字段的结构体
	t2 := T2{age: 1, name: "go", f1: func(s string) string {
		return "结构体中的函数"
	}}
	fmt.Println(t2.age)    // 1
	fmt.Println(t2.name)   // go
	fmt.Println(t2.f1("")) // 结构体中的函数

	// 结构体嵌套
	author := Author{name: "作者1"}
	book := Book{name: "golang学习", author: author}
	fmt.Println(author) // {作者1}
	fmt.Println(book)   // {golang学习 {作者1}}

	// 嵌套结构体（匿名）
	book2 := Book2{name: "golang学习", Author2: Author2{name: "作者1", age: 39}}
	fmt.Println("作者年龄", book2.age) // 通过book2直接访问Author2的信息

	// 结构体比较
	c1 := T1{age: 1, string: "go"}
	c2 := T1{age: 1, string: "go"}
	if c1 == c2 {
		fmt.Println("c1 == c2") // c1 == c2
	} else {
		fmt.Println("c1 != c2")
	}

	// 结构体切片
	var a1 []T1
	k1 := append(a1, t1)
	fmt.Println(k1) // [{golang 1}]

}

type T1 struct {
	string // 该字段就是匿名字段
	age    int
}

// Func1 定义函数类型
type Func1 func(s string) string

// T2 函数作为结构体字段
type T2 struct {
	age  int
	name string
	f1   Func1
}

// Author 用户结构体
type Author struct {
	name string
}

//Book 书籍结构体
type Book struct {
	name   string
	author Author
}

// Author2 用户结构体
type Author2 struct {
	name string
	age  int
}

//Book2 书籍结构体
type Book2 struct {
	name string
	Author2
}
