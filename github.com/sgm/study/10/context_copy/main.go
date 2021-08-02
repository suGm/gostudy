package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 为什么需要context

var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
}

func f2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("func f()")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
}

func main() {
	// 如何通知子goroutine退出
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)

	cancel()

	wg.Wait()

}
