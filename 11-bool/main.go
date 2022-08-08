package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 定义布尔类型
	var b bool
	fmt.Print("布尔类型的零值是", b) // false

	fmt.Println()
	// 布尔运算符
	i := 1
	j := 2
	fmt.Println(i == j) // false
	fmt.Println(i != j) //true
	fmt.Println(i >= j) // false
	fmt.Println(i <= j) // true
	fmt.Println(i < j)  // true
	fmt.Println(i > j)  //false

	// 字符串和bool相互转化
	// ParseBool返回两个值，第一个是具体的值，第二个是error。
	b1, _ := strconv.ParseBool("true")
	fmt.Println(b1) // true
	// 如果字符串不是true或者false呢
	b2, _ := strconv.ParseBool("能正常转化成功吗")
	fmt.Println(b2) // false, 这是因为我忽略错误。也可以不忽略错误
	//处理错误
	b3, err := strconv.ParseBool("能正常转化成功吗")
	if err != nil {
		fmt.Println("转化有错误")
	} else {
		fmt.Println(b3)
	}
	b4, _ := strconv.ParseBool("True")
	fmt.Println(b4)

	// 也可以将bool转化为字符串
	fmt.Println(strconv.FormatBool(false)) // false
	fmt.Println(strconv.FormatBool(true))  // true

}
