package main

import "fmt"

func main(){
	m := make(map[int]int)
	m[0] = 1
	d, ok := m[1]
	fmt.Println(d,ok)
}
