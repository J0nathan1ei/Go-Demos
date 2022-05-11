package main

import "fmt"

func init() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	for i := 10; i < 13; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

	for i := 20; i < 23; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}

	fmt.Println("888")

	fmt.Println(test())
}

func test() (i int) { // 规则三
	defer func() { i++ }()
	return 555
}
