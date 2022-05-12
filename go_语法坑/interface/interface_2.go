package main

import (
	"fmt"
)

type myInt64 int64 // 自定义int64

func init() {
	var a myInt64 = 32
	var b int64 = 12
	var ok bool

	a = b // TODO ERROR: cannot use b (type int64) as type myInt64 in assignment
	b = a // TODO ERROR: cannot use a (type myInt64) as type int64 in assignment

	a = myInt64(b) // 强制类型转换

	var c interface{} = a
	a = myInt64(c) // TODO ERROR: cannot convert c (type interface {}) to type myInt64: need type assertion

	// NOTE interface的类型断言
	b, ok = c.(int64) // 0 false
	fmt.Println(b, ok)

	a, ok = c.(myInt64) // 12 true
	fmt.Println(a, ok)

	str, ok := c.(string) // "" false
	fmt.Println(str, ok)
}
