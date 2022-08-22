package main

import "fmt"

func main() {
	// 定义数组
	var a1 [2]int
	// 不初始化打印输出
	fmt.Println(a1) // [0 0]

	// 初始化
	a1 = [2]int{1, 2}
	fmt.Println(a1) // [1 2]

	// 获取数组元素（通过索引下标）
	fmt.Println(a1[0]) // 1

	//通过索引更新
	a1[0] = 100
	fmt.Println(a1[0]) // 100

	// 初始化还有个特别的地方，数组具体值数目可以不等于定义的数组长度，那么未设置的地方就会以零填充
	var a2 [3]int = [3]int{1}
	fmt.Println(a2) //[1 0 0]

	// 迭代数组
	for i := range a1 {
		fmt.Println(i) // 0 1
	}

	// 数组的长度和容量分别通过内置方法len()和cap()获取
	fmt.Println("a1的长度是", len(a1), ", 容量是", cap(a1)) // a1的长度是 2 , 容量是 2

	// 二维数组
	arr := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(arr) // [[1 2] [3 4]]

	// 多维数组
	arr2 := [2][2][2]int{
		{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}},
	}
	fmt.Println(arr2) // [[[1 2] [3 4]] [[5 6] [7 8]]]

	// 数组长度不可变
	//a3 := [1]int{1}
	//a3[1] = 1 // 报错，超出边界
}
