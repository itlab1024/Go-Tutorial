package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 序列化
	book := Book{Name: "计算机导论", Author: Author{Name: "谁写的", Age: 10}}
	result, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result)) // {"Name":"计算机导论","Author":{"Age":10,"Name":"谁写的"}}
	// 反序列化
	var b1 Book
	err = json.Unmarshal(result, &b1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b1) // {计算机导论 {10 谁写的}}

	// 忽略空字段
	b2 := Book2{Author: Author2{Name: "zhangsan", Age: 1}}
	j, _ := json.Marshal(b2)
	fmt.Println(string(j)) //{"Author":{"Age":1,"authorName":"zhangsan"}}

	// 跳过字段
	b3 := Book2{Name: "go", Author: Author2{Name: "itlab", Age: 1}}
	k, _ := json.Marshal(b3)
	fmt.Println(string(k)) //{"name":"go"}
}

type Book struct {
	Name   string
	Author Author
}
type Author struct {
	Age  int
	Name string
}

type Book2 struct {
	Name   string  `json:"name,omitempty"`
	Author Author2 `json:"-"`
}
type Author2 struct {
	Age  int
	Name string `json:"authorName"`
}
