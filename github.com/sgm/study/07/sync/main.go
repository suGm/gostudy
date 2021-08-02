package main

import "sync"

var wg sync.WaitGroup

// 某个操作只想执行一次
var once sync.Once

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}

	f := func() {
		close(ch2)
	}
	once.Do(f)
}

func main() {

}
