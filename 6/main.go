package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 未指定类型 默认是float64
	var f0 = 3.33
	var f1 float32 = 1.2
	var f2 float64 = 3.4
	fmt.Printf("不指定类型，go编译器自动推断小数类型为%T\n", f0)
	fmt.Printf("f1的类型是%T,值=%v, %d字节\n", f1, f1, unsafe.Sizeof(f1))
	fmt.Printf("f2的类型是%T,值=%v, %d字节\n", f2, f2, unsafe.Sizeof(f2))

	//精度丢失
	var f32 float32
	var f64 float64

	f64 = 1.234567890123
	f32 = float32(f64)

	fmt.Println(f32, f64)
}
