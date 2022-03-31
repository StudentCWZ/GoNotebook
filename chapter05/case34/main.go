package main

import (
	"fmt"
	"unsafe"
)

/*
	1. 内存布局：不管结构体包含多少字段，其内存总是一次性分配的，各字段在相邻的地址空间按定义顺序排列。
	2. 当然，对于引用类型、字符串和指针，结构内存中只包含其基本（头部）数据。还有，所有匿名字段成员也包含在内。
	3. 借助 unsafe 包的相关函数，可输出所有字段的偏移量和长度。
*/

type point struct {
	x, y int
}

type value struct {
	id    int    // 基本类型
	name  string // 字符串
	data  []byte // 引用类型
	next  *value // 指针类型
	point        // 匿名字段
}

func main() {
	v := value{
		id:    1,
		name:  "test",
		data:  []byte{1, 2, 3, 4},
		point: point{x: 100, y: 200},
	}

	s := `
		v: %p ~ %x, size: %d, aligin: %d


		field		address					offset	  					size
		------+----------------------+---------------------------------+---------------------------------------------------
		id		%p			%d		  %d
		name		%p			%d		  %d
		data		%p			%d		  %d
		next		%p			%d		  %d
		x		%p			%d		  %d
		y		%p			%d		  %d
	`
	fmt.Printf(s,
		&v, uintptr(unsafe.Pointer(&v))+unsafe.Sizeof(v), unsafe.Sizeof(v), unsafe.Alignof(v),
		&v.id, unsafe.Offsetof(v.id), unsafe.Sizeof(v.id),
		&v.name, unsafe.Offsetof(v.name), unsafe.Sizeof(v.name),
		&v.data, unsafe.Offsetof(v.data), unsafe.Sizeof(v.data),
		&v.next, unsafe.Offsetof(v.next), unsafe.Sizeof(v.next),
		&v.x, unsafe.Offsetof(v.x), unsafe.Sizeof(v.x),
		&v.y, unsafe.Offsetof(v.y), unsafe.Sizeof(v.y),
	)
}
