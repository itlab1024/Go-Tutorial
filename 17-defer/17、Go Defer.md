![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202207281322806.png)

# 什么是Defer？

defer是Go中的一个关键字，是延迟、推迟执行的意思，Go中使用defer定义的语句会在函数（或者方法）return前执行。

# 示例

先看一个非defer的例子

```go

package main

import "fmt"

func main() {
  printMsg()
  fmt.Println("main程序执行完毕")
}
func printMsg() {
  fmt.Println("做点什么...")
}
```

运行结果如下：

```tex
做点什么...
main程序执行完毕
```

如果在printMsg()调用的前面加上defer关键字在执行

```go
package main

import "fmt"

func main() {
	//printMsg()
	defer printMsg()
	fmt.Println("main程序执行完毕")
}
func printMsg() {
	fmt.Println("做点什么...")
}
```

执行结果如下：

```tex
main程序执行完毕
做点什么...
```

执行顺序发生类变化。

# 使用场景

通过上面的例子知道defer就是将函数放到最后执行。

这在实际使用中有很多例子，比如读取文件，打开文件流，是不是出现过忘记关闭，访问数据库，连接是否忘记关闭，反正我是做过这样的错事！

这时候defer可能就很有用，因为我们可以打开流后，紧接着就是用defer关闭。防止写代码写着写着，嘎~忘了。

比如如下读取文件的例子。

```go

package main

import (
  "fmt"
  "io"
  "os"
)

func main() {
  contents, _ := Contents("1.txt")
  fmt.Println(contents)
}

// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
  f, err := os.Open(filename)
  if err != nil {
    return "", err
  }
  defer f.Close() // f.Close will run when we're finished.

  var result []byte
  buf := make([]byte, 100)
  for {
    n, err := f.Read(buf[0:])
    result = append(result, buf[0:n]...) // append is discussed later.
    if err != nil {
      if err == io.EOF {
        break
      }
      return "", err // f will be closed if we return here.
    }
  }
  return string(result), nil // f will be closed if we return here.
}
```

# 原理解析

Defer是将defer的内容压入栈来执行的。什么是栈？其特点就是后进先出（First In Last Out ）FILO。

看如下例子

```go
package main

import "fmt"

func main() {
  var i = 0
  defer fmt.Println("第一次定义的defer，i=", i)
  i++
  defer fmt.Println("第二次定义的defer，i=", i)
  i++
  fmt.Println("i=", i)
}
```

运行结果是：

```tex
i= 2

第二次定义的defer，i= 1

第一次定义的defer，i= 0
```

栈过程如下

![图片](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208091733960.png)

需要注意的是压入栈的不只是方法，还有其中的变量值等。通过运行结果可以看出，i的值是压入栈的时候的值 ，而不是程序运行到最后的值。