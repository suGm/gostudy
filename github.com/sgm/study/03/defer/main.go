package main

import "fmt"

// go函数的return不是原子操作，是分为两步进行操作的
// 第一步，返回值赋值
// 第二步，真正的RET返回
// 如果存在defer那么defer执行的时机是在第一步和第二步之间

func main() {
	//defer fmt.Println("main")
	//deferDemo()
	//fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	//fmt.Println(f4())
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}

func f1() int {
	x := 5
	defer func() {
		x++
	}() // 由于这里返回的是一个int类型所以当return x的时候就已经把5赋值返回值中，在把x++就与返回值无关了，直接返回5
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}() // 由于这里返回的是一个x引用类型的值，所以当return x的时候就已经把5赋予在x的地址上，再x++则会改变x地址的数值，所以返回x地址后，输出的就是6（引用传值）
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 由于这里返回的是y的地址，是赋值传参 即x=5 y=x的副本，所以改变了x不改变y的地址，输出的就是5
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 由于这里的defer中的x是赋值传参，即返回的是一个x的副本，所以匿名函数内的x++与外界的x无关，所以返回5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
