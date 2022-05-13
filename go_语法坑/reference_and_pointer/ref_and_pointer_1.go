package main

import (
	"fmt"
	"unsafe"
)

type Tt1 struct {
	ss string
}

func p1() Tt1 {
	var t1 Tt1

	t1 = Tt1{
		"dfsg",
	}
	fmt.Println("1.1", t1, &t1)

	t2 := t1
	fmt.Println("1.2", t2, &t2)
	//	var t3 *Tt1
	t3 := &t1
	fmt.Println("1.3", t3, &t3, *t3)

	t4 := t3
	fmt.Println("1.4", t4, &t4, *t4)

	t5 := &t3
	fmt.Println("1.5", t5, &t5, *t5, **t5)

	t7 := unsafe.Pointer(&t1)
	fmt.Println("1.7", t7, &t7)

	return t1

}
func p2() *Tt1 {

	//var t1 *Tt1

	t1 := &Tt1{
		"dfsg",
	}
	fmt.Println("2.1", t1, &t1, *t1)

	return t1

}
func p3() {
	//普通变量
	d1 := "sdfasf"
	fmt.Println("3.1", d1, &d1)

	d2 := 123
	fmt.Println("3.2", d2, &d2)

	// 非普通变量

	d3 := [4]byte{1, 3, 4}
	fmt.Println("3.3", d3, &d3)

	d4 := []byte{1, 3, 4}
	fmt.Println("3.4", d4, &d4)

}

func p4() {
	A := "hello"
	Z1 := &A
	Z2 := &Z1
	Z3 := &Z2
	fmt.Println("4.1", A, Z1, Z2, Z3, *Z3, **Z3, ***Z3)

}

func main() {

	t1 := p1()
	fmt.Println("1", t1, &t1)

	t2 := p2()
	fmt.Println("2", t2, &t2)

	p3()

	p4()
}
