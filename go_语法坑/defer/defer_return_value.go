package main

import (
	"fmt"
)

func f() int {
	var i = 0
	defer func() {
		i++
		fmt.Printf("defer i = %d\n", i)
	}()
	fmt.Printf("main i = %d\n", i)
	return i
}

func main() {
	d := f()
	fmt.Println(d)
}
