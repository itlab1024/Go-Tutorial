package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func f(v *int64, wg *sync.WaitGroup) {
	for i := 0; i < 3000; i++ {
		atomic.AddInt64(v, 1)
	}
	wg.Done()
}

func main() {
	var v int64 = 100
	var wg sync.WaitGroup
	wg.Add(2)
	go f(&v, &wg)
	go f(&v, &wg)
	wg.Wait()

	fmt.Println(v) // 始终是6100
}
