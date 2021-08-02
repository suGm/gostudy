package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	// 使用匿名嵌套结构体
	address
}

type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "1",
		age:  1,
		address: address{
			province: "2",
			city:     "3",
		},
	}

	fmt.Println(p1.province)
}
