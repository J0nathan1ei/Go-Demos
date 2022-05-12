package main

import (
	"encoding/json"
	"fmt"
)

type Temp struct {
	M   int     `json:"test"` // NOTE 接收和输出参数为 test
	T   int     `json:"-"`    // NOTE 不接受 和 不输出
	n   string  `json:"n"`    // TODO 不被解析 因为是小写字母开头
	Str *string `json:"str"`  // NOTE 使用指针来解析json的话 就能确认是否传输此值
}

func initJson1() {
	buff := []byte(`{"test":123, "n":"number", "str":"string" }`)
	test := func() {
		var t Temp
		err := json.Unmarshal(buff, &t) // NOTE json解析到当前结构体
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s --- %+v\n", buff, t)
	}
	test()

	buff = []byte(`{"n":"number"}`)
	test()

	buff = []byte(`{"TEST":123, "Number":"number", "str":"string", "key": {"test":1, "count":111} }`)
	var m map[string]interface{}
	err := json.Unmarshal(buff, &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s --- %+v\n", buff, m)
}


func main(){
	initJson1()
}