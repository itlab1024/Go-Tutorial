# 什么是等待组Waitgroups?
协程等待是很常见的，比如sleep方式，以及使用channel等待。
但是sleep不合理，如果不需要传递数据也不用使用channel，此时如何等待呢？
等待组就提供了一种简单的方式。
# 代码示例
```go
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
```
运行结果：
![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208172013514.png)
正常执行，限制性了子协程，主协程等待子协程执行完毕后继续执行。

# 具有匿名函数的等待组
```go
package main
 
import (
    "fmt"
    "sync"
)
 
func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
		fmt.Println("子协程执行完毕")
        wg.Done()
    }()
    wg.Wait()
	fmt.Println("主协程执行完毕")
}
```