package main

import (
	"fmt"
	"sync"
)

func main() {
	// 定义一个等待组
	var wg sync.WaitGroup
	// 一个子协程，所以设置为1
	wg.Add(1)
	// 调用协程，传递等待组的指针
	go f1(&wg)
	// 主协程等待
	wg.Wait()
	// 子协程执行完毕后执行
	fmt.Println("主协程执行完毕")
}
func f1(wg *sync.WaitGroup) {
	fmt.Println("子协程执行完毕")
	wg.Done()
}
