package main

import "fmt"

func main() {
	person := Person{Name: "高斯林"}
	// 使用接受器调用方法
	person.speak()
	// 调用普通方法
	speak(person)
	// 指针类型的方法
	person.setName("华为")
}

type Person struct {
	Name string
}

// 定义一个方法
func (person Person) speak() {
	fmt.Println(person.Name + "说您好!")
}

// 定义一个函数
func speak(person Person) {
	fmt.Println(person.Name + "说您好!")
}

// 定义一个指针类型的方法
func (person *Person) setName(Name string) {
	person.Name = Name
}
