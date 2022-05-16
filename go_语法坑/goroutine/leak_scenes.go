package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("before goroutines: ", runtime.NumGoroutine())
	block1()
	time.Sleep(time.Second * 1)
	fmt.Println("after goroutines: ", runtime.NumGoroutine())
}

// channel未初始化阻塞
func block1() {
	var ch chan int
	for i := 0; i < 10; i++ {
		go func() {
			<-ch
		}()
	}
}

// channel接收者大于发送者
func block2() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			ch <- 1
		}()
	}
}

// channel发送者大于接收者
func block3() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			<-ch
		}()
	}
}

// http 泄漏场景
var wg1 = sync.WaitGroup{}
func requestWithNoClose() {
	_, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("error occurred while fetching page, error: %s", err.Error())
	}
}

func requestWithClose() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("error occurred while fetching page, error: %s", err.Error())
		return
	}
	defer resp.Body.Close()
}

func block4() {
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			requestWithNoClose()
		}()
	}
}


// 加锁未解锁
func block5() {
	var mutex sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			mutex.Lock()
		}()
	}
}


// WaitGroup add和done不匹配
func block6() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(2)
			wg.Done()
			wg.Wait()
		}()
	}
}