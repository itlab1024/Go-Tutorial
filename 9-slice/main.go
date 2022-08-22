package main

import (
	"fmt"
)

func main() {
	// 定义切片
	var slice1 []int
	//slice1[0] = 1                            // 编译通过，但是没有初始化所以执行报错，触发panic
	fmt.Println("切片的零值是nil?", slice1 == nil) //切片的零值是nil? true

	// 这根数组定义有些类似，千万不要搞混
	// 比如如下是数组定义
	// var array1 [1]int

	// 切片的初始化
	// 第一种
	slice1 = []int{1, 2, 3}
	fmt.Println("切片slice1的长度是", len(slice1), "容量是", cap(slice1)) //切片slice1的长度是 3 容量是 3

	// 第二种使用new, new通常用来创建结构体指针，所以以下代码虽然没有问题，但是却不常用，常使用make函数来创建切片
	var s1 = new([]int)
	*s1 = append(*s1, 1, 2, 3)
	fmt.Println(*s1) // [1 2 3]

	// 第三种 使用make函数初始化切片
	// 指定类型和长度
	var s2 = make([]int, 3)
	fmt.Println("s2的长度是", len(s2), "容量是", cap(s2)) //s2的长度是 3 容量是 3

	// 指定类型，长度=0，容量=3
	s3 := make([]int, 0, 3)
	fmt.Println("s3的长度是", len(s3), "容量是", cap(s3)) //s3的长度是 0 容量是 3

	//第四种：通过重新切片初始化切片
	s4 := make([]int, 2, 8)
	fmt.Println("s4的长度是", len(s4), "容量是", cap(s4)) //s4的长度是 2 容量是 8
	// s4已经是切片了，还可以对其进行再次切片
	//s5 := s4[2:9] // 9是不允许的，因为s4的cap=8，也就是说再次切片的最大容量不能超过8
	s5 := s4[2:8]
	fmt.Println("s5的长度是", len(s5), "容量是", cap(s5)) //s5的长度是 6 容量是 6

	// 第五种使用数组初始化切片
	var array [3]int = [3]int{1, 2, 3}
	// 通过array数组初始化切片
	var sl1 = array[1:2]
	for _, v := range sl1 {
		fmt.Println(v) // 结果是3，数组索引从1，到2（不包含2，左闭右开）结果就是2
	}
	//针对这种[strartIndex:endIndex]切片方式，两个index是可以省略的，strartIndex的默认值就是0，endIndex默认值就是最大索引，如果省略就是使用默认值。
	sl2 := array[:]
	sl3 := array[0:]
	sl4 := array[:2]
	fmt.Println(sl2, sl3, sl4) //[1 2 3] [1 2 3] [1 2]

	// 追加元素
	var ap = []int{1, 2, 3}
	after := append(ap, 4)
	fmt.Println(ap)    //[1 2 3]
	fmt.Println(after) // [1 2 3 4]

	// 循环
	for index, value := range after {
		fmt.Print(index, "=", value, ",") // 0=1,1=2,2=3,3=4,
	}
	fmt.Println()
	//一个接收默认是index
	for index := range after {
		fmt.Print(index, ",") // 0,1,2,3,
	}
	fmt.Println()
	// 可以使用_忽略某个值
	for _, value := range after {
		fmt.Print(value, ",") //1,2,3,4,
	}
	fmt.Println("------")

	//拷贝
	cp1 := []int{1, 2, 3, 4, 5}
	cp2 := []int{5, 4, 3}
	copy(cp2, cp1)   // 只会复制cp1的前3个元素到cp2中,cp1结果还是
	fmt.Println(cp1) //[1 2 3 4 5]
	fmt.Println(cp2) //[1 2 3]
	fmt.Println("------")

	copy(cp1, cp2)   // 只会复制cp2的3个元素到cp1的前3个位置
	fmt.Println(cp1) //[1 2 3 4 5]
	fmt.Println(cp2) //[1 2 3]
}
