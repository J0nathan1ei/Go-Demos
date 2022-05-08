package main

import "fmt"

func main() {
	x := make([]int, 3)
	for i := 0; i < 3; i++ {
		x[i] = i
	}
	firstPtr := &x[0]
	for i := 3; i < 6; i++ {
		x = append(x, i)
	}
	*firstPtr = 100

	fmt.Println(x)
}
