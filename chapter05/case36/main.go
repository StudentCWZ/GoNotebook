package main

import (
	"fmt"
	"unsafe"
)

/*
	1. 比较特殊的是空结构类型字段。如果它是最后一个字段，那么编译器将其当作长度为 1 的类型做对齐处理，以便其地址不会越界，避免引发垃圾回收错误。
*/

func main() {
	v := struct {
		a struct{}
		b int
		c struct{}
	}{}
	s := `
		v: %p ~ %x, size: %d, align: %d
		
		field	address		offset	size
		------+-------------+------+--------
		a	%p	%d	  %d
		b	%p	%d	  %d
		c	%p	%d 	  %d
	`
	fmt.Printf(s,
		&v, uintptr(unsafe.Pointer(&v))+unsafe.Sizeof(v), unsafe.Sizeof(v), unsafe.Alignof(v),
		&v.a, unsafe.Offsetof(v.a), unsafe.Sizeof(v.a),
		&v.b, unsafe.Offsetof(v.b), unsafe.Sizeof(v.b),
		&v.c, unsafe.Offsetof(v.c), unsafe.Sizeof(v.c),
	)
}
