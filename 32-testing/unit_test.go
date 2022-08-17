package testing

import "testing"

// 定义一个普通函数
func WrongMin(x, y float64) float64 {
	if x > y {
		return x
	} else {
		return y
	}
}

// 定义一个测试函数，用于测试WrongMin函数
func TestMathBasics(t *testing.T) {
	v := WrongMin(10, 0)
	if v != 0 {
		t.Error("测试失败")
	}
}
