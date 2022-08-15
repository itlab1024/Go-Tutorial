package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// 读取文件
	file, err := os.Open("data/go.txt")
	if err != nil {
		return
	}
	// 使用defer关闭文件
	defer file.Close()
	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Println(fileInfo)                // &{go.txt 58 420 {678331806 63796076733 0x114b580} {16777220 33188 1 3566130 501 20 0 [0 0 0 0] {1660479934 31636101} {1660479933 678331806} {1660479933 678331806} {1660479912 307934053} 58 8 4096 0 0 0 [0 0]}}
	fmt.Println("文件名", fileInfo.Name())  // 文件名 go.txt
	fmt.Println("模式", fileInfo.Mode())   //模式 -rw-r--r--
	fmt.Println("文件大小", fileInfo.Size()) // 文件大小 58(单位字节)
	// 使用ioutil读取文件
	bytes, err := ioutil.ReadFile("data/go.txt")
	if err != nil {
		return
	}
	fmt.Println(string(bytes)) // 一行一行的打印出文件内的所有内容

	// 分块读取
	file2, err := os.Open("data/go.txt")
	defer file2.Close()
	if err != nil {
		return
	}
	const MaxSize = 4
	b := make([]byte, MaxSize)
	for {
		readTotal, err := file2.Read(b)
		if err != nil {
			if err != io.EOF {
				log.Fatalln(err)
			}
			break
		}
		fmt.Println(string(b[:readTotal]))
	}

	// 使用bufio进行逐行读取
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
