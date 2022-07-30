package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 复数
	var n1 complex64 = 1.2
	var n2 complex128 = 2.3
	n3 := complex(5, 7)
	fmt.Printf("n1的类型是%T,值=%v, %d字节\n", n1, n1, unsafe.Sizeof(n1))
	fmt.Printf("n2的类型是%T,值=%v, %d字节\n", n2, n2, unsafe.Sizeof(n2))
	fmt.Printf("n3的类型是%T,值=%v, %d字节\n", n3, n3, unsafe.Sizeof(n3))

	n4 := 1 + 2i
	// 获取实部和虚部
	f := real(n4)
	f2 := imag(n4)
	println(f)  // +1.000000e+000
	println(f2) // +2.000000e+000

	// 复数运算
	c1 := complex(2, 3)
	c2 := 4 + 5i
	c3 := c1 + c2
	fmt.Println("加法: ", c3)
	re := real(c3)
	im := imag(c3)
	fmt.Println(re, im) // 6 8
}
