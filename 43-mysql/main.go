package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

type Book struct {
	ID     int
	Name   string
	Author string
}

func main() {
	// 创建连接属性
	cfg := mysql.Config{
		User:   "root",
		Passwd: "qwe!@#123",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "go-tutorial",
	}
	// 使用上面的配置创建连接
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalln("数据库连接异常,请检查配置")
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln("数据库Ping失败,请检查配置")
	}
	defer db.Close()
	fmt.Println("恭喜您成功连接数据库")
	// 创建表
	sql := `CREATE TABLE IF NOT EXISTS BOOK
(
    ID     INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    NAME   VARCHAR(200) NULL COMMENT '书籍名称',
    AUTHOR VARCHAR(200) NULL COMMENT '作者名称'

) COMMENT '主键ID';
`
	r, err := db.Exec(sql)
	if err != nil {
		return
	}
	fmt.Println("Book表创建成功")
	// 新增数据
	insertSQL := `insert into book (name, author) values ("软件开发基础", "张飞")`
	r, err = db.Exec(insertSQL)
	if err != nil {
		return
	}
	id, _ := r.LastInsertId()
	fmt.Println("成功插入数据，主键ID是", id)
	// 查询
	stmt, err := db.Prepare("SELECT * FROM BOOK WHERE ID = ?")
	if err != nil {
		return
	}
	rows, _ := stmt.Query(id)
	defer rows.Close()
	for rows.Next() {
		var book = Book{}
		err = rows.Scan(&book.ID, &book.Name, &book.Author)
		if err != nil {
			return
		}
		fmt.Println(book)
	}
}
