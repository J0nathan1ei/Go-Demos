package main

import (
	"fmt"
)


// 子切片改变了父本的数值
func main() {
	var a [3]int = [3]int{1, 2, 3} // var 数组变量名 [数组长度]数组类型
	fmt.Printf("a=%v len(a)=%v \n", a, len(a))
	// a = 1 2 3 len(a) = 3

	var s []int                                                         // var 切片变量名 []切片类型
	s = a[:1]                                                           // 截取 下标 [0~1)	---> [1]
	fmt.Printf("s=%v len(s)=%v cap(s)=%v \n", s, len(s), cap(s))        //
	// s = 1 len(s) = 1 cap(s) = 1

	s = append(s, 1)                                                    // 追加一个元素
	fmt.Printf("s=%v len(s)=%v cap(s)=%v a=%v\n", s, len(s), cap(s), a) //
	// s = 1 1 	len(s) = 2	 cap(s) = 2  a = 1 2 3

	// s[2] = 2                                                            // panic

	s = append(s, 22, 33)                                               // 追加一个元素
	fmt.Printf("s=%v len(s)=%v cap(s)=%v a=%v\n", s, len(s), cap(s), a) //
	// s = 1 1 22 33 len(s) = 4  cap(s) = 4    a = 1 2 3

	ns := make([]int, 1, 6)                                             // 创建一个新的切片 len=1 cap=6
	fmt.Printf("ns=%v len(ns)=%v cap(ns)=%v \n", ns, len(ns), cap(ns))
	// ns = [] len(ns) = 1  cap(ns) = 6
}
