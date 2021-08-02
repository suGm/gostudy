package main

import "fmt"

//变量的作用域

var x = 100

func main() {
	f1()
}

func f1() {
	fmt.Println(x)
}
