package main

import (
	"fmt"
	"os"
)

// 学生管理

var allStudent = make(map[int]*student, 50) // 变量声明

type student struct {
	id   int
	name string
}

func main() {
	// 1、打印菜单
	// 2、等待选择
	// 3、执行对应操作
	for {

		fmt.Println("---------------------")
		fmt.Println(`
			1、查看所有学生
			2、新增学生
			3、删除学生
		`)
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("输入了%d\n", choice)

		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			delStudent()
		case 4:
			os.Exit(4)
		default:
			fmt.Println("滚")
		}
	}
}

func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}

//
//func newStudent(id int, name string) *student {
//	return &student{
//		id: id,
//		name:name,
//	}
//}

func addStudent() {
	var (
		id   int
		name string
	)
	fmt.Print("输入:id")
	fmt.Scanln(&id)
	fmt.Print("输入:name")
	fmt.Scanln(&name)
	stu := newStudent(id, name)
	allStudent[id] = stu
}

func delStudent() {
	// 1、输入要删除的学生的学号
	// 2、去allStudent的map中去删除该学号为输入学号的学生
	var (
		deleteId int
	)

	fmt.Println("输入需要删除的学号")
	fmt.Scanln(&deleteId)
	delete(allStudent, deleteId)

}
