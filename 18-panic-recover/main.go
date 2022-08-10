package main

import "fmt"

func main() {
	// 会触发panic
	startWeb(nil)
}
func startWeb(dbUrl *string) {
	//defer fmt.Println("执行defer语句")
	//if dbUrl == nil {
	//	panic("数据库地址是nil，无法启动web应用")
	//}
	//fmt.Println("数据库地址是", *dbUrl, " 程序正常启动")

	// recover
	fmt.Println("测试开始")
	fmt.Println(testPanic())
	fmt.Println("测试结束")
}

func testPanic() (ret int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			ret = -1
		}
	}()
	var i = 10
	var j = 0
	var t = i / j
	fmt.Println(t)
	return t
}
