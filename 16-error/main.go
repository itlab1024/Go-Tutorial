package main

import (
	"fmt"
	"log"
)

func main() {
	i := 1
	j := 1
	i2, err := division(i, j)
	if err != nil {
		log.Fatal(err) // 不能为0
		return
	}
	fmt.Println(i2)

	// 自定义错误的使用
	name := ""
	r, err := NameCheck(name)
	if err != nil {
		log.Fatal(err) // Name不能为
		return
	}
	fmt.Println(r)
}

func division(i int, j int) (int, error) {
	if j == 0 {
		//return 0, errors.New("不能为0")
		return 0, fmt.Errorf("不能为%d", j)
	}
	return i / j, nil
}

// NameEmptyError 自定义错误
type NameEmptyError struct {
	Name string
}

func (e *NameEmptyError) Error() string {
	return "Name不能为空"
}

func NameCheck(name string) (bool, error) {
	if name == "" {
		return false, &NameEmptyError{name}
	}
	return true, nil
}
