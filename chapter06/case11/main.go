/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/6 14:40
*/

package main

import "fmt"

// 当 method 作为参数时，会复制含 receiver 在内的整个 method value。

type N int

func (n N) test() {
	fmt.Printf("test.n: %p, %v\n", &n, n)
}

func call(m func()) {
	m()
}

func main() {
	var n N = 100
	p := &n

	fmt.Printf("main.n: %p, %v\n", p, n)

	n++
	call(n.test)

	n++
	call(p.test)
}
