package main

import (
	"fmt"
	"unsafe"
)

/*
	1. 某些时候，转换操作会拖累算法性能，可尝试用非安全方法进行改善
	2. 该方法利用了 []byte 和 string 头结构部分相同，以非安全的指针类型转换来实现类型变更，从而避免了底层数组的复制。
	3. 利用 append 函数，可将 string 直接追加到 []byte 内
	4. 考虑到字符串只读特征，转换时复制数据到新分配内存是可以理解的。当然，性能同样重要，编译器会为某些场合进行专门优化，避免额外分配和复制操作。
		- 将 []byte 转换为 string key，去 map[string] 查询的时候
		- 将 string 转换为 []byte，进行 for range 迭代时，直接取字节赋值给局部变量。
*/

func toString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func main() {
	bs := []byte("hello, world!")
	s := toString(bs)

	fmt.Printf("bs: %p\n", &bs)
	fmt.Printf("s: %p\n", &s)

	// append 函数
	var b_slice []byte
	b_slice = append(b_slice, "abc"...)
	fmt.Println(b_slice)

	// 编译器会为某些场合进行专门优化
	m := map[string]int{
		"abc": 123,
	}
	key := []byte("abc")
	x, ok := m[string(key)]
	println(x, ok)
}
