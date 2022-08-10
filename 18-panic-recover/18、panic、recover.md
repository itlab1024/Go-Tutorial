![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202207281322806.png)

# Panic

panic的意思是恐慌，go推荐使用error来处理错误，程序中一般使用错误（Error）来处理异常情况，但是有些情况当异常发生时程序不应该继续运行（比如数据库连接无法建立，那么Web程序不应该被启动），这个时候我们就使用Panic，当函数执行Panic时候，他会中止运行，程序控制会返回到函数的调用者（如果有defer函数会先执行defer函数），接下来会打印panic的堆栈信息。

举例说明：

```go
package main

import "fmt"

func main() {
	// 会触发panic
	startWeb(nil)
}
func startWeb(dbUrl *string) {
	defer fmt.Println("执行defer语句")
	if dbUrl == nil {
		panic("数据库地址是nil，无法启动web应用")
	}
	fmt.Println("数据库地址是", *dbUrl, " 程序正常启动")
}
```

运行结果：

```tex
执行defer语句
panic: 数据库地址是nil，无法启动web应用

goroutine 1 [running]:
main.startWeb(0x0)
        /Users/itlab/workspace/github/go-tutorial/18-panic-recover/main.go:12 +0x21d
main.main()
        /Users/itlab/workspace/github/go-tutorial/18-panic-recover/main.go:7 +0x1b
```

# Recover

recover 是一个内建函数，用于重新获得 panic 协程的控制。



如何理解：上面的例子 panic将异常抛给了调用者。此时控制权在调用者。使用recover可以保持控制者为当前函数。



Go推荐尽可能使用Error，只有遇到真正的异常的时候，才用到Painc，Go使用defer, panic, recover三者合一来进行异常处理。



如何使用？

recover必须使用defer内部。

例如：

```go
func testPanic() (ret int) {
  defer func() {
    if err := recover(); err != nil {
      fmt.Println(err)
      ret = -1
    }
  }()
  var i = 10
  var j = 0
  var t = i / j
  fmt.Println(t)
  return t
}
```

运行结果：

```tex
测试开始
runtime error: integer divide by zero
-1
测试结束
```

当程序执行到14行的时候，因为除数是0，会引发异常，第6行defer中使用了recover捕获了异常，并且对异常进行了处理t=-1，最终返回给main的值是-1
