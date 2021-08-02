package main

import (
	"fmt"
	"github.com/sgm/study/08/split_string"
)

func main() {
	ret := split_string.Split("babcbef", "b")
	fmt.Printf("%#v\n", ret)
	ret1 := split_string.Split("bbb", "b")
	fmt.Printf("%#v\n", ret1)
	ret2 := split_string.Split("jdkajlji", "b")
	fmt.Printf("%#v\n", ret2)
}
