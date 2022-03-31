package demo

import "testing"

/*
	1. 切片只是很小的结构体对象，用来代替数组传参可避免复制开销。
	2. 还有 make 函数允许在运行期动态指定数组长度，绕开了数组类型必须使用编译期常量的限制。
	3. 并非所有时候都适合用切片代替数组，因为切片底层数组可能会在堆上分配内存。而且小数组在栈拷贝的消耗也未必比 make 大
*/

func array() [1024]int {
	var x [1024]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func slice() []int {
	x := make([]int, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}
