package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
	1. 以切片语法（起始和结束索引号）返回子串时候，其内部依旧指向原字节数组
*/

func StringIndex() {
	s := "abcdefg"
	s1 := s[:3]  // 从头开始，仅指定结束索引位置
	s2 := s[1:4] // 指定开始和结束位置，返回 [start, end]
	s3 := s[2:]  // 指定开始位置，返回后面全部内容

	println(s1, s2, s3)
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
}

func main() {
	StringIndex()
}
