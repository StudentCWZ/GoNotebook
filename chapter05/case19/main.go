package main

import "fmt"

/*
	1. 在两个切片对象间复制数据，允许只想同一底层数组，允许目标区间重叠。最终所复制长度以较短的切片长度为准。
	2. 还可直接从字符串中复制数据到 []byte
*/

func copySlice() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[5:8]
	n := copy(s[4:], s1) // 在同一底层数组的不同区间复制
	fmt.Println(n, s1, s)

	s2 := make([]int, 6) // 在不同数组间复制
	n = copy(s2, s)
	fmt.Println(n, s2, s)
}

func copySliceByte() {
	b := make([]byte, 3)
	n := copy(b, "abcde")
	fmt.Println(n, b)
}

func main() {
	copySlice()
	copySliceByte()
}
