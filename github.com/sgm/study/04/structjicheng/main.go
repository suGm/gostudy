package main

import "fmt"

// 结构体摸你实现其他语言中的继承

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s会动!", a.name)
}

type dog struct {
	foot uint8
	animal
}

func (d dog) wang() {
	fmt.Printf("%s汪汪汪", d.name)
}

func main() {
	d1 := dog{
		foot: 4,
		animal: animal{
			name: "xxx",
		},
	}

	d1.wang()
	d1.move()
}
