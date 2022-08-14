package main

import "fmt"

func main() {
	// 简单使用
	i := 1
	switch i {
	case 1:
		fmt.Println("结果是1") // 结果是1
	default:
		fmt.Println("结果不是1")
	}
	// 简单实用二，多条件
	switch i {
	case 1, 2:
		fmt.Println("结果是1或者2") // 结果是1或者2
	default:
		fmt.Println("结果不是1也不是2")
	}

	// break
	switch i {
	case 1, 2:
		fmt.Println("结果是1或者2") // 结果是1或者2
		break
	default:
		fmt.Println("结果不是1也不是2")
		break
	}

	// fallthrough
	switch i {
	case 1, 2:
		fmt.Println("结果是1或者2") // 结果是1或者2
		fallthrough
	default:
		fmt.Println("结果不是1也不是2") // 结果不是1也不是2
	}
	//变量初始化
	switch v := 1; v {
	case 1:
		fmt.Println("值等于1")
	default:
	}

	// 条件case语句
	switch {
	case i > 0:
		fmt.Println("i > 0") // i > 0
	case i <= 0:
		fmt.Println("i <= 0")
	}
}
