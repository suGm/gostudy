package main

import "fmt"

// 构造函数

type person struct {
	name string
	age  int
}

// 构造函数返回的是结构体还是结构体指针是有考量的，当结构体内容比较小，直接返回结构体是没问题的，但是如果结构体内容比较大，直接返回结构体会导致消耗大量内存，建议用结构体指针

// 构造函数 约定俗成用new开头
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("aa", 10)
	p2 := newPerson("bb", 20)
	fmt.Println(p1, p2)
}
