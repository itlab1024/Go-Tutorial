package main

import (
	htmlTemplate "html/template"
	"os"
	"text/template"
)

func main() {
	// 文本模板
	// 预定义模板{.name}正在学习golang
	user := User{Name: "乔布斯"}
	temp := "{{.Name}}正在学习golang"
	t, err := template.New("message").Parse(temp)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, user) //乔布斯正在学习

	if err != nil {
		panic(err)
	}

	// 读取html模板
	ht, _ := htmlTemplate.ParseFiles("15-template/templates/html-template.html")
	_ = ht.Execute(os.Stdout, user)

	//模板验证
	ht2 := htmlTemplate.Must(htmlTemplate.ParseFiles("15-template/templates/html-template.html"))
	_ = ht2.Execute(os.Stdout, user)
}

type User struct {
	Name string
}
