package main

import "fmt"

func main() {
	r := sum(1, 2)
	fmt.Println(r)
	x, y, z := f4()
	fmt.Println(x, y, z)
}

func sum(x, y int) (ret int) {
	return x + y
}

func f1(x, y int) {
	fmt.Println(x + y)
}

func f2() {
	fmt.Println(1)
}

func f3() int {
	return 3
}

func f4() (int, string, bool) {
	return 1, "s", true
}
