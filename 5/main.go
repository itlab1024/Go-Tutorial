package main

func main() {
	// 大范围赋值给小范围类型。
	//var i int8 = 128 // 编译错误 ，因为int8的范围是（-127-127）
	var j int8 = 127
	var k int8 = j + 1
	println(k) // 结果溢出，并不是128，而是-128

	// 类型转换
	var l int32 = int32(k)
	println(l)
}
