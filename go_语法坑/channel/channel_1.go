package main

import (
	"fmt"
	"time"
)

func init1() {
	var c chan int = make(chan int) //  var 管道名 chan 传输类型 = make(chan 传输类型， 缓存个数)
	defer close(c)                  // 使用完毕后关闭，关闭后不可在使用
	// c <- 11                      // TODO fatal error: all goroutines are asleep - deadlock!

	for i := 0; i < 5; i++ {
		go prod(c, i+1)() // 调用prod生成的函数
	}

breakFor:
	for {
		select { // 				多路复用 哪个分支的case 先触发就先执行哪个
		case m := <-c: // 			c 接收数据触发
			fmt.Println(m)
		case <-time.After(time.Second): // 超时1秒 触发
			fmt.Println("time out") //
			break breakFor          // NOTE 跳出for循环
		}
	}

}

func prod(c chan int, n int) func() {
	return func() {
		for i := 0; i < n; i++ {
			c <- i
		}
	}
}

func main(){
	init1()
}
