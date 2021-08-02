package main

import "fmt"

// 使用值接收者和指针接收者的区别

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet uint8
}

func (c *cat) move() {
	fmt.Println("猫步")
}

func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
	var a1 animal

	c1 := cat{"tom", 4}   // cat
	c2 := &cat{"tom1", 4} // *cat

	a1 = &c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
