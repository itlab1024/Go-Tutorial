package main

import (
	"flag"
	"fmt"
)

func main() {
	// 通过os获取命令行参数
	//args := os.Args
	//s0 := args[0]
	//s1 := args[1]
	//fmt.Println("s0=", s0, "s1=", s1)

	//使用flag编辑参数
	var s0 string
	var s1 string
	flag.StringVar(&s0, "s0", "s0默认值", "s0的描述信息")
	flag.StringVar(&s1, "s1", "s1默认值", "s1的描述信息")
	flag.Parse() // 转化
	fmt.Println("s0=", s0, "s1=", s1)
}
