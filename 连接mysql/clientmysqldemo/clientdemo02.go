package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //定义一个全局db对象

type employee struct {
	id         int
	name       string
	occupation string
	age        int
}

func initDB() (err error) {
	//链接
	dsn := "root:123456@tcp(192.168.1.94:3306)/demo?charset=utf8mb4&parseTime=True"

	//
	db, err = sql.Open("mysql", dsn) //Open函数可能只是验证其参数格式是否正确，实际上并不创建与数据库的连接。
	if err != nil {
		return err
	}

	//尝试与数据库建立连接 (校验dsn 链接是否正确)
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

//插入数据：func (db *DB) Exec(query string, args ...interface{}) (Result, error)

func insertEmployee() {
	sqlStr := "insert into employee_table(name, occupation, age) values (?, ?, ?)"
	ret, err := db.Exec(sqlStr, "啊而今", "go开发", "22")
	if err != nil {
		fmt.Printf("insert failed err : %v\n", err)
		return
	}
	theID, err := ret.LastInsertId() //新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err : %v\n", err)
		return
	}
	fmt.Printf("insert success, this id is %d.\n", theID)
}

//单行查询:单行查询db.QueryRow()执行一次查询，并期望返回最多一行结果（即Row）。QueryRow总是返回非nil的值，
//直到返回值的Scan方法被调用时，才会返回被延迟的错误。（如：未找到结果）
//func (db *DB) QueryRow(query string, args ...interface{}) *Row

func queryRow() {
	sqlStr := "select id, name, occupation, age from employee_table where id =?" //根据id查信息
	var u employee
	err := db.QueryRow(sqlStr, 2).Scan(&u.id, &u.name, &u.occupation, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err : %v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s occupation:%s age:%d\n", u.id, u.name, u.occupation, u.age)
}

// 多行查询
func queryMultiRow() {
	sqlStr := "select id, name, occupation, age from employee_table where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("scan multi failed err : %v\n", err)
	}

	// 关闭rows，释放持有的数据库连接
	defer rows.Close()

	//循环读取结果集中的数据
	for rows.Next() {
		var u employee
		err := rows.Scan(&u.id, &u.name, &u.occupation, &u.age) //scan函数可以将查询的数据存储到指定数据类型中
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
			return
		}
		fmt.Printf("id: %d name: %s occupation: %s age: %d\n", u.id, u.name, u.occupation, u.age)
	}
}

func updateRow() {
	sqlStr := "update employee_table set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 40, 3)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected() //该操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d\n", n)
}

// 删除数据
func deleteRow() {
	sqlStr := "delete from employee_table where id = ?"
	ret, err := db.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("delete failed err : %v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作影响行数
	if err != nil {
		fmt.Printf("get rowsaffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows: %d\n", n)
}

// 预处理查询
func prepareQuery() {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db.Prepare(sqlStr) // Prepare() 方法返回一个 stmt 结构体对象
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0) // Query方法返回rows 对象
	if err != nil {
		fmt.Printf("query failed, err: %v \n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u employee
		err := rows.Scan(&u.id, &u.name, &u.occupation, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.occupation, u.age)
	}
}

// 预处理插入、更新和删除操作的预处理十分类似，这里以插入操作的预处理为例：
func prepareInsert() {
	sqlStr := "insert into employee_table(name, occupation, age) values (?, ?, ?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("sql prepare failed err : %v \n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("曹阿蛮", "丞相", "57")
	if err != nil {
		fmt.Printf(" insert failed, err: %v \n", err)
		return
	}
	_, err = stmt.Exec("王阳明", "GO开发工程师", "45")
	if err != nil {
		fmt.Printf("insert failed, err : %v \n", err)
		return
	}
	fmt.Println("insert success.")
}

//sql 注入问题
//我们任何时候都不应该自己拼接SQL语句！

// GO实现mysql事务
func transactionDemo() {
	tx, err := db.Begin() //开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() //回滚
		}
		fmt.Printf("begin trans failed, err: %v\n", err)
		return
	}

	sqlStr1 := "Update employee_table set age=30 where id =?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() //回滚
		fmt.Printf("exec sql1 failed, err: %v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err : %v\n", err)
		return
	}

	sqlStr2 := "Update employee_table set age=40 where id=?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback() //回滚
		fmt.Printf("exec sql2 failed, err: %v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err: %v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交。。。")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚了")
	}

	fmt.Println("exec trans success")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db err: %v\n", err)
	}
	//insertEmployee() //插入数据
	//queryRow()
	//queryMultiRow()
	//updateRow()
	deleteRow()
}
