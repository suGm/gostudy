package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

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

type user struct {
	ID   int
	Name string
	Age  int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init Db failed, err:%v\n", err)
		return
	}

	sqlStr1 := `select id,name,age from tb_emp1 where id=42`
	var u user
	err = db.Get(&u, sqlStr1)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("u:%v\n", u)

	var userList = make([]user, 0, 10)
	sqlStr2 := `select id,name,age from tb_emp1 where id < 45`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("userList:%#v\n", userList)
}
