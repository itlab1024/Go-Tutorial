package main

import (
	"fmt"
	"sync"
)

func f(v *int, wg *sync.WaitGroup, m *sync.Mutex) {
	// 加锁
	m.Lock()
	for i := 0; i < 3000; i++ {
		*v++
	}
	// 解锁
	m.Unlock()
	wg.Done()
}

func main() {
	v := 100
	var wg sync.WaitGroup
	wg.Add(2)
	// 定义锁
	var m sync.Mutex
	go f(&v, &wg, &m)
	go f(&v, &wg, &m)
	wg.Wait()

	fmt.Println(v) // 6006
}
