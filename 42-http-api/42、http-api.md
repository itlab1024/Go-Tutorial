# 说明
我会创建一个结构体书籍信息。然后提供一个接口来查询书籍信息。
# 服务创建以及接口定义
源码如下：
```go
package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/books/getBookById", getBookById)
	http.ListenAndServe("", nil)
}

func getBookById(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	writer.Header().Add("content-type", "application/json;charset=UTF-8")
	if id == "" {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		writer.WriteHeader(http.StatusOK)
		book := Book{Name: "学习Go", Author: "诸葛亮", PubDate: time.Now()}
		json.NewEncoder(writer).Encode(book)
	}
}

type Book struct {
	Name    string
	Author  string
	PubDate time.Time
}
```
使用浏览器访问：http://localhost/books/getBookById?id=1
![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208181335138.png)
测试接口我跟还是喜欢实用POSTMAN，ApiFox、httpie等工具
```shell
➜  ~ http http://localhost/books/getBookById\?id\=1
HTTP/1.1 200 OK
Content-Length: 86
Content-Type: application/json;charset=UTF-8
Date: Thu, 18 Aug 2022 05:39:09 GMT

{
    "Author": "诸葛亮",
    "Name": "学习Go",
    "PubDate": "2022-08-18T13:39:09.334689+08:00"
}
```