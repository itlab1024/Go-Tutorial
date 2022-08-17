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

	// 使用kind返回类型的种类
	fmt.Println(reflect.TypeOf(T{Name: ""}))        // main.T
	fmt.Println(reflect.TypeOf(T{Name: ""}).Kind()) // struct
	fmt.Println(reflect.TypeOf(1))                  // int
	fmt.Println(reflect.TypeOf(1).Kind())           // int
}

type T struct {
	Name string
	age  int
}

func (receiver T) GetName() {

}
func (receiver T) getAge() {

}
