package main

import (
	"fmt"
	"strconv"
	"sync"
)

// go 内置的map不是并发安全的

var m = make(map[string]int)
var lock sync.Mutex

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

//func main() {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 100; i++ {
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n)
//			lock.Lock()
//			set(key, n)
//			lock.Unlock()
//			fmt.Println(get(key))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

// 并发安全的syncMap
var m2 = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)        //存值必须要使用sync中的Store方法设置键值对
			value, _ = m2.Load(key) // 必须使用sync.Map提供的Load方法根据key取值
			fmt.Println(key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
