package main

import "fmt"

/*
	1. 要分清指针数组和数组指针的区别，指针数组是指元素为指针类型的数组，数组指针是获取数组变量的地址。
	2. 数组指针可以直接用来操作元素
*/

func ptrArrayOrArrayPtr() {
	x, y := 10, 20
	a := [...]*int{&x, &y} // 元素为指针的指针数组
	p := &a                // 存储数组地址的指针
	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", p, p)
}

func ptrArray() {
	a := [...]int{1, 2}
	println(&a, &a[0], &a[1])
	p := &a

	p[1] += 10
	println(p[1])
}

func main() {
	ptrArrayOrArrayPtr()
	ptrArray()
}
