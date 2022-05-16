package main

import (
	"fmt"
	"time"
)

func loop2(ch chan int) {
	for {
		select {
		case i := <-ch:
			fmt.Println("this  value of unbuffer channel", i)
		}
	}
}

func main() {
	ch := make(chan int,3)
	ch <- 1
	ch <- 2
	ch <- 3
	go loop2(ch)
	ch <- 4
	time.Sleep(1 * time.Millisecond)
}