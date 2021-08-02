package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	x += 1
	fmt.Println(x)
	time.Sleep(5 * time.Millisecond)
	lock.Unlock()
}

func add() {
	for i := 0; i < 50000; i++ {
		// 互斥锁 锁住资源
		lock.Lock()
		x += 1 // 两个goroutine让x同时新增一个数导致x永远不会加到1000000
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	//wg.Add(2)
	//go add()
	//go add()
	//wg.Wait()
	//fmt.Println(x)
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}

	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
