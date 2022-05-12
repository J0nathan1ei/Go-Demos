package Struct

import "fmt"

type Info struct { // type 类型名 结构体明细
	str   string
	count int
}

type Detail struct {
	detail struct { // 匿名结构体
		num int
		str string
	}
	nums []int
	sum  int32
	Info        // 整体包含
	str  string // NOTE 与 Info 结构体中str 重名
}

func initStruct1() {
	t := &Detail{}
	t.str = "str"
	t.count = 1
	t.detail.str = "str1"
	t.detail.num = 2
	t.sum = 3
	t.nums = []int{1, 2, 3}
	t.Info.str = "infoStr" // NOTE 赋值Info.str 内容
	fmt.Println(t)

	t = &Detail{
		sum: 111,
		detail: struct {
			num int
			str string
		}{
			num: 8,
			str: "str3",
		},
		nums: []int{4, 5, 6},
		Info: Info{
			count: 7,
			str:   "str2",
		},
	}
	fmt.Println(t)
}

func main(){
	initStruct1()
}