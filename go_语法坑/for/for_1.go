package main

import "fmt"

func main(){
	d := []int{1,2,3}

	// 仅修改迭代器，不会影响到原切片内容
	for _, v := range d{
		v += 1
	}
	fmt.Println(d)

	// 通过索引修改，才能修改原切片内容
	for k, _ := range d{
		d[k] += 1
	}
	fmt.Println(d)
}
