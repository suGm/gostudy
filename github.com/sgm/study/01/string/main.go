package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := "\\gostudy\\src\\github.com\\sgm\\study\\01\\bool" // 双引号定义字符串
	s2 := "+"                                                // 单引号定义字符

	// 切割字符串
	ret := strings.Split(s1, "\\")
	fmt.Println(ret)
	// 包含字符串
	fmt.Println(strings.Contains(s1, "gostudy"))
	// 前缀
	fmt.Println(strings.HasPrefix(s1, "\\"))
	// 后缀
	fmt.Println(strings.HasSuffix(s1, "bool"))
	// 字符串位置
	fmt.Println(strings.Index(s1, "01"))
	// 最后出现位置
	fmt.Println(strings.LastIndex(s1, "s"))
	// 拼接
	fmt.Println(strings.Join(ret, s2))

	// byte(abcd等字节)和rune(中文韩文等字符)类型
	// uint8类型，或者叫byte类型，代表ASCII码的一个字符
	// rune类型(实际是int32)，代表一个UTF-8字符
	// 当要处理中文，日本等复合字符时，需要用到rune类型
	// go使用了rune类型来处理unicode,当然也可以使用byte类型来处理

	s3 := "白萝卜"
	s4 := []rune(s3) // 将字符串强制转换成rune切片
	s4[0] = '红'
	fmt.Println(string(s4)) // 把rune切片强制转换成string

	c1 := "红"
	c2 := '红'
	fmt.Printf("%T, %T\n", c1, c2)
	fmt.Printf("%T", byte(c2))

	// 类型转换 bool类型不能强制转换
	n := 10 // int
	var f = float64(n)
	fmt.Printf("%T", f)

	// 获取下面中汉字的数量
	var str string
	str = "hello , 苏光淼最屌"

	num := 0
	for _, str2 := range str {
		if unicode.Is(unicode.Han, str2) {
			num++
		}
	}

	fmt.Printf("\n%d\n", num)

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d\t", j, i)
		}
		fmt.Println()
	}

}
