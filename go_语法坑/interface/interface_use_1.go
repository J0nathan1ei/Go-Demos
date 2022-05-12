package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}
type People interface { // NOTE 包含 Animal接口
	Animal
	Name(string) string
}

type Dog struct{} // 全部实现 Animal 接口内函数

func (d *Dog) Speak() string { return "Woof!" }

type Cat struct{} // 全部实现 Animal 接口内函数

func (c *Cat) Speak() string { return "Meow!" }

type Human struct{} // 全部实现 Animal和People 接口内函数

func (l *Human) Speak() string        { return "你好！" }
func (l *Human) Name(v string) string { return "詹姆斯" + v }

func initUse1() {
	aniamls := []Animal{&Dog{}, &Cat{}, &Human{}}
	for _, v := range aniamls {
		fmt.Println(v.Speak())
	}

	var p People = &Human{} // NOTE Human
	fmt.Println(p.Name("test"), p.Speak())

	//	p = &Dog{} // TODO ERROR: *Dog does not implement People (missing Name method)

	var tt *string = nil
	temp(tt)
	temp(nil)
}

func temp(i interface{}) {
	fmt.Println(i == nil, i)

	if i != nil {
		v, ok := i.(*string)                            // NOTE 类型断言
		fmt.Printf("\ti=%v value=%v ok=%v\n", i, v, ok) // 输出interface内值
	}
}

func main(){
	initUse1()
}
