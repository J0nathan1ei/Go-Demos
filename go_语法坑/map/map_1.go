package main

import (
	"fmt"
	"sync"
)

func initMap1() {
	// map 结构是并发不安全 需要加锁 sync.RWMutex
	var m map[string]string        // var 名称 map[key类型]value类型
	m = make(map[string]string, 1) // 初始化 make(map类型， 初始大小)
	m["test"] = "asdf"             // 写入
	fmt.Println(m["test"])         // 读取

	// 并发安全 无需加锁 无需初始化
	var nm sync.Map
	nm.Store("test", "asdf")
	if tmp, ok := nm.Load("test"); ok { // tmp 读取接口类型，  ok  是否读取成功 bool类型
		value, ok := tmp.(string) // 类型断言 value 成功转换的string类型，  ok  是否转换成功 bool类型
		fmt.Println(value, ok)
	}

	{
		key := 123
		nm.Store(key, 3.1415)
		if tmp, ok := nm.Load(key); ok { // tmp 读取接口类型，  ok  是否读取成功 bool类型
			value, ok := tmp.(float64) // 类型断言 float64
			fmt.Println(value, ok)
		}
	}
}


func main(){
	initMap1()
}