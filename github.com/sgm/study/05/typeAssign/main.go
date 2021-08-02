package main

import "fmt"

// 类型断言

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Printf("%T\n", a)
	}
	fmt.Println(str)
}

func main() {
	assign(100)
}
