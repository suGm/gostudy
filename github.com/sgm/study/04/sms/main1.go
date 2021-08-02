package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int
	name string
}

type sms struct {
	allStudent map[int]*student
}

func newStudent(id int, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func (s *sms) showAllStudent() {
	fmt.Println("查看所有学生")
	for k, v := range s.allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
	fmt.Println("-------end--------")
	fmt.Printf("%v\n", s)
}

func (s *sms) addStudent() {
	fmt.Println("新增学生")

	var (
		id   int
		name string
	)
	fmt.Print("输入id:")
	fmt.Scanln(&id)
	fmt.Print("输入name:")
	fmt.Scanln(&name)
	var stu = newStudent(id, name)
	s.allStudent[id] = stu
}

func (s *sms) delStudent() {
	fmt.Println("删除学生")

	var (
		deleteId int
	)

	fmt.Print("输入要删除的id:")
	fmt.Scanln(&deleteId)
	delete(s.allStudent, deleteId)
}

// 构造函数
func newSms(allStudent map[int]*student) sms {
	return sms{
		allStudent: allStudent,
	}
}

func main() {
	s := newSms(make(map[int]*student))
	for {
		fmt.Println("---------------")
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
			s.showAllStudent()
		case 2:
			s.addStudent()
		case 3:
			s.delStudent()
		case 4:
			os.Exit(4)
		default:
			fmt.Println("滚")
		}

	}
}
