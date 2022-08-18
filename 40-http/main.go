package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 先定义一个handler,这里使用/定义的是主页
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer, "http服务主页")
	//})
	// 创建服务并监听端口
	//http.ListenAndServe("", nil)
	//http.ListenAndServe("goserver:8080", nil)

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
