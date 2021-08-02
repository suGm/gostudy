package main

import "fmt"

func main() {
	f1 := 1.23456
	fmt.Printf("%T\n", f1) // 默认用的float64
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)

	// f1 不能直接 赋值给 f2 类型不一样
}
