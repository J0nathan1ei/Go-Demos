package main

import (
	"fmt"
)

func initPanic1() {
	defer func() {
		if err := recover(); err != nil { // 判断是否有panic
			fmt.Println("init", err)
		}
		fmt.Println("init over")
	}()

	panic1()
	panic2()
}

func panic1() {
	defer func() { // 挂起一个函数
		if err := recover(); err != nil { // 判断是否有panic
			fmt.Println("panic1", err)
		}
		fmt.Println("panic1 over")
	}()

	go panic("Test panic1") // TODO 失控的panic
}

func panic2() {
	defer func() { // 挂起一个函数
		if err := recover(); err != nil { // 判断是否有panic
			fmt.Println("panic2", err)
		}
		fmt.Println("panic2 over")
	}()

	panic("Test panic2")
}


func main(){
	initPanic1()
}