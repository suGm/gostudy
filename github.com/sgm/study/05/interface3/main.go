package main

import "fmt"

// 空接口

type xx interface {
}

// 空接口没有必要起名字，通常定义成 interface{}
// 所有的类型都实现了空接口，也就是任意类型的变量都能保存到空接口中
//func Println(a ...interface{})  {
//
//}

// 空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("type:%T\n", a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "xx"
	m1["age"] = 1
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)
}
