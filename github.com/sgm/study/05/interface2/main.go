package main

import "fmt"

// 接口还可以嵌套
type animal interface {
	mover
	eater
}

// 同一个结构体可以实现多个接口
type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet uint8
}

func (c *cat) move() {
	fmt.Println("猫步")
}

func (c *cat) eat(foot string) {
	fmt.Printf("eat %s...\n", foot)
}

func main() {

}
