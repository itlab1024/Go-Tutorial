package main

import "fmt"

func main() {
	// 接口的零值
	var i1 interface{}
	fmt.Println(i1) // <nil>

	// 接口调用
	var p = Human{Name: "tom"}
	fmt.Printf("类型是%T", p)
	fmt.Println(p.greet())

	isAPerson(&p) // 如果Human结构体没有定义greet方法，那么此处会编译错误，比如你去修改Person的greet方法，改个名，就出错了。
	var m Man = &p
	fmt.Println(m.sayHello()) // tom say hello

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
}

// Person 定义一个接口
type Person interface {
	greet() string
}

// 定义第二个接口
type Man interface {
	sayHello() string
}

// Human实现sayHello接口。
func (h *Human) sayHello() string {
	return h.Name + " say hello"
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

// 接口合并
type Merge interface {
	Person
	Human
}
type MergeImpl struct {
}
