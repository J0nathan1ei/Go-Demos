package main

import (
	"fmt"
)

// TODO 运行前猜测结果
/*

 */
func main() {
	exec3()
}

func exec3() {
	const max = 3
	var ci [max]chan int

	for i := 0; i < max; i++ {
		go func(i int) {
			ci[i] <- i * 100
		}(i)
	}

	for i := 0; i < max; i++ {
		fmt.Println(<-ci[i])
	}
}
