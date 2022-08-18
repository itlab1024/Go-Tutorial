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
