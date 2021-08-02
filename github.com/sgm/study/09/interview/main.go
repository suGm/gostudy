package main

import "fmt"

// 如何判断一个链表有没有闭环
// 定义两个变量 x和y x每次走一步 y每次都2步， 如果某一时刻x=y。则说明这个链表有闭环

// 台阶问题 一次可以迈一个台阶，一次可以迈2个台阶，现在有n个台阶，有多少种解法
// 计算最后一次 迈一步的次数+迈两步的次数

type a struct {
	val  int
	next *a
}

func f(n int) int {
	if n == 1 || n == 0 {
		return n
	}

	a := 2
	b := 1
	c := 0

	for i := 1; i <= n; i++ {
		a = b + c
		c = b
		b = a
	}

	return a
}

func main() {
	fmt.Println(f(6))
}
