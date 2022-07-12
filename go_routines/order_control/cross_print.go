package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)
// 控制goroutine的打印顺序，主要是通过channel实现
// 使用空结构体的channel控制顺序，一个channel为入参，一个channel为出参

func crossPrint(){
	var wg sync.WaitGroup

	numChan := make(chan struct{}, 1)
	charChan := make(chan struct{}, 1)
	wg.Add(2)
	charChan <- struct{}{}
	go printNum(charChan, numChan, &wg)
	go printChar(charChan, numChan, &wg)
	wg.Wait()
}

func printNum(charChan, numChan chan struct{}, wg *sync.WaitGroup){
	var cnt int64 = 1
	for cnt <= 28{
		<-charChan
		fmt.Print(cnt)
		fmt.Print(cnt+1)
		atomic.AddInt64(&cnt, 2)
		numChan<- struct{}{}
	}
	wg.Done()
}

func printChar(charChan, numChan chan struct{}, wg *sync.WaitGroup){
	var cnt int64 = 0
	for cnt < 26{
		<-numChan
		fmt.Printf("%c", 'A' + cnt)
		fmt.Printf("%c", 'A' + cnt + 1)
		atomic.AddInt64(&cnt, 2)
		charChan<- struct{}{}
	}
	wg.Done()
}


func main(){
	crossPrint()
}