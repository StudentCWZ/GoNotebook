package main

import (
	"fmt"
	"unsafe"
)

/*
	1. 在分配内存时，字段须做对齐处理，通常以所有字段中最长的基础类型宽度为标准
*/

func main() {
	v1 := struct {
		a byte
		b byte
		c int32 // 对齐宽度 4
	}{}
	v2 := struct {
		a byte
		b byte // 对齐宽度 1
	}{}
	v3 := struct {
		a byte
		b []int // 基础类型 int，对齐宽度 8
		c byte
	}{}
	fmt.Printf("v1: %d, %d\n", unsafe.Alignof(v1), unsafe.Sizeof(v1))
	fmt.Printf("v2: %d, %d\n", unsafe.Alignof(v2), unsafe.Sizeof(v2))
	fmt.Printf("v3: %d, %d\n", unsafe.Alignof(v3), unsafe.Sizeof(v3))
}
