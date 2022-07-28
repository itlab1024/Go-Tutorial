package main

import "fmt"

func main() {
	// 声明v1变量
	var v1 int
	// 初始化
	v1 = 1
	fmt.Println(v1)
	// 声明多个变量
	var (
		name string = "go"
		age  int    = 10
	)
	fmt.Println(name, age)
	// 声明多个变量第二种方式
	var a, b = 1, 2
	fmt.Println(a, b)

	// 常量声明
	const PI = 3.1415926

	// 声明多个常量
	const (
		HOUR = 60
		DAY  = 24 * 60
	)

	// 声明多个常量第二种方式
	const c, d = 1, 2
	fmt.Println(c, d)

	//HOUR = 70 // 这是不允许的，常量不可变

	// 类型推断，省略类型
	var v2 = 1
	fmt.Println(v2)

	// 简短定义赋值
	v3 := 1
	fmt.Println(v3)
}
