package main

import (
	"fmt"
	"time"
)

// TODO 运行前猜测结果

func main() {
	deferTest()
	deferTest1()
	deferTest2()
}

func deferTest() {
	defer fmt.Println("1")
	fmt.Println("2")
	defer fmt.Println("3")
	defer func() {
		fmt.Println("4")
	}()
	fmt.Println("5")
}

func deferTest1() {
	m := 100

	defer fmt.Println("1", m)
	defer func() {
		m += 1
		fmt.Println("2", m)
	}()
	defer fmt.Println("3", m)

	fmt.Println("4", m)
	m += 1

	defer func(m int) {
		m += 1
		fmt.Println("5", m)
	}(m)
}

func deferTest2() {
	{
		defer fmt.Println("test")
	}

	time.Sleep(time.Second)

	func() {
		defer fmt.Println("test222")
	}()

	time.Sleep(time.Second)

	fmt.Println("test11")

}
