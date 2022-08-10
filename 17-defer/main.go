package main

import "fmt"

func main() {
	//printMsg()
	defer printMsg()
	fmt.Println("main程序执行完毕")
}
func printMsg() {
	fmt.Println("做点什么...")
}
