package main

import (
	"fmt"
)

func f1() int {
	var i = 0
	defer A1()
	panic("test defer function's order")
	defer A2()
	return i
}

func A1() {
	fmt.Println("A1")
}

func A2() {
	fmt.Println("A2")
}

func main() {
	f1()
}
