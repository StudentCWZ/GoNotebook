/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/6 14:46
*/

package main

import "fmt"

/*
	1. 如果目标方法的 receiver 是指针类型，那么被复制的仅是指针。
*/

type N int

func (n *N) test() {
	fmt.Printf("test.n: %p, %v\n", n, *n)
}

func main() {
	var n N = 100
	p := &n

	n++
	f1 := n.test // 因为 test 方法的 receiver 是 *N 类型。所以复制 &n

	n++
	f2 := p.test // 复制 p 指针

	n++
	fmt.Printf("main.n: %p, %v\n", p, n)

	f1() // 延迟调用，n == 103
	f2()
}
