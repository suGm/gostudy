package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	str := "10000"
	//ret1 := int64(str)
	ret1, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println(ret1)

	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)

	// 从字符串中解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%t %T\n", boolValue, boolValue)

	i := int32(2000)

	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", ret2)
}
