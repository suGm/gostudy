package main

import "fmt"

// 声明变量
//var name string
//var age int
//var isOk bool

// 批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "sgm"
	age = 18
	isOk = true
	fmt.Print(isOk)
}
