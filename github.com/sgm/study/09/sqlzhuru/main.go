package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type user struct {
	ID   int
	Name string
	Age  int
}

// sql注入

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:abcdefg@tcp(127.0.0.1:33306)/test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}

	db.SetMaxOpenConns(10) // 设置数据库连接池最大连接数
	db.SetMaxIdleConns(5)  // 设置最大空闲连接数
	return
}

func sqlInject(name string) {
	// 自己拼接sql语句的字符串
	sqlStr := fmt.Sprintf("select id,name,age from tb_emp1 where name='%s'", name)

	fmt.Printf("SQL:%s\n", sqlStr)

	var users []user
	err := db.Select(&users, sqlStr)

	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}

	for _, u := range users {
		fmt.Printf("user:%#v\n", u)
	}

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init Db failed, err:%v\n", err)
		return
	}

	// SQL注入的几种示例

	sqlInject("sgm")

}
