package main

import (
	"fmt"
)

func initRange1() {
	s := make([]int, 10)

	// for 初始化; 循环判断条件 ; 增量
	for i := 0; i < 10; i++ { // i从0到9循环10次
		fmt.Println(s[i])
	}

	// 死循环
	// for { }
	// for true { }

	for k, v := range s { // k---下标0到9   v---对应下标变量
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	for k := range s { // k---下标0到9
		fmt.Println(k)
	}

	for k, v := range s {
		fmt.Printf("k=%v &k=%v &v=%v &s[k]=%v \n", k, &k, &v, &s[k])
	}
}

func main(){
	initRange1()
}