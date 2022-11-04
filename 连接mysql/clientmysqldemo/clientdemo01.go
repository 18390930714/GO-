package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@tcp(192.168.1.194:3306)/demo"
	db, err := sql.Open("mysql", dsn) //Open函数可能只是验证其参数格式是否正确，实际上并不创建与数据库的连接。
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
