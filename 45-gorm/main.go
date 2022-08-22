package main

import (
	"fmt"
	"gorm.io/gorm"
)
import "gorm.io/driver/mysql"

type Data struct {
	gorm.Model
	Name string
}

func main() {
	// 建立连接
	db, err := gorm.Open(mysql.Open("root:qwe!@#123@tcp(127.0.0.1:3306)/go-tutorial?charset=utf8mb4&charset=utf8&parseTime=true"), &gorm.Config{})
	if err != nil {
		return
	}
	// 迁移 schema，如果数据库中没有这个schema会自动创建
	db.AutoMigrate(&Data{})

	// C
	//g := Data{Name: "grom"}
	//db.Create(&g)
	// R
	var f Data
	db.First(&f, 1)
	fmt.Println("查询结果是", f)
	// 可以通过条件查询
	db.First(&f, "name = ?", "gorm")
	fmt.Println("通过条件查询结果是", f)

	// U
	db.Model(&f).UpdateColumn("name", "grom2")
	// 更新多个字段
	db.Model(&f).UpdateColumns(Data{Name: "刘备"})
	// 更新多个字段(Map)
	db.Model(&f).UpdateColumns(map[string]interface{}{"id": 2, "name": "grom3"})

	//D
	db.Delete(&f)
}
