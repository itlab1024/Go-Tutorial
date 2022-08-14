package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// int8转化为int64
	var i int8 = 3
	j := int64(i)
	fmt.Println(j)
	// strconv  字符串转化为int
	s := "1"
	i1, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(i1) // 1
	// 将int转化为字符串
	s2 := strconv.Itoa(12)
	fmt.Println(s2) // 12
	// int和float转化
	f := 12.34567
	i2 := int(f)    // 丢失精度
	fmt.Println(i2) // 12

	ii := 34
	f1 := float64(ii)
	fmt.Println(f1) // 34

	//字符串和字节转换
	s3 := "我在学go"
	b1 := []byte(s3)
	fmt.Println(b1) //[230 136 145 229 156 168 229 173 166 103 111]
	// 字节转化为字符串使用string()函数
	fmt.Println(string(b1)) //我在学go

	// 除法操作的隐式转换
	f2 := 6.3 / 3   // float类型除以int类型，结果是float类型的f2
	fmt.Println(f2) // 2.1
}
