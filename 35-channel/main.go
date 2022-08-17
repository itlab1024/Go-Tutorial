package main

import "fmt"

func main() {
	// 定义通道
	var ch chan int
	fmt.Println("零值是", ch) // 零值是 <nil>
	// 初始化
	ch = make(chan int)
	go func() {
		ch <- 1
		fmt.Println("子协程执行完毕")
		// 关闭通道
		close(ch)
	}()
	//v := <-ch
	//fmt.Println("子协程写入通道的数据是", v)
	for v := range ch {
		fmt.Println("子协程写入通道的数据是", v) //子协程写入通道的数据是 1
	}
	fmt.Println("main函数执行完毕")

	// 非阻塞通道定义
	ch1 := make(chan int, 2)
	fmt.Println(ch1)

}
