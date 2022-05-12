package main

import (
	"fmt"
)

// TODO 运行前猜测结果
/*
5
5
5
5
5
111
*/
func main() {
	exec4()
}

func exec4() {
	const max = 5
	cc := make(chan *int, max)
	defer close(cc)
	defer fmt.Println(111)

	ss := make([]int, max)
	for i := 0; i < max; i++ {
		ss[i] = i + 1
	}

	for _, v := range ss {
		cc <- &v
	}

	for i := 0; i < max; i++ {
		fmt.Println(*<-cc)
	}
}
