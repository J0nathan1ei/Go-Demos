package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup

	testStrings := make([]string, 16)

	for i := 0; i < len(testStrings); i++{
		testStrings[i] = "test" + strconv.Itoa(i)
	}

	for _, t := range testStrings{
		go func(){
			wg.Add(1)
			fmt.Println(t)
			wg.Done()
		}()
	}


	// 解决办法，赋予协程参数，不让它形成闭包
	// 先sleep 1s，让前面的协程执行完
	time.Sleep(1*time.Second)
	fmt.Println("No Closure: ")
	for _, t2 := range testStrings{
		go func(s string){
			wg.Add(1)
			fmt.Println(s)
			wg.Done()
		}(t2)
	}

	wg.Wait()
}
