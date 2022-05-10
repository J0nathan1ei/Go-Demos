package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan string)

	//缓冲通道预先放置10个消息
	messages := make(chan int, 10)
	defer close(messages)
	for i := 0; i < 10; i++ {
		messages <- i
	}
	//启动3个子协程消费messages消息
	for i := 1; i <= 3; i++ {
		go child(i, done, messages)
	}

	time.Sleep(3 * time.Second) //等待子协程接收一半的消息
	close(done)                 //结束前通知子协程
	time.Sleep(2 * time.Second) //等待所有的子协程输出
	fmt.Println("主协程结束")
}

//从messages通道获取信息，当收到结束信号的时候不再接收
func child(i int, done <-chan string, messages <-chan int) {
Consume:
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-done:
			fmt.Printf("[%d]被主协程通知结束...\n", i)
			break Consume
		default:
			fmt.Printf("[%d]接收消息: %d\n", i, <-messages)
		}
	}
}
