package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

// https://www.jianshu.com/p/7161a8e17da5
// channel 使用完后未关闭导致go-routine泄漏
func bug(_ http.ResponseWriter, _ *http.Request) {
	taskChan := make(chan int, 100)

	for i := 0; i < 100; i++ {
		taskChan <- i
	}

	consumer := func() {
		for task := range taskChan {
			fmt.Println(task)
		}
	}

	for i := 0; i < 100; i++ {
		go consumer()
	}
}

func main() {
	http.HandleFunc("/bug", bug)
	http.ListenAndServe(":8000", nil)
}
