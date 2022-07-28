![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202207281322806.png)

go提供os包，可以获取命令行参数

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 通过os获取命令行参数
	args := os.Args
	s0 := args[0]
	s1 := args[1]
	fmt.Println("s0=", s0, "s1=", s1)
}
```

先编译

```shell
➜  3 git:(main) ✗ go build main.go
```

执行编译后的文件，并传递两个参数

```
➜  3 git:(main) ✗ ./main hello go
s0= ./main s1= hello
```

可以看到已经接受到了参数。

这样传递参数有一个问题，必须按照顺序传递参数，能否不按照顺序，通过标志来实现呢，比如我使用`-s0 hello -s1 go`。这需要使用`flag`包的StringVar函数。

```
// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVar(p *string, name string, value string, usage string) {
	CommandLine.Var(newStringValue(value, p), name, usage)
}
```

实现代码如下。

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// 通过os获取命令行参数
	//args := os.Args
	//s0 := args[0]
	//s1 := args[1]
	//fmt.Println("s0=", s0, "s1=", s1)

	//使用flag编辑参数
	var s0 string
	var s1 string
	flag.StringVar(&s0, "s0", "s0默认值", "s0的描述信息")
	flag.StringVar(&s1, "s1", "s1默认值", "s1的描述信息")
	flag.Parse() // 转化
	fmt.Println("s0=", s0, "s1=", s1)
}
```

编译并运行

```shell
#编译
➜  3 git:(main) ✗ go build main.go
#传递参数,s0,s1参数顺序无所谓。
➜  3 git:(main) ✗ ./main -s0 hello -s1 go
s0= hello s1= go
#不传递参数,则打印默认值
➜  3 git:(main) ✗ ./main
s0= s0默认值 s1= s1默认
#如果指定传递s1但是没有设置值，则会报错。并且提示错误信息（请注意自己设置的第四个参数，使用说明）
➜  3 git:(main) ✗ ./main -s0 hello -s1   
flag needs an argument: -s1
Usage of ./main:
  -s0 string
    	s0的描述信息 (default "s0默认值")
  -s1 string
    	s1的描述信息 (default "s1默认值")
```

