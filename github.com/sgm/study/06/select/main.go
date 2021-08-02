package main

import "fmt"

// 同一时刻有多个通道要操作的场景，使用select

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:

		}
	}
}
