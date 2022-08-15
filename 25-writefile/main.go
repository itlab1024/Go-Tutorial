package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

func main() {
	// 创建文件
	file, err := os.Create("data/write.txt")
	if err != nil {
		return
	}
	defer file.Close()
	// 向文件写入内容
	ioutil.WriteFile("data/write.txt", []byte("哈哈哈"), fs.ModeAppend)

	// 传统写入
	file.WriteString("传统写入")

	// 追加内容
	f, err := os.OpenFile("data/write.txt", os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "保留原有文件内容，追加新内容"
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
	}
}
