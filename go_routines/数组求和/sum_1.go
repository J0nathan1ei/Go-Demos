package main

import "fmt"

var Size int = 80000

func main() {
	demo := make([]int, Size)
	for i := 0; i < Size; i++ {
		demo[i] = i
	}
	res := sum(demo)

	fmt.Println(res)
}

func sum(data []int) int {
	res := 0
	length := len(data)
	const N = 5
	seg := length / N
	var chs [N]chan int

	for i := 0; i < N; i++ {
		chs[i] = worker(data[i*seg : (i+1)*seg])
	}

	for i := 0; i < N; i++ {
		res += <-chs[i]
	}

	return res
}

func worker(s []int) chan int {
	res := make(chan int)

	go func() {
		sum := 0
		for i := 0; i < len(s); i++ {
			sum += s[i]
		}
		res <- sum
	}()
	return res
}
