package main

import (
	"fmt"
	"sync"
)

var Size2 int = 80000

func main() {
	demo := make([]int, Size2)
	for i := 0; i < Size2; i++ {
		demo[i] = i
	}
	res := sum2(demo)

	fmt.Println(res)
}

func sum2(data []int) int {
	res := 0
	length := len(data)
	const N = 5
	seg := length / N

	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(N) // 直接加N个

	for i := 0; i < N; i++ {
		go func(j int) {
			tmpS := data[j*seg : (j+1)*seg]
			mu.Lock()
			for i := 0; i < len(tmpS); i++ {
				res += tmpS[i]
			}
			mu.Unlock()
			wg.Done() // 一个goroutine运行完
		}(i)
	}
	wg.Wait() // 等N个goroutine都运行完

	return res
}
