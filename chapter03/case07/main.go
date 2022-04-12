/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/12 09:59
*/

package main

/*
	1. 指针类型支持相等运算符，但不能做加减法运算和类型转换。如果两个指针指向同一地址，或都为 nil，那么它们相等。
*/

func main() {
	x := 10
	p := &x

	// p++	// 报错
	// var p *int = p + 1 // 报错
	p2 := &x
	println(p == p2)
}
