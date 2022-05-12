package function

import (
	"fmt"
)

type Temp struct {
	n    int
	M    int
	next *Temp
}

func (t *Temp) func1(mm int) (nn int) {
	fmt.Println(t.n, t.M, t.next)
	t.n = mm
	nn = t.n
	return
}

func func1() { // 无输入参数 无返回值 函数
	t := &Temp{
		n:    11,
		next: nil,
	}
	fmt.Println(t.func1(11))
}

func func2(i int) int { // 输入一个参数 返回一个参数
	return 0
}

// TODO func func3(num1 ,num2 *int, str *string) { }
func func3() (int, int, string) { // 返回多个参数
	return 1, 2, ""
}

func func4() (p *int) { // 返回局部变量地址
	n := 2
	return &n
}

func func5(n int) func() { // 返回一个函数
	return func() {
		for i := 0; i < n; i++ {
			fmt.Println(i)
		}
	}
}
