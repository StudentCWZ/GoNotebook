/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/6 14:21
*/

package main

import "fmt"

/*
	1. 基于实例或指针引用的 method value，参数签名不会改变，依旧按正常方式调用。
	2. 但当 method value 被赋值给变量或作为参数传递时，会立即计算并复制该方法执行所需的 receiver 对象，与其绑定，以便在稍后执行时，能隐式
       传入 receiver 参数。
*/

type N int

func (n N) test() {
	fmt.Printf("test.n: %p, %v\n", &n, n)
}

func main() {
	var n N = 100
	p := &n

	n++
	f1 := n.test // 因为 test 方法的 receiver 是 N 类型，所以复制 n，等于 101

	n++
	f2 := p.test // 复制 *P，等于 102

	n++
	fmt.Printf("main.n: %p, %v\n", p, n)

	f1()
	f2()
}
