package main

import "fmt"

// copy

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	a3 := make([]int, 3, 4)
	copy(a3, a1)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	// 删除切片元素
	// 删除索引为1的元素
	a1 = append(a1[:1], a1[2:]...)

	x1 := [...]int{1, 3, 5}
	s1 := x1[:] // len 3 cap 3
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(x1) // 1 5 5
	fmt.Println(s1) // 1 5
}
