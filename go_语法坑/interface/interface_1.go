package main

import "fmt"

func testInter1()interface{}{
	var x *int = nil
	return x
}

func testInter2()interface{}{
	return nil
}

func main(){
	fmt.Println(testInter1() == nil)
	fmt.Println(testInter2() == nil)
}
