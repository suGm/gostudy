package main

import (
	"fmt"
)

// goroutine 的本质是GMP
// 把m个goroutine分配给n个操作系统线程
// goroutine 是用户态的线程，比内核态的线程更加轻量级，初始时只占用2kb空间
// 可以轻松开启数十万的goroutine也不会崩内存
// runtime.GOMAXPROCS 在go1.5之后默认就是操作系统的逻辑核心数，默认跑满CPU

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 100; i++ {
		go func(i int) { // goroutine对应的函数执行完了，goroutine也就结束了
			fmt.Println(i)
		}(i)
	}
	//go hello() // 开启一个单独的goroutine去执行hello函数
	fmt.Println("main")
	// main函数结束了 由main函数启动的goroutine都结束
	//time.Sleep(time.Second * 5)

}
