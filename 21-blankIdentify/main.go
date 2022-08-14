package main

import "fmt"
import _ "sync"

func main() {
	value, _ := multiReturnFunc()
	fmt.Println(value)
	_ = value + 10
}
func multiReturnFunc() (int, error) {
	return 1, nil
}
