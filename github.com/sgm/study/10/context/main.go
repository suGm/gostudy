package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要context

var wg sync.WaitGroup
var exiyChan = make(chan bool, 1)

func f() {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("func f()")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exiyChan:
			break LOOP
		default:

		}
	}
}

func main() {
	// 如何通知子goroutine退出
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	exiyChan <- true
	wg.Wait()

}
