package main

import (
	"errors"
	"fmt"
)

func main() {
	i := 1
	j := 0
	i2, err := division(i, j)
	if err != nil {
		return
	}
	fmt.Println(i2)
}

func division(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("不能为0")
	}
	return i / j, nil
}
