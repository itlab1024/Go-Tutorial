package main

import "fmt"
import _ "go-tutorial/27-init2"
import _ "go-tutorial/27-init1"

func main() {
	fmt.Println("main函数执行")
}
func init() {
	fmt.Println("初始化函数1")
}
func init() {
	fmt.Println("初始化函数2")
}
