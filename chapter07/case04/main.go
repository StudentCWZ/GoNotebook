/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/7 09:39
*/

package main

/*
	1. 超集接口变量可隐式转换成子集，反过来不行。
*/

type stringer interface {
	string() string
}

type tester interface {
	stringer
	test()
}

type data struct{}

func (*data) test() {}

func (data) string() string {
	return ""
}

func pp(a stringer) {
	println(a.string())
}

func main() {
	var d data

	var t tester = &d
	pp(t) // 隐式转换为子集接口

	var a stringer = t
	println(a.string())

	// var t2 tester = s // 报错
}
