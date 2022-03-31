package main

import (
	"fmt"
	"unsafe"
)

/*
	1. 空结构（struct{}）是指没有字段的结构类型。它比较特殊，因为无论是其自身，还是作为数组元素类型，其长度都为零。
	2. 尽管没有分配数组内存，但依然可以操作元素，对应切片 len、cap 属性很正常。
	3. 实际上，这类长度为零的对象通常都指向 runtime.zerobase 变量。
	4. 空结构可作为通道元素类型，用于事件通知。
*/

func emptyStruct() {
	var a struct{}
	var b [100]struct{}
	println(unsafe.Sizeof(a), unsafe.Sizeof(b))
}

func emptyStructOperate() {
	var d [100]struct{}
	s := d[:]

	d[1] = struct{}{}
	s[2] = struct{}{}

	fmt.Println(s[3], len(s), cap(s))
}

func emptyStructPtr() {
	a := [10]struct{}{}
	b := a[:] // 底层数组指向 runtime。zerobase 变量
	c := [0]int{}

	fmt.Printf("%p, %p, %p\n", &a[0], &b[0], &c)
}

func emptyStructAsChan() {
	exit := make(chan struct{})
	go func() {
		println("hello, world!")
		exit <- struct{}{}
	}()
	<-exit
	println("end ...")
}

func main() {
	emptyStruct()
	emptyStructOperate()
	emptyStructPtr()
	emptyStructAsChan()
}
