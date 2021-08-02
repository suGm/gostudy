package main

import "fmt"

func main() {
	fmt.Println(taijie(4))
}

// 阶乘
func f(n int) int {
	if n-1 > 0 {
		return n * f(n-1)
	}

	return n
}

func taijie(n uint64) uint64 {
	if n <= 2 {
		return n
	}
	return taijie(n-1) + taijie(n-2)
}
