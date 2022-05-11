package baseType

import (
	"fmt"
)

type BaseType struct {
	isOk  bool            // 对外包不可见	首字母小写
	Str   string          // 对外包可见		首字母大写
	Num   int             // 可 int32 可 int64 随着运行环境变化
	Num32 int32           // 确定范围 32位
	Num64 int64           // 确定范围 64位
	next  *BaseType       // 链表常用
	str   string          // 底层 []byte
	c     chan bool       // channel 需要初始化
	m     map[string]bool // map     需要初始化 	并发不安全
	s     []interface{}   // 切片
}

func test(arr []byte) {} // 传递指针

func test1(arr [2]byte) {} // 传递整个数组

func test2(arr string) {} // 传递整个字符串 这种时候可用 []byte

func init() {
	var m map[string]map[string]chan int

	type myType map[string]chan int // 自定义类型
	var m1 map[string]myType

	fmt.Println(m, m1)

	// TODO m,m1 完全使用的时候 需要初始化几次？
}

func test3() {
	var m map[string]map[string]chan int
	m = make(map[string]map[string]chan int)
	m["test"] = make(map[string]chan int)
	m["test"]["asdf"] = make(chan int, 1)
	m["test"]["asdf"] <- 2
	fmt.Println(<-m["test"]["asdf"])
}
