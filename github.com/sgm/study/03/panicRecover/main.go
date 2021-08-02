package main

import "fmt"

func funcA() {
	fmt.Println("a")
}

func funcB() {
	defer func() {
		//fmt.Println("释放连接") // 报错之前可以释放连接
		err := recover() // 这里会打印err，但是不会讲程序整个退出 recover必须搭配defer使用
		fmt.Println(err)
	}()
	panic("未知错误") // 程序崩溃退出
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()

	// 获取用户输入
	//var s string
	//fmt.Scan(&s)
	//fmt.Println(s)

	var (
		name  string
		age   int
		class string
	)
	fmt.Scanf("%s %d %s \n", &name, &age, &class)
	fmt.Println(name, age, class)
}
