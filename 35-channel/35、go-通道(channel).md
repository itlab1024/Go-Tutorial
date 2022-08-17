# 什么是通道(channel)
channel是go中的一种数据类型，是协程通信的媒介。
# 通道的定义
go中的通道使用chan定义，需要使用make函数初始化，否则零值是nil，无法使用。
```go
// 定义通道
var ch chan int
fmt.Println("零值是", ch) // 零值是 <nil>
```
# 如何通信
刚刚讲解协程的时候，我使用sleep来保证子协程能够正常执行，现在我通过channel来实现下。
```go
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
	}()
	v := <-ch
	fmt.Println("子协程写入通道的数据是", v)
	fmt.Println("main函数执行完毕")
}
```
这里要说明下：
ch <- 1：代表向通道中写入数据
v := <-ch：代表从通道中读取数据
![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208171933902.png)
通过打印日志可以看到，正确执行了主协程和子协程。
# 阻塞通道和非阻塞通道
上面使用的通道是阻塞通道，默认就是阻塞通道，因为没有设置其缓冲区。
反之设置缓冲区的通道就是非阻塞通道
```go
// 非阻塞通道定义
ch1 := make(chan int, 2)
fmt.Println(ch1)
```
第二个参数大于0的通道就是非阻塞通道，缓冲区大小就是2。
# 通道的方向
通道可以分为双向通道和单向通道，单向通道又分为读通道和写通道。
双向通道就像上面那样定义，使用chan
只读通道使用<-chan 定义
只写通道使用chan->定义
设计双向通道的意义主要是防止通道滥用，针对只需要读数据的通道就要使用只读通道。
# 通道关闭
通道发送完毕数据后可以关闭通道,只读通道是无需关闭的，编译也会报错。
```go
go func() {
    ch <- 1
    fmt.Println("子协程执行完毕")
    // 关闭通道
    close(ch)
}()
```
# 使用range读取通值
可以使用range关键字循环接收通道数据。
```go
for v := range ch {
    fmt.Println("子协程写入通道的数据是", v) //子协程写入通道的数据是 1
}
```