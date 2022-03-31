package main

import "fmt"

/*
	与 C 数组变量隐式作为变量指针使用不同，Go 数组是值类型，赋值和传参操作都会复制整个数组数据。
*/

func test(x [2]int) { // 值传递
	fmt.Printf("x: %p, %v\n", &x, x)
}

func testTwo(x *[2]int) { // 指针传递
	fmt.Printf("x: %p, %v\n", x, *x)
	x[1] += 100
}

func copyValue() {
	a := [2]int{10, 20}
	var b [2]int
	b = a

	fmt.Printf("a: %p, %v\n", &a, a)
	fmt.Printf("b: %p, %v\n", &b, b)
	test(a)
}

func copyPtr() {
	a := [2]int{10, 20}
	testTwo(&a)
	fmt.Printf("a: %p, %v\n", &a, a)
}

func main() {
	copyValue()
	copyPtr()
}
