package main

import (
	"fmt"
	"sync"
)

func initSync1() {
	var l sync.Mutex // 互斥锁
	l.Lock()         // 互斥锁加锁 其他加锁操作等待解锁
	l.Unlock()       // 互斥锁解锁

	var rl sync.RWMutex // 读写锁
	func() {
		rl.RLock()         // 读写锁 读锁加锁 其他读锁可叠加锁，其他写锁等待所有锁解开后加锁
		defer rl.RUnlock() // 读写锁 读锁解锁
	}()

	func() {
		rl.Lock()         // 读写锁 写锁加锁 其他锁需要等待写锁解锁
		defer rl.Unlock() // 读写锁 写锁解锁
	}()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) // 加一
		go func(n int) {
			fmt.Println(n)
			wg.Done() // 减一
		}(i)
	}
	wg.Wait() // 等待清零

	var 只执行一次 sync.Once
	只执行一次.Do(test)
	只执行一次.Do(test)
	只执行一次.Do(test)
	只执行一次.Do(test)
}

func test() {
	fmt.Println("test")
}


func main(){
	initSync1()
}