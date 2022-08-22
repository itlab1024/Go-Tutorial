# Gorm
[Gorm](https://gorm.io/zh_CN/)是Go语言的一套ORM框架，类似于Java中的JPA规范的实现（比如Hibernate，EclipseLink等）
# 安装
使用go get安装
```shell
➜  45-gorm git:(main) ✗ go get -u gorm.io/gorm
go: downloading gorm.io/gorm v1.23.8
go: downloading github.com/jinzhu/inflection v1.0.0
go: downloading github.com/jinzhu/now v1.1.4
go: downloading github.com/jinzhu/now v1.1.5
go: added github.com/jinzhu/inflection v1.0.0
go: added github.com/jinzhu/now v1.1.5
go: added gorm.io/gorm v1.23.8
```
还需要引入gorm MySQL的依赖
```shell
➜  45-gorm git:(main) ✗ go get gorm.io/driver/mysql
go: downloading gorm.io/driver/mysql v1.3.6
go: added github.com/go-sql-driver/mysql v1.6.0
go: added gorm.io/driver/mysql v1.3.6
```
# 使用示例
首先我要定义一个Gorm的结构体，该结构体就是数据库Gorm的映射，然后使用gorm进行CRUD操作。
```go
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
}

```
数据库中会创建名字叫做data的表。
![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208221408304.png)
图中绿色的是gorm框架自动增加的列。
# CRUD
上面已经创建出来表，接下来实现crud。
```go
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
```
运行中截图：
![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208221409150.png)
上面绿色的deleted_at也作为了查询条件。但是实际我只想通过ID更新，与我预期不符合。这应该就需要自定义了，我这里就不一一尝试了，
官网有详细说明： https://gorm.io/zh_CN/docs/models.html
