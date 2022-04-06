/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/6 09:55
*/

package main

import "fmt"

/*
	1. 可使用实例值或指针调用方法，编译器会根据方法 receiver 类型自动在基础类型和指针类型间转换
	2. 不能用多级指针调用方法
*/

type N int

func (n N) value() {
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Printf("p: %p, %v\n", n, *n)
}

func test() {
	//var a N = 25
	//p := &a
	//p2 := &p

	//p2.value()	// 报错：多级指针调用方法
	//p2.pointer()
}

func main() {
	var a N = 25
	p := &a
	a.value()
	a.pointer()

	p.value()
	p.pointer()
}
