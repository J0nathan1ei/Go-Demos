package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

// 利用有缓冲channel满了就阻塞的特性，来限制goroutine的数量
func main() {
	// 模拟用户请求数量
	requestCount := 10
	fmt.Println("goroutine_num", runtime.NumGoroutine())
	// 管道长度即最大并发数
	ch := make(chan bool, 3)
	for i := 0; i < requestCount; i++ {
		wg.Add(1)
		ch <- true
		go Read(ch, i)
	}

	wg.Wait()
}

func Read(ch chan bool, i int) {
	fmt.Printf("goroutine_num: %d, go func: %d", runtime.NumGoroutine(), i)
	<-ch
	wg.Done()
}