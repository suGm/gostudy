package main

import (
	"fmt"
	"strings"
)

// 1、函数可以作为返回值
// 2、函数内部查找变量的顺序，现在自己内部找，找不到往外层找

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func f1(f func()) {
	fmt.Println("f1")
	f()
}

func f2(x, y int) {
	fmt.Println("f2")
	fmt.Println(x + y)
}

func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	//ret := adder(100)
	//ret2 := ret(200)
	//fmt.Println(ret2)

	f1(f3(f2, 100, 200))
	jpegFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpegFunc("text"))
	fmt.Println(txtFunc("text"))

	f5, f6 := calc(10)

	fmt.Println(f5(1), f6(2)) // 11 9
	fmt.Println(f5(3), f6(4)) // 12 8
}
