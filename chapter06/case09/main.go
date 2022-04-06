/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/6 13:52
*/

package main

import "fmt"

/*
	1. 方法和函数一样，除直接调用外，还可以赋值给变量，或作为参数传递。依据具体引用方式不同，可分为 expression 和 value 两种状态。
*/

// Method Expression
// 通过类型引用的 method expression 会被还原为普通函数样式，receiver 是第一参数，调用时须显式传参。
// 至于类型可以是 T 或者 *T，只要目标方法存在于该类型方法集中即可。

type N int

func (n N) test() {
	fmt.Printf("test.n: %p, %d\n", &n, n)
}

func main() {
	var n N = 25
	fmt.Printf("main.n: %p, %d\n", &n, n)

	f1 := N.test // func(n N)
	f1(n)

	f2 := (*N).test // func(n *N)
	f2(&n)          //按方法集中的签名传递正确类型的参数

	println("----------------")
	// 当然也可以直接以表达式调用。
	N.test(n)
	(*N).test(&n) // 注意：*N 须使用括号，以免语法解析错误
}





