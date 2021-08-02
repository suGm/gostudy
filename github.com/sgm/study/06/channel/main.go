package main

import (
	"fmt"
	"sync"
)

// 通过channel实现多个goroutine之间的通信
// CSP：通过通信来共享内存
// channel是一种类型，一种引用类型。make函数初始化之后才能使用。（slice、map、channel）
// channel的申明:var ch chan 元素类型
// channel的初始化:ch = make(chan 元素类型， [缓冲区大小])
// channel的操作：
// 发送 ch <- 100
// 接收 x := <- ch
// 关闭 close(ch)
// 带缓冲区的通道和无缓冲区的通道:

var a []int
var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup
var x sync.Once

func noBufChannel() {
	b = make(chan int) // 通道初始化,通道必须初始化才能使用,通道必须使用make 不带缓冲区的初始化
	//b = make(chan int, 16) // 通道初始化,通道必须初始化才能使用,通道必须使用make 带缓冲区的初始化
	// 通道的操作
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println(x)
	}()

	// 1、发送值
	b <- 10
	// 2、接收值
	//fmt.Println(x)

	wg.Wait()
}

func bufChannel() {
	b = make(chan int, 16)
	b <- 10
	b <- 20
	x := <-b // 10
	y := <-b // 20
	fmt.Println(x, y)
}

// channel练习
// 1、启动一个goroutine，生成100个数发送到ch1
// 2、启动一个goroutine，从ch1中取值，计算其平方放到ch2中
// 3、在main中从ch2中取值打印出来

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok { // 通道关闭的时候ok为false
			break
		}
		ch2 <- x * x
	}
	x.Do(func() { close(ch2) })
}

func channelWork() {
	a := make(chan int, 50)
	b := make(chan int, 100)
	wg.Add(2)
	go f1(a)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}

func main() {
	//noBufChannel()
	//bufChannel()
	channelWork()
}
