package main

import "fmt"

const pi = 3.1415926

// 每个const声明都会让iota计数器清0
// 每一行常量声明都会让iota+1,不管这一行常量有几个
const (
	n1, n2 = iota + 1, iota + 2 // 1 2
	n3, n4 = iota + 1, iota + 2 // 2 3
	n5, n6 = iota + 1, iota + 2 // 3 4
)

func main() {
	fmt.Println(n1, n2, n3, n4, n5, n6)
}
