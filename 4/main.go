package main

import "fmt"

func main() {
	sayHello()
	// 测试函数参数传递是否是值传递
	i := 1
	fmt.Println(&i)  // 0xc000069f00
	prepareChange(i) // 设置i=100
	fmt.Println(i)   // 结果还是1

	j := 1
	prepareChange2(&j)
	fmt.Println(j) // j=100

	// 空标识符,比如下面是正常接收
	//r1, r2 := addSub(1, 2)
	// 如果我不关心r2的结果则可使用空标识符
	r1, _ := addSub(1, 2)
	fmt.Println(r1)

	// 匿名函数
	func1 := func() {
		fmt.Println("匿名函数体")
	}
	func1()

	// 匿名函数立即调用
	func() {
		fmt.Println("匿名函数体")
	}()

	// 调用调用返回类型是函数的函数
	f := returnFunc()
	f(1)

	// 函数作为参数
	biz(i, callback)

	// 闭包
	f2 := cumulative()
	i2 := f2(1)
	i3 := f2(1)
	i4 := f2(1)
	fmt.Println(i2, i3, i4) // 1 2 3
	f2 = cumulative()
	i5 := f2(1)
	i6 := f2(1)
	i7 := f2(1)
	fmt.Println(i5, i6, i7) // 1 2 3
}

// 定义了一个无参数，返回空类型的函数
func sayHello() {
	fmt.Println("Hello go")
}

// 函数的参数(标准)
func param1(i, j int) {
	fmt.Println(i, j)
}

// 函数返回值
func add(x, y int) int {
	return x + y
}

// 多返回值
func addSub(x, y int) (int, int) {
	return x + y, x - y
}

// 命名返回值
func namedReturn(x, y int) (a int, b int) {
	a = x + y
	b = x - y
	return
}

// 函数的值传递
func prepareChange(i int) {
	fmt.Println(&i) // 0xc000069ee0
	i = 100
}

// 函数的引用传递
func prepareChange2(i *int) {
	*i = 100
}

// 函数作为返回值
func returnFunc() func(i int) int {
	return func(i int) int {
		return i + 1
	}
}

func biz(i int, c func(i int)) {
	c(i)
}

func callback(i int) {
	fmt.Println("回调成功", i)
}

// 闭包函数，累加功能
func cumulative() func(i int) int {
	var result int
	return func(i int) int {
		result += i
		return result
	}
}
