# HTTP服务
Go原生支持http服务，使用net/http包来实现。
# 基本使用
http包下有ListenAndServe函数来创建http服务，指定端口。
通过HandleFunc等方法来具体定义api端点以及具体业务内容
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 先定义一个handler,这里使用/定义的是主页
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "http服务主页")
	})
	// 创建服务并监听端口
	http.ListenAndServe(":80", nil)
}
```
这就创建了一个基本的http服务，端口是80。
运行http://localhost截图
![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208181109835.png)

# Addr设置
上面的例子中http.ListenAndServe(":80", nil)第一个参数是设置host和端口。
格式是host:port,也可以传递空字符串，这就会使用默认80端口。
可以自定义
```go
http.ListenAndServe("goserver:8080", nil)
```
这个goserver是我在hosts中自定义的一个域名。
运行访问http://goserver:8080/
![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208181111189.png)
# 多路复用
http.ListenAndServe("goserver:8080", nil)这段代码第二个参数是nil。他其实就是多路复用参数。
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 多路复用
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/welcome", welcome)
	http.ListenAndServe(":80", mux)
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "http服务主页")
}

func welcome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "http服务欢迎页")
}
```
主页：
![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208181122555.png)
welcome页
![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208181122173.png)