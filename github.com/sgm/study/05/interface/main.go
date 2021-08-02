package main

import "fmt"

// 引出接口实例

type speaker interface {
	speak() // 只要实现了speak方法的类型都是speaker类型
}

type cat struct {
}

type dog struct {
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func da(x speaker) {
	// 接收一个参数，传进来什么就da什么
	x.speak()
}

// 定义一个car接口类型
// 不管是什么结构体 只要有run方法都能称为car类型
type car interface {
	run()
}

type car1 struct {
	brand string
}

type car2 struct {
	brand string
}

func (c1 car1) run() {
	fmt.Println("car1")
}

func (c2 car2) run() {
	fmt.Println("car2")
}

func drive(c car) {
	c.run()
}

func main() {
	//var c1 cat
	//var d1 dog
	//da(c1)
	//da(d1)

	var c1 = car1{brand: "c1"}
	var c2 = car2{brand: "c2"}

	c1.run()
	c2.run()
	drive(c1)
	fmt.Printf("%v\n", c1)
}
