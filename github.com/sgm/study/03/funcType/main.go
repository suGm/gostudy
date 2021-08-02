package main

import "fmt"

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	f3(f2)

}

func f1() {
	fmt.Println("hellow")
}

func f2() int {
	return 10
}

func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func ff(x func() int) func() int {
	return func() int {
		return 1
	}
}

// 闭包
func f4(f func()) {
	fmt.Println("f4")
	f()
}

func f5(x, y int) {
	fmt.Println("f5")
	fmt.Println(x + y)
}
