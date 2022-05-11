package main

import (
	"fmt"
)

func f() int {
	var i = 0
	defer func() {
		i++
	}()
	return i
}

func f0()(i int){
	defer func(){
		i++
	}()
	return 0
}

func main() {
	d := f()
	d0 := f0()
	fmt.Println(d)
	fmt.Println(d0)
}
