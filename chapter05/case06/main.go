package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
	1. 要修改字符串，须将其转换为可变类型（[]rune 或 []byte），待完成后再转换回来。但不管如何转换，都须重新分配内存，并复制数据。
*/

func pp(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format, *h)
}

func main() {
	s := "hello, world!"
	pp("s: %x\n", &s)

	bs := []byte(s)
	s2 := string(bs)

	pp("string to []byte, bs: %x\n", &bs)
	pp("[]byte to string, s2: %x\n", &s2)

	rs := []rune(s)
	s3 := string(rs)

	pp("string to []rune, rs: %x\n", &rs)
	pp("[]rune to string, s2: %x\n", &s3)
}
