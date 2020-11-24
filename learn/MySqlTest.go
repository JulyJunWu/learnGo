package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 这个的作用是注册一个mysql的驱动包
)

// 数据库连接
func main() {
	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ws")
	defer db.Close()
	//ConnectionTest()
	//TransactionTest()
	one := QueryOne(1, db)
	fmt.Println(one)
}

// 查询单条
func QueryOne(id int, db *sql.DB) User {
	prepare, _ := db.Prepare("select * from user where id = ?")
	queryRow := prepare.QueryRow(id)
	var u User
	_ = queryRow.Scan(&u.Id, &u.Name, &u.Age, &u.Sex)
	return u
}

// 开启事务
func TransactionTest() {
	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ws")
	defer db.Close()
	// 开启一个事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	prepare, err := tx.Prepare("update user set name = '西瓜' where id = ?")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 设置参数并执行
	result, _ := prepare.Exec(1)
	fmt.Println(result)
	// 提交事务
	_ = tx.Commit()
}

func ConnectionTest() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ws")
	if err != nil {
		fmt.Println(err)
	}
	query, _ := db.Query("select * from user ")
	defer query.Close()
	fmt.Println(query)
	users := make([]User, 0, 100)
	for query.Next() {
		_, _ = query.Columns()
		var u = User{}
		err := query.Scan(&u.Id, &u.Name, &u.Age, &u.Sex)
		if err == nil {
			// 封装为对应的User
			users = append(users, u)
		}
	}
	fmt.Println(len(users), cap(users))
	fmt.Println(users)
}

type User struct {
	Id   int
	Name string
	Age  int
	Sex  string
}
