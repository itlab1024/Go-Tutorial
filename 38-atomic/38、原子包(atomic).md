# 原子操作的意义
原子操作在多线程中的意义是非常大的，他保证一段业务要么一起执行，要么不执行。如果不使用原子操作，多个线程操作统一数据的时候就出现了安全的问题，
运算结果可能是不正确的。
# 非原子操作代码
如下代码是v初始值是100，两个协程，每个协程加3000，结果应该是6100.
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
	v := 42
	var wg sync.WaitGroup
	wg.Add(2)
	go f(&v, &wg)
	go f(&v, &wg)
	wg.Wait()

	fmt.Println(v) // 6006
}
```
但是实际运行结果却是不确定的，比如我的运行结果就是6006。这就产生了安全问题。
v++操作其实是分步骤的，第一步读v的数据，第二步+1，第三步再写数据。
# 原子操作
上面的例子，通过原子类修改下
```go
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
```
最终结果永远是6100。