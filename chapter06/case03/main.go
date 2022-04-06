/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/6 09:44
*/

package main

import "fmt"

/*
	1. 方法可看作特殊的函数，那么 receiver 的类型自然可以是基础类型或指针类型。这会关系到调用时对象实例是否会被复制。
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

func main() {
	var a N = 25
	a.value()
	a.pointer()
	fmt.Printf("a: %p, %v\n", &a, a)
}
