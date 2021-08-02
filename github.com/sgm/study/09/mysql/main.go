package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 一个数据库的连接池对象
var db *sql.DB

func initDB() (err error) {
	// 连接数据库
	dsn := "root:abcdefg@tcp(127.0.0.1:33306)/test"

	// 连接数据库 不会校验用户名密码是否正确，但是会验证dsn格式
	db, err = sql.Open("mysql", dsn)

	// dsn 格式不正确的时候会报错
	if err != nil {
		return err
	}

	// 尝试连接数据库，是否成功
	err = db.Ping()
	if err != nil {
		return err
	}

	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	// 设置数据库连接池的最大闲置连接数
	db.SetMaxIdleConns(5)
	return nil
}

type user struct {
	id   int
	name string
	age  int
}

func query(n int) {
	var u1 user
	// 写查询单条语句
	sqlStr := `select id,name,age from tb_emp1 where id=?;`

	//执行 并 解析结果
	db.QueryRow(sqlStr, n).Scan(&u1.id, &u1.name, &u1.age) // 必须对rowObj对象调用scan方法，因为该方法会释放连接

	// 打印结果
	fmt.Printf("u1:%#v\n", u1)
}

func queryMore(n int) {
	// 1.SQL语句
	sqlStr := `select id,name,age from tb_emp1 where id < ?;`
	// 2.执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%v\n", sqlStr, err)
		return
	}
	// 3。一定要关闭rows
	defer rows.Close()

	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
		} else {
			fmt.Printf("u1:%#v\n", u1)
		}
	}
}

func insert() {
	// 1、写sql语句
	sqlStr := `insert into tb_emp1(name, age) values("sgm", 11);`

	// 2、exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed err:%#v\n", err)
		return
	}

	// 3、如果是插入数据的操作，能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get last failed,err:%v\n", err)
		return
	}

	fmt.Println("id:", id)
}

// 更新操作
func updateRow(id, newAge int) {
	sqlStr := `update tb_emp1 set age=? where id = ?;`

	ret, err := db.Exec(sqlStr, newAge, id)

	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}

	fmt.Printf("更新了%d行数据\n", n)

}

func deleteRow(id int) {
	sqlStr := `delete from tb_emp1 where id=?;`

	ret, err := db.Exec(sqlStr, id)

	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected()

	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}

	fmt.Printf("删除了%d行数据\n", n)
}

// 预处理方式插入多条数据
func prepareInsert() {
	sqlStr := `insert into tb_emp1(name, age) values(?, ?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}

	defer stmt.Close()

	// 后续只需要拿到stmt去执行一些操作
	var m = map[string]int{
		"1": 30,
		"2": 10,
	}

	for k, v := range m {
		stmt.Exec(k, v)
	}
}

// 事务
func transaction() {
	// 开启事务
	tx, err := db.Begin()

	if err != nil {
		fmt.Printf("start begin failed,err:%v\n", err)
		return
	}

	// 执行多个sql操作
	sqlStr1 := `update tb_emp1 set age = age + 2 where id = 1;`
	sqlStr2 := `update tb_emp1 set age = age - 2 where id = 43;`

	_, err = tx.Exec(sqlStr1)

	if err != nil {
		// 回滚
		tx.Rollback()
		fmt.Printf("执行sql1出错回滚:%#v\n", err)
		return
	}

	_, err = tx.Exec(sqlStr2)
	if err != nil {
		fmt.Printf("执行sql2出错回滚:%#v\n", err)
		// 回滚
		tx.Rollback()
		return
	}

	// 执行成功 提交本次事务
	err = tx.Commit()
	fmt.Println("操作成功")
}

// Go连接mysql示例
func main() {
	err := initDB()

	if err != nil {
		fmt.Printf("open failed, err:%v\n", err)
		return
	}

	fmt.Println("连接数据库成功!")

	//query(1009854)

	//queryMore(50)

	//insert()

	//updateRow(1009854, 26)

	//deleteRow(1009854)

	//prepareInsert()

	transaction()
}
