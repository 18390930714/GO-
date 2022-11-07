package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 更强大的sqlx
//安装sqlx ：go get github.com/jmoiron/sqlx

var db *sqlx.DB

type employee struct {
	Id         int
	Name       string
	Occupation string
	Age        int
}

func initDB() (err error) {
	dsn := "root:123456@tcp(192.168.1.94:3306)/demo?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func queryRow() {
	sqlStr := "select id, name, occupation, age from employee_table where id=?"
	var u employee
	err := db.Get(&u, sqlStr, 2) //db.Get 需要访问 employee结构体 u 中的各个字段，所以结构体各个字段首字母要大写
	if err != nil {
		fmt.Printf("get failed, err: %v \n", err)
		return
	}
	fmt.Printf("id: %d name: %s occupation:%s age: %d\n", u.Id, u.Name, u.Occupation, u.Age)
}

func queryMultiRow() {
	sqlStr := "select id, name, occupation, age from employee_table where id > ?"
	var ps []employee
	err := db.Select(&ps, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
		return
	}
	fmt.Printf("people: %#v\n", ps)
}

// 插入
func insertRow() {
	sqlStr := "insert into employee_table(name, occupation, age) values (?, ?, ?)"
	ret, err := db.Exec(sqlStr, "服部半藏", "GO开发", "109")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() //新插入数据的ID
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d \n", theID)
}

// 更新
func updateRow() {
	sqlStr := "update employee_table set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 34, 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update suceess, affected rows:%d\n", n)
}

// 删除
func deleteRow() {
	sqlStr := "delete from employee_table where id = ?"
	ret, err := db.Exec(sqlStr, 4)
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows: %d\n", n)
}

// DB.NameExec 方法用来绑定SQL语句与结构体或者map中的同名字段.
func insertEmployee() (err error) {
	sqlStr := "INSERT INTO employee_table (name, occupation, age) VALUES (:name, :occupation, :age)"
	_, err = db.NamedExec(sqlStr,
		map[string]interface{}{
			"name":       "伍中豪",
			"occupation": "GO高级开发",
			"age":        27,
		})
	return
}

// NamedQuery，与DB.NamedExec同理，这里是支持查询。
func nameQuery() {
	sqlStr := "SELECT * FROM employee_table WHERE name=:name"
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name", "伍中豪"})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u employee
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err :%v\n", err)
			continue
		}
		fmt.Printf("user: %#v\n", u)
	}

	u := employee{
		Name: "伍中豪",
	}
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NameQuery failed, err: %v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u employee
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err : %v\n", err)
	}
	//queryRow()
	//queryMultiRow()
	//insertRow()
	//updateRow()
	deleteRow()
}

//
