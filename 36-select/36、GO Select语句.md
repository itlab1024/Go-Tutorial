# Go中的select是什么？

select类似于switch语句，select是专门针对channel设计。
换句话说它只能处理channel数据。

# 语法
语法也类似switch
```go
select {
    case case1:
        // case 1...
    case case2:
        // case 2...
    case case3:
        // case 3...
    case case4:
        // case 4...
    default:
        // default case...
}
```
# 使用示例
```go
package main

import (
	"fmt"
)

func g1(ch chan int) {
	ch <- 1
}

func g2(ch chan int) {
	ch <- 2
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go g1(ch1)
	go g2(ch2)

	select {
	case v1 := <-ch1:
		fmt.Println("获取到数据: ", v1)
	case v2 := <-ch2:
		fmt.Println("获取到数据: ", v2)
	}
}
```
运行结果：
多次运行看似每次都打印"获取到数据:  2"，但实际上结果是不确定的，这取决于具体协程的执行时间等等因素。如果所有case条件都成立他会选择任一个输出。
# 默认分支
如果没有其他case准备好执行，则执行default case。这可以防止 select 阻塞主协程，因为默认情况下操作是阻塞的。
# 使用场景
go的使用场景，我目前还真的不清楚，我也查了些资料，目前在k8s中使用到了。
比如k8s中需要寻找节点部署应用的时候，通过select去查找第一个符合条件的节点进行部署。

