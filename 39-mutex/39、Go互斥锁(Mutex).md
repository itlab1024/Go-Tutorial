# 互斥条件
当多个协程访问共享数据的时候，就会出现互斥条件，导致他们无法正确更新数据结果。
# 代码示例
之前学习原子包的时候提到一个无法正确计算结果的例子
```go
package main

import (
	"fmt"
	"sync"
)

func f(v *int, wg *sync.WaitGroup) {
	for i := 0; i < 3000; i++ {
		*v++
	}
	wg.Done()
}

func main() {
	v := 100
	var wg sync.WaitGroup
	wg.Add(2)
	go f(&v, &wg)
	go f(&v, &wg)
	wg.Wait()

	fmt.Println(v) // 6006
}
```
当时使用原子int来实现的并发正确执行。那么这里使用锁是同样可以实现的。
看如下代码
```go
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
```
运行结果也是没有问题的。
# 互斥锁和原子包
互斥锁与原子操作非常相似，但它们比这要复杂得多，这可能会让人感到困惑。原子利用 CPU 指令，而互斥锁利用锁定机制。因此，当更新整数等共享变量时，原子更快。但是当同时处理复杂的数据结构时，互斥锁的真正威力就来了。然后它是唯一的选择，因为原子不支持它。