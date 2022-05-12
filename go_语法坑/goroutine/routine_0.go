package main

import (
	"fmt"
	"sync"
	"time"
)

var l sync.Mutex

func init0() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 10; i < 15; i++ {
		go fmt.Println("go ", i) // 协程
	}

	for i := 100; i < 105; i++ {
		go test(i) // 协程
	}

	time.Sleep(time.Second / 10) // NOTE 等待协程执行完毕
}

func test(key int) {
	l.Lock()         // 加锁
	defer l.Unlock() // 解锁
	fmt.Println("go ", key)
}

func main(){
	init0()
}

