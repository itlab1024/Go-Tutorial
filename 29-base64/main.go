package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 加密
	em := base64.StdEncoding.EncodeToString([]byte("go"))
	fmt.Println(em) // Z28=
	// 解密
	dm, err := base64.StdEncoding.DecodeString(em)
	if err != nil {
		return
	}
	fmt.Println(string(dm)) // go

}
