package main

import (
	"fmt"
	"time"
)

func main(){
	// go 协程捕获了外部变量i
	// 协程创建这个事件和协程执行代码是分离的，
	// 他可以全部创建完再执行，而且主线程和协程是同时运行的(并发)，
	// 也有可能主线程执行完了，协程还没执行
	for i := 0; i < 3; i++ {
		fmt.Printf("第一次 i 产生变化中 %v \n", i)
		go func() {
			fmt.Printf("第一次输出： %v\n", i)
		}()
	}
	time.Sleep(time.Second)


	// 给协程赋予变量，不让协程捕获外部变量
	fmt.Println("不让闭包形成：")
	for i := 0; i < 3; i++ {
		fmt.Printf("第一次 i 产生变化中 %v \n", i)
		go func(num int) {
			fmt.Printf("第一次输出： %v\n", num)
		}(i)
	}
	time.Sleep(time.Second)

}
