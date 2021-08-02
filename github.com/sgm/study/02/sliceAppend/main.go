package main

import "fmt"

// append() 为切片追加元素
func main() {
	s1 := []string{"北京", "上海", "深圳"}
	// 调用append函数必须用原来的切片变量接收
	s1 = append(s1, "广州") // append追加元素，原来的底层数组放不下时，go底层就会把底层数组换一个新的所以上面两个s1指向的地址不同
	fmt.Println(s1)

	s2 := []string{"武汉", "成都"}
	s1 = append(s1, s2...)
}
