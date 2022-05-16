package main

import (
	"fmt"
	"time"
)

func loop(ch chan int) {
	for {
		select {
		case i := <-ch:
			fmt.Println("this  value of unbuffer channel", i)
		}
	}
}

// this will deadlock, because read is after write.
//func main() {
//	ch := make(chan int)
//	ch <- 1
//	go loop(ch)
//	time.Sleep(1 * time.Millisecond)
//}

// correct operation, read before write
func main() {
	ch := make(chan int)
	go loop(ch)
	ch <- 1
	time.Sleep(1 * time.Millisecond)
}

