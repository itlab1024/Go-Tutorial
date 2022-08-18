# HTML模板
html模板的使用之前已经讲过，这里不做讲解。
# 示例
请看如下示例：学生列表，每个学生下面有书籍列表。实现功能，在页面上展示出来。
```go
package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("", nil)
}

// Student 学生
type Student struct {
	Name  string
	Books []Book
}

// Book 书籍
type Book struct {
	Name string
}

func index(writer http.ResponseWriter, request *http.Request) {
	// 组装数据
	books := []Book{{Name: "计算机基础"}, {Name: "Go教程"}}
	students := make([]Student, 0)
	for i := 0; i < 10; i++ {
		student := Student{Name: "学生" + strconv.Itoa(i), Books: books}
		students = append(students, student)
	}
	t, _ := template.ParseFiles("41-http-html/index.html")
	t.Execute(writer, students)
}
```
定义一个html文件。
```go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>学生信息列表</title>
</head>
<body>
<ul>
    {{ range . }}
    <li>{{ .Name }}:
        <ul>
            {{ range .Books}}
            <li>{{.Name}}</li>
            {{ end}}
        </ul>
    </li>
    {{ end}}
</ul>
</body>
</html>
```
这里我使用range循环渲染。
运行结果：
![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208181307201.png)
# 使用场景
其实目前一般这种前后端融合在一起情况不多了，小型项目还是可以使用这种方式的。
不过当今，微服务思想非常流行。前后端分离更加常用。
后端通过API接口给前端使用是更常用的方式。