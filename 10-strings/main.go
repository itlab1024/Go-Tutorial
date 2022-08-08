package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 字符串定义
	var s string = "I am learning go language"
	fmt.Println(s) // I am learning go language

	// 按照索引访问
	fmt.Printf("%c\n", s[0]) // I
	fmt.Printf("%x\n", s[0]) // 49

	// 字符串的长度
	fmt.Println(len(s))    // 25
	fmt.Println(len("哈哈")) // 6
	// 字符串比较
	a := "go"
	b := "golang"
	fmt.Println(a == b) // false
	fmt.Println(a >= b) // false
	fmt.Println(a <= b) // true
	fmt.Println(a > b)  // false
	fmt.Println(a < b)  // true
	fmt.Println(a != b) // true

	// 遍历字符串
	var s1 = "我在学习go语言"
	for index, value := range s1 {
		fmt.Println("索引", index, "的值是", value)
	}
	// 使用rune输出
	runeSlice := []rune(s1)
	for index, value := range runeSlice {
		fmt.Printf("索引%d的值是%c,", index, value) // 索引0的值是我,索引1的值是在,索引2的值是学,索引3的值是习,索引4的值是g,索引5的值是o,索引6的值是语,索引7的值是言,
	}

	// 打印字符串长度
	fmt.Println()
	fmt.Println(utf8.RuneCountInString(s1)) // 8

	// 字符串是不可变的
	//s1[0] = "1" // 编译就不通过。

	// 构造字符串
	byteSlice := []byte{72, 101, 108, 108, 111}
	s3 := string(byteSlice)
	fmt.Println(s3) // "Hello"

	// 也可以使用unicode
	runeSlice2 := []rune{0x0048, 0x0065, 0x006C, 0x006C, 0x006F}
	s4 := string(runeSlice2)
	fmt.Println(s4) // Hello
}
