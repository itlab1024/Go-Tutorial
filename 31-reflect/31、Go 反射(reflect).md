# Go中的反射
Go中的反射跟Java中的反射功能类似，能获取到类型、值、属性、方法等内容。
# 获取类型和值
通过reflect.TypeOf获取类型，通过reflect.ValueOf获取值。
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 1
	// 获取类型
	t := reflect.TypeOf(i)
	fmt.Println(t) // int
	// 获取值
	value := reflect.ValueOf(t)
	fmt.Println(value) // 1
}
```
# 获取字段数
可以通过反射获取结构体中的字段数目
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 1
	// 获取类型
	t := reflect.TypeOf(i)
	fmt.Println(t) // int
	// 获取值
	value := reflect.ValueOf(t)
	fmt.Println(value) // 1

	// 获取结构体的字段数目
	num := reflect.ValueOf(T{Name: "张三", age: 1}).NumField()
	fmt.Println(num) // 2
	num = reflect.ValueOf(T{Name: "张三", age: 1}).NumMethod()
	fmt.Println(num) // 1

}

type T struct {
	Name string
	age  int
}

func (receiver T) GetName() {

}
func (receiver T) getAge() {

}
```
这里特别要注意的是，NumField会统计所有导出和位导出的属性，NumMethod只会统计导出的方法。
在go中首字母大写的属性或者方法是导出的，可以在其他包中方法。
# 获取类型种类
比如我们定义一个struct，T，如果使用TypeOf会返回T类型，如果使用Kind会返回struct。对于非包装类型，结果相同。
```go
// 使用kind返回类型的种类
fmt.Println(reflect.TypeOf(T{Name: ""}))        // main.T
fmt.Println(reflect.TypeOf(T{Name: ""}).Kind()) // struct
fmt.Println(reflect.TypeOf(1)) // int
fmt.Println(reflect.TypeOf(1).Kind()) // int
```