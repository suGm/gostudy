package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      = 0
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

// 读操作加读锁
func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
}

func write() {
	defer wg.Done()
	rwlock.Lock()
	x += 1
	time.Sleep(time.Millisecond * 5)
	rwlock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
