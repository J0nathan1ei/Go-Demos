package main

import "fmt"

type TestRange struct {
	test string
	temp string
}

// TODO 运行前猜测结果
/*
1
2
3

4a -
5a -
6a -

*/
func main() {
	t := []TestRange{
		TestRange{"1", ""},
		TestRange{"2", ""},
		TestRange{"3", ""},
	}

	for k, v := range t {
		v.test += "a"
		v.temp = changeDefaultString(&v.temp)
		fmt.Println(t[k].test, t[k].temp) // NOTE 猜测输出内容 为什么？
	}

	tt := []*TestRange{
		&TestRange{"4", ""},
		&TestRange{"5", ""},
		&TestRange{"6", ""},
	}

	for k, v := range tt {
		v.test += "a"
		v.temp = changeDefaultString(&v.temp)
		fmt.Println(tt[k].test, tt[k].temp) // NOTE 猜测输出内容 为什么？
	}
}

func changeDefaultString(src *string) string {
	test := "-"
	if src == nil || len(*src) < 1 {
		return test
	}
	return *src
}
