package main

import (
	"fmt"
)

func A3(){
	fmt.Println("A3 recover panic")
	recover()
}

func main() {
	defer A3()
	panic("unknown err")
}
