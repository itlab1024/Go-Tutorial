package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("协程方法执行完毕") // 协程方法执行完毕
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main方法执行完毕") // main方法执行完毕
}
