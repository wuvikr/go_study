package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID   int    `sqlxDB:"id"`
	Age  int    `sqlxDB:"age`
	Name string `sqlxDB:"name`
}

var sqlxDB *sqlx.DB

func initSqlxDB() (err error) {
	// 数据库连接信息
	user := "root"
	password := "123456"
	host := "127.0.0.1"
	port := "3306"
	database := "go_test"

	// 连接数据库
	sqlxDB, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database))

	if err != nil {
		log.Fatalln(err)
	}

	// 设置连接池配置
	sqlxDB.SetMaxOpenConns(100)            // 最大连接数
	sqlxDB.SetMaxIdleConns(10)             // 最大空闲连接数
	sqlxDB.SetConnMaxLifetime(time.Minute) // 连接最大生存时间

	return err
}

// 插入数据
func insert() {
	user := User{Name: "Tom", Age: 20}
	sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
	result, err := sqlxDB.NamedExec(sqlStr, user)

	if err != nil {
		log.Fatalln(err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}

	userCounts, _ := result.RowsAffected()
	fmt.Printf("插入了 %d 条数据，新插入的数据的 ID 是 %d\n", userCounts, userID)
}

// 查询数据
func findData() {
	var users []User

	sqlStr := "SELECT * FROM user"
	err := sqlxDB.Select(&users, sqlStr)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("查询到了 %d 条数据：\n", len(users))

	for _, u := range users {
		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d\n", u.ID, u.Name, u.Age)
	}
}

// 更新数据
func update() {
	user := User{ID: 1, Name: "Jerry", Age: 25}

	sqlStr := "UPDATE user SET name=:name, age=:age WHERE id=:id"
	result, err := sqlxDB.NamedExec(sqlStr, user)
	if err != nil {
		log.Fatalln(err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("更新了 %d 条数据\n", rowsAffected)
}

// 删除数据
func delete() {
	sqlStr := "DELETE FROM user WHERE id=?"
	result, err := sqlxDB.Exec(sqlStr, 1)
	if err != nil {
		log.Fatalln(err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("删除了 %d 条数据\n", rowsAffected)
}

func main() {
	err := initSqlxDB()

	if err != nil {
		panic(err)
	}

	defer sqlxDB.Close()

	// // 创建表
	// _, err = sqlxDB.Exec(`
	// 	CREATE TABLE IF NOT EXISTS user (
	// 		id INT(11) NOT NULL AUTO_INCREMENT,
	// 		name VARCHAR(50) NOT NULL,
	// 		age INT(11) NOT NULL,
	// 		PRIMARY KEY (id)
	//   	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	// `)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// update()

	delete()

}
