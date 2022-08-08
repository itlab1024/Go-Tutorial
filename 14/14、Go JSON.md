![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202207281322806.png)

# 什么是JSON？

JSON 是**JavaScript Object Notation**的缩写，一种广泛使用的数据交换格式。

# 默认包

Go中的json处理使用的是`encoding/json`包。

序列化使用`json.Marshal`，反序列化使用`json.UnMarshal`

# 基本使用

定义一个结构体，序列化为json，再反序列化

```go
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
}

type Book struct {
	Name   string
	Author Author
}
type Author struct {
	Age  int
	Name string
}
```

# 自定义属性值

比如上线的作者名字我想序列化的时候名字更改为authorName，则需要修改结构体

```go
type Author struct {
	Age  int
	Name string `json:"authorName", omitempty`
}
```

序列化的结果变为如下内容

```tex
{"Name":"计算机导论","Author":{"Age":10,"authorName":"谁写的"}}
```

# 忽略空字段

比如书籍如果没有设置作者信息，则忽略该字段，需要使用` omitempty`标志

```go
type Book2 struct {
	Name   string `json:"name,omitempty"`
	Author Author2
}
type Author2 struct {
	Age  int
	Name string `json:"authorName"`
}


// 忽略空字段
b2 := Book2{Author: Author2{Name: "zhangsan", Age: 1}}
j, _ := json.Marshal(b2)
fmt.Println(string(j)) //{"Author":{"Age":1,"authorName":"zhangsan"}}
```

# 跳过字段

跳过字段使用`-`

```go
type Book2 struct {
	Name   string  `json:"name,omitempty"`
	Author Author2 `json:"-"`
}
type Author2 struct {
	Age  int
	Name string `json:"authorName"`
}
// 跳过字段
b3 := Book2{Name: "go", Author: Author2{Name: "itlab", Age: 1}}
k, _ := json.Marshal(b3)
fmt.Println(string(k)) //{"name":"go"}
```

作者字段不为空，但是被跳过了。