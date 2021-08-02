package main

import (
	"fmt"

	calc "github.com/sgm/study/05/package"
)

func init() {
	fmt.Println("自动调用1")
}

func main() {
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}
