package main

import "fmt"

func main() {
	//切片定义
	var s1 []int    //定义一个存在int类型的切片
	var s2 []string //定义一个存放string 类型的切片
	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"aaa", "bbb", "ccc"}

	// 长度和容量
	fmt.Printf("len[s1]:%d cap(s2):%d", len(s1), cap(s2))

	// 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13} // 长度7 容量7
	s3 := a1[0:4]                         // 基于一个数组切割 左闭右开 长度4 容量7 切片的容量指的是从底层数组第一个元素指向最后一个元素的数量
	fmt.Println(s3)

	s4 := s3[3:]
	// 长度和容量 切片是引用类型
	fmt.Printf("len[s4]:%d cap(s4):%d", len(s4), cap(s4))
}
