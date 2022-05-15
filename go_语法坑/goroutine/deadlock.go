package main

import (
	"fmt"
	"time"
)

// 只写不读
func deadlock1() {
	ch := make(chan int)
	ch <- 3 //  这里会发生一直阻塞的情况，执行不到下面一句
}

// 写在读后面
func deadlock2() {
	ch := make(chan int)
	ch <- 3  //  这里会发生一直阻塞的情况，执行不到下面一句
	num := <-ch
	fmt.Println("num=", num)
}

// 写在读后面
func deadlock22() {
	ch := make(chan int)
	ch <- 100 //  这里会发生一直阻塞的情况，执行不到下面一句
	go func() {
		num := <-ch
		fmt.Println("num=", num)
	}()
	time.Sleep(time.Second)
}

// 超出channel缓冲区
func deadlock3() {
	ch := make(chan int, 3)
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6  //  这里会发生一直阻塞的情况
}

// 空读
func deadlock4() {
	ch := make(chan int)
	// ch := make(chan int, 1)
	fmt.Println(<-ch)  //  这里会发生一直阻塞的情况
}

// 互相等待，类似于Linux的线程死锁
func deadlock5() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 互相等对方造成死锁
	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("num=", num)
				ch2 <- 100
			}
		}
	}()

	for {
		select {
		case num := <-ch2:
			fmt.Println("num=", num)
			ch1 <- 300
		}
	}
}

func main(){
	deadlock5()
}