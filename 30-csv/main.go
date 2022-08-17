package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data/read.csv")
	if err != nil {
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	fmt.Println(records) // 返回的是二维数组切片。[[姓名 职位] [诸葛亮 军师] [刘备 左将军] [曹操 丞相]]
	// 写文件
	file, err = os.Create("data/write.csv")
	defer file.Close()
	if err != nil {
		return
	}
	writer := csv.NewWriter(file)
	// 准备数据
	data := [][]string{
		{"姓名", "职位"},
		{"庞统", "军师"},
	}
	writer.WriteAll(data)
}
