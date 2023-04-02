package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 空导入，注册mysql驱动对象
)

// 定义一个全局对象db
var db *sql.DB

// 定义一个user结构体
type user struct {
	id       int
	username string
	password string
}

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN: Data Source Name
	dsn := "root:123456@tcp(10.249.213.185:3306)/go?charset=utf8mb4&parseTime=true"

	// 给全局变量赋值，Open函数并不会校验账户密码是否正确
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 尝试和数据库建立连接
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

// 插入数据
func insertData() {
	sqlStr := "insert into user_tbl(username,password) values(?,?)"

	ret, err := db.Exec(sqlStr, "王五", "223314")
	if err != nil {
		fmt.Printf("insert faild, err: %v\n", err)
	}

	id, err2 := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsertID failed,err2: %v\n", err2)
	}
	fmt.Printf("insert success, the id is : %v\n", id)
}

// 查询单行数据
func queryRow() {
	sqlStr := "select * from user_tbl where id = ?"

	var u user
	// 确保调用QueryRow方法后调用Scan方法，否则持有的数据库连接不会被释放
	err := db.QueryRow(sqlStr, 3).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}
	fmt.Printf("id:%d, username: %s, password: %s", u.id, u.username, u.password)
}

// 查询多行
func queryMultiRow() {
	sqlStr := "select id, username, password from user_tbl where id > ?"

	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
	}

	// 重要，关闭rows，释放数据库连接
	defer rows.Close()

	for rows.Next() {
		var u user
		err2 := rows.Scan(&u.id, &u.username, &u.password)
		if err2 != nil {
			fmt.Printf("scan failed, err2: %v\n", err2)
			return
		}
		fmt.Printf("id:%d\tusername: %s\tpassword: %s\n", u.id, u.username, u.password)
	}
}

// 更新数据
func update() {
	sqlStr := "update user_tbl set username = ?, password = ? where id = ?"
	r, err := db.Exec(sqlStr, "刘六", "999999", 2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	} else {
		i, _ := r.RowsAffected() // 操作影响的行数
		fmt.Printf("update success, affected rows:%d\n", i)
	}

}

func delete() {
	sqlStr := "delete from user_tbl where id = ?"
	ret, err := db.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("delete sql exec failed! err: %v\n", err)
		return
	}
	row, err2 := ret.RowsAffected()
	if err2 != nil {
		fmt.Printf("delete row failed, err2: %v\n", err2)
		return
	}
	fmt.Printf("delete success, delete row: %d\n", row)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("数据库连接失败！ err: %v\n", err)
		return
	} else {
		fmt.Println("数据库初始化成功！")
	}

	// db.SetConnMaxLifetime(time.Minute * 3) // 最大连接时长
	// db.SetMaxOpenConns(10)                 // 最大连接数
	// db.SetMaxIdleConns(10)                 // 最大空闲连接数

	// print(db)

	// insertData()

	// queryRow()
	// queryMultiRow()

	// update()

	delete()
}
