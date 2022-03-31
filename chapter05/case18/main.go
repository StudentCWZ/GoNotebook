package main

import "fmt"

/*
	1. 向切片尾部 slice[len] 添加数据，返回新的切片对象。
	2. 下面案例是超出切片 cap 限制，而非底层数组长度限制，因为 cap 可小于数组长度。新分配数组长度是原 cap 的两倍，而非原数组的两倍。
	3. 并非总是两倍，对于较大切片，会尝试扩容 1/4，以节约内存。
	4. 向 nil 切片追加数据的时，会为其分配底层数组内存
	5. 正是因为存在重新分配底层数组的缘故，在某些场合建议预留足够多的空间，避免中途内存分配和数据复制开销。
*/

func sliceAppend() {
	s := make([]int, 0, 5)
	s1 := append(s, 10)
	s2 := append(s1, 20, 30)       // 追加多个数据
	fmt.Println(s, len(s), cap(s)) // 不会修改原 slice 属性
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Printf("s 的 ptr: %p s1 的 ptr: %p s2 的 ptr: %p\n", &s, &s1, &s2)
}

func sliceCap() {
	s := make([]int, 0, 100)
	s1 := s[:2:4]
	s2 := append(s1, 1, 2, 3, 4, 5, 6) // 超出 s1 cap 限制，分配新底层数组
	fmt.Printf("s1: %p: %v\n", &s1[0], s1)
	fmt.Printf("s2: %p: %v\n", &s2[0], s2)                   // 数组地址不同，确认新分配
	fmt.Printf("s data: %v\n", s[:10])                       // append 并未向原数组写入部分数据
	fmt.Printf("s1 cap: %d, s2 cap: %d\n", cap(s1), cap(s2)) // 新数组是原 cap 两倍
}

func nilSliceAppend() {
	var s []int
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func main() {
	sliceAppend()
	sliceCap()
	nilSliceAppend()
}
