/*
   @Author: StudentCWZ
   @Description:
   @File: main.go
   @Software: GoLand
   @Project: GoNotebook
   @Date: 2022/4/9 21:24
*/

package main

import (
	"fmt"
	"unsafe"
)

/*
	1. 常量值也可以是某些编译器能计算出结果的表达式，如 unsafe.Sizeof、len、cap 等。
*/

func main() {
	const (
		ptrSize = unsafe.Sizeof(uintptr(0))
		strSize = len("hello, world!")
	)
	fmt.Println(ptrSize, strSize)
}
