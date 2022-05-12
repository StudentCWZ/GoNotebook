/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/5/12 09:58
*/

package main

import (
	"fmt"
	"unsafe"
)

/*
	1. 缓存区仅是内部属性，不属于类型组成部分。另外通道变量本身就是指针，可用相等操作符判断是否为同一对象或 nil。
	2. 虽然可传递指针来避免数据复制，但须额外注意数据并发安全。
*/

func main() {
	var a, b = make(chan int, 3), make(chan int)
	var c chan bool
	println(a == b)
	println(c == nil)

	fmt.Printf("%p, %d\n", a, unsafe.Sizeof(a))
}
