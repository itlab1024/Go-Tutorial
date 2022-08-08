![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202207281322806.png)

# 模板类型

Go中的模板主要分为文本模板和HTML模板，分别位于`text/template`和`html/template`。

# 替换规则

文本替换主要使用{{}}来替换。

比如{{.}}代表根

{{.Name}}代表Name属性。

{{/*注释*/}}

条件控制语句：{{ if .condition }} {{ else }} {{ end }}

range：{{ range .Items }} {{ end }

# 文本模板

请看如下代码

```go
package main

import (
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
}

type User struct {
	Name string
}
```



# 模板读取

使用`html/template读取`

```go
import (
	htmlTemplate "html/template"
	"os"
	"text/template"
)
// 读取html模板
ht, _ := htmlTemplate.ParseFiles("15-template/templates/html-template.html")
_ = ht.Execute(os.Stdout, user)
```

结果是：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>HTML模板</title>
</head>
<body>
    乔布斯
</body>
</html>
```

# 模板验证

使用`Must方法`验证

```go
//模板验证
ht2 := htmlTemplate.Must(htmlTemplate.ParseFiles("15-template/templates/html-template.html"))
_ = ht2.Execute(os.Stdout, user);
```

