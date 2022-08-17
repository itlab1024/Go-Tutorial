package main

import (
	"fmt"
	"sort"
)

func main() {
	// int排序
	intSlice := []int{1, 4, 6, 3, 0, 5}
	sort.Ints(intSlice)
	fmt.Println(intSlice) // [0 1 3 4 5 6]
	// float64排序
	float64s := []float64{1, 2, 4.3, 5.6, 1.8}
	sort.Float64s(float64s)
	fmt.Println(float64s) //[1 1.8 2 4.3 5.6]
	// string 排序
	strings := []string{"a", "c", "b"}
	sort.Strings(strings)
	fmt.Println(strings) // [a b c]

	// 检查切片是否有效
	sorted := sort.IntsAreSorted(intSlice)
	fmt.Println("是否是有序的", sorted) // true

	// 切片反转
	i := []int{1, 4, 6, 3, 0, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(i)))
	fmt.Println(i) // [6 5 4 3 1 0]

	order1 := Order{Name: "c"}
	order2 := Order{Name: "b"}
	orders := []Order{order1, order2}
	sort.Sort(Orders(orders))
	fmt.Println(orders) //[{b} {c}]
}

type Order struct {
	Name string
}
type Orders []Order

func (orders Orders) Len() int {
	return len(orders)
}

func (orders Orders) Less(i, j int) bool {
	return orders[i].Name < orders[j].Name
}

func (orders Orders) Swap(i, j int) {
	orders[i], orders[j] = orders[j], orders[i]
}
