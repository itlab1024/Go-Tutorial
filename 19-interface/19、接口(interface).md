# 接口

接口是一个抽象概念，这跟java中的接口在概念上没有区别。go通过接口实现了多态性。

# 声明接口

go使用如下方式声明接口

```go
type 接口名 interface{}
```

# 零值

接口的零值是nil

```go
package main

import "fmt"

func main() {
	// 接口的零值
	var i1 interface{}
	fmt.Println(i1) // <nil>
}
```

# 接口实现

Java中的接口实现使用的是`implement`关键字，Go中采用的是隐式方式。

```go
package main

import "fmt"

func main() {
	// 接口调用
	var p = Human{Name: "tom"}
	fmt.Printf("类型是%T", p)
	fmt.Println(p.greet())
	
	isAPerson(&p) // 如果Human结构体没有定义greet方法，那么此处会编译错误，比如你去修改Person的greet方法，改个名，就出错了。
}

// Person 定义一个接口
type Person interface {
	greet() string
}

// Human 定义结构体
type Human struct {
	Name string
}

// 定义Human的方法
func (h *Human) greet() string {
	return "Hi, I am " + h.Name
}

func isAPerson(h Person) {
	fmt.Println(h.greet())
}
```

# 实现多接口

go中支持实现多接口，同样是隐士实现。需要实现所有方法。

```go
// 定义第二个接口
type Man interface {
	sayHello() string
}
// Human实现sayHello接口。
func (h *Human) sayHello() string {
	return h.Name + " say hello"
}
//main方法
var m Man = &p
fmt.Println(m.sayHello()) // tom say hello

```

# 接口断言

类型断言是一种获取接口持有的底层值的方法。这意味着如果为接口变量分配了一个字符串，那么它所保存的基础值就是该字符串。

接口断言通过      **接口变量.(结构体类型名)** 来判断

```go
// 接口断言
var i interface{} = Human{Name: "zhangsan"}
fmt.Println(i.(Human)) // {zhangsan}

// 上面的代码是正确的，因为i就是Human类型，当如果断言是其他类型的时候就会触发panic。
// fmt.Println(i.(string))
// i.(string)有两个返回值，可以先判断是否是string类型再获取值
s, b := i.(string)
if !b {
  fmt.Println("i不是字符串类型") // i不是字符串类型
} else {
  fmt.Println(s)
}
```

![类型断言错误触发panic](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208141042319.png)

# 接口合并

多个接口可以合并为一个接口。

```go
// 接口合并
type Merge interface {
	Person
	Human
}
type MergeImpl struct {
}
```

# 通过switch判断接口类型

go中的switch能够判断接口底层持有的类型

```go
// 通过switch判断接口类型
switch i.(type) {
  case int:
  fmt.Println("int类型")
  case string:
  fmt.Println("string类型")
  case Human:
  fmt.Println("human类型") // human类型
  default:
  fmt.Println("其他类型")
}
```

# 接口的使用

在需要多态性的 Go 中使用接口。在可以传递多种类型的函数中，可以使用接口。接口允许 Go 具有多态性。

接口是 Go 中的一个很棒的特性，推荐使用。
