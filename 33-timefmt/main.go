package main

import (
	"fmt"
	"time"
)

func main() {
	// 默认时间
	now := time.Now()
	fmt.Println(now) // 2022-08-17 18:09:06.332117 +0800 CST m=+0.000641201
	// 格式化
	format := now.Format("3:04PM") // 6:09PM
	fmt.Println(format)
	fmt.Println(now.Format("2006年01月02日15时01分05秒")) // 2022年08月17日18时08分06秒
	// 字符串转化为时间
	t, _ := time.Parse("3:04PM", "12:04AM")
	fmt.Println(t) // 0000-01-01 00:04:00 +0000 UTC
}
