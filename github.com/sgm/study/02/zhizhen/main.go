package main

import "fmt"

// 指针 pointer

func main() {
	// 1、&:取地址
	// 2、*:根据地址取值
	n := 10
	p := &n
	fmt.Println(&p)
	fmt.Printf("%T\n", p)
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	//var a *int // nil
	//*a = 100 // 报错，没有内存地址
	//fmt.Println(*a)

	// 上面的解决方法
	var a = new(int) // 开辟了一块内存空间
	*a = 100         // 通过地址赋值
	fmt.Println(*a)

	//make 和 new 都是用来申请内存的
	// new 很少用，基本上是给基本数据类型申请内存空间的
	// make是用来给slice、map、chan申请内存的

}
