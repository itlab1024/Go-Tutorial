package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// strings.split切割字符串
	s := "I,am,learning,go,language"
	slice := strings.Split(s, ",")
	fmt.Println(slice) //[I am learning go language]
	n := strings.SplitN(s, ",", 2)
	fmt.Println(n) // [I am,learning,go,language]

	// 使用空格切割字符串
	s1 := "i am learning go language"
	fmt.Println(strings.Split(s1, " ")) // [i am learning go language]
	fmt.Println(strings.Fields(s1))     // [i am learning go language]
	// 不删除分隔符的字符串切分
	s2 := "I,am,learning,go"
	fmt.Println(strings.SplitAfter(s2, ",")) //[I, am, learning, go]
	// 正则表达式
	s3 := "itlab1024@163.com"
	t := regexp.MustCompile(`@`)
	v := t.Split(s3, -1)
	fmt.Println(v) // [itlab1024 163.com]
}
